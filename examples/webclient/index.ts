declare global {
  interface Window {
    movePlayer(changeX: number, changeY: number, player: number): void;
    addItemToPlayer(item: number, newName: string): string;
    spawnZoneItems(items: string[]): string;
  }
}

interface AddItemToPlayerResponse {
  playerPath: string;
}

interface SpawnZoneItemsResponse {
  newZoneItemPaths: string;
}

enum ReferencedDataStatus {
  ReferencedDataModified = "MODIFIED",
  ReferencedDataUnchanged = "UNCHANGED",
}

enum OperationKind {
  OperationKindDelete = "DELETE",
  OperationKindUpdate = "UPDATE",
  OperationKindUnchanged = "UNCHANGED",
}

enum ElementKind {
  ElementKindAttackEvent = "AttackEvent",
  ElementKindEquipmentSet = "EquipmentSet",
  ElementKindGearScore = "GearScore",
  ElementKindItem = "Item",
  ElementKindPlayer = "Player",
  ElementKindPosition = "Position",
  ElementKindZone = "Zone",
  ElementKindZoneItem = "ZoneItem",
}

interface ElementReference {
  OperationKind: OperationKind;
  ElementID: number;
  ElementKind: ElementKind;
  ReferencedDataStatus: ReferencedDataStatus;
  ElementPath: string;
}

interface ZoneItem {
  ID: number;
  Item?: Item;
  Position?: Position;
  OperationKind: OperationKind;
  ElementKind: ElementKind;
}

interface Item {
  ID: number;
  BoundTo?: ElementReference;
  GearScore?: GearScore;
  Name?: string;
  Origin?: Player | Position;
  OperationKind: OperationKind;
  ElementKind: ElementKind;
}

interface AttackEvent {
  ID: number;
  Target?: ElementReference;
  OperationKind: OperationKind;
  ElementKind: ElementKind;
}

interface EquipmentSet {
  ID: number;
  Equipment?: { [id: number]: ElementReference };
  Name?: string;
  OperationKind: OperationKind;
  ElementKind: ElementKind;
}

interface Position {
  ID: number;
  X?: number;
  Y?: number;
  OperationKind: OperationKind;
  ElementKind: ElementKind;
}

interface GearScore {
  ID: number;
  Level?: number;
  Score?: number;
  OperationKind: OperationKind;
  ElementKind: ElementKind;
}

interface Player {
  ID: number;
  Action?: { [id: number]: AttackEvent };
  EquipmentSets?: { [id: number]: ElementReference };
  GearScore?: GearScore;
  GuildMembers?: { [id: number]: ElementReference };
  Items?: { [id: number]: Item };
  Position?: Position;
  Target?: ElementReference;
  TargetedBy?: { [id: number]: ElementReference };
  OperationKind: OperationKind;
  ElementKind: ElementKind;
}

interface Zone {
  ID: number;
  Interactables?: { [id: number]: Item | Player | ZoneItem };
  Items?: { [id: number]: ZoneItem };
  Players?: { [id: number]: Player };
  Tags?: string[];
  OperationKind: OperationKind;
  ElementKind: ElementKind;
}

interface Tree {
  AttackEvent?: { [id: number]: AttackEvent };
  EquipmentSet?: { [id: number]: EquipmentSet };
  GearScore?: { [id: number]: GearScore };
  Item?: { [id: number]: Item };
  Player?: { [id: number]: Player };
  Position?: { [id: number]: Position };
  Zone?: { [id: number]: Zone };
  ZoneItem?: { [id: number]: ZoneItem };
}

const currentState: Tree = {};

function RECEIVEUPDATE(json: string) {
  importUpdate(JSON.parse(json) as Tree);
}

function importUpdate(update: Tree) {
  const current = currentState;
  if (update.EquipmentSet != null && update.EquipmentSet != undefined) {
    if (current.EquipmentSet == null || current.EquipmentSet == undefined) {
      current.EquipmentSet = {};
    }
    for (const id in update.EquipmentSet) {
      current.EquipmentSet[id] = importEquipmentSet(current.EquipmentSet[id], update.EquipmentSet[id]);
    }
  }
  if (update.GearScore != null && update.GearScore != undefined) {
    if (current.GearScore == null || current.GearScore == undefined) {
      current.GearScore = {};
    }
    for (const id in update.GearScore) {
      current.GearScore[id] = importGearScore(current.GearScore[id], update.GearScore[id]);
    }
  }
  if (update.Item != null && update.Item != undefined) {
    if (current.Item == null || current.Item == undefined) {
      current.Item = {};
    }
    for (const id in update.Item) {
      current.Item[id] = importItem(current.Item[id], update.Item[id]);
    }
  }
  if (update.Player != null && update.Player != undefined) {
    if (current.Player == null || current.Player == undefined) {
      current.Player = {};
    }
    for (const id in update.Player) {
      current.Player[id] = importPlayer(current.Player[id], update.Player[id]);
    }
  }
  if (update.Position != null && update.Position != undefined) {
    if (current.Position == null || current.Position == undefined) {
      current.Position = {};
    }
    for (const id in update.Position) {
      current.Position[id] = importPosition(current.Position[id], update.Position[id]);
    }
  }
  if (update.Zone != null && update.Zone != undefined) {
    if (current.Zone == null || current.Zone == undefined) {
      current.Zone = {};
    }
    for (const id in update.Zone) {
      current.Zone[id] = importZone(current.Zone[id], update.Zone[id]);
    }
  }
  if (update.ZoneItem != null && update.ZoneItem != undefined) {
    if (current.ZoneItem == null || current.ZoneItem == undefined) {
      current.ZoneItem = {};
    }
    for (const id in update.ZoneItem) {
      current.ZoneItem[id] = importZoneItem(current.ZoneItem[id], update.ZoneItem[id]);
    }
  }
}

function importEquipmentSet(current: EquipmentSet | null | undefined, update: EquipmentSet): EquipmentSet {
  if (current == null || current == undefined) {
    current = { ID: update.ID, ElementKind: update.ElementKind, OperationKind: update.OperationKind };
  }
  if (update.Equipment != null && update.Equipment != undefined) {
    if (current.Equipment == null || current.Equipment == undefined) {
      current.Equipment = {};
    }
    for (const id in update.Equipment) {
      current.Equipment[id] = importElementReference(current.Equipment[id], update.Equipment[id]);
    }
  }
  if (update.Name != null && update.Name != undefined) {
    current.Name = update.Name;
  }
  return current;
}

function importGearScore(current: GearScore | null | undefined, update: GearScore): GearScore {
  if (current == null || current == undefined) {
    current = { ID: update.ID, ElementKind: update.ElementKind, OperationKind: update.OperationKind };
  }
  if (update.Level != null && update.Level != undefined) {
    current.Level = update.Level;
  }
  if (update.Score != null && update.Score != undefined) {
    current.Score = update.Score;
  }
  return current;
}

function importItem(current: Item | null | undefined, update: Item): Item {
  if (current == null || current == undefined) {
    current = { ID: update.ID, ElementKind: update.ElementKind, OperationKind: update.OperationKind };
  }
  if (update.BoundTo != null && update.BoundTo != undefined) {
    current.BoundTo = importElementReference(current.BoundTo, update.BoundTo);
  }
  if (update.GearScore != null && update.GearScore != undefined) {
    current.GearScore = importGearScore(current.GearScore, update.GearScore);
  }
  if (update.Name != null && update.Name != undefined) {
    current.Name = update.Name;
  }
  if (update.Origin != null && update.Origin != undefined) {
    if (update.ElementKind == ElementKind.ElementKindPlayer) {
      current.Origin = importPlayer(current.Origin as Player, update.Origin);
    }
    if (update.ElementKind == ElementKind.ElementKindPosition) {
      current.Origin = importPosition(current.Origin as Position, update.Origin);
    }
  }
  return current;
}

function importPosition(current: Position | null | undefined, update: Position): Position {
  if (current == null || current == undefined) {
    current = { ID: update.ID, ElementKind: update.ElementKind, OperationKind: update.OperationKind };
  }
  if (update.X != null && update.X != undefined) {
    current.X = update.X;
  }
  if (update.Y != null && update.Y != undefined) {
    current.Y = update.Y;
  }
  return current;
}

function importPlayer(current: Player | null | undefined, update: Player): Player {
  if (current == null || current == undefined) {
    current = { ID: update.ID, ElementKind: update.ElementKind, OperationKind: update.OperationKind };
  }
  if (update.EquipmentSets != null && update.EquipmentSets != undefined) {
    if (current.EquipmentSets == null || current.EquipmentSets == undefined) {
      current.EquipmentSets = {};
    }
    for (const id in update.EquipmentSets) {
      current.EquipmentSets[id] = importElementReference(current.EquipmentSets[id], update.EquipmentSets[id]);
    }
  }
  if (update.GearScore != null && update.GearScore != undefined) {
    current.GearScore = importGearScore(current.GearScore, update.GearScore);
  }
  if (update.GuildMembers != null && update.GuildMembers != undefined) {
    if (current.GuildMembers == null || current.GuildMembers == undefined) {
      current.GuildMembers = {};
    }
    for (const id in update.GuildMembers) {
      current.GuildMembers[id] = importElementReference(current.GuildMembers[id], update.GuildMembers[id]);
    }
  }
  if (update.Items != null && update.Items != undefined) {
    if (current.Items == null || current.Items == undefined) {
      current.Items = {};
    }
    for (const id in update.Items) {
      current.Items[id] = importItem(current.Items[id], update.Items[id]);
    }
  }
  if (update.Position != null && update.Position != undefined) {
    current.Position = importGearScore(current.Position, update.Position);
  }
  if (update.Target != null && update.Target != undefined) {
    current.Target = importElementReference(current.Target, update.Target);
  }
  if (update.TargetedBy != null && update.TargetedBy != undefined) {
    if (current.TargetedBy == null || current.TargetedBy == undefined) {
      current.TargetedBy = {};
    }
    for (const id in update.TargetedBy) {
      current.TargetedBy[id] = importElementReference(current.TargetedBy[id], update.TargetedBy[id]);
    }
  }
  return current;
}

function importZone(current: Zone | null | undefined, update: Zone): Zone {
  if (current == null || current == undefined) {
    current = { ID: update.ID, ElementKind: update.ElementKind, OperationKind: update.OperationKind };
  }
  if (update.Interactables != null && update.Interactables != undefined) {
    if (current.Interactables == null || current.Interactables == undefined) {
      current.Interactables = {};
    }
    for (const id in update.Interactables) {
      if (update.Interactables[id].ElementKind == ElementKind.ElementKindItem) {
        current.Interactables[id] = importItem(current.Interactables[id] as Item, update.Interactables[id]);
      }
      if (update.Interactables[id].ElementKind == ElementKind.ElementKindZoneItem) {
        current.Interactables[id] = importPosition(current.Interactables[id] as ZoneItem, update.Interactables[id]);
      }
    }
  }
  if (update.Items != null && update.Items != undefined) {
    if (current.Items == null || current.Items == undefined) {
      current.Items = {};
    }
    for (const id in update.Items) {
      current.Items[id] = importZoneItem(current.Items[id], update.Items[id]);
    }
  }
  if (update.Players != null && update.Players != undefined) {
    if (current.Players == null || current.Players == undefined) {
      current.Players = {};
    }
    for (const id in update.Players) {
      current.Players[id] = importPlayer(current.Players[id], update.Players[id]);
    }
  }
  if (update.Tags != null && update.Tags != undefined) {
    current.Tags = update.Tags;
  }
  return current;
}

function importZoneItem(current: ZoneItem | null | undefined, update: ZoneItem): ZoneItem {
  if (current == null || current == undefined) {
    current = { ID: update.ID, ElementKind: update.ElementKind, OperationKind: update.OperationKind };
  }
  if (update.Item != null && update.Item != undefined) {
    current.Item = importGearScore(current.Item, update.Item);
  }
  if (update.Position != null && update.Position != undefined) {
    current.Position = importGearScore(current.Position, update.Position);
  }
  return current;
}

function importElementReference(current: ElementReference | null | undefined, update: ElementReference): ElementReference {
  current != null; // prevents unused param
  return update;
}

export function MovePlayer(changeX: number, changeY: number, player: number) {
  window.movePlayer(changeX, changeY, player);
}

export function AddItemToPlayer(item: number, newName: string): AddItemToPlayerResponse {
  const responseString = window.addItemToPlayer(item, newName);
  return JSON.parse(responseString);
}

export function SpawnZoneItems(items: string[]): SpawnZoneItemsResponse {
  const responseString = window.spawnZoneItems(items);
  return JSON.parse(responseString);
}

export {};
