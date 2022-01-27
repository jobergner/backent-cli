package enginefactory

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWritePools(t *testing.T) {
	t.Run("writes pools", func(t *testing.T) {
		sf := newStateFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writePools()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			equipmentSetCheckPool_type,
			equipmentSetIDSlicePool_type,
			gearScoreCheckPool_type,
			gearScoreIDSlicePool_type,
			itemCheckPool_type,
			itemIDSlicePool_type,
			playerCheckPool_type,
			playerIDSlicePool_type,
			positionCheckPool_type,
			positionIDSlicePool_type,
			zoneCheckPool_type,
			zoneIDSlicePool_type,
			zoneItemCheckPool_type,
			zoneItemIDSlicePool_type,
			equipmentSetEquipmentRefCheckPool_type,
			equipmentSetEquipmentRefIDSlicePool_type,
			itemBoundToRefCheckPool_type,
			itemBoundToRefIDSlicePool_type,
			playerEquipmentSetRefCheckPool_type,
			playerEquipmentSetRefIDSlicePool_type,
			playerGuildMemberRefCheckPool_type,
			playerGuildMemberRefIDSlicePool_type,
			playerTargetRefCheckPool_type,
			playerTargetRefIDSlicePool_type,
			playerTargetedByRefCheckPool_type,
			playerTargetedByRefIDSlicePool_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
