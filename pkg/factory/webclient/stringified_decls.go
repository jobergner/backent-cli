package webclient

const class_Client = `export class Client {
  private ws: WebSocket;
  private responseEmitter: EventEmitter;
  private id: string;
  constructor(url: string, onUpdate: (update: Tree) => void = () => null) {
    this.id = "";
    this.ws = new WebSocket(url);
    this.responseEmitter = new EventEmitter();
    this.ws.addEventListener("open", () => {
      console.log("WebSocket connection opened");
    });
    this.ws.addEventListener("message", (event) => {
      const message = JSON.parse(event.data);
      switch (message.kind) {
        case MessageKind.ID:
          this.id = message.content;
          break;
        case MessageKind.Update:
        case MessageKind.CurrentState:
          const update = JSON.parse(message.content) as Tree;
          emit_Update(update);
          onUpdate(update);
          import_Update(currentState, update);
          break;
        case MessageKind.Error:
          console.log(message);
          break;
        default:
          this.responseEmitter.emit(message.id, JSON.parse(message.content));
          break;
      }
    });
    this.ws.addEventListener("close", () => {
      console.log("WebSocket connection closed");
    });
  }
  public getID() {
    return this.id;
  }
  public addItemToPlayer(item: number, newName: string): Promise<AddItemToPlayerResponse> {
    const messageID = generateID();
    const message: WebSocketMessage = {
      id: messageID,
      kind: MessageKind.ActionAddItemToPlayer,
      content: JSON.stringify({item, newName}),
    };
    setTimeout(() => {
      this.ws.send(JSON.stringify(message));
    }, 0);
    return new Promise((resolve, reject) => {
      this.responseEmitter.on(messageID, (response: AddItemToPlayerResponse) => {
        resolve(response);
      });
      setTimeout(() => {
        reject(ErrResponseTimeout);
      }, responseTimeout);
    });
  }
  public movePlayer(changeX: number, changeY: number, player: number) {
    const messageID = generateID();
    const message: WebSocketMessage = {
      id: messageID,
      kind: MessageKind.ActionMovePlayer,
      content: JSON.stringify({changeX, changeY, player}),
    };
    setTimeout(() => {
      this.ws.send(JSON.stringify(message));
    }, 0);
  }
  public spawnZoneItems(items: number[]): Promise<SpawnZoneItemsResponse> {
    const messageID = generateID();
    const message: WebSocketMessage = {
      id: messageID,
      kind: MessageKind.ActionSpawnZoneItems,
      content: JSON.stringify({items}),
    };
    setTimeout(() => {
      this.ws.send(JSON.stringify(message));
    }, 0);
    return new Promise((resolve, reject) => {
      this.responseEmitter.on(messageID, (response: SpawnZoneItemsResponse) => {
        resolve(response);
      });
      setTimeout(() => {
        reject(ErrResponseTimeout);
      }, responseTimeout);
    });
  }
}`
const type_EventListener = `type EventListener = (arg: any) => void;`
const class_EventEmitter = `class EventEmitter {
  private readonly listeners = new Map<number, Set<EventListener>>();
  public on(event: number, listener: EventListener): void {
    let listeners = this.listeners.get(event);
    if (!listeners) {
      listeners = new Set<EventListener>();
      this.listeners.set(event, listeners);
    }
    listeners.add(listener);
  }
  public off(event: number, listener?: EventListener): void {
    const listeners = this.listeners.get(event);
    if (!listeners) {
      return;
    }
    if (listener) {
      listeners.delete(listener);
      if (listeners.size === 0) {
        this.listeners.delete(event);
      }
    } else {
      this.listeners.delete(event);
    }
  }
  public emit(event: number, arg: any): void {
    const listeners = this.listeners.get(event);
    if (listeners) {
      listeners.forEach((listener) => listener(arg));
    }
  }
}`
const const_ErrResponseTimeout = `const ErrResponseTimeout = "ErrResponseTimeout";`
const const_responseTimeout = `const responseTimeout = 1000;`
const const_elementRegistrar = `export const elementRegistrar: { [id: number]: boolean } = {};`
const const_eventEmitter = `export const eventEmitter = new EventEmitter();`
const interface_AddItemToPlayerResponse = `export interface AddItemToPlayerResponse {
  playerPath: string;
}`
const interface_SpawnZoneItemsResponse = `export interface SpawnZoneItemsResponse {
  newZoneItemPaths: string[];
}`
const enum_MessageKind = `export enum MessageKind {
  ID = "id",
  Error = "error",
  Update = "update",
  CurrentState = "currentState",
  ActionAddItemToPlayer = "addItemToPlayer",
  ActionMovePlayer = "movePlayer",
  ActionSpawnZoneItems = "spawnZoneItems",
}`
const enum_ReferencedDataStatus = `export enum ReferencedDataStatus {
  ReferencedDataModified = "MODIFIED",
  ReferencedDataUnchanged = "UNCHANGED",
}`
const enum_OperationKind = `export enum OperationKind {
  OperationKindDelete = "DELETE",
  OperationKindUpdate = "UPDATE",
  OperationKindCreate = "Create",
  OperationKindUnchanged = "UNCHANGED",
}`
const enum_ElementKind = `export enum ElementKind {
  ElementKind_Root = "root",
  ElementKindAttackEvent = "AttackEvent",
  ElementKindEquipmentSet = "EquipmentSet",
  ElementKindGearScore = "GearScore",
  ElementKindItem = "Item",
  ElementKindPlayer = "Player",
  ElementKindPosition = "Position",
  ElementKindZone = "Zone",
  ElementKindZoneItem = "ZoneItem",
}`
const interface_ElementReference = `export interface ElementReference {
  id: number;
  operationKind: OperationKind;
  elementID: number;
  elementKind: ElementKind;
  referencedDataStatus: ReferencedDataStatus;
  elementPath: string;
}`
const interface_ZoneItem = `export interface ZoneItem {
  id: number;
  item?: Item;
  position?: Position;
  operationKind: OperationKind;
  elementKind: ElementKind;
}`
const interface_Item = `export interface Item {
  id: number;
  boundTo?: ElementReference;
  gearScore?: GearScore;
  name?: string;
  origin?: Player | Position;
  operationKind: OperationKind;
  elementKind: ElementKind;
}`
const interface_AttackEvent = `export interface AttackEvent {
  id: number;
  target?: ElementReference;
  operationKind: OperationKind;
  elementKind: ElementKind;
}`
const interface_EquipmentSet = `export interface EquipmentSet {
  id: number;
  equipment?: { [id: number]: ElementReference };
  name?: string;
  operationKind: OperationKind;
  elementKind: ElementKind;
}`
const interface_Position = `export interface Position {
  id: number;
  x?: number;
  y?: number;
  operationKind: OperationKind;
  elementKind: ElementKind;
}`
const interface_GearScore = `export interface GearScore {
  id: number;
  level?: number;
  score?: number;
  operationKind: OperationKind;
  elementKind: ElementKind;
}`
const interface_Player = `export interface Player {
  id: number;
  action?: { [id: number]: AttackEvent };
  equipmentSets?: { [id: number]: ElementReference };
  gearScore?: GearScore;
  guildMembers?: { [id: number]: ElementReference };
  items?: { [id: number]: Item };
  position?: Position;
  target?: ElementReference;
  targetedBy?: { [id: number]: ElementReference };
  operationKind: OperationKind;
  elementKind: ElementKind;
}`
const interface_Zone = `export interface Zone {
  id: number;
  interactables?: { [id: number]: Item | Player | ZoneItem };
  items?: { [id: number]: ZoneItem };
  players?: { [id: number]: Player };
  tags?: string[];
  operationKind: OperationKind;
  elementKind: ElementKind;
}`
const interface_Tree = `export interface Tree {
  attackEvent?: { [id: number]: AttackEvent };
  equipmentSet?: { [id: number]: EquipmentSet };
  gearScore?: { [id: number]: GearScore };
  item?: { [id: number]: Item };
  player?: { [id: number]: Player };
  position?: { [id: number]: Position };
  zone?: { [id: number]: Zone };
  zoneItem?: { [id: number]: ZoneItem };
}`
const const_currentState = `export const currentState: Tree = {};`
const function_import_Update = `export function import_Update(current: Tree, update: Tree) {
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
}`
const function_importEquipmentSet = `function importEquipmentSet(current: EquipmentSet | null | undefined, update: EquipmentSet): EquipmentSet {
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
}`
const function_importGearScore = `function importGearScore(current: GearScore | null | undefined, update: GearScore): GearScore {
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
}`
const function_importItem = `function importItem(current: Item | null | undefined, update: Item): Item {
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
    if (update.elementKind === ElementKind.ElementKindPlayer) {
      current.origin = importPlayer(current.origin as Player, update.origin);
    }
    if (update.elementKind === ElementKind.ElementKindPosition) {
      current.origin = importPosition(current.origin as Position, update.origin);
    }
  }
  return current;
}`
const function_importPosition = `function importPosition(current: Position | null | undefined, update: Position): Position {
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
}`
const function_importPlayer = `function importPlayer(current: Player | null | undefined, update: Player): Player {
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
}`
const function_importZone = `function importZone(current: Zone | null | undefined, update: Zone): Zone {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.interactables !== null && update.interactables !== undefined) {
    if (current.interactables === null || current.interactables === undefined) {
      current.interactables = {};
    }
    for (const id in update.interactables) {
      if (update.interactables[id].elementKind === ElementKind.ElementKindItem) {
        if (update.interactables[id].operationKind === OperationKind.OperationKindDelete) {
          delete current.interactables[id];
        } else {
          current.interactables[id] = importItem(current.interactables[id] as Item, update.interactables[id]);
        }
      }
      if (update.interactables[id].elementKind === ElementKind.ElementKindPlayer) {
        if (update.interactables[id].operationKind === OperationKind.OperationKindDelete) {
          delete current.interactables[id];
        } else {
          current.interactables[id] = importPlayer(current.interactables[id] as Player, update.interactables[id]);
        }
      }
      if (update.interactables[id].elementKind === ElementKind.ElementKindZoneItem) {
        if (update.interactables[id].operationKind === OperationKind.OperationKindDelete) {
          delete current.interactables[id];
        } else {
          current.interactables[id] = importZoneItem(current.interactables[id] as ZoneItem, update.interactables[id]);
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
}`
const function_importZoneItem = `function importZoneItem(current: ZoneItem | null | undefined, update: ZoneItem): ZoneItem {
  if (current === null || current === undefined) {
    return update;
  }
  if (update.item !== null && update.item !== undefined) {
    current.item = importItem(current.item, update.item);
  }
  if (update.position !== null && update.position !== undefined) {
    current.position = importPosition(current.position, update.position);
  }
  return current;
}`
const function_importElementReference = `function importElementReference(current: ElementReference | null | undefined, update: ElementReference): ElementReference {
  return update;
}`
const function_emit_Update = `export function emit_Update(update: Tree) {
  if (update.attackEvent !== null && update.attackEvent !== undefined) {
    for (const id in update.attackEvent) {
      emitAttackEvent(update.attackEvent[id]);
    }
  }
  if (update.equipmentSet !== null && update.equipmentSet !== undefined) {
    for (const id in update.equipmentSet) {
      emitEquipmentSet(update.equipmentSet[id]);
    }
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    for (const id in update.gearScore) {
      emitGearScore(update.gearScore[id]);
    }
  }
  if (update.item !== null && update.item !== undefined) {
    for (const id in update.item) {
      emitItem(update.item[id]);
    }
  }
  if (update.player !== null && update.player !== undefined) {
    for (const id in update.player) {
      emitPlayer(update.player[id]);
    }
  }
  if (update.position !== null && update.position !== undefined) {
    for (const id in update.position) {
      emitPosition(update.position[id]);
    }
  }
  if (update.zone !== null && update.zone !== undefined) {
    for (const id in update.zone) {
      emitZone(update.zone[id]);
    }
  }
  if (update.zoneItem !== null && update.zoneItem !== undefined) {
    for (const id in update.zoneItem) {
      emitZoneItem(update.zoneItem[id]);
    }
  }
}`
const function_emitAttackEvent = `function emitAttackEvent(update: AttackEvent) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  if (update.target !== null && update.target !== undefined) {
    emitElementReference(update.target);
  }
  eventEmitter.emit(update.id, update);
}`
const function_emitEquipmentSet = `function emitEquipmentSet(update: EquipmentSet) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  if (update.equipment !== null && update.equipment !== undefined) {
    for (const id in update.equipment) {
      emitElementReference(update.equipment[id]);
    }
  }
  eventEmitter.emit(update.id, update);
}`
const function_emitGearScore = `function emitGearScore(update: GearScore) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  eventEmitter.emit(update.id, update);
}`
const function_emitItem = `function emitItem(update: Item) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  if (update.boundTo !== null && update.boundTo !== undefined) {
    emitElementReference(update.boundTo);
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    emitGearScore(update.gearScore);
  }
  if (update.origin !== null && update.origin !== undefined) {
    if (update.elementKind === ElementKind.ElementKindPlayer) {
      emitPlayer(update.origin);
    }
    if (update.elementKind === ElementKind.ElementKindPosition) {
      emitPosition(update.origin);
    }
  }
  eventEmitter.emit(update.id, update);
}`
const function_emitPosition = `function emitPosition(update: Position) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  eventEmitter.emit(update.id, update);
}`
const function_emitPlayer = `function emitPlayer(update: Player) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  if (update.action !== null && update.action !== undefined) {
    for (const id in update.action) {
      emitAttackEvent(update.action[id]);
    }
  }
  if (update.equipmentSets !== null && update.equipmentSets !== undefined) {
    for (const id in update.equipmentSets) {
      emitElementReference(update.equipmentSets[id]);
    }
  }
  if (update.gearScore !== null && update.gearScore !== undefined) {
    emitGearScore(update.gearScore);
  }
  if (update.guildMembers !== null && update.guildMembers !== undefined) {
    for (const id in update.guildMembers) {
      emitElementReference(update.guildMembers[id]);
    }
  }
  if (update.items !== null && update.items !== undefined) {
    for (const id in update.items) {
      emitItem(update.items[id]);
    }
  }
  if (update.position !== null && update.position !== undefined) {
    emitPosition(update.position);
  }
  if (update.target !== null && update.target !== undefined) {
    emitElementReference(update.target);
  }
  if (update.targetedBy !== null && update.targetedBy !== undefined) {
    for (const id in update.targetedBy) {
      emitElementReference(update.targetedBy[id]);
    }
  }
  eventEmitter.emit(update.id, update);
}`
const function_emitZone = `function emitZone(update: Zone) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  if (update.interactables !== null && update.interactables !== undefined) {
    for (const id in update.interactables) {
      if (update.interactables[id].elementKind === ElementKind.ElementKindItem) {
        emitItem(update.interactables[id]);
      }
      if (update.interactables[id].elementKind === ElementKind.ElementKindPlayer) {
        emitPlayer(update.interactables[id]);
      }
      if (update.interactables[id].elementKind === ElementKind.ElementKindZoneItem) {
        emitZoneItem(update.interactables[id]);
      }
    }
  }
  if (update.items !== null && update.items !== undefined) {
    for (const id in update.items) {
      emitZoneItem(update.items[id]);
    }
  }
  if (update.players !== null && update.players !== undefined) {
    for (const id in update.players) {
      emitPlayer(update.players[id]);
    }
  }
  eventEmitter.emit(update.id, update);
}`
const function_emitZoneItem = `function emitZoneItem(update: ZoneItem) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  if (update.item !== null && update.item !== undefined) {
    emitItem(update.item);
  }
  if (update.position !== null && update.position !== undefined) {
    emitPosition(update.position);
  }
  eventEmitter.emit(update.id, update);
}`
const function_emitElementReference = `function emitElementReference(update: ElementReference) {
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    return;
  }
  if (update.operationKind === OperationKind.OperationKindDelete && elementRegistrar[update.id] !== undefined) {
    delete elementRegistrar[update.id];
  }
  if (update.operationKind !== OperationKind.OperationKindDelete && elementRegistrar[update.id] === undefined) {
    update.operationKind = OperationKind.OperationKindCreate;
    elementRegistrar[update.id] = true;
  }
  eventEmitter.emit(update.id, update);
}`
const interface_WebSocketMessage = `export interface WebSocketMessage {
  id: number;
  kind: string;
  content: string;
}`
const function_generateID = `function generateID(): number {
  const max = 10 ** 10;
  let n = 0;
  for (let i = 0; i < 10; i++) {
    n = n * 10 + Math.floor(Math.random() * 10);
  }
  return n % max;
}`

var decl_to_string_decl_collection = map[string]string{
	"class_Client":                      class_Client,
	"class_EventEmitter":                class_EventEmitter,
	"const_ErrResponseTimeout":          const_ErrResponseTimeout,
	"const_currentState":                const_currentState,
	"const_elementRegistrar":            const_elementRegistrar,
	"const_eventEmitter":                const_eventEmitter,
	"const_responseTimeout":             const_responseTimeout,
	"enum_ElementKind":                  enum_ElementKind,
	"enum_MessageKind":                  enum_MessageKind,
	"enum_OperationKind":                enum_OperationKind,
	"enum_ReferencedDataStatus":         enum_ReferencedDataStatus,
	"function_emitAttackEvent":          function_emitAttackEvent,
	"function_emitElementReference":     function_emitElementReference,
	"function_emitEquipmentSet":         function_emitEquipmentSet,
	"function_emitGearScore":            function_emitGearScore,
	"function_emitItem":                 function_emitItem,
	"function_emitPlayer":               function_emitPlayer,
	"function_emitPosition":             function_emitPosition,
	"function_emitZone":                 function_emitZone,
	"function_emitZoneItem":             function_emitZoneItem,
	"function_emit_Update":              function_emit_Update,
	"function_generateID":               function_generateID,
	"function_importElementReference":   function_importElementReference,
	"function_importEquipmentSet":       function_importEquipmentSet,
	"function_importGearScore":          function_importGearScore,
	"function_importItem":               function_importItem,
	"function_importPlayer":             function_importPlayer,
	"function_importPosition":           function_importPosition,
	"function_importZone":               function_importZone,
	"function_importZoneItem":           function_importZoneItem,
	"function_import_Update":            function_import_Update,
	"interface_AddItemToPlayerResponse": interface_AddItemToPlayerResponse,
	"interface_AttackEvent":             interface_AttackEvent,
	"interface_ElementReference":        interface_ElementReference,
	"interface_EquipmentSet":            interface_EquipmentSet,
	"interface_GearScore":               interface_GearScore,
	"interface_Item":                    interface_Item,
	"interface_Player":                  interface_Player,
	"interface_Position":                interface_Position,
	"interface_SpawnZoneItemsResponse":  interface_SpawnZoneItemsResponse,
	"interface_Tree":                    interface_Tree,
	"interface_WebSocketMessage":        interface_WebSocketMessage,
	"interface_Zone":                    interface_Zone,
	"interface_ZoneItem":                interface_ZoneItem,
	"type_EventListener":                type_EventListener,
}
