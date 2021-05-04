package state

const (
	itemsIdentifier         int = -1
	gearScoreIdentifier     int = -2
	positionIdentifier      int = -3
	targetIdentifier        int = -4
	playersIdentifier       int = -5
	interactablesIdentifier int = -6
	itemIdentifier          int = -7
	originIdentifier        int = -8
)

type path []int

func newPath(elementIdentifier, id int) path {
	return []int{elementIdentifier, id}
}

func newEmptyPath() path {
	var p path
	return p
}

func (p path) items() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, itemsIdentifier)
	return newPath
}

func (p path) gearScore() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, gearScoreIdentifier)
	return newPath
}

func (p path) position() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, positionIdentifier)
	return newPath
}

func (p path) target() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, targetIdentifier)
	return newPath
}

func (p path) players() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, playersIdentifier)
	return newPath
}

func (p path) interactables() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, interactablesIdentifier)
	return newPath
}

func (p path) item() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, itemIdentifier)
	return newPath
}

func (p path) origin() path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, originIdentifier)
	return newPath
}

func (p path) index(i int) path {
	newPath := make([]int, len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, i)
	return newPath
}

func (p path) equals(parentPath path) bool {
	if len(p) != len(parentPath) {
		return false
	}

	for i, segment := range parentPath {
		if segment != p[i] {
			return false
		}
	}

	return true
}
