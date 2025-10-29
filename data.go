package gonumfmt

import "strings"

// localeData хранит встроенные данные CLDR для поддерживаемых локалей
var localeData = map[string]*LocaleData{
	"en": {
		DecimalSeparator: ".",
		GroupSeparator:   ",",
		PercentSymbol:    "%",
		NegativePattern:  "-{number}",
		PositivePattern:  "{number}",
		PercentPattern:   "{number}%",
		MinusSign:        "-",
		PlusSign:         "+",
		Exponential:      "E",
		CurrencyFormats: map[string]*CurrencyData{
			"USD": {Symbol: "$", Name: "US Dollar", Format: "{symbol}{number}"},
			"EUR": {Symbol: "€", Name: "Euro", Format: "{symbol} {number}"},
			"GBP": {Symbol: "£", Name: "British Pound", Format: "{symbol}{number}"},
			"JPY": {Symbol: "¥", Name: "Japanese Yen", Format: "{symbol}{number}"},
			"CNY": {Symbol: "¥", Name: "Chinese Yuan", Format: "{symbol}{number}"},
			"RUB": {Symbol: "₽", Name: "Russian Ruble", Format: "{number} {symbol}"},
		},
		CompactPatterns: map[CompactRange]*CompactPattern{
			Thousand: {Short: "0K", Long: "0 thousand"},
			Million:  {Short: "0M", Long: "0 million"},
			Billion:  {Short: "0B", Long: "0 billion"},
			Trillion: {Short: "0T", Long: "0 trillion"},
		},
	},
	"ru": {
		DecimalSeparator: ",",
		GroupSeparator:   " ",
		PercentSymbol:    "%",
		NegativePattern:  "-{number}",
		PositivePattern:  "{number}",
		PercentPattern:   "{number}%",
		MinusSign:        "-",
		PlusSign:         "+",
		Exponential:      "E",
		CurrencyFormats: map[string]*CurrencyData{
			"USD": {Symbol: "$", Name: "доллар США", Format: "{number} {symbol}"},
			"EUR": {Symbol: "€", Name: "евро", Format: "{number} {symbol}"},
			"GBP": {Symbol: "£", Name: "фунт стерлингов", Format: "{number} {symbol}"},
			"JPY": {Symbol: "¥", Name: "иена", Format: "{number} {symbol}"},
			"CNY": {Symbol: "¥", Name: "юань", Format: "{number} {symbol}"},
			"RUB": {Symbol: "₽", Name: "российский рубль", Format: "{number} {symbol}"},
		},
		CompactPatterns: map[CompactRange]*CompactPattern{
			Thousand: {Short: "0 тыс.", Long: "0 тысяч"},
			Million:  {Short: "0 млн", Long: "0 миллионов"},
			Billion:  {Short: "0 млрд", Long: "0 миллиардов"},
			Trillion: {Short: "0 трлн", Long: "0 триллионов"},
		},
	},
	"de": {
		DecimalSeparator: ",",
		GroupSeparator:   ".",
		PercentSymbol:    "%",
		NegativePattern:  "-{number}",
		PositivePattern:  "{number}",
		PercentPattern:   "{number}%",
		MinusSign:        "-",
		PlusSign:         "+",
		Exponential:      "E",
		CurrencyFormats: map[string]*CurrencyData{
			"USD": {Symbol: "$", Name: "US-Dollar", Format: "{number} {symbol}"},
			"EUR": {Symbol: "€", Name: "Euro", Format: "{number} {symbol}"},
			"GBP": {Symbol: "£", Name: "Britisches Pfund", Format: "{number} {symbol}"},
			"JPY": {Symbol: "¥", Name: "Japanischer Yen", Format: "{number} {symbol}"},
		},
		CompactPatterns: map[CompactRange]*CompactPattern{
			Thousand: {Short: "0 Tsd.", Long: "0 Tausend"},
			Million:  {Short: "0 Mio.", Long: "0 Millionen"},
			Billion:  {Short: "0 Mrd.", Long: "0 Milliarden"},
			Trillion: {Short: "0 Bio.", Long: "0 Billionen"},
		},
	},
	"fr": {
		DecimalSeparator: ",",
		GroupSeparator:   " ",
		PercentSymbol:    "%",
		NegativePattern:  "-{number}",
		PositivePattern:  "{number}",
		PercentPattern:   "{number}%",
		MinusSign:        "-",
		PlusSign:         "+",
		Exponential:      "E",
		CurrencyFormats: map[string]*CurrencyData{
			"USD": {Symbol: "$", Name: "dollar américain", Format: "{number} {symbol}"},
			"EUR": {Symbol: "€", Name: "euro", Format: "{number} {symbol}"},
			"GBP": {Symbol: "£", Name: "livre sterling", Format: "{number} {symbol}"},
		},
		CompactPatterns: map[CompactRange]*CompactPattern{
			Thousand: {Short: "0 k", Long: "0 mille"},
			Million:  {Short: "0 M", Long: "0 million"},
			Billion:  {Short: "0 Md", Long: "0 milliard"},
			Trillion: {Short: "0 bn", Long: "0 billion"},
		},
	},
	"ja": {
		DecimalSeparator: ".",
		GroupSeparator:   ",",
		PercentSymbol:    "%",
		NegativePattern:  "-{number}",
		PositivePattern:  "{number}",
		PercentPattern:   "{number}%",
		MinusSign:        "-",
		PlusSign:         "+",
		Exponential:      "E",
		CurrencyFormats: map[string]*CurrencyData{
			"USD": {Symbol: "$", Name: "アメリカドル", Format: "{symbol}{number}"},
			"EUR": {Symbol: "€", Name: "ユーロ", Format: "{symbol}{number}"},
			"GBP": {Symbol: "£", Name: "英ポンド", Format: "{symbol}{number}"},
			"JPY": {Symbol: "¥", Name: "日本円", Format: "{symbol}{number}"},
			"CNY": {Symbol: "元", Name: "中国人民元", Format: "{symbol}{number}"},
		},
		CompactPatterns: map[CompactRange]*CompactPattern{
			Thousand: {Short: "0千", Long: "0千"},
			Million:  {Short: "0百万", Long: "0百万"},
			Billion:  {Short: "0十億", Long: "0十億"},
			Trillion: {Short: "0兆", Long: "0兆"},
		},
	},
	"zh": {
		DecimalSeparator: ".",
		GroupSeparator:   ",",
		PercentSymbol:    "%",
		NegativePattern:  "-{number}",
		PositivePattern:  "{number}",
		PercentPattern:   "{number}%",
		MinusSign:        "-",
		PlusSign:         "+",
		Exponential:      "E",
		CurrencyFormats: map[string]*CurrencyData{
			"USD": {Symbol: "US$", Name: "美元", Format: "{symbol}{number}"},
			"EUR": {Symbol: "€", Name: "欧元", Format: "{symbol}{number}"},
			"GBP": {Symbol: "£", Name: "英镑", Format: "{symbol}{number}"},
			"JPY": {Symbol: "¥", Name: "日元", Format: "{symbol}{number}"},
			"CNY": {Symbol: "¥", Name: "人民币", Format: "{symbol}{number}"},
		},
		CompactPatterns: map[CompactRange]*CompactPattern{
			Thousand: {Short: "0千", Long: "0千"},
			Million:  {Short: "0百万", Long: "0百万"},
			Billion:  {Short: "0十亿", Long: "0十亿"},
			Trillion: {Short: "0兆", Long: "0兆"},
		},
	},
}

// SupportedLocales возвращает список поддерживаемых локалей
func SupportedLocales() []string {
	locales := make([]string, 0, len(localeData))
	for locale := range localeData {
		locales = append(locales, locale)
	}
	return locales
}

// IsLocaleSupported проверяет поддержку локали
func IsLocaleSupported(locale string) bool {
	locale = normalizeLocale(locale)

	if _, exists := localeData[locale]; exists {
		return true
	}

	// Проверяем базовую локаль
	base := strings.Split(locale, "-")[0]
	_, exists := localeData[base]
	return exists
}
