package httpx

import (
	"pineal/sensor/i18n"
	"testing"
)

func TestDefaultI18n(t *testing.T) {
	defaultConfig, err := i18n.LoadI18nConfigFromBuffer([]byte(DefaultI18nConfig))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(defaultConfig)
	}
}
