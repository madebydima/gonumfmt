package gonumfmt

// Style определяет стиль форматирования
type Style int

const (
	Decimal Style = iota
	Currency
	Percent
	Scientific
	Compact
)

// CurrencyDisplay определяет как отображать валюту
type CurrencyDisplay int

const (
	CurrencySymbol CurrencyDisplay = iota
	CurrencyCode
	CurrencyName
)

// CompactDisplay определяет тип компактного отображения
type CompactDisplay int

const (
	Short CompactDisplay = iota
	Long
)

// Notation определяет нотацию форматирования
type Notation int

const (
	Standard Notation = iota
	ScientificNotation
	Engineering
)

// SignDisplay определяет отображение знака
type SignDisplay int

const (
	SignAuto SignDisplay = iota
	SignAlways
	SignNever
	SignExceptZero
)

// RoundingMode определяет режим округления
type RoundingMode int

const (
	RoundHalfEven RoundingMode = iota
	RoundHalfUp
	RoundHalfDown
	RoundCeiling
	RoundFloor
	RoundDown
	RoundUp
)

// Options содержит все настройки форматирования
type Options struct {
	Locale                string
	Style                 Style
	Currency              string
	CurrencyDisplay       CurrencyDisplay
	UseGrouping           bool
	MinimumIntegerDigits  int
	MinimumFractionDigits int
	MaximumFractionDigits int
	RoundingMode          RoundingMode
	CompactDisplay        CompactDisplay
	CompactPrecision      int
	Notation              Notation
	SignDisplay           SignDisplay
	TrimTrailingZeros     bool
}

// FormatterOption функция для настройки форматирования
type FormatterOption func(*Options)

// DefaultOptions возвращает настройки по умолчанию
func DefaultOptions() Options {
	return Options{
		Locale:                getSystemLocale(),
		Style:                 Decimal,
		UseGrouping:           true,
		MinimumIntegerDigits:  1,
		MinimumFractionDigits: 0,
		MaximumFractionDigits: 3,
		RoundingMode:          RoundHalfEven,
		CompactDisplay:        Short,
		CompactPrecision:      2,
		Notation:              Standard,
		SignDisplay:           SignAuto,
		TrimTrailingZeros:     true,
	}
}

// WithLocale устанавливает локаль
func WithLocale(locale string) FormatterOption {
	return func(o *Options) {
		o.Locale = locale
	}
}

// WithStyle устанавливает стиль форматирования
func WithStyle(style Style) FormatterOption {
	return func(o *Options) {
		o.Style = style
	}
}

// WithCurrency устанавливает валюту
func WithCurrency(currency string) FormatterOption {
	return func(o *Options) {
		o.Currency = currency
		o.Style = Currency
	}
}

// WithCurrencyDisplay устанавливает отображение валюты
func WithCurrencyDisplay(display CurrencyDisplay) FormatterOption {
	return func(o *Options) {
		o.CurrencyDisplay = display
	}
}

// WithGrouping включает/выключает группировку цифр
func WithGrouping(useGrouping bool) FormatterOption {
	return func(o *Options) {
		o.UseGrouping = useGrouping
	}
}

// WithPrecision устанавливает точность
func WithPrecision(minFraction, maxFraction int) FormatterOption {
	return func(o *Options) {
		o.MinimumFractionDigits = minFraction
		o.MaximumFractionDigits = maxFraction
	}
}

// WithFixedPrecision устанавливает фиксированную точность
func WithFixedPrecision(precision int) FormatterOption {
	return func(o *Options) {
		o.MinimumFractionDigits = precision
		o.MaximumFractionDigits = precision
	}
}

// WithIntegerDigits устанавливает минимальное количество целых цифр
func WithIntegerDigits(digits int) FormatterOption {
	return func(o *Options) {
		o.MinimumIntegerDigits = digits
	}
}

// WithRoundingMode устанавливает режим округления
func WithRoundingMode(mode RoundingMode) FormatterOption {
	return func(o *Options) {
		o.RoundingMode = mode
	}
}

// WithCompactDisplay устанавливает компактное отображение
func WithCompactDisplay(display CompactDisplay) FormatterOption {
	return func(o *Options) {
		o.CompactDisplay = display
		o.Style = Compact
	}
}

// WithCompactPrecision устанавливает точность для компактной записи
func WithCompactPrecision(precision int) FormatterOption {
	return func(o *Options) {
		o.CompactPrecision = precision
	}
}

// WithNotation устанавливает нотацию
func WithNotation(notation Notation) FormatterOption {
	return func(o *Options) {
		o.Notation = notation
	}
}

// WithSignDisplay устанавливает отображение знака
func WithSignDisplay(signDisplay SignDisplay) FormatterOption {
	return func(o *Options) {
		o.SignDisplay = signDisplay
	}
}

// WithTrailingZeroRemoval включает/выключает удаление нулей
func WithTrailingZeroRemoval(trim bool) FormatterOption {
	return func(o *Options) {
		o.TrimTrailingZeros = trim
	}
}
