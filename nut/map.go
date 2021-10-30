package nut

import (
	"encoding/json"
	"fmt"
)

type Map map[string]interface{}

func (m Map) HasKeyValue(key string, value interface{}) bool {
	return m[key] == value
}

func (m Map) HasKey(key string) bool {
	return m[key] != nil
}

func (m Map) GetString(key string) (string, error) {
	if m[key] == nil {
		return "", fmt.Errorf("string '%s' not found", key)
	}
	return m[key].(string), nil
}

func (m Map) GetInt64(key string) (int64, error) {
	if m[key] == nil {
		return 0, fmt.Errorf("int64 '%s' not found", key)
	}
	return int64(m[key].(float64)), nil
}

func (m Map) GetArrayString(key string) ([]string, error) {
	if m[key] == nil {
		return nil, fmt.Errorf("array string '%s' not found", key)
	}
	return m[key].([]string), nil
}

func (m Map) GetArrayInt64(key string) ([]int64, error) {
	if m[key] == nil {
		return nil, fmt.Errorf("array interger '%s' not found", key)
	}
	return m[key].([]int64), nil
}

func (m Map) ToString() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}