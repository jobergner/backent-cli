import {eventEmitter, OperationKind, ElementKind, Tree, Client, WebSocketMessage, MessageKind} from "./index";
import {Server} from "mock-socket";

test("client handles actions", async () => {
  const fakeURL = "ws://localhost:8080";
  const mockServer = new Server(fakeURL);

  const addItemToPlayerResponseContent = {playerPath: "foo.bar"};
  const spawnZoneItemsResponseContent = {newZoneItemPaths: ["foo.bar"]};

  mockServer.on("connection", (socket) => {
    socket.on("message", (data) => {
      const message = JSON.parse(data as string) as WebSocketMessage;
      const content = JSON.parse(message.content);

      switch (message.kind) {
        case MessageKind.ActionAddItemToPlayer:
          expect(message.kind).toEqual(MessageKind.ActionAddItemToPlayer);
          expect(content).toEqual({item: 123, newName: "bar"});
          const addItemToPlayerResponse = {
            id: message.id,
            kind: MessageKind.ActionAddItemToPlayer,
            content: JSON.stringify(addItemToPlayerResponseContent),
          };
          socket.send(JSON.stringify(addItemToPlayerResponse));
          break;
        case MessageKind.ActionMovePlayer:
          expect(message.kind).toEqual(MessageKind.ActionMovePlayer);
          expect(content).toEqual({changeX: 1, changeY: 2, player: 123});
          break;
        case MessageKind.ActionSpawnZoneItems:
          expect(message.kind).toEqual(MessageKind.ActionSpawnZoneItems);
          expect(content).toEqual({items: [123]});
          const spawnZoneItemsResponse = {
            id: message.id,
            kind: MessageKind.ActionSpawnZoneItems,
            content: JSON.stringify(spawnZoneItemsResponseContent),
          };
          socket.send(JSON.stringify(spawnZoneItemsResponse));
          break;
        default:
          fail("unknown message kind");
      }
    });
  });

  const client = new Client(fakeURL);

  const responseAddItemToPlayer = await client.addItemToPlayer(123, "bar");
  expect(responseAddItemToPlayer).toEqual(addItemToPlayerResponseContent);

  client.movePlayer(1, 2, 123);

  const responseSpawnZoneItems = await client.spawnZoneItems([123]);
  expect(responseSpawnZoneItems).toEqual(spawnZoneItemsResponseContent);

  mockServer.close();
});

test("client triggers updates", async () => {
  const fakeURL = "ws://localhost:8080";
  const mockServer = new Server(fakeURL);

  new Client(fakeURL);

  const updateTree: Tree = {
    equipmentSet: {
      1: {
        id: 1,
        name: "foo",
        operationKind: OperationKind.OperationKindUpdate,
        elementKind: ElementKind.ElementKindEquipmentSet,
      },
    },
  };

  const update: WebSocketMessage = {
    id: 1234567890,
    kind: MessageKind.Update,
    content: JSON.stringify(updateTree),
  };

  const emit_equipmentSet = jest.fn();
  eventEmitter.on(1, emit_equipmentSet);

  mockServer.emit("message", JSON.stringify(update));

  expect(emit_equipmentSet).toHaveBeenCalledTimes(1);
  mockServer.close();
});
