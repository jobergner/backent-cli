package state

type assemblePlanner struct {
	updatedPaths          map[int]path
	updatedReferencePaths map[int]path
	updatedElementPaths   map[int]path
	includedElements      map[int]bool
	equipmentSetPath      map[int]path
	gearScorePath         map[int]path
	itemPath              map[int]path
	playerPath            map[int]path
	positionPath          map[int]path
	zonePath              map[int]path
	zoneItemPath          map[int]path
}

func newAssemblePlanner() *assemblePlanner {
	return &assemblePlanner{
		updatedPaths:          make(map[int]path),
		updatedElementPaths:   make(map[int]path),
		updatedReferencePaths: make(map[int]path),
		includedElements:      make(map[int]bool),
		equipmentSetPath:      make(map[int]path),
		gearScorePath:         make(map[int]path),
		itemPath:              make(map[int]path),
		playerPath:            make(map[int]path),
		positionPath:          make(map[int]path),
		zonePath:              make(map[int]path),
		zoneItemPath:          make(map[int]path),
	}
}

func (a *assemblePlanner) clear() {
	for key := range a.updatedPaths {
		delete(a.updatedPaths, key)
	}
	for key := range a.updatedElementPaths {
		delete(a.updatedElementPaths, key)
	}
	for key := range a.updatedReferencePaths {
		delete(a.updatedReferencePaths, key)
	}
	for key := range a.includedElements {
		delete(a.includedElements, key)
	}
	for key := range a.equipmentSetPath {
		delete(a.equipmentSetPath, key)
	}
	for key := range a.gearScorePath {
		delete(a.gearScorePath, key)
	}
	for key := range a.itemPath {
		delete(a.itemPath, key)
	}
	for key := range a.playerPath {
		delete(a.playerPath, key)
	}
	for key := range a.positionPath {
		delete(a.positionPath, key)
	}
	for key := range a.zonePath {
		delete(a.zonePath, key)
	}
	for key := range a.zoneItemPath {
		delete(a.zoneItemPath, key)
	}
}
