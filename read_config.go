package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/validator"
)

type jsonConfig struct {
	State     map[string]interface{} `json:"state"`
	Actions   map[string]interface{} `json:"actions"`
	Responses map[string]interface{} `json:"responses"`
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
		return fmt.Errorf("\"state\" field in json not found but is required")
	}
	return nil
}

func newAST() (*ast.AST, []error) {
	configFile, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	jc := jsonConfig{}
	err = json.Unmarshal(configFile, &jc)
	if err != nil {
		panic(err)
	}

	err = validateJSONConfig(jc)
	if err != nil {
		panic(err)
	}

	state := makeAmbiguous(jc.State)
	actions := makeAmbiguous(jc.Actions)
	responses := makeAmbiguous(jc.Responses)

	if errs := validator.ValidateStateConfig(state); len(errs) != 0 {
		return nil, errs
	}
	if errs := validator.ValidateActionsConfig(state, actions); len(errs) != 0 {
		return nil, errs
	}
	if errs := validator.ValidateResponsesConfig(state, actions, responses); len(errs) != 0 {
		return nil, errs
	}

	return ast.Parse(state, actions, responses), nil
}
