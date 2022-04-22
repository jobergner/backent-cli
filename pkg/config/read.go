// config provides utility methods for reading and validating a config
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/jobergner/backent-cli/pkg/validator"
)

var (
	ErrInvalidConfig = errors.New("invalid config")
)

type jsonConfig struct {
	State     map[string]interface{} `json:"state"`
	Actions   map[string]interface{} `json:"actions"`
	Responses map[string]interface{} `json:"responses"`
}

// Read validates and parses config file at `path`
func Read(path string) (state, actions, responses map[interface{}]interface{}, err error) {
	config, err := readConfig(path)
	if err != nil {
		return nil, nil, nil, err
	}

	ok, errs := validate(*config)
	if !ok {
		for _, err := range errs {
			fmt.Println(err.Error())
		}
		return nil, nil, nil, ErrInvalidConfig
	}

	state, actions, responses = config.prepare()

	return state, actions, responses, nil
}

func validate(config jsonConfig) (bool, []error) {
	state, actions, responses := config.prepare()

	if errs := validator.ValidateStateConfig(state); len(errs) != 0 {
		return false, errs
	}
	if errs := validator.ValidateActionsConfig(state, actions); len(errs) != 0 {
		return false, errs
	}
	if errs := validator.ValidateResponsesConfig(state, actions, responses); len(errs) != 0 {
		return false, errs
	}

	return true, nil
}

func readConfig(path string) (*jsonConfig, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	jc := new(jsonConfig)
	err = json.Unmarshal(configFile, jc)
	if err != nil {
		return nil, err
	}

	return jc, nil
}

// prepare converts the type of each field because validator and ast want map[interface{}]interface{}
func (j jsonConfig) prepare() (state, actions, responses map[interface{}]interface{}) {
	return makeAmbiguous(j.State), makeAmbiguous(j.Actions), makeAmbiguous(j.Responses)
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
