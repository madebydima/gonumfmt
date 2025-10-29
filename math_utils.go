// math_utils.go - исправленная версия
package gonumfmt

import (
	"math"
	"strconv"
	"strings"
)

// isVerySmallNumber проверяет, является ли число очень маленьким
func isVerySmallNumber(number float64) bool {
	return number != 0 && math.Abs(number) < 1e-15
}

// isExtremelySmallNumber проверяет экстремально маленькие числа
func isExtremelySmallNumber(number float64) bool {
	return number != 0 && math.Abs(number) < 1e-100
}

// formatExtremelySmallNumber форматирует экстремально маленькие числа
func (f *Formatter) formatExtremelySmallNumber(number float64) string {
	if number == 0 {
		return "0"
	}

	// Для экстремально маленьких чисел используем научную нотацию
	// Убрали неиспользуемые переменные sign и absNumber
	return f.formatScientific(number)
}

// formatVerySmallNumber форматирует очень маленькие числа
func (f *Formatter) formatVerySmallNumber(number float64) string {
	if number == 0 {
		return "0"
	}

	if isExtremelySmallNumber(number) {
		return f.formatExtremelySmallNumber(number)
	}

	sign := f.getSign(number)
	absNumber := math.Abs(number)

	// Используем точное строковое представление с увеличенной точностью
	str := f.formatExactString(absNumber)

	// Применяем форматирование локали
	result := f.applyLocaleFormatting(str)

	return f.applySignPattern(result, sign)
}

// formatExactString возвращает точное строковое представление числа
func (f *Formatter) formatExactString(number float64) string {
	// Для очень маленьких чисел увеличиваем точность
	if isVerySmallNumber(number) {
		precision := f.options.MaximumFractionDigits + countLeadingZeros(number) + 10
		if precision > 100 {
			precision = 100
		}
		// Используем формат 'f' для избежания научной нотации
		return strconv.FormatFloat(number, 'f', precision, 64)
	}
	return strconv.FormatFloat(number, 'f', -1, 64)
}

// countLeadingZeros считает количество ведущих нулей в дробной части
func countLeadingZeros(number float64) int {
	if number == 0 {
		return 0
	}

	absNumber := math.Abs(number)
	if absNumber >= 1 {
		return 0
	}

	count := 0
	for absNumber < 1 {
		absNumber *= 10
		count++
	}
	return count - 1
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

	// Обрабатываем дробную часть
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

// roundToPrecision округляет число до указанной точности
func roundToPrecision(number float64, precision int, mode RoundingMode) float64 {
	if precision <= 0 {
		return math.Round(number)
	}

	scale := math.Pow10(precision)
	scaled := number * scale

	var rounded float64
	switch mode {
	case RoundHalfUp:
		rounded = math.Round(scaled)
	case RoundHalfDown:
		if scaled-math.Floor(scaled) == 0.5 {
			rounded = math.Floor(scaled)
		} else {
			rounded = math.Round(scaled)
		}
	case RoundHalfEven:
		rounded = roundHalfEven(scaled)
	case RoundCeiling:
		rounded = math.Ceil(scaled)
	case RoundFloor:
		rounded = math.Floor(scaled)
	case RoundUp:
		if number > 0 {
			rounded = math.Ceil(scaled)
		} else {
			rounded = math.Floor(scaled)
		}
	case RoundDown:
		if number > 0 {
			rounded = math.Floor(scaled)
		} else {
			rounded = math.Ceil(scaled)
		}
	default:
		rounded = math.Round(scaled)
	}

	return rounded / scale
}

// roundHalfEven реализует банковское округление
func roundHalfEven(number float64) float64 {
	intPart, fracPart := math.Modf(number)
	if math.Abs(fracPart) == 0.5 {
		// Банковское округление: к ближайшему четному
		if int64(intPart)%2 == 0 {
			return math.Floor(number)
		}
		return math.Ceil(number)
	}
	return math.Round(number)
}
