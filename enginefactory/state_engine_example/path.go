package state

const (
	equipmentSetIdentifier int = -1
	gearScoreIdentifier    int = -2
	itemIdentifier         int = -3
	playerIdentifier       int = -4
	positionIdentifier     int = -5
	zoneIdentifier         int = -6
	zoneItemIdentifier     int = -7
)

type path []int

func newPath(elementIdentifier, id int) []int {
	return []int{elementIdentifier, id}
}

func (p path) equipmentSet() []int {
	tmp := make([]int, len(p)+1)
	copy(tmp, p)
	tmp = append(tmp, equipmentSetIdentifier)
	return tmp
}

func (p path) gearScore() []int {
	tmp := make([]int, len(p)+1)
	copy(tmp, p)
	tmp = append(tmp, gearScoreIdentifier)
	return tmp
}

func (p path) item() []int {
	tmp := make([]int, len(p)+1)
	copy(tmp, p)
	tmp = append(tmp, itemIdentifier)
	return tmp
}

func (p path) player() []int {
	tmp := make([]int, len(p)+1)
	copy(tmp, p)
	tmp = append(tmp, playerIdentifier)
	return tmp
}

func (p path) position() []int {
	tmp := make([]int, len(p)+1)
	copy(tmp, p)
	tmp = append(tmp, positionIdentifier)
	return tmp
}

func (p path) zone() []int {
	tmp := make([]int, len(p)+1)
	copy(tmp, p)
	tmp = append(tmp, zoneIdentifier)
	return tmp
}

func (p path) zoneItem() []int {
	tmp := make([]int, len(p)+1)
	copy(tmp, p)
	tmp = append(tmp, zoneItemIdentifier)
	return tmp
}

func (p path) index(i int) []int {
	tmp := make([]int, len(p)+1)
	copy(tmp, p)
	tmp = append(tmp, i)
	return tmp
}
