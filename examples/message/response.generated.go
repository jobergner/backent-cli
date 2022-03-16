package message

type AddItemToPlayerResponse struct {
	PlayerPath string `json:"playerPath"`
}

type SpawnZoneItemsResponse struct {
	NewZoneItemPaths []string `json:"newZoneItemPaths"`
}
