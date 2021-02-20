package statemachine

type Parentage []ParentInfo
type ParentInfo struct {
	Kind EntityKind
	ID   int
}
