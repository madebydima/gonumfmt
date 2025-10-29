package gonumfmt

import (
	"math"
	"testing"
)

func TestFormatter_Decimal(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   string
		expected string
	}{
		{"English large number", 1234567.890, "en", "1,234,567.89"},
		{"English small number", 1234.56, "en", "1,234.56"},
		{"English integer", 12345.000, "en", "12,345"},
		{"Russian large number", 1234567.890, "ru", "1 234 567,89"},
		{"Russian small number", 1234.56, "ru", "1 234,56"},
		{"Russian integer", 12345.000, "ru", "12 345"},
		{"German number", 1234567.890, "de", "1.234.567,89"},
		{"French number", 1234567.890, "fr", "1 234 567,89"},
		{"Japanese number", 1234567.890, "ja", "1,234,567.89"},
		{"Chinese number", 1234567.890, "zh", "1,234,567.89"},
		{"Very small number", 0.0000000000000000000000000000000000000000000000000023000, "en", "0"},
		{"Very small number default", 0.0000000000000000000000000000000000000000000000000023000, "en", "0"},
		{"Zero", 0.0, "en", "0"},
		{"Negative number", -1234.56, "en", "-1,234.56"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(WithLocale(tt.locale))
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Format(%f) with locale %s = %s, expected %s",
					tt.number, tt.locale, result, tt.expected)
			}
		})
	}
}

func TestFormatter_Currency(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   string
		currency string
		expected string
	}{
		{"USD English", 1234.56, "en", "USD", "$1,234.56"},
		{"USD Russian", 1234.56, "ru", "USD", "1 234,56 $"},
		{"EUR English", 99.99, "en", "EUR", "€99.99"},
		{"EUR German", 99.99, "de", "EUR", "99,99 €"},
		{"RUB Russian", 1234.56, "ru", "RUB", "1 234,56 ₽"},
		{"GBP English", 1234.56, "en", "GBP", "£1,234.56"},
		{"JPY Japanese", 1234.56, "ja", "JPY", "¥1,234.56"},
		{"CNY Chinese", 1234.56, "zh", "CNY", "¥1,234.56"},
		{"Unknown Currency", 123.45, "en", "XYZ", "XYZ123.45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithLocale(tt.locale),
				WithCurrency(tt.currency),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Currency format(%f) %s in %s = %s, expected %s",
					tt.number, tt.currency, tt.locale, result, tt.expected)
			}
		})
	}
}

func TestFormatter_Percent(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   string
		expected string
	}{
		{"English percent", 0.1567, "en", "15.67%"},
		{"Russian percent", 0.1567, "ru", "15,67%"},
		{"German percent", 0.1567, "de", "15,67%"},
		{"French percent", 0.1567, "fr", "15,67%"},
		{"Japanese percent", 0.1567, "ja", "15.67%"},
		{"Chinese percent", 0.1567, "zh", "15.67%"},
		{"Large percent", 1.5, "en", "150%"},
		{"Small percent", 0.001, "en", "0.1%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithLocale(tt.locale),
				WithStyle(Percent),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Percent format(%f) in %s = %s, expected %s",
					tt.number, tt.locale, result, tt.expected)
			}
		})
	}
}

func TestFormatter_Compact(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   string
		expected string
	}{
		{"English thousand", 1500.0, "en", "1.5K"},
		{"English million", 1500000.0, "en", "1.5M"},
		{"English billion", 1500000000.0, "en", "1.5B"},
		{"English trillion", 1500000000000.0, "en", "1.5T"},
		{"Russian thousand", 1500.0, "ru", "1,5 тыс."},
		{"Russian million", 1500000.0, "ru", "1,5 млн"},
		{"Russian billion", 1500000000.0, "ru", "1,5 млрд"},
		{"Russian trillion", 1500000000000.0, "ru", "1,5 трлн"},
		{"German million", 1500000.0, "de", "1,5 Mio."},
		{"French million", 1500000.0, "fr", "1,5 M"},
		{"Japanese million", 1500000.0, "ja", "1.5百万"},
		{"Chinese million", 1500000.0, "zh", "1.5百万"},
		{"Small number compact", 999.0, "en", "999"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithLocale(tt.locale),
				WithStyle(Compact),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Compact format(%f) in %s = %s, expected %s",
					tt.number, tt.locale, result, tt.expected)
			}
		})
	}
}

func TestFormatter_Precision(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		minFrac  int
		maxFrac  int
		trim     bool
		expected string
	}{
		{"Fixed precision", 123.456, 2, 2, false, "123.46"},
		{"Variable precision", 123.456, 1, 3, false, "123.456"},
		{"Trim zeros enabled", 123.450, 0, 3, true, "123.45"},
		{"Trim zeros disabled", 123.450, 0, 3, false, "123.450"},
		{"More decimals than needed", 123.4, 0, 5, true, "123.4"},
		{"Minimum fraction digits", 123.0, 2, 2, false, "123.00"},
		{"Very small number", 0.000000001, 0, 10, true, "0.000000001"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithPrecision(tt.minFrac, tt.maxFrac),
				WithTrailingZeroRemoval(tt.trim),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Precision format(%f) [%d-%d, trim:%v] = %s, expected %s",
					tt.number, tt.minFrac, tt.maxFrac, tt.trim, result, tt.expected)
			}
		})
	}
}

func TestFormatter_SignDisplay(t *testing.T) {
	tests := []struct {
		name        string
		number      float64
		signDisplay SignDisplay
		expected    string
	}{
		{"Auto positive", 123.45, SignAuto, "123.45"},
		{"Auto negative", -123.45, SignAuto, "-123.45"},
		{"Auto zero", 0.0, SignAuto, "0"},
		{"Always positive", 123.45, SignAlways, "+123.45"},
		{"Always negative", -123.45, SignAlways, "-123.45"},
		{"Always zero", 0.0, SignAlways, "+0"},
		{"Never positive", 123.45, SignNever, "123.45"},
		{"Never negative", -123.45, SignNever, "123.45"},
		{"Never zero", 0.0, SignNever, "0"},
		{"ExceptZero positive", 123.45, SignExceptZero, "123.45"},
		{"ExceptZero negative", -123.45, SignExceptZero, "-123.45"},
		{"ExceptZero zero", 0.0, SignExceptZero, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(WithSignDisplay(tt.signDisplay))
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("SignDisplay format(%f) with %v = %s, expected %s",
					tt.number, tt.signDisplay, result, tt.expected)
			}
		})
	}
}

func TestFormatter_Scientific(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		expected string
	}{
		{"Large number", 1234567.89, "1.23456789E6"},
		{"Small number", 0.000123, "1.23E-4"},
		{"Very small number", 0.0000000000000000000000000000000000000000000000000023, "2.3E-51"},
		{"One", 1.0, "1E0"},
		{"Zero", 0.0, "0E0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithStyle(Scientific),
				WithPrecision(2, 8),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Scientific format(%f) = %s, expected %s",
					tt.number, result, tt.expected)
			}
		})
	}
}

func TestFormatter_VerySmallNumbers(t *testing.T) {
	tests := []struct {
		name      string
		number    float64
		locale    string
		precision int
		expected  string
	}{
		{"Very small number English", 0.0000000000000000000000000000000000000000000000000023, "en", 60, "0.0000000000000000000000000000000000000000000000000023"},
		{"Very small number Russian", 0.0000000000000000000000000000000000000000000000000023, "ru", 60, "0,0000000000000000000000000000000000000000000000000023"},
		{"Small number with precision", 0.000000001, "en", 10, "0.000000001"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithLocale(tt.locale),
				WithPrecision(0, tt.precision),
				WithTrailingZeroRemoval(false),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Very small number format(%e) with precision %d = %s, expected %s",
					tt.number, tt.precision, result, tt.expected)
			}
		})
	}
}

func TestUtilities(t *testing.T) {
	t.Run("Format utility", func(t *testing.T) {
		result := Format(1234.56)
		if result != "1,234.56" {
			t.Errorf("Format(1234.56) = %s, expected 1,234.56", result)
		}
	})

	t.Run("FormatInt utility", func(t *testing.T) {
		result := FormatInt(1234567)
		if result != "1,234,567" {
			t.Errorf("FormatInt(1234567) = %s, expected 1,234,567", result)
		}
	})

	t.Run("FormatCurrency utility", func(t *testing.T) {
		result := FormatCurrency(1234.56, "USD")
		if result != "$1,234.56" {
			t.Errorf("FormatCurrency(1234.56, 'USD') = %s, expected $1,234.56", result)
		}
	})

	t.Run("FormatPercent utility", func(t *testing.T) {
		result := FormatPercent(0.1567)
		if result != "15.67%" {
			t.Errorf("FormatPercent(0.1567) = %s, expected 15.67%%", result)
		}
	})

	t.Run("FormatCompact utility", func(t *testing.T) {
		result := FormatCompact(1500000.0)
		if result != "1.5M" {
			t.Errorf("FormatCompact(1500000) = %s, expected 1.5M", result)
		}
	})
}

func TestSystemLocale(t *testing.T) {
	// Тестируем определение системной локали
	locale := getSystemLocale()
	if locale == "" {
		t.Error("System locale should not be empty")
	}
	t.Logf("System locale: %s", locale)
}

func TestCurrencyFormats(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   string
		currency string
		display  CurrencyDisplay
		expected string
	}{
		{"USD Symbol US", 1234.56, "en", "USD", CurrencySymbol, "$1,234.56"},
		{"USD Code US", 1234.56, "en", "USD", CurrencyCode, "USD1,234.56"},
		{"USD Name US", 1234.56, "en", "USD", CurrencyName, "US Dollar1,234.56"},
		{"EUR Symbol DE", 99.99, "de", "EUR", CurrencySymbol, "99,99 €"},
		{"RUB Symbol RU", 1234.56, "ru", "RUB", CurrencySymbol, "1 234,56 ₽"},
		{"JPY Symbol JP", 1234.56, "ja", "JPY", CurrencySymbol, "¥1,234.56"},
		{"Unknown Currency", 123.45, "en", "XYZ", CurrencySymbol, "XYZ123.45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithLocale(tt.locale),
				WithCurrency(tt.currency),
				WithCurrencyDisplay(tt.display),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Currency format failed: got %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestCompactPrecision(t *testing.T) {
	tests := []struct {
		name      string
		number    float64
		precision int
		expected  string
	}{
		{"Default precision", 1234567, 2, "1.23M"},
		{"Zero precision", 1234567, 0, "1M"},
		{"High precision", 1234567, 4, "1.2346M"},
		{"Large number precision", 1234567890123, 1, "1.2T"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithStyle(Compact),
				WithCompactPrecision(tt.precision),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Compact precision %d failed: got %s, expected %s",
					tt.precision, result, tt.expected)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		options  []FormatterOption
		expected string
	}{
		{"NaN", math.NaN(), nil, "NaN"},
		{"Positive infinity", math.Inf(1), nil, "∞"},
		{"Negative infinity", math.Inf(-1), nil, "-∞"},
		{"Very large number", 1.7976931348623157e+100, nil, "1.7976931348623157E100"},
		{"Grouping disabled", 1234567.89, []FormatterOption{WithGrouping(false)}, "1234567.89"},
		{"High precision", 123.456789, []FormatterOption{WithPrecision(0, 6)}, "123.456789"},
		{"No fraction digits", 123.456, []FormatterOption{WithFixedPrecision(0)}, "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(tt.options...)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Edge case %s = %s, expected %s",
					tt.name, result, tt.expected)
			}
		})
	}
}

func TestAllExampleCases(t *testing.T) {
	// Русские примеры
	ruFormatter := NewFormatter(
		WithLocale("ru"),
		WithPrecision(0, 4),
		WithTrailingZeroRemoval(false),
	)

	cases := []struct {
		input    float64
		expected string
	}{
		{1234567.890, "1 234 567,89"},
		{1234.56, "1 234,56"},
		{0.1567, "0,1567"},
		{1234567, "1 234 567"},
		{12345.000, "12 345"},
	}

	for i, tc := range cases {
		result := ruFormatter.Format(tc.input)
		if result != tc.expected {
			t.Errorf("Russian case %d failed: input %f, got %s, expected %s",
				i, tc.input, result, tc.expected)
		}
	}

	// Английские примеры
	enFormatter := NewFormatter(
		WithLocale("en"),
		WithPrecision(0, 4),
		WithTrailingZeroRemoval(false),
	)

	enCases := []struct {
		input    float64
		expected string
	}{
		{1234567.890, "1,234,567.89"},
		{1234.56, "1,234.56"},
		{0.1567, "0.1567"},
		{1234567, "1,234,567"},
		{12345.000, "12,345"},
	}

	for i, tc := range enCases {
		result := enFormatter.Format(tc.input)
		if result != tc.expected {
			t.Errorf("English case %d failed: input %f, got %s, expected %s",
				i, tc.input, result, tc.expected)
		}
	}
}

func TestPercentFormatting(t *testing.T) {
	tests := []struct {
		name     string
		number   float64
		locale   string
		expected string
	}{
		{"English percent", 0.1567, "en", "15.67%"},
		{"Russian percent", 0.1567, "ru", "15,67%"},
		{"100%", 1.0, "en", "100%"},
		{"More than 100%", 1.5, "en", "150%"},
		{"Very small percent", 0.0001, "en", "0.01%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFormatter(
				WithLocale(tt.locale),
				WithStyle(Percent),
			)
			result := f.Format(tt.number)
			if result != tt.expected {
				t.Errorf("Percent formatting failed: got %s, expected %s", result, tt.expected)
			}
		})
	}
}
