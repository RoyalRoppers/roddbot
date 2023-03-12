package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/movitz-s/roddbot/internal/permissions"
	"go.uber.org/zap"
)

type RoleMapPayload struct {
	AdminRoleID  *string `optname:"admin-role"`
	PlayerRoleID *string `optname:"player-role"`
}

func (b *bot) handlePermissionMap(m *discordgo.InteractionCreate, p *RoleMapPayload) {

	x, err := b.sess.GuildRoles(m.GuildID)
	if err != nil {
		b.log.Error("could not get roles", zap.Error(err))
		return
	}

	xx := make(map[string]struct{})
	for _, r := range x {
		xx[r.ID] = struct{}{}
	}

	if p.AdminRoleID != nil {
		_, ok := xx[*p.AdminRoleID]
		if !ok {
			b.reply(m.Interaction, "admin role id, not an actual role id :/")
			return
		}
	}
	if p.PlayerRoleID != nil {
		_, ok := xx[*p.PlayerRoleID]
		if !ok {
			b.reply(m.Interaction, "player role id, not an actual role id :/")
			return
		}
	}

	err = b.perm.MapRoles(m.GuildID, p.AdminRoleID, p.PlayerRoleID)
	if err != nil {
		b.log.Error("could not map roles", zap.Error(err))
		b.reply(m.Interaction, "something went wrong")
		return
	}

	b.reply(m.Interaction, "Role mappings updated!")

}

func (b *bot) hasPermission(m *discordgo.InteractionCreate, p permissions.Permission) bool {
	mapping := b.perm.GetMappings(m.GuildID)

	// note: the @everyone role is wierd, should not be used here

	permms := []permissions.Permission{}
	if m.Member.Permissions&discordgo.PermissionAdministrator != 0 {
		permms = append(permms, permissions.Roles.Admin...)
	}
	for _, v := range m.Member.Roles {
		if mapping.AdminRoleID.String == v {
			permms = append(permms, permissions.Roles.Admin...)
		}
		if !mapping.PlayerRoleID.Valid || mapping.PlayerRoleID.String == v {
			permms = append(permms, permissions.Roles.Player...)
		}
	}

	for _, p2 := range permms {
		if p2 == p {
			return true
		}
	}
	return false
}
