package utils

import "gopkg.in/ini.v1"

type Ini struct {
	f *ini.File
}

func LoadIniConf(path string) (*Ini, error) {
	f, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	return &Ini{f: f}, nil
}

func (i *Ini) GetString(section, key, _default string) string {
	k, err := i.getValue(section, key)
	if err != nil {
		return _default
	}
	v := k.String()
	if v == "" {
		return _default
	}
	return v
}

func (i *Ini) GetInt(section, key string, _default int) int {
	k, err := i.getValue(section, key)
	if err != nil {
		return _default
	}
	v, err := k.Int()
	if err != nil {
		return _default
	}
	return v
}

func (i *Ini) getValue(section, key string) (*ini.Key, error) {
	s, err := i.f.GetSection(section)
	if err != nil {
		return nil, err
	}
	return s.Key(key), nil
}
