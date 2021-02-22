package statemachine

type Parentage []ParentInfo
type ParentInfo struct {
	Kind EntityKind `json:"kind"`
	ID   int        `json:"id"`
}
