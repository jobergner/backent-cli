package typescript

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {

	prepare := func(s string) string {
		// ignore whitespace and semicolon, may result in logical differences in syntax
		// in some cases but is enough for this test
		s = strings.ReplaceAll(s, "	", "") // tab
		s = strings.ReplaceAll(s, " ", "") // space
		s = strings.ReplaceAll(s, ";", "")

		return s
	}

	t.Run("writes function", func(t *testing.T) {

		actual := Function("importEquipmentSet").Param(Param{"current", Id("EquipmentSet").OrType("null").OrType("undefined")}, Param{"update", Id("EquipmentSet")}).ReturnType("EquipmentSet").Block(
			If(Id("current").Equals().Id("null").Or().Id("current").Equals().Id("undefined")).Block(
				Id("current").Id("=").Object(ObjectField{"id", Id("update").Dot("id")}, ObjectField{"elementKind", Id("update").Dot("elementKind")}, ObjectField{"operationKind", Id("update").Dot("operationKind")}),
			),
			If(Id("update").Dot("equipment").EqualsNot().Id("null").And().Id("update").Dot("equipment").EqualsNot().Id("undefined")).Block(
				If(Id("current").Dot("equipment").Equals().Id("null").Or().Id("current").Dot("equipment").Equals().Id("undefined")).Block(
					Id("current").Dot("equipment").Id("=").Object(),
				),
				ForIn(Id("const").Id("id"), Id("update").Dot("equipment")).Block(
					Id("current").Dot("equipment").Index("id").Id("=").Id("importElementReference").Call(Id("current").Dot("equipment").Index("id"), Id("update").Dot("equipment").Index("id")),
				),
			),
			If(Id("update").Dot("name").EqualsNot().Id("null").And().Id("update").Dot("name").EqualsNot().Id("undefined")).Block(
				Id("current").Dot("name").Id("=").Id("update").Dot("name"),
			),
			Return("current"),
		)

		expected := `function importEquipmentSet(current: EquipmentSet | null | undefined, update: EquipmentSet): EquipmentSet {
  if (current === null || current === undefined) {
    current = { id: update.id, elementKind: update.elementKind, operationKind: update.operationKind };
  }
  if (update.equipment !== null && update.equipment !== undefined) {
    if (current.equipment === null || current.equipment === undefined) {
      current.equipment = {};
    }
    for (const id in update.equipment) {
      current.equipment[id] = importElementReference(current.equipment[id], update.equipment[id]);
    }
  }
  if (update.name !== null && update.name !== undefined) {
    current.name = update.name;
  }
  return current;
}`

		assert.Equal(t, prepare(expected), prepare(actual.toString()))
	})
}
