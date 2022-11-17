package unmarshal_test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/movitz-s/roddbot/internal/unmarshal"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	type X struct {
		Name        string  `optname:"name"`
		Description *string `optname:"desc"`
		Age         int     `optname:"alder"`
		NonAge      *int    `optname:"age"`
		Huh         bool    `optname:"huh"`
	}
	x := X{}

	hej := []*discordgo.ApplicationCommandInteractionDataOption{
		{
			Name:  "name",
			Type:  discordgo.ApplicationCommandOptionString,
			Value: "hej",
		},
		{
			Name:  "alder",
			Type:  discordgo.ApplicationCommandOptionInteger,
			Value: 123.0, // wierd discordgo behaviour
		},
		{
			Name:  "desc",
			Type:  discordgo.ApplicationCommandOptionString,
			Value: "asdaidjf",
		},
		{
			Name:  "huh",
			Type:  discordgo.ApplicationCommandOptionBoolean,
			Value: true,
		},
	}
	unmarshal.Unmarshal(hej, &x)

	assert.Equal(t, x.Name, "hej")
	assert.Equal(t, x.Age, 123)
	assert.Equal(t, *x.Description, "asdaidjf")
	assert.Nil(t, x.NonAge)
	assert.Equal(t, x.Huh, true)
}
