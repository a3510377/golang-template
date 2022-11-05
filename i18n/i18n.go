package i18n

func Init(configLocale string) string {
	if configLocale != "" {
		if locale := findMatchLocale(configLocale); locale != "" {
			setLocale(locale)
			return locale
		}
	}

	if osLocale := GetLocaleFromOS(); osLocale != "" {
		if locale := findMatchLocale(osLocale); locale != "" {
			setLocale(locale)
			return locale
		}
	}

	setLocale("en")
	return "en"
}
