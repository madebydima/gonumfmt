package gonumfmt

import (
	"os"
	"runtime"
	"strings"
)

// getSystemLocale определяет системную локаль
func getSystemLocale() string {
	// Пробуем разные методы в зависимости от ОС
	switch runtime.GOOS {
	case "windows":
		return getWindowsLocale()
	case "darwin":
		return getMacOSLocale()
	default:
		return getUnixLocale()
	}
}

// getUnixLocale получает локаль на Unix-системах
func getUnixLocale() string {
	// Пробуем переменные окружения в порядке приоритета
	envVars := []string{"LC_ALL", "LC_NUMERIC", "LANG", "LANGUAGE"}

	for _, envVar := range envVars {
		if locale := os.Getenv(envVar); locale != "" {
			return normalizeSystemLocale(locale)
		}
	}

	return "en" // fallback
}

// getWindowsLocale получает локаль на Windows
func getWindowsLocale() string {
	// На Windows используем GetUserDefaultLocaleName из win32 API
	// Временная реализация - пробуем переменные окружения
	if locale := os.Getenv("LC_ALL"); locale != "" {
		return normalizeSystemLocale(locale)
	}
	if locale := os.Getenv("LANG"); locale != "" {
		return normalizeSystemLocale(locale)
	}

	return "en" // fallback
}

// getMacOSLocale получает локаль на macOS
func getMacOSLocale() string {
	// На macOS также используем переменные окружения
	return getUnixLocale()
}

// normalizeSystemLocale нормализует системную локаль
func normalizeSystemLocale(locale string) string {
	// Примеры входных форматов:
	// en_US.UTF-8, ru_RU.UTF-8, en-US.UTF-8, es_ES, etc.

	// Убираем кодировку
	if dotIndex := strings.Index(locale, "."); dotIndex != -1 {
		locale = locale[:dotIndex]
	}

	// Заменяем подчеркивания на дефисы
	locale = strings.Replace(locale, "_", "-", -1)

	// Приводим к нижнему регистру для базовой части
	parts := strings.Split(locale, "-")
	if len(parts) > 0 {
		parts[0] = strings.ToLower(parts[0])
	}
	if len(parts) > 1 {
		parts[1] = strings.ToUpper(parts[1])
	}

	return strings.Join(parts, "-")
}
