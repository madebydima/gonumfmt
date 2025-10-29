package gonumfmt

import (
	"math"
	"strconv"
	"strings"
)

// isVerySmallNumber проверяет, является ли число очень маленьким
func isVerySmallNumber(number float64) bool {
	return number != 0 && math.Abs(number) < 1e-10
}

// formatVerySmallNumber форматирует очень маленькие числа
func (f *Formatter) formatVerySmallNumber(number float64) string {
	if number == 0 {
		return "0"
	}

	sign := f.getSign(number)
	absNumber := math.Abs(number)

	// Для очень маленьких чисел используем точное строковое представление с большой точностью
	requiredPrecision := f.options.MaximumFractionDigits + countLeadingZeros(absNumber) + 2
	if requiredPrecision > 100 {
		requiredPrecision = 100
	}

	// Используем точное строковое представление
	str := strconv.FormatFloat(absNumber, 'f', requiredPrecision, 64)

	// Применяем форматирование локали
	result := f.applyLocaleFormatting(str)

	return f.applySignPattern(result, sign)
}

// countLeadingZeros считает количество ведущих нулей
func countLeadingZeros(number float64) int {
	if number == 0 {
		return 0
	}

	absNumber := math.Abs(number)
	if absNumber >= 1 {
		return 0
	}

	// Используем научную нотацию для определения порядка
	str := strconv.FormatFloat(number, 'e', -1, 64)

	// Парсим экспоненту (формат: "2.3e-51")
	parts := strings.Split(str, "e-")
	if len(parts) != 2 {
		return 0
	}

	exponent, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}

	return exponent - 1
}

// applyLocaleFormatting применяет форматирование локали к строковому числу
func (f *Formatter) applyLocaleFormatting(numberStr string) string {
	parts := strings.Split(numberStr, ".")
	intPart := parts[0]
	fracPart := ""
	if len(parts) > 1 {
		fracPart = parts[1]
	}

	// Применяем группировку к целой части
	if f.options.UseGrouping {
		intPart = f.applyGrouping(intPart)
	}

	// Обрабатываем дробную часть согласно настройкам
	if fracPart != "" {
		// Обрезаем до максимального количества знаков
		if len(fracPart) > f.options.MaximumFractionDigits {
			fracPart = fracPart[:f.options.MaximumFractionDigits]
		}
		// Добавляем минимальное количество знаков
		for len(fracPart) < f.options.MinimumFractionDigits {
			fracPart += "0"
		}
		// Удаляем trailing zeros если нужно
		if f.options.TrimTrailingZeros {
			fracPart = strings.TrimRight(fracPart, "0")
		}
	}

	if fracPart != "" {
		return intPart + f.locale.DecimalSeparator + fracPart
	}
	return intPart
}
