package gonumfmt

import (
	"math"
	"strconv"
	"strings"
)

// Formatter основной тип для форматирования чисел
type Formatter struct {
	options Options
	locale  *LocaleData
}

// NewFormatter создает новый форматтер с указанными опциями
func NewFormatter(opts ...FormatterOption) *Formatter {
	options := DefaultOptions()
	for _, opt := range opts {
		opt(&options)
	}

	// Загружаем данные локали
	locale := GetLocaleData(options.Locale)
	if locale == nil {
		// Fallback на английскую локаль
		locale = GetLocaleData("en")
	}

	return &Formatter{
		options: options,
		locale:  locale,
	}
}

// Format форматирует число в строку
func (f *Formatter) Format(number float64) string {
	// Проверка специальных значений
	if math.IsNaN(number) {
		return "NaN"
	}
	if math.IsInf(number, 1) {
		return "∞"
	}
	if math.IsInf(number, -1) {
		return "-∞"
	}

	switch f.options.Style {
	case Decimal:
		return f.formatDecimal(number)
	case Currency:
		return f.formatCurrency(number)
	case Percent:
		return f.formatPercent(number)
	case Scientific:
		return f.formatScientific(number)
	case Compact:
		return f.formatCompact(number)
	default:
		return f.formatDecimal(number)
	}
}

// FormatInt форматирует целое число
func (f *Formatter) FormatInt(number int64) string {
	return f.Format(float64(number))
}

// formatDecimal форматирует число в десятичном формате
func (f *Formatter) formatDecimal(number float64) string {
	// Обработка очень маленьких чисел
	if isVerySmallNumber(number) {
		return f.formatVerySmallNumber(number)
	}

	// Определяем знак
	sign := f.getSign(number)
	absNumber := math.Abs(number)

	// Округляем число
	rounded := f.roundNumber(absNumber)

	// Разделяем на целую и дробную части
	intPart, fracPart := f.splitNumber(rounded)

	// Форматируем целую часть с группировкой
	formattedInt := f.formatIntegerPart(intPart)

	// Форматируем дробную часть
	formattedFrac := f.formatFractionalPart(fracPart)

	// Собираем результат
	var result string
	if formattedFrac != "" {
		result = formattedInt + f.locale.DecimalSeparator + formattedFrac
	} else {
		result = formattedInt
	}

	// Применяем шаблон знака
	return f.applySignPattern(result, sign)
}

// formatCurrency форматирует число как валюту
func (f *Formatter) formatCurrency(number float64) string {
	decimalStr := f.formatDecimal(number)

	if f.options.Currency == "" {
		return decimalStr
	}

	currencyData, exists := f.locale.CurrencyFormats[f.options.Currency]
	if !exists {
		// Fallback для неизвестной валюты - код валюты перед числом
		return f.options.Currency + decimalStr
	}

	var currencyDisplay string
	switch f.options.CurrencyDisplay {
	case CurrencySymbol:
		currencyDisplay = currencyData.Symbol
	case CurrencyCode:
		currencyDisplay = f.options.Currency
	case CurrencyName:
		currencyDisplay = currencyData.Name
	default:
		currencyDisplay = currencyData.Symbol
	}

	// Применяем формат валюты
	format := currencyData.Format
	format = strings.ReplaceAll(format, "{symbol}", currencyDisplay)
	format = strings.ReplaceAll(format, "{code}", f.options.Currency)
	format = strings.ReplaceAll(format, "{number}", decimalStr)

	return format
}

// formatPercent форматирует число как процент
func (f *Formatter) formatPercent(number float64) string {
	// Умножаем на 100 для процентов
	percentNumber := number * 100
	decimalStr := f.formatDecimal(percentNumber)

	// Применяем шаблон процентов
	format := f.locale.PercentPattern
	format = strings.ReplaceAll(format, "{number}", decimalStr)
	format = strings.ReplaceAll(format, "{symbol}", f.locale.PercentSymbol)

	return format
}

// formatScientific форматирует число в научной нотации
func (f *Formatter) formatScientific(number float64) string {
	if number == 0 {
		return "0" + f.locale.Exponential + "0"
	}

	sign := f.getSign(number)
	absNumber := math.Abs(number)

	// Определяем экспоненту
	exponent := 0
	if absNumber >= 1 {
		for absNumber >= 10 {
			absNumber /= 10
			exponent++
		}
	} else {
		for absNumber < 1 && absNumber > 0 {
			absNumber *= 10
			exponent--
		}
	}

	// Округляем мантиссу
	mantissa := f.roundNumber(absNumber)

	// Форматируем мантиссу как десятичное число
	mantissaStr := f.formatDecimal(mantissa)

	// Форматируем экспоненту без лишних нулей
	exponentStr := strconv.Itoa(exponent)
	if exponent >= 0 {
		// Убираем + для положительной экспоненты
		exponentStr = strconv.Itoa(exponent)
	} else {
		exponentStr = "-" + strconv.Itoa(-exponent)
	}

	result := mantissaStr + f.locale.Exponential + exponentStr
	return f.applySignPattern(result, sign)
}

// formatCompact форматирует число в компактной записи
func (f *Formatter) formatCompact(number float64) string {
	absNumber := math.Abs(number)
	sign := f.getSign(number)

	// Определяем диапазон
	var rangeType CompactRange
	var divisor float64

	switch {
	case absNumber >= 1e12:
		rangeType = Trillion
		divisor = 1e12
	case absNumber >= 1e9:
		rangeType = Billion
		divisor = 1e9
	case absNumber >= 1e6:
		rangeType = Million
		divisor = 1e6
	case absNumber >= 1e3:
		rangeType = Thousand
		divisor = 1e3
	default:
		// Число слишком маленькое для компактной записи
		return f.formatDecimal(number)
	}

	// Вычисляем компактное значение
	compactValue := number / divisor
	absCompactValue := math.Abs(compactValue)

	// Округляем до нужной точности
	roundedValue := f.roundNumber(absCompactValue)

	// Форматируем число
	numberStr := f.formatDecimal(roundedValue)

	// Получаем шаблон для компактной записи
	pattern := f.getCompactPattern(rangeType)

	// Заменяем шаблон
	result := strings.ReplaceAll(pattern, "0", numberStr)

	return f.applySignPattern(result, sign)
}

// getCompactPattern возвращает шаблон для компактной записи
func (f *Formatter) getCompactPattern(rangeType CompactRange) string {
	patternData, exists := f.locale.CompactPatterns[rangeType]
	if !exists {
		return ""
	}

	if f.options.CompactDisplay == Long {
		return patternData.Long
	}
	return patternData.Short
}

// getSign возвращает знак числа и как его отображать
func (f *Formatter) getSign(number float64) string {
	if number < 0 {
		return f.locale.MinusSign
	}

	// Для положительных чисел и нуля
	switch f.options.SignDisplay {
	case SignAlways:
		return f.locale.PlusSign
	case SignExceptZero:
		if number > 0 {
			return f.locale.PlusSign
		}
	}

	return ""
}

// applySignPattern применяет шаблон знака к отформатированному числу
func (f *Formatter) applySignPattern(numberStr, sign string) string {
	// Для SignNever игнорируем все знаки
	if f.options.SignDisplay == SignNever {
		return numberStr
	}

	if sign == "" {
		return numberStr
	}

	pattern := f.locale.PositivePattern
	if sign == f.locale.MinusSign {
		pattern = f.locale.NegativePattern
	}

	result := strings.ReplaceAll(pattern, "{number}", numberStr)
	result = strings.ReplaceAll(result, "{sign}", sign)

	return result
}

// roundNumber округляет число согласно настройкам
func (f *Formatter) roundNumber(number float64) float64 {
	// Для очень маленьких чисел возвращаем как есть
	if number == 0 || math.Abs(number) < 1e-324 {
		return number
	}

	scale := math.Pow10(f.options.MaximumFractionDigits)
	var rounded float64

	switch f.options.RoundingMode {
	case RoundHalfUp:
		rounded = math.Round(number*scale) / scale
	case RoundHalfDown:
		rounded = f.roundHalfDown(number, scale)
	case RoundHalfEven:
		rounded = f.roundHalfEven(number, scale)
	case RoundCeiling:
		rounded = math.Ceil(number*scale) / scale
	case RoundFloor:
		rounded = math.Floor(number*scale) / scale
	case RoundUp:
		rounded = f.roundUp(number, scale)
	case RoundDown:
		rounded = f.roundDown(number, scale)
	default:
		rounded = math.Round(number*scale) / scale
	}

	return rounded
}

// Вспомогательные методы для округления
func (f *Formatter) roundHalfDown(number, scale float64) float64 {
	scaled := number * scale
	floor := math.Floor(scaled)
	ceil := math.Ceil(scaled)

	if scaled-floor > 0.5 {
		return ceil / scale
	}
	return floor / scale
}

func (f *Formatter) roundHalfEven(number, scale float64) float64 {
	scaled := number * scale
	floor := math.Floor(scaled)
	ceil := math.Ceil(scaled)

	if scaled-floor == 0.5 {
		// Банковское округление: к ближайшему четному
		if int64(floor)%2 == 0 {
			return floor / scale
		}
		return ceil / scale
	}

	if scaled-floor > 0.5 {
		return ceil / scale
	}
	return floor / scale
}

func (f *Formatter) roundUp(number, scale float64) float64 {
	if number > 0 {
		return math.Ceil(number*scale) / scale
	}
	return math.Floor(number*scale) / scale
}

func (f *Formatter) roundDown(number, scale float64) float64 {
	if number > 0 {
		return math.Floor(number*scale) / scale
	}
	return math.Ceil(number*scale) / scale
}

// splitNumber разделяет число на целую и дробную части
func (f *Formatter) splitNumber(number float64) (string, string) {
	// Используем точное строковое представление для избежания ошибок округления
	str := strconv.FormatFloat(number, 'f', -1, 64)

	parts := strings.Split(str, ".")
	intPart := parts[0]
	fracPart := ""
	if len(parts) > 1 {
		fracPart = parts[1]
	}

	return intPart, fracPart
}

// formatIntegerPart форматирует целую часть числа
func (f *Formatter) formatIntegerPart(intPart string) string {
	// Добавляем ведущие нули если нужно
	for len(intPart) < f.options.MinimumIntegerDigits {
		intPart = "0" + intPart
	}

	if !f.options.UseGrouping {
		return intPart
	}

	return f.applyGrouping(intPart)
}

// formatFractionalPart форматирует дробную часть
func (f *Formatter) formatFractionalPart(fracPart string) string {
	// Если дробная часть пустая, но нужно минимальное количество знаков
	if fracPart == "" && f.options.MinimumFractionDigits > 0 {
		fracPart = strings.Repeat("0", f.options.MinimumFractionDigits)
	}

	// Обрезаем или дополняем нулями до нужной длины
	if len(fracPart) > f.options.MaximumFractionDigits {
		fracPart = fracPart[:f.options.MaximumFractionDigits]
	} else if len(fracPart) < f.options.MinimumFractionDigits {
		fracPart += strings.Repeat("0", f.options.MinimumFractionDigits-len(fracPart))
	}

	// Удаляем trailing zeros если нужно
	if f.options.TrimTrailingZeros {
		fracPart = strings.TrimRight(fracPart, "0")
	}

	return fracPart
}

// applyGrouping применяет группировку цифр
func (f *Formatter) applyGrouping(number string) string {
	if len(number) <= 3 {
		return number
	}

	var result strings.Builder
	groupSize := 3
	firstGroupSize := len(number) % groupSize

	if firstGroupSize == 0 {
		firstGroupSize = groupSize
	}

	result.WriteString(number[:firstGroupSize])

	for i := firstGroupSize; i < len(number); i += groupSize {
		result.WriteString(f.locale.GroupSeparator)
		result.WriteString(number[i : i+groupSize])
	}

	return result.String()
}
