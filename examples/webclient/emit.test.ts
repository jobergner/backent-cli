import { eventEmitter, emit_Update, ReferencedDataStatus, ElementKind, Tree, OperationKind } from "./index";

test("emit updates", () => {
  const update: Tree = {
    player: {
      1: {
        id: 1,
        items: {
          4: {
            id: 4,
            name: "newName",
            gearScore: {
              id: 5,
              score: 3,
              operationKind: OperationKind.OperationKindUpdate,
              elementKind: ElementKind.ElementKindGearScore,
            },
            operationKind: OperationKind.OperationKindUpdate,
            elementKind: ElementKind.ElementKindItem,
          },
        },
        position: {
          id: 2,
          x: 2,
          operationKind: OperationKind.OperationKindUpdate,
          elementKind: ElementKind.ElementKindPosition,
        },
        target: {
          operationKind: OperationKind.OperationKindUpdate,
          elementID: 99,
          elementKind: ElementKind.ElementKindPlayer,
          referencedDataStatus: ReferencedDataStatus.ReferencedDataModified,
          elementPath: "$.player[99]",
        },
        operationKind: OperationKind.OperationKindUpdate,
        elementKind: ElementKind.ElementKindPlayer,
      },
    },
  };

  const emit_player = jest.fn();
  eventEmitter.on("player", emit_player);

  const emitPlayer_items = jest.fn();
  eventEmitter.on("player_items", emitPlayer_items);

  const emitItem_gearScore = jest.fn();
  eventEmitter.on("item_gearScore", emitItem_gearScore);

  const emitPlayer_position = jest.fn();
  eventEmitter.on("player_position", emitPlayer_position);

  const emitPlayer_target = jest.fn();
  eventEmitter.on("player_target", emitPlayer_target);

  emit_Update(update);

  expect(emit_player).toHaveBeenCalledTimes(1);
  expect(emit_player).toHaveBeenCalledWith(update?.player?.[1]);

  expect(emitPlayer_items).toHaveBeenCalledTimes(1);
  expect(emitPlayer_items).toHaveBeenCalledWith(update?.player?.[1].items?.[4]);

  expect(emitItem_gearScore).toHaveBeenCalledTimes(1);
  expect(emitItem_gearScore).toHaveBeenCalledWith(update?.player?.[1].items?.[4].gearScore);

  expect(emitPlayer_position).toHaveBeenCalledTimes(1);
  expect(emitPlayer_position).toHaveBeenCalledWith(update?.player?.[1].position);

  expect(emitPlayer_target).toHaveBeenCalledTimes(1);
  expect(emitPlayer_target).toHaveBeenCalledWith(update?.player?.[1].target);
});
