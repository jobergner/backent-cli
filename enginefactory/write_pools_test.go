package enginefactory

import (
	"strings"
	"testing"

	"github.com/Java-Jonas/bar-cli/testutils"
)

func TestWritePools(t *testing.T) {
	t.Run("writes pools", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writePools()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
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
