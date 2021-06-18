package ast

import "sort"

func newAction(name string) Action {
	return Action{
		Name:   name,
		Params: make(map[string]Field),
	}
}

type Action struct {
	Name     string
	Params   map[string]Field
	Response map[string]Field
}

func (a *Action) RangeParams(fn func(field Field)) {
	var keys []string
	for key := range a.Params {
		keys = append(keys, key)
	}
	sort.Slice(keys, caseInsensitiveSort(keys))
	for _, key := range keys {
		fn(a.Params[key])
	}
}

func (a *Action) RangeResponse(fn func(field Field)) {
	var keys []string
	for key := range a.Response {
		keys = append(keys, key)
	}
	sort.Slice(keys, caseInsensitiveSort(keys))
	for _, key := range keys {
		fn(a.Response[key])
	}
}
