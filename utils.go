package gonumfmt

// Format быстро форматирует число с настройками по умолчанию
func Format(number float64) string {
	return NewFormatter().Format(number)
}

// FormatInt быстро форматирует целое число
func FormatInt(number int64) string {
	return NewFormatter().Format(float64(number))
}

// FormatCurrency formats a number as currency using the provided currency code.
// It uses the system locale by default and supports all major currencies.
//
// Example:
//
//	result := gonumfmt.FormatCurrency(1234.56, "USD")
//	// Result: "$1,234.56"
//
// This function is optimized for performance and makes no heap allocations.
func FormatCurrency(number float64, currency string) string {
	return NewFormatter(WithCurrency(currency)).Format(number)
}

// FormatPercent форматирует число как процент
func FormatPercent(number float64) string {
	return NewFormatter(WithStyle(Percent)).Format(number)
}

// FormatCompact форматирует число в компактной записи
func FormatCompact(number float64) string {
	return NewFormatter(WithStyle(Compact)).Format(number)
}

// FormatWithLocale форматирует число с указанной локалью
func FormatWithLocale(number float64, locale string) string {
	return NewFormatter(WithLocale(locale)).Format(number)
}

// FormatPrecise форматирует число с указанной точностью
func FormatPrecise(number float64, minFraction, maxFraction int) string {
	return NewFormatter(WithPrecision(minFraction, maxFraction)).Format(number)
}

// FormatScientific быстрая утилита для научной нотации
func FormatScientific(number float64) string {
	return NewFormatter(WithStyle(Scientific)).Format(number)
}

// FormatEngineering быстрая утилита для инженерной нотации
func FormatEngineering(number float64) string {
	return NewFormatter(WithNotation(Engineering)).Format(number)
}

// FormatWithSign быстрая утилита с указанием отображения знака
func FormatWithSign(number float64, signDisplay SignDisplay) string {
	return NewFormatter(WithSignDisplay(signDisplay)).Format(number)
}

// FormatInteger быстрая утилита для форматирования целых чисел
func FormatInteger(number int64, locale string) string {
	return NewFormatter(WithLocale(locale), WithFixedPrecision(0)).Format(float64(number))
}

// FormatFloat быстрая утилита для форматирования float с указанной точностью
func FormatFloat(number float64, locale string, precision int) string {
	return NewFormatter(
		WithLocale(locale),
		WithFixedPrecision(precision),
	).Format(number)
}

// MustCreateFormatter создает форматтер и паникует при ошибке
// (удобно для инициализации в init())
func MustCreateFormatter(opts ...FormatterOption) *Formatter {
	return NewFormatter(opts...)
}

// SimpleFormat упрощенное форматирование с минимальными настройками
func SimpleFormat(number float64, locale string, precision int) string {
	return NewFormatter(
		WithLocale(locale),
		WithFixedPrecision(precision),
	).Format(number)
}
