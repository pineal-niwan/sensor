package i18n

import "testing"

func TestLoadI18nConfigFromFile(t *testing.T) {
	i18nConfig, err := LoadI18nConfigFromFile("./i18n_config_test.yaml")
	if err != nil {
		t.Errorf("load err:%+v", err)
	} else {
		t.Logf("config:%+v", i18nConfig)
	}
}
