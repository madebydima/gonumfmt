package gonumfmt

import (
	"strings"
	"sync"
)

// LocaleData содержит данные для форматирования в конкретной локали
type LocaleData struct {
	DecimalSeparator       string
	GroupSeparator         string
	PercentSymbol          string
	CurrencyFormats        map[string]*CurrencyData
	NegativePattern        string
	PositivePattern        string
	PercentPattern         string
	CompactPatterns        map[CompactRange]*CompactPattern
	MinusSign              string
	PlusSign               string
	Exponential            string
	SuperscriptingExponent bool
	NumberingSystem        string
}

// CurrencyData содержит данные о валюте
type CurrencyData struct {
	Symbol  string
	Name    string
	Format  string
	Spacing string
}

// CompactRange представляет диапазон для компактной записи
type CompactRange int

const (
	Thousand CompactRange = iota
	Million
	Billion
	Trillion
)

// CompactPattern содержит шаблоны для компактной записи
type CompactPattern struct {
	Short string
	Long  string
}

// localeCache кэширует загруженные локали
var localeCache = make(map[string]*LocaleData)
var cacheMutex sync.RWMutex

// GetLocaleData возвращает данные для локали
func GetLocaleData(locale string) *LocaleData {
	cacheMutex.RLock()
	if data, exists := localeCache[locale]; exists {
		cacheMutex.RUnlock()
		return data
	}
	cacheMutex.RUnlock()

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// Проверяем еще раз на случай параллельной загрузки
	if data, exists := localeCache[locale]; exists {
		return data
	}

	data := loadLocaleData(locale)
	localeCache[locale] = data
	return data
}

// loadLocaleData загружает данные локали из встроенных данных CLDR
func loadLocaleData(locale string) *LocaleData {
	// Нормализуем локаль
	locale = normalizeLocale(locale)

	// Пробуем загрузить точное соответствие
	if data := getExactLocaleData(locale); data != nil {
		return data
	}

	// Пробуем загрузить базовую локаль (без региона)
	baseLocale := strings.Split(locale, "-")[0]
	if data := getExactLocaleData(baseLocale); data != nil {
		return data
	}

	// Fallback на английскую локаль
	return getExactLocaleData("en")
}

// getExactLocaleData возвращает данные для конкретной локали
func getExactLocaleData(locale string) *LocaleData {
	if data, exists := localeData[locale]; exists {
		// Обогащаем данные валют
		for currencyCode := range data.CurrencyFormats {
			if extendedData := getCurrencyData(currencyCode); extendedData != nil {
				data.CurrencyFormats[currencyCode] = extendedData
			}
		}
		return data
	}
	return nil
}

// normalizeLocale нормализует строку локали
func normalizeLocale(locale string) string {
	return strings.ReplaceAll(strings.ToLower(locale), "_", "-")
}
