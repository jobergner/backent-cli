package typescript

import (
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestCode(t *testing.T) {

	t.Run("writes function", func(t *testing.T) {

		actual := Function("importEquipmentSet").Param(Param{"current", Id("EquipmentSet").OrType("null").OrType("undefined")}, Param{"update", Id("EquipmentSet")}).ReturnType("EquipmentSet").Block(
			If(Id("current").Equals().Id("null").Or().Id("current").Equals().Id("undefined")).Block(
				Id("current").Assign().Object(ObjectField{Id("id"), Id("update").Dot("id")}, ObjectField{Id("elementKind"), Id("update").Dot("elementKind")}, ObjectField{Id("operationKind"), Id("update").Dot("operationKind")}).Sc(),
			),
			If(Id("update").Dot("equipment").EqualsNot().Id("null").And().Id("update").Dot("equipment").EqualsNot().Id("undefined")).Block(
				If(Id("current").Dot("equipment").Equals().Id("null").Or().Id("current").Dot("equipment").Equals().Id("undefined")).Block(
					Id("current").Dot("equipment").Assign().Object().Sc(),
				),
				ForIn(Id("const").Id("id"), Id("update").Dot("equipment")).Block(
					Id("current").Dot("equipment").Index(Id("id")).Assign().Id("importElementReference").Call(Id("current").Dot("equipment").Index(Id("id")), Id("update").Dot("equipment").Index(Id("id"))).Sc(),
				),
			),
			If(Id("update").Dot("name").EqualsNot().Id("null").And().Id("update").Dot("name").EqualsNot().Id("undefined")).Block(
				Id("current").Dot("name").Assign().Id("update").Dot("name").Sc(),
			),
			Return("current").Sc(),
		).String()

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

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
