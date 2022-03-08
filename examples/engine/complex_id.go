package state

import (
	"bytes"
	"fmt"
	"strconv"
)

// ComplexID is used for entities which are not real elements
// like references and any-containers. This way we do not have
// to generate IDs for these entities, which would require
// the server to do it. This makes more methods available for broadcasting.
// easyjson:skip
type ComplexID struct {
	// despite the fact that a slice of references cannot hold multiple references of the same element,
	// we need the field identifier because otherwise an element with multiple reference fields
	// may be able to contain references with the same ParentID-ChildID combination.
	Field    treeFieldIdentifier `json:"field"`
	ParentID int                 `json:"parentID"`
	// ChildID describes the next true element => references of any-containers will not have the any-container ID as ChildID
	ChildID int `json:"childID"`
	// when a reference references an any-container both would have the same ParentID && ChildID
	// IsMediator simply acts as a differentiation between the two, so each ID is guaranteed to be unique
	IsMediator bool `json:"isMediator"`
}

var (
	complexIDStructCache = make(map[string]ComplexID)
	ComplexIDStringCache = make(map[ComplexID][]byte)
	complexIDZeroString  = []byte("0-0-0-0")
)

func (c ComplexID) MarshalJSON() ([]byte, error) {
	if cachedString, ok := ComplexIDStringCache[c]; ok {
		return cachedString, nil
	}

	var isMediatorBin int
	if c.IsMediator {
		isMediatorBin = 1
	}

	newS := []byte(fmt.Sprintf("\"%d-%d-%d-%d\"", c.Field, c.ParentID, c.ChildID, isMediatorBin))
	ComplexIDStringCache[c] = newS
	return newS, nil
}

func (c *ComplexID) UnmarshalJSON(s []byte) error {

	if bytes.Equal(s, complexIDZeroString) {
		return nil
	}

	asString := string(s)

	if cachedID, ok := complexIDStructCache[asString]; ok {
		c.Field = cachedID.Field
		c.ParentID = cachedID.ParentID
		c.ChildID = cachedID.ChildID
		c.IsMediator = cachedID.IsMediator
		return nil
	}

	idSegments := bytes.Split(s[1:len(s)-1], []byte{'-'})

	ident, _ := strconv.Atoi(string(idSegments[0]))
	c.Field = treeFieldIdentifier(ident)

	c.ParentID, _ = strconv.Atoi(string(idSegments[1]))

	c.ChildID, _ = strconv.Atoi(string(idSegments[2]))

	isMediatorBin, _ := strconv.Atoi(string(idSegments[3]))
	if isMediatorBin == 1 {
		c.IsMediator = true
	}

	complexIDStructCache[asString] = *c

	return nil
}

func (x AttackEventTargetRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *AttackEventTargetRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = AttackEventTargetRefID(temp)
	return nil
}

func (x PlayerGuildMemberRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *PlayerGuildMemberRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = PlayerGuildMemberRefID(temp)
	return nil
}

func (x ItemBoundToRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *ItemBoundToRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = ItemBoundToRefID(temp)
	return nil
}

func (x EquipmentSetEquipmentRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *EquipmentSetEquipmentRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = EquipmentSetEquipmentRefID(temp)
	return nil
}

func (x PlayerEquipmentSetRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *PlayerEquipmentSetRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = PlayerEquipmentSetRefID(temp)
	return nil
}

func (x AnyOfItem_Player_ZoneItemID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *AnyOfItem_Player_ZoneItemID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = AnyOfItem_Player_ZoneItemID(temp)
	return nil
}

func (x AnyOfPlayer_ZoneItemID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *AnyOfPlayer_ZoneItemID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = AnyOfPlayer_ZoneItemID(temp)
	return nil
}

func (x AnyOfPlayer_PositionID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *AnyOfPlayer_PositionID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = AnyOfPlayer_PositionID(temp)
	return nil
}

func (x PlayerTargetRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *PlayerTargetRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = PlayerTargetRefID(temp)
	return nil
}

func (x PlayerTargetedByRefID) MarshalJSON() ([]byte, error) {
	return ComplexID(x).MarshalJSON()
}
func (x *PlayerTargetedByRefID) UnmarshalJSON(s []byte) error {
	temp := ComplexID(*x)
	temp.UnmarshalJSON(s)
	*x = PlayerTargetedByRefID(temp)
	return nil
}
