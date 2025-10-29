package gonumfmt

// currencyData содержит расширенные данные о валютах из CLDR
var currencyData = map[string]*CurrencyData{
	"USD": {Symbol: "$", Name: "US Dollar", Format: "{symbol}{number}", Spacing: ""},
	"EUR": {Symbol: "€", Name: "Euro", Format: "{number} {symbol}", Spacing: " "},
	"GBP": {Symbol: "£", Name: "British Pound", Format: "{symbol}{number}", Spacing: ""},
	"JPY": {Symbol: "¥", Name: "Japanese Yen", Format: "{symbol}{number}", Spacing: ""},
	"CNY": {Symbol: "¥", Name: "Chinese Yuan", Format: "{symbol}{number}", Spacing: ""},
	"RUB": {Symbol: "₽", Name: "Russian Ruble", Format: "{number} {symbol}", Spacing: " "},
	"INR": {Symbol: "₹", Name: "Indian Rupee", Format: "{symbol}{number}", Spacing: ""},
	"BRL": {Symbol: "R$", Name: "Brazilian Real", Format: "{symbol} {number}", Spacing: " "},
	"CAD": {Symbol: "CA$", Name: "Canadian Dollar", Format: "{symbol}{number}", Spacing: ""},
	"AUD": {Symbol: "A$", Name: "Australian Dollar", Format: "{symbol}{number}", Spacing: ""},
	"CHF": {Symbol: "CHF", Name: "Swiss Franc", Format: "{number} {symbol}", Spacing: " "},
	"SEK": {Symbol: "kr", Name: "Swedish Krona", Format: "{number} {symbol}", Spacing: " "},
	"NOK": {Symbol: "kr", Name: "Norwegian Krone", Format: "{number} {symbol}", Spacing: " "},
	"DKK": {Symbol: "kr", Name: "Danish Krone", Format: "{number} {symbol}", Spacing: " "},
	"PLN": {Symbol: "zł", Name: "Polish Zloty", Format: "{number} {symbol}", Spacing: " "},
	"TRY": {Symbol: "₺", Name: "Turkish Lira", Format: "{symbol}{number}", Spacing: ""},
	"KRW": {Symbol: "₩", Name: "South Korean Won", Format: "{symbol}{number}", Spacing: ""},
	"MXN": {Symbol: "MX$", Name: "Mexican Peso", Format: "{symbol}{number}", Spacing: ""},
	"SAR": {Symbol: "ر.س", Name: "Saudi Riyal", Format: "{number} {symbol}", Spacing: " "},
	"AED": {Symbol: "د.إ", Name: "UAE Dirham", Format: "{number} {symbol}", Spacing: " "},
	// Добавляем больше валют по мере необходимости
}

// getCurrencyData возвращает данные о валюте
func getCurrencyData(currencyCode string) *CurrencyData {
	if data, exists := currencyData[currencyCode]; exists {
		return data
	}

	// Fallback для неизвестных валют
	return &CurrencyData{
		Symbol:  currencyCode,
		Name:    currencyCode,
		Format:  "{number} {symbol}",
		Spacing: " ",
	}
}
