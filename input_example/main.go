package inputexample

type position struct {
	X float64
	Y float64
}

type gearScore struct {
	level int
	score int
}

type item struct {
	gearScore gearScore
}

type player struct {
	items     []item
	gearScore gearScore
	position  position
}

type zoneItem struct {
	position position
	item     item
}

type zone struct {
	players []player
	items   []zoneItem
}
