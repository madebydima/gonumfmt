package gonumfmt_test

import (
	"fmt"

	"github.com/madebydima/gonumfmt"
)

func Example_basicUsage() {
	// Простое форматирование
	result := gonumfmt.Format(1234567.890)
	fmt.Println(result)
	// Output: 1,234,567.89
}

func Example_differentLocales() {
	// Английская локаль
	en := gonumfmt.NewFormatter(gonumfmt.WithLocale("en"))
	fmt.Println(en.Format(1234567.890))

	// Русская локаль
	ru := gonumfmt.NewFormatter(gonumfmt.WithLocale("ru"))
	fmt.Println(ru.Format(1234567.890))

	// Немецкая локаль
	de := gonumfmt.NewFormatter(gonumfmt.WithLocale("de"))
	fmt.Println(de.Format(1234567.890))

	// Output:
	// 1,234,567.89
	// 1 234 567,89
	// 1.234.567,89
}

func Example_currencyFormatting() {
	// Доллары в американском формате
	usd := gonumfmt.NewFormatter(
		gonumfmt.WithLocale("en"),
		gonumfmt.WithCurrency("USD"),
	)
	fmt.Println(usd.Format(1234.56))

	// Евро в немецком формате
	eur := gonumfmt.NewFormatter(
		gonumfmt.WithLocale("de"),
		gonumfmt.WithCurrency("EUR"),
	)
	fmt.Println(eur.Format(99.99))

	// Рубли в русском формате
	rub := gonumfmt.NewFormatter(
		gonumfmt.WithLocale("ru"),
		gonumfmt.WithCurrency("RUB"),
	)
	fmt.Println(rub.Format(1234.56))

	// Output:
	// $1,234.56
	// 99,99 €
	// 1 234,56 ₽
}

func Example_percentFormatting() {
	// Процентное форматирование
	percent := gonumfmt.NewFormatter(gonumfmt.WithStyle(gonumfmt.Percent))
	fmt.Println(percent.Format(0.1567))

	// С высокой точностью
	precisePercent := gonumfmt.NewFormatter(
		gonumfmt.WithStyle(gonumfmt.Percent),
		gonumfmt.WithPrecision(1, 3),
	)
	fmt.Println(precisePercent.Format(0.1234))

	// Output:
	// 15.67%
	// 12.34%
}

func Example_compactFormatting() {
	// Компактная запись
	compact := gonumfmt.NewFormatter(gonumfmt.WithStyle(gonumfmt.Compact))

	fmt.Println(compact.Format(1500))       // 1.5K
	fmt.Println(compact.Format(1500000))    // 1.5M
	fmt.Println(compact.Format(1500000000)) // 1.5B

	// Output:
	// 1.5K
	// 1.5M
	// 1.5B
}

func Example_precisionControl() {
	// Фиксированная точность
	fixed := gonumfmt.NewFormatter(gonumfmt.WithFixedPrecision(2))
	fmt.Println(fixed.Format(123.456))

	// Диапазон точности
	rangePrecision := gonumfmt.NewFormatter(gonumfmt.WithPrecision(1, 4))
	fmt.Println(rangePrecision.Format(123.456))

	// Удаление нулей
	trimZeros := gonumfmt.NewFormatter(
		gonumfmt.WithPrecision(0, 3),
		gonumfmt.WithTrailingZeroRemoval(true),
	)
	fmt.Println(trimZeros.Format(123.450))

	// Output:
	// 123.46
	// 123.456
	// 123.45
}

func Example_signDisplay() {
	// Всегда показывать знак
	alwaysSign := gonumfmt.NewFormatter(gonumfmt.WithSignDisplay(gonumfmt.SignAlways))

	fmt.Println(alwaysSign.Format(123.45))
	fmt.Println(alwaysSign.Format(-123.45))
	fmt.Println(alwaysSign.Format(0.0))

	// Output:
	// +123.45
	// -123.45
	// +0
}

func Example_utilityFunctions() {
	// Быстрые утилиты
	fmt.Println(gonumfmt.Format(1234.56))
	fmt.Println(gonumfmt.FormatCurrency(1234.56, "USD"))
	fmt.Println(gonumfmt.FormatPercent(0.1567))
	fmt.Println(gonumfmt.FormatCompact(1500000))

	// Output:
	// 1,234.56
	// $1,234.56
	// 15.67%
	// 1.5M
}

func Example_scientificNotation() {
	// Научная нотация
	scientific := gonumfmt.NewFormatter(gonumfmt.WithStyle(gonumfmt.Scientific))

	fmt.Println(scientific.Format(1234567.89))
	fmt.Println(scientific.Format(0.000000123))

	// Output:
	// 1.23456789E6
	// 1.23E-7
}

func Example_customConfiguration() {
	// Создание кастомного форматтера
	customFormatter := gonumfmt.NewFormatter(
		gonumfmt.WithLocale("ru-RU"),
		gonumfmt.WithCurrency("USD"),
		gonumfmt.WithPrecision(2, 4),
		gonumfmt.WithSignDisplay(gonumfmt.SignAlways),
		gonumfmt.WithGrouping(true),
	)

	result := customFormatter.Format(1234.5678)
	fmt.Println(result)

	// Output:
	// +1 234,5678 $
}
