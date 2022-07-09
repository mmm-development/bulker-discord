package locale

import "github.com/bwmarrin/discordgo"

type LocaleDict map[discordgo.Locale]map[string]string

func (ld LocaleDict) Get(locale discordgo.Locale, key string) string {
	if v, ok := ld[locale][key]; ok {
		return v
	}
	return ld[DefLocale][key]
}

func (ld LocaleDict) LocaleMap(key string) *map[discordgo.Locale]string {
	to_return := make(map[discordgo.Locale]string)
	for locale, translations := range ld {
		if v, ok := translations[key]; ok {
			to_return[locale] = v
		}
	}
	return &to_return
}

var (
	L         LocaleDict       = make(LocaleDict)
	DefLocale discordgo.Locale = discordgo.Russian
)
