package statemachine

type parentage []parentInfo
type parentInfo struct {
	kind entityKind
	id   int
}
