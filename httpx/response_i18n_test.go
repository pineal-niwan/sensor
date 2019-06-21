package httpx

import (
	"github.com/pineal-niwan/sensor/i18n"
	"testing"
)

func TestDefaultYamlI18n(t *testing.T) {
	defaultConfig, err := i18n.LoadI18nConfigFromYamlBuffer([]byte(DefaultI18nYamlConfig))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(defaultConfig)
	}
}

func TestDefaultTomlI18n(t *testing.T) {
	defaultConfig, err := i18n.LoadI18nConfigFromTomlBuffer([]byte(DefaultI18nTomlConfig))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(defaultConfig)
	}
}