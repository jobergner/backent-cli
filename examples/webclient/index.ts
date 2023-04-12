import {EventEmitter} from "events";

const ErrResponseTimeout = "ErrResponseTimeout"
const responseTimeout = 1000

export const eventEmitter = new EventEmitter();

export interface AddItemToPlayerResponse {
  playerPath: string;
}

export interface SpawnZoneItemsResponse {
  newZoneItemPaths: string;
}

export enum ReferencedDataStatus {
  ReferencedDataModified = "MODIFIED",
  ReferencedDataUnchanged = "UNCHANGED",
}

export enum OperationKind {
  OperationKindDelete = "DELETE",
  OperationKindUpdate = "UPDATE",
  OperationKindUnchanged = "UNCHANGED",
}

export enum ElementKind {
  ElementKind_Root = "root",
  ElementKindAttackEvent = "AttackEvent",
  ElementKindEquipmentSet = "EquipmentSet",
  ElementKindGearScore = "GearScore",
  ElementKindItem = "Item",
  ElementKindPlayer = "Player",
  ElementKindPosition = "position",
  ElementKindZone = "Zone",
  ElementKindZoneItem = "ZoneItem",
}

interface ElementReference {
  operationKind: OperationKind;
  elementID: number;
  elementKind: ElementKind;
  referencedDataStatus: ReferencedDataStatus;
  elementPath: string;
}

interface ZoneItem {
  id: number;
  item?: Item;
  position?: Position;
  operationKind: OperationKind;
  elementKind: ElementKind;
}

interface Item {
  id: number;
  boundTo?: ElementReference;
  gearScore?: GearScore;
  name?: string;
  origin?: Player | Position;
  operationKind: OperationKind;
  elementKind: ElementKind;
}

interface AttackEvent {
  id: number;
  target?: ElementReference;
  operationKind: OperationKind;
  elementKind: ElementKind;
}

interface EquipmentSet {
  id: number;
  equipment?: {[id: number]: ElementReference};
  name?: string;
  operationKind: OperationKind;
  elementKind: ElementKind;
}

interface Position {
  id: number;
  x?: number;
  y?: number;
  operationKind: OperationKind;
  elementKind: ElementKind;
}

interface GearScore {
  id: number;
  level?: number;
  score?: number;
  operationKind: OperationKind;
  elementKind: ElementKind;
}

interface Player {
  id: number;
  action?: {[id: number]: AttackEvent};
  equipmentSets?: {[id: number]: ElementReference};
  gearScore?: GearScore;
  guildMembers?: {[id: number]: ElementReference};
  items?: {[id: number]: Item};
  position?: Position;
  target?: ElementReference;
  targetedBy?: {[id: number]: ElementReference};
  operationKind: OperationKind;
  elementKind: ElementKind;
}

interface Zone {
  id: number;
  interactables?: {[id: number]: Item | Player | ZoneItem};
  items?: {[id: number]: ZoneItem};
  players?: {[id: number]: Player};
  tags?: string[];
  operationKind: OperationKind;
  elementKind: ElementKind;
}

export interface Tree {
  attackEvent?: {[id: number]: AttackEvent};
  equipmentSet?: {[id: number]: EquipmentSet};
  gearScore?: {[id: number]: GearScore};
  item?: {[id: number]: Item};
  player?: {[id: number]: Player};
  position?: {[id: number]: Position};
  zone?: {[id: number]: Zone};
  zoneItem?: {[id: number]: ZoneItem};
}

export const currentState: Tree = {};

function RECEIVEUPDATE(json: string) {
  const update = JSON.parse(json) as Tree;
  emit_Update(update);
  import_Update(currentState, update);
}

export function import_Update(current: Tree, update: Tree) {
  if (update.equipmentSet !== null && update.equipmentSet !== undefined) {
    if (current.equipmentSet === null || current.equipmentSet === undefined) {
      current.equipmentSet = {};
    }
    for (const id in update.equipmentSet) {
      if (update.equipmentSet[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.equipmentSet[id];
      } else {
        current.equipmentSet[id] = importEquipmentSet(current.equipmentSet[id], update.equipmentSet[id]);
      }
    }
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    if (current.gearScore === null || current.gearScore === undefined) {
      current.gearScore = {};
    }
    for (const id in update.gearScore) {
      if (update.gearScore[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.gearScore[id];
      } else {
        current.gearScore[id] = importGearScore(current.gearScore[id], update.gearScore[id]);
      }
    }
  }
  if (update.item !== null && update.item !== undefined) {
    if (current.item === null || current.item === undefined) {
      current.item = {};
    }
    for (const id in update.item) {
      if (update.item[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.item[id];
      } else {
        current.item[id] = importItem(current.item[id], update.item[id]);
      }
    }
  }
  if (update.player !== null && update.player !== undefined) {
    if (current.player === null || current.player === undefined) {
      current.player = {};
    }
    for (const id in update.player) {
      if (update.player[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.player[id];
      } else {
        current.player[id] = importPlayer(current.player[id], update.player[id]);
      }
    }
  }
  if (update.position !== null && update.position !== undefined) {
    if (current.position === null || current.position === undefined) {
      current.position = {};
    }
    for (const id in update.position) {
      if (update.position[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.position[id];
      } else {
        current.position[id] = importPosition(current.position[id], update.position[id]);
      }
    }
  }
  if (update.zone !== null && update.zone !== undefined) {
    if (current.zone === null || current.zone === undefined) {
      current.zone = {};
    }
    for (const id in update.zone) {
      if (update.zone[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.zone[id];
      } else {
        current.zone[id] = importZone(current.zone[id], update.zone[id]);
      }
    }
  }
  if (update.zoneItem !== null && update.zoneItem !== undefined) {
    if (current.zoneItem === null || current.zoneItem === undefined) {
      current.zoneItem = {};
    }
    for (const id in update.zoneItem) {
      if (update.zoneItem[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.zoneItem[id];
      } else {
        current.zoneItem[id] = importZoneItem(current.zoneItem[id], update.zoneItem[id]);
      }
    }
  }
}

function importEquipmentSet(current: EquipmentSet | null | undefined, update: EquipmentSet): EquipmentSet {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.equipment !== null && update.equipment !== undefined) {
    if (current.equipment === null || current.equipment === undefined) {
      current.equipment = {};
    }
    for (const id in update.equipment) {
      if (update.equipment[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.equipment[id];
      } else {
        current.equipment[id] = importElementReference(current.equipment[id], update.equipment[id]);
      }
    }
  }
  if (update.name !== null && update.name !== undefined) {
    current.name = update.name;
  }
  return current;
}

function importGearScore(current: GearScore | null | undefined, update: GearScore): GearScore {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.level !== null && update.level !== undefined) {
    current.level = update.level;
  }
  if (update.score !== null && update.score !== undefined) {
    current.score = update.score;
  }
  return current;
}

function importItem(current: Item | null | undefined, update: Item): Item {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.boundTo !== null && update.boundTo !== undefined) {
    if (update.boundTo.operationKind === OperationKind.OperationKindDelete) {
      delete current.boundTo;
    } else {
      current.boundTo = importElementReference(current.boundTo, update.boundTo);
    }
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    current.gearScore = importGearScore(current.gearScore, update.gearScore);
  }
  if (update.name !== null && update.name !== undefined) {
    current.name = update.name;
  }
  if (update.origin !== null && update.origin !== undefined) {
    if (update.elementKind == ElementKind.ElementKindPlayer) {
      current.origin = importPlayer(current.origin as Player, update.origin);
    }
    if (update.elementKind == ElementKind.ElementKindPosition) {
      current.origin = importPosition(current.origin as Position, update.origin);
    }
  }
  return current;
}

function importPosition(current: Position | null | undefined, update: Position): Position {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.x !== null && update.x !== undefined) {
    current.x = update.x;
  }
  if (update.y !== null && update.y !== undefined) {
    current.y = update.y;
  }
  return current;
}

function importPlayer(current: Player | null | undefined, update: Player): Player {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.equipmentSets !== null && update.equipmentSets !== undefined) {
    if (current.equipmentSets === null || current.equipmentSets === undefined) {
      current.equipmentSets = {};
    }
    for (const id in update.equipmentSets) {
      if (update.equipmentSets[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.equipmentSets[id];
      } else {
        current.equipmentSets[id] = importElementReference(current.equipmentSets[id], update.equipmentSets[id]);
      }
    }
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    current.gearScore = importGearScore(current.gearScore, update.gearScore);
  }
  if (update.guildMembers !== null && update.guildMembers !== undefined) {
    if (current.guildMembers === null || current.guildMembers === undefined) {
      current.guildMembers = {};
    }
    for (const id in update.guildMembers) {
      if (update.guildMembers[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.guildMembers[id];
      } else {
        current.guildMembers[id] = importElementReference(current.guildMembers[id], update.guildMembers[id]);
      }
    }
  }
  if (update.items !== null && update.items !== undefined) {
    if (current.items === null || current.items === undefined) {
      current.items = {};
    }
    for (const id in update.items) {
      if (update.items[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.items[id];
      } else {
        current.items[id] = importItem(current.items[id], update.items[id]);
      }
    }
  }
  if (update.position !== null && update.position !== undefined) {
    current.position = importPosition(current.position, update.position);
  }
  if (update.target !== null && update.target !== undefined) {
    if (update.target.operationKind === OperationKind.OperationKindDelete) {
      delete current.target;
    } else {
      current.target = importElementReference(current.target, update.target);
    }
  }
  if (update.targetedBy !== null && update.targetedBy !== undefined) {
    if (current.targetedBy === null || current.targetedBy === undefined) {
      current.targetedBy = {};
    }
    for (const id in update.targetedBy) {
      if (update.targetedBy[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.targetedBy[id];
      } else {
        current.targetedBy[id] = importElementReference(current.targetedBy[id], update.targetedBy[id]);
      }
    }
  }
  return current;
}

function importZone(current: Zone | null | undefined, update: Zone): Zone {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.interactables !== null && update.interactables !== undefined) {
    if (current.interactables === null || current.interactables === undefined) {
      current.interactables = {};
    }
    for (const id in update.interactables) {
      if (update.interactables[id].elementKind == ElementKind.ElementKindItem) {
        if (update.interactables[id].operationKind === OperationKind.OperationKindDelete) {
          delete current.interactables[id];
        } else {
          current.interactables[id] = importItem(current.interactables[id] as Item, update.interactables[id]);
        }
      }
      if (update.interactables[id].elementKind == ElementKind.ElementKindZoneItem) {
        if (update.interactables[id].operationKind === OperationKind.OperationKindDelete) {
          delete current.interactables[id];
        } else {
          current.interactables[id] = importPosition(current.interactables[id] as ZoneItem, update.interactables[id]);
        }
      }
    }
  }
  if (update.items !== null && update.items !== undefined) {
    if (current.items === null || current.items === undefined) {
      current.items = {};
    }
    for (const id in update.items) {
      if (update.items[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.items[id];
      } else {
        current.items[id] = importZoneItem(current.items[id], update.items[id]);
      }
    }
  }
  if (update.players !== null && update.players !== undefined) {
    if (current.players === null || current.players === undefined) {
      current.players = {};
    }
    for (const id in update.players) {
      if (update.players[id].operationKind === OperationKind.OperationKindDelete) {
        delete current.players[id];
      } else {
        current.players[id] = importPlayer(current.players[id], update.players[id]);
      }
    }
  }
  if (update.tags !== null && update.tags !== undefined) {
    current.tags = update.tags;
  }
  return current;
}

function importZoneItem(current: ZoneItem | null | undefined, update: ZoneItem): ZoneItem {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.item !== null && update.item !== undefined) {
    current.item = importGearScore(current.item, update.item);
  }
  if (update.position !== null && update.position !== undefined) {
    current.position = importGearScore(current.position, update.position);
  }
  return current;
}

function importElementReference(current: ElementReference | null | undefined, update: ElementReference): ElementReference {
  current !== null; // prevents unused param
  return update;
}


enum MessageKind {
  ID = "id",
  Error = "error",
  Update = "update",
  CurrentState = "currentState",
  ActionAddItemToPlayer = "addItemToPlayer",
  ActionMovePlayer = "movePlayer",
  ActionSpawnZoneItems = "spawnZoneItems",
}

interface WebSocketMessage {
  id: string;
  kind: string;
  content: string;
}

export class Client {
  private ws: WebSocket;
  private responseEmitter: EventEmitter;
  private id: string;

  constructor(url: string) {
    this.id = ""

    this.ws = new WebSocket(url);

    this.responseEmitter = new EventEmitter()

    this.ws.addEventListener('open', () => {
      console.log('WebSocket connection opened');
    });

    this.ws.addEventListener('message', (event) => {
      console.log('WebSocket message received:', event.data);

      const message = JSON.parse(event.data) as WebSocketMessage;

      switch (message.kind) {
        case MessageKind.ID:
          this.id = message.content
          break;
        case MessageKind.Update:
        case MessageKind.CurrentState:
          RECEIVEUPDATE(message.content)
          break;
        case MessageKind.Error:
          console.log(message)
          break;
        default:
          this.responseEmitter.emit(message.id, JSON.parse(message.content))
          break;
      }
    });

    this.ws.addEventListener('close', () => {
      console.log('WebSocket connection closed');
    });
  }

  public getID() {
    return this.id;
  }

  public movePlayer(changeX: number, changeY: number, player: string) {
    const messageID = this.generateID()

    const message: WebSocketMessage = {
      id: messageID,
      kind: MessageKind.ActionMovePlayer,
      content: JSON.stringify({changeX, changeY, player})
    };

    this.ws.send(JSON.stringify(message));
  }

  public addItemToPlayer(item: string, newName: string) {
    const messageID = this.generateID()

    const message: WebSocketMessage = {
      id: messageID,
      kind: MessageKind.ActionAddItemToPlayer,
      content: JSON.stringify({item, newName})
    };

    setTimeout(() => {
      this.ws.send(JSON.stringify(message));
    }, 0)

    return new Promise((resolve, reject) => {
      this.responseEmitter.on(messageID, (response: AddItemToPlayerResponse) => {
        resolve(response)
      })
      setTimeout(() => {
        reject(ErrResponseTimeout)
      }, responseTimeout)
    })
  }

  public spawnZoneItems(items: string[]) {
    const messageID = this.generateID()

    const message: WebSocketMessage = {
      id: messageID,
      kind: MessageKind.ActionSpawnZoneItems,
      content: JSON.stringify({items})
    };

    setTimeout(() => {
      this.ws.send(JSON.stringify(message));
    }, 0)

    return new Promise((resolve, reject) => {
      this.responseEmitter.on(messageID, (response: SpawnZoneItemsResponse) => {
        resolve(response)
      })
      setTimeout(() => {
        reject(ErrResponseTimeout)
      }, responseTimeout)
    })
  }

  private generateID(): string {
    return Date.now().toString() + Math.random().toString(36).slice(2, 5);
  }
}

export function emit_Update(update: Tree) {
  if (update.attackEvent !== null && update.attackEvent !== undefined) {
    for (const id in update.attackEvent) {
      emitEquipmentSet(update.attackEvent[id], ElementKind.ElementKind_Root, ElementKind.ElementKind_Root);
    }
  }
  if (update.equipmentSet !== null && update.equipmentSet !== undefined) {
    for (const id in update.equipmentSet) {
      emitEquipmentSet(update.equipmentSet[id], ElementKind.ElementKind_Root, ElementKind.ElementKind_Root);
    }
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    for (const id in update.gearScore) {
      emitGearScore(update.gearScore[id], ElementKind.ElementKind_Root, ElementKind.ElementKind_Root);
    }
  }
  if (update.item !== null && update.item !== undefined) {
    for (const id in update.item) {
      emitItem(update.item[id], ElementKind.ElementKind_Root, ElementKind.ElementKind_Root);
    }
  }
  if (update.player !== null && update.player !== undefined) {
    for (const id in update.player) {
      emitPlayer(update.player[id], ElementKind.ElementKind_Root, ElementKind.ElementKind_Root);
    }
  }
  if (update.position !== null && update.position !== undefined) {
    for (const id in update.position) {
      emitPosition(update.position[id], ElementKind.ElementKind_Root, ElementKind.ElementKind_Root);
    }
  }
  if (update.zone !== null && update.zone !== undefined) {
    for (const id in update.zone) {
      emitZone(update.zone[id], ElementKind.ElementKind_Root, ElementKind.ElementKind_Root);
    }
  }
  if (update.zoneItem !== null && update.zoneItem !== undefined) {
    for (const id in update.zoneItem) {
      emitZoneItem(update.zoneItem[id], ElementKind.ElementKind_Root, ElementKind.ElementKind_Root);
    }
  }
}

function emitEquipmentSet(update: EquipmentSet, parent: ElementKind, fieldName: string) {
  if (update.equipment !== null && update.equipment !== undefined) {
    for (const id in update.equipment) {
      emitElementReference(update.equipment[id], ElementKind.ElementKindPlayer, "equipment");
    }
  }
  if (parent === ElementKind.ElementKind_Root) {
    eventEmitter.emit("equipmentSet", update);
  }
}

function emitGearScore(update: GearScore, parent: ElementKind, fieldName: string) {
  if (parent === ElementKind.ElementKind_Root) {
    eventEmitter.emit("gearScore", update);
  }
  if (parent === ElementKind.ElementKindPlayer) {
    if (fieldName === "gearScore") {
      eventEmitter.emit("player_gearScore", update);
    }
  }
  if (parent === ElementKind.ElementKindItem) {
    if (fieldName === "gearScore") {
      eventEmitter.emit("item_gearScore", update);
    }
  }
}

function emitItem(update: Item, parent: ElementKind, fieldName: string) {
  if (update.boundTo !== null && update.boundTo !== undefined) {
    emitElementReference(update.boundTo, ElementKind.ElementKindPlayer, "boundTo");
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    emitGearScore(update.gearScore, ElementKind.ElementKindItem, "gearScore");
  }
  if (update.origin !== null && update.origin !== undefined) {
    if (update.elementKind === ElementKind.ElementKindPlayer) {
      emitPlayer(update.origin, ElementKind.ElementKindItem, "origin");
    }
    if (update.elementKind === ElementKind.ElementKindPosition) {
      emitPosition(update.origin, ElementKind.ElementKindItem, "origin");
    }
  }
  if (parent === ElementKind.ElementKind_Root) {
    eventEmitter.emit("item", update);
  }
  if (parent === ElementKind.ElementKindPlayer) {
    if (fieldName === "items") {
      eventEmitter.emit("player_items", update);
    }
  }
  if (parent === ElementKind.ElementKindZone) {
    if (fieldName === "interactables") {
      eventEmitter.emit("zone_interactables", update);
    }
  }
  if (parent === ElementKind.ElementKindZoneItem) {
    if (fieldName === "item") {
      eventEmitter.emit("zoneItem_item", update);
    }
  }
}

function emitPosition(update: Position, parent: ElementKind, fieldName: string) {
  if (parent === ElementKind.ElementKind_Root) {
    eventEmitter.emit("position", update);
  }
  if (parent === ElementKind.ElementKindPlayer) {
    if (fieldName === "position") {
      eventEmitter.emit("player_position", update);
    }
  }
  if (parent === ElementKind.ElementKindZoneItem) {
    if (fieldName === "position") {
      eventEmitter.emit("zoneItem_position", update);
    }
  }
  if (parent === ElementKind.ElementKindItem) {
    if (fieldName === "origin") {
      eventEmitter.emit("item_origin", update);
    }
  }
}

function emitPlayer(update: Player, parent: ElementKind, fieldName: string) {
  if (update.equipmentSets !== null && update.equipmentSets !== undefined) {
    for (const id in update.equipmentSets) {
      emitElementReference(update.equipmentSets[id], ElementKind.ElementKindPlayer, "equipmentSets");
    }
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    emitGearScore(update.gearScore, ElementKind.ElementKindPlayer, "gearScore");
  }
  if (update.guildMembers !== null && update.guildMembers !== undefined) {
    for (const id in update.guildMembers) {
      emitElementReference(update.guildMembers[id], ElementKind.ElementKindPlayer, "guildMembers");
    }
  }
  if (update.items !== null && update.items !== undefined) {
    for (const id in update.items) {
      emitItem(update.items[id], ElementKind.ElementKindPlayer, "items");
    }
  }
  if (update.position !== null && update.position !== undefined) {
    emitPosition(update.position, ElementKind.ElementKindPlayer, "position");
  }
  if (update.target !== null && update.target !== undefined) {
    emitElementReference(update.target, ElementKind.ElementKindPlayer, "target");
  }
  if (update.targetedBy !== null && update.targetedBy !== undefined) {
    for (const id in update.targetedBy) {
      emitElementReference(update.targetedBy[id], ElementKind.ElementKindPlayer, "targetedBy");
    }
  }
  if (parent === ElementKind.ElementKind_Root) {
    eventEmitter.emit("player", update);
  }
  if (parent === ElementKind.ElementKindZone) {
    if (fieldName === "player") {
      eventEmitter.emit("zone_players", update);
    }
    if (fieldName === "interactables") {
      eventEmitter.emit("zone_interactables", update);
    }
  }
  if (parent === ElementKind.ElementKindItem) {
    if (fieldName === "origin") {
      eventEmitter.emit("item_origin", update);
    }
  }
}

function emitZone(update: Zone, parent: ElementKind, fieldName: string) {
  if (update.interactables !== null && update.interactables !== undefined) {
    for (const id in update.interactables) {
      emitPosition(update.interactables[id], ElementKind.ElementKindZone, "interactables");
    }
  }
  if (update.items !== null && update.items !== undefined) {
    for (const id in update.items) {
      emitZoneItem(update.items[id], ElementKind.ElementKindZone, "items");
    }
  }
  if (update.players !== null && update.players !== undefined) {
    for (const id in update.players) {
      emitPlayer(update.players[id], ElementKind.ElementKindZone, "players");
    }
  }
  if (parent === ElementKind.ElementKind_Root) {
    eventEmitter.emit("zone", update);
  }
}

function emitZoneItem(update: ZoneItem, parent: ElementKind, fieldName: string) {
  if (update.item !== null && update.item !== undefined) {
    emitGearScore(update.item, ElementKind.ElementKindZoneItem, "item");
  }
  if (update.position !== null && update.position !== undefined) {
    emitGearScore(update.position, ElementKind.ElementKindZoneItem, "position");
  }
  if (parent === ElementKind.ElementKind_Root) {
    eventEmitter.emit("zoneItem", update);
  }
  if (parent === ElementKind.ElementKindZone) {
    if (fieldName === "item") {
      eventEmitter.emit("zone_items", update);
    }
    if (fieldName === "interactables") {
      eventEmitter.emit("zone_interactables", update);
    }
  }
}

function emitElementReference(update: ElementReference, parent: ElementKind, fieldName: string) {
  if (parent === ElementKind.ElementKindItem) {
    if (fieldName === "boundTo") {
      eventEmitter.emit("item_boundTo", update);
    }
  }
  if (parent === ElementKind.ElementKindEquipmentSet) {
    if (fieldName === "equipment") {
      eventEmitter.emit("equipmentSet_equipment", update);
    }
  }
  if (parent === ElementKind.ElementKindPlayer) {
    if (fieldName === "equipmentSets") {
      eventEmitter.emit("player_equipmentSets", update);
    }
    if (fieldName === "guildMembers") {
      eventEmitter.emit("player_guildMembers", update);
    }
    if (fieldName === "target") {
      eventEmitter.emit("player_target", update);
    }
    if (fieldName === "targetedBy") {
      eventEmitter.emit("player_targetedBy", update);
    }
  }
}
