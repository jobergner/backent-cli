package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

type config struct {
	State   map[interface{}]interface{}
	Actions map[interface{}]interface{}
}
type jsonConfig struct {
	State   map[string]interface{} `json:"state"`
	Actions map[string]interface{} `json:"actions"`
}

func makeAmbiguous(a map[string]interface{}) map[interface{}]interface{} {
	b := make(map[interface{}]interface{})
	for k, v := range a {
		if s, ok := isString(v); ok {
			b[k] = s
			continue
		}
		if m, ok := isMap(v); ok {
			tmp := make(map[interface{}]interface{})
			for k_, v_ := range m {
				tmp[k_] = v_
			}
			b[k] = tmp
			continue
		}
		if v == nil {
			b[k] = nil
			continue
		}
	}
	return b
}

func isString(unknown interface{}) (string, bool) {
	v := reflect.ValueOf(unknown)
	if v.Kind() == reflect.String {
		valueString := fmt.Sprintf("%v", unknown)
		return valueString, true
	}
	return "", false
}

func isMap(unknown interface{}) (map[string]interface{}, bool) {
	v := reflect.ValueOf(unknown)
	if v.Kind() == reflect.Map {
		if mapValue, ok := unknown.(map[string]interface{}); ok {
			return mapValue, true
		} else {
			return nil, false
		}
	}
	return nil, false
}

func validateJSONConfig(jc jsonConfig) error {
	if len(jc.State) == 0 {
		return fmt.Errorf("no configuration found for \"state\"")
	}
	return nil
}

func readConfig() (*config, error) {
	configFile, err := ioutil.ReadFile(*configNameFlag)
	if err != nil {
		panic(err)
	}
	jc := jsonConfig{}
	err = json.Unmarshal(configFile, &jc)
	if err != nil {
		return nil, err
	}
	err = validateJSONConfig(jc)
	if err != nil {
		return nil, err
	}
	c := &config{
		State:   makeAmbiguous(jc.State),
		Actions: makeAmbiguous(jc.Actions),
	}
	return c, nil
}
