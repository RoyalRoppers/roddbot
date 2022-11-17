package unmarshal

import (
	"reflect"

	"github.com/bwmarrin/discordgo"
)

func Unmarshal(m []*discordgo.ApplicationCommandInteractionDataOption, dest any) {
	destType := reflect.TypeOf(dest)
	val := reflect.ValueOf(&dest)

	fields := reflect.VisibleFields(destType.Elem())

	// cachable : )
	typeMap := make(map[string]string)
	for _, sf := range fields {
		typeMap[sf.Tag.Get("optname")] = sf.Name
	}

	for _, opt := range m {
		x, ok := typeMap[opt.Name]
		if !ok {
			continue
		}
		field := reflect.Indirect(val).Elem().Elem().FieldByName(x)

		switch opt.Type {
		case discordgo.ApplicationCommandOptionString:
			x := opt.StringValue()
			if field.Type().Kind() == reflect.Pointer {
				field.Set(reflect.ValueOf(&x))
			} else {
				field.SetString(x)
			}
		case discordgo.ApplicationCommandOptionInteger:
			x := opt.IntValue()
			if field.Type().Kind() == reflect.Pointer {
				y := int(x) // hack
				field.Set(reflect.ValueOf(&y))
			} else {
				field.SetInt(x)
			}

		case discordgo.ApplicationCommandOptionBoolean:
			x := opt.BoolValue()
			if field.Type().Kind() == reflect.Pointer {
				field.Set(reflect.ValueOf(&x))
			} else {
				field.SetBool(x)
			}
		default:
			panic("type not supported")
		}
	}
}
