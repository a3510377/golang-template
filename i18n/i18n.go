package i18n

func Init(configLocale string) string {
	if configLocale != "" {
		if locale := findMatchLocale(configLocale); locale != "" {
			return setLocale(locale)
		}
	}

	if osLocale := GetLocaleFromOS(); osLocale != "" {
		if locale := findMatchLocale(osLocale); locale != "" {
			return setLocale(locale)
		}
	}

	return setLocale("en")
}

func Tr(msg string, args ...any) string {
	// return po.Get(msg, args...)
	return msg
}
