package translate

type TranslateService interface {
	LanguageList() string
	LanguageCode(lang string) string
	DetectLanguage(sentece string) string
}
