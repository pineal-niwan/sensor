package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

//i18n配置
type I18nConfigHash map[string]map[string]string

//load i18n配置
func LoadI18nConfigFromYamlFile(filePath string) (I18nConfigHash, error) {
	var i18nConfig I18nConfigHash

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buf, &i18nConfig)
	return i18nConfig, err
}

//load i18n配置
func LoadI18nConfigFromYamlBuffer(buf []byte) (I18nConfigHash, error) {
	var i18nConfig I18nConfigHash
	err := yaml.Unmarshal(buf, &i18nConfig)
	return i18nConfig, err
}

//load i18n配置
func LoadI18nConfigFromTomlFile(filePath string) (I18nConfigHash, error) {
	var i18nConfig I18nConfigHash

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(buf, &i18nConfig)
	return i18nConfig, err
}

//load i18n配置
func LoadI18nConfigFromTomlBuffer(buf []byte) (I18nConfigHash, error) {
	var i18nConfig I18nConfigHash
	err := toml.Unmarshal(buf, &i18nConfig)
	return i18nConfig, err
}
