package i18n

import (
	"embed"
	"os"
	"strings"

	"test/utils"

	"github.com/leonelquinteros/gotext"
)

var (
	po *gotext.Po

	//go:embed data/*.po
	contents embed.FS
)

func init() {
	po = new(gotext.Po)
}

func FormatLocale(locale string) string {
	return strings.ReplaceAll(strings.ToLower(locale), "-", "_")
}

func GetLocaleFromOS() string {
	local := getLocale()

	if local == "" {
		local = getLocalFromEnv()
	}

	return FormatLocale(local)
}

func getLocalFromEnv() string {
	locale := os.Getenv("LC_ALL")

	if locale == "" {
		locale = os.Getenv("LANG")
	}
	return strings.Split(locale, ",")[0]
}

func SupportedLocales() []string {
	var locales []string

	files, _ := contents.ReadDir("data")

	for _, file := range files {
		locales = append(locales, strings.TrimSuffix(file.Name(), ".po"))
	}

	return locales
}

func findMatchLocale(locale string) string {
	locale = FormatLocale(locale)

	var matchLocale string
	var matchLocales []string

	for _, l := range SupportedLocales() {
		if l == locale {
			matchLocale = l
		} else if strings.HasPrefix(l, locale) {
			matchLocales = append(matchLocales, l)
		}
	}

	if matchLocale != "" {
		return matchLocale
	} else if len(matchLocales) > 0 {
		return matchLocales[0]
	}

	return ""
}

func setLocale(locale string) {
	defer utils.Recover()

	poFile, err := contents.ReadFile("data/" + locale + ".po")
	if err != nil {
		panic(err)
	}

	po = new(gotext.Po)
	po.Parse(poFile)
}
