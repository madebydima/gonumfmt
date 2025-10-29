package gonumfmt

import (
	"testing"
)

func BenchmarkFormatDecimal(b *testing.B) {
	f := NewFormatter(WithLocale("en-US"))
	number := 1234567.890

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Format(number)
	}
}

func BenchmarkFormatCurrency(b *testing.B) {
	f := NewFormatter(
		WithLocale("en-US"),
		WithCurrency("USD"),
	)
	number := 1234567.890

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Format(number)
	}
}

func BenchmarkFormatPercent(b *testing.B) {
	f := NewFormatter(
		WithLocale("en-US"),
		WithStyle(Percent),
	)
	number := 0.1567

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Format(number)
	}
}

func BenchmarkFormatCompact(b *testing.B) {
	f := NewFormatter(
		WithLocale("en-US"),
		WithStyle(Compact),
	)
	number := 1500000.0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Format(number)
	}
}

func BenchmarkFormatWithHighPrecision(b *testing.B) {
	f := NewFormatter(
		WithLocale("en-US"),
		WithPrecision(0, 10),
	)
	number := 3.1415926535

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Format(number)
	}
}

func BenchmarkUtilityFunctions(b *testing.B) {
	numbers := []float64{1234.56, 789012.34, 0.1567, 1500000.0}

	b.Run("Format", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, num := range numbers {
				Format(num)
			}
		}
	})

	b.Run("FormatCurrency", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, num := range numbers {
				FormatCurrency(num, "USD")
			}
		}
	})

	b.Run("FormatPercent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, num := range numbers {
				FormatPercent(num)
			}
		}
	})

	b.Run("FormatCompact", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, num := range numbers {
				FormatCompact(num)
			}
		}
	})
}

func BenchmarkDifferentLocales(b *testing.B) {
	locales := []string{"en", "ru", "de", "fr", "ja", "zh"}
	number := 1234567.890

	for _, locale := range locales {
		b.Run(locale, func(b *testing.B) {
			f := NewFormatter(WithLocale(locale))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				f.Format(number)
			}
		})
	}
}

func BenchmarkFormatterCreation(b *testing.B) {
	b.Run("Simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewFormatter()
		}
	})

	b.Run("Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewFormatter(
				WithLocale("ru-RU"),
				WithCurrency("RUB"),
				WithPrecision(2, 4),
				WithSignDisplay(SignAlways),
			)
		}
	})
}
