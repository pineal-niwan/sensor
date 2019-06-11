package i18n

import "sync"

//翻译
type langStringHash struct {
	stringHash map[string]string
}

//初始化
func (l *langStringHash) init(localHash map[string]string) {
	l.stringHash = localHash
}

//获取
func (l *langStringHash) get(key string) (val string, ok bool) {
	if len(l.stringHash) > 0 {
		val, ok = l.stringHash[key]
	} else {
		//没有找到，返回key
		val, ok = key, false
	}
	return val, ok
}

//翻译组
type LangStringGroup struct {
	langHash    map[string]*langStringHash
	defaultLang string
	sync.RWMutex
}

//初始化
func (l *LangStringGroup) Init(defaultLang string) {
	l.Lock()
	l.defaultLang = defaultLang
	l.langHash = make(map[string]*langStringHash)
	l.Unlock()
}

//设置缺省语言
func (l *LangStringGroup) SetDefaultLang(defaultLang string) {
	l.Lock()
	l.defaultLang = defaultLang
	l.Unlock()
}

//获取缺省语言
func (l *LangStringGroup) GetDefaultLang() string {
	l.RLock()
	defaultLang := l.defaultLang
	l.RUnlock()
	return defaultLang
}

//为一种语言添加新的翻译
func (l *LangStringGroup) AddLangString(lang string, localHash map[string]string) {
	var newLangHash map[string]string

	l.RLock()
	currentLang, ok := l.langHash[lang]
	l.RUnlock()

	if ok {
		// 已经有对应的语言包
		oldLangHash := currentLang.stringHash
		newLangHash = make(map[string]string)
		for k, v := range oldLangHash {
			newLangHash[k] = v
		}
		for k, v := range localHash {
			newLangHash[k] = v
		}
	} else {
		newLangHash = localHash
	}

	l.Lock()
	local := &langStringHash{}
	local.init(newLangHash)
	l.langHash[lang] = local
	l.Unlock()
}

//通过配置设置lang
func (l *LangStringGroup) ConfigBy(configHash I18nConfigHash) {
	for lang, pack := range configHash {
		l.AddLangString(lang, pack)
	}
}

//获取语言翻译
func (l *LangStringGroup) GetLocale(lang string, key string) (string, bool) {
	l.RLock()
	langHash, exist := l.langHash[lang]
	if !exist {
		// 没有对应的语言包
		defaultHash, exist := l.langHash[l.defaultLang]
		if !exist {
			// 没有缺省的语言包
			l.RUnlock()
			return key, false
		} else {
			// 有缺省的语言包
			l.RUnlock()
			return defaultHash.get(key)
		}
	} else {
		// 在指定的语言包中获取
		val, ok := langHash.get(key)
		if !ok {
			// 语言包中没有对应的key
			// 在缺省语言包中获取
			defaultHash, exist := l.langHash[l.defaultLang]
			if !exist {
				// 没有缺省的语言包
				l.RUnlock()
				return key, false
			} else {
				// 有缺省的语言包
				l.RUnlock()
				return defaultHash.get(key)
			}
		} else {
			l.RUnlock()
			return val, ok
		}
	}
}
