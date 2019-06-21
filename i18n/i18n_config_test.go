package i18n

import "testing"

func TestLoadI18nConfigFromYamlFile(t *testing.T) {
	i18nConfig, err := LoadI18nConfigFromYamlFile("./i18n_config_test.yaml")
	if err != nil {
		t.Errorf("load err:%+v", err)
	} else {
		t.Logf("config:%+v", i18nConfig)
	}
}

func TestLoadI18nConfigFromTomlFile(t *testing.T) {
	i18nConfig, err := LoadI18nConfigFromTomlFile("./i18n_config_test.toml")
	if err != nil {
		t.Errorf("load err:%+v", err)
	} else {
		t.Logf("config:%+v", i18nConfig)
	}
}