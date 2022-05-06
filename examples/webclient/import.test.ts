import { import_Update, ReferencedDataStatus, ElementKind, currentState, Tree, OperationKind } from "./index";

test("import updates", () => {
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

  const current = {
    player: {
      1: {
        id: 1,
        items: {
          4: {
            id: 4,
            name: "oldName",
            gearScore: {
              id: 5,
              score: 1,
              level: 1,
              operationKind: OperationKind.OperationKindUpdate,
              elementKind: ElementKind.ElementKindGearScore,
            },
            operationKind: OperationKind.OperationKindUpdate,
            elementKind: ElementKind.ElementKindItem,
          },
        },
        position: {
          id: 2,
          x: 1,
          y: 1,
          operationKind: OperationKind.OperationKindUpdate,
          elementKind: ElementKind.ElementKindPosition,
        },
        operationKind: OperationKind.OperationKindUpdate,
        elementKind: ElementKind.ElementKindPlayer,
      },
    },
  };

  import_Update(current, update);

  const expected: Tree = {
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
              level: 1,
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
          y: 1,
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

  expect(current).toEqual(expected);
});

test("import creates", () => {
  const update: Tree = {
    player: {
      1: {
        id: 1,
        items: {
          4: {
            id: 4,
            name: "name",
            gearScore: {
              id: 5,
              score: 1,
              level: 1,
              operationKind: OperationKind.OperationKindUpdate,
              elementKind: ElementKind.ElementKindGearScore,
            },
            operationKind: OperationKind.OperationKindUpdate,
            elementKind: ElementKind.ElementKindItem,
          },
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

  const current = {
    player: {
      1: {
        id: 1,
        position: {
          id: 2,
          x: 1,
          y: 1,
          operationKind: OperationKind.OperationKindUpdate,
          elementKind: ElementKind.ElementKindPosition,
        },
        operationKind: OperationKind.OperationKindUpdate,
        elementKind: ElementKind.ElementKindPlayer,
      },
    },
  };

  import_Update(current, update);

  const expected: Tree = {
    player: {
      1: {
        id: 1,
        items: {
          4: {
            id: 4,
            name: "name",
            gearScore: {
              id: 5,
              score: 1,
              level: 1,
              operationKind: OperationKind.OperationKindUpdate,
              elementKind: ElementKind.ElementKindGearScore,
            },
            operationKind: OperationKind.OperationKindUpdate,
            elementKind: ElementKind.ElementKindItem,
          },
        },
        position: {
          id: 2,
          x: 1,
          y: 1,
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

  expect(current).toEqual(expected);
});

test("import deletes", () => {
  const update: Tree = {
    player: {
      1: {
        id: 1,
        items: {
          4: {
            id: 4,
            gearScore: {
              id: 5,
              operationKind: OperationKind.OperationKindDelete,
              elementKind: ElementKind.ElementKindGearScore,
            },
            operationKind: OperationKind.OperationKindDelete,
            elementKind: ElementKind.ElementKindItem,
          },
        },
        operationKind: OperationKind.OperationKindUpdate,
        elementKind: ElementKind.ElementKindPlayer,
        target: {
          operationKind: OperationKind.OperationKindDelete,
          elementID: 99,
          elementKind: ElementKind.ElementKindPlayer,
          referencedDataStatus: ReferencedDataStatus.ReferencedDataUnchanged,
          elementPath: "$.player[99]",
        },
      },
    },
  };

  const current = {
    player: {
      1: {
        id: 1,
        items: {
          4: {
            id: 4,
            gearScore: {
              id: 5,
              operationKind: OperationKind.OperationKindUpdate,
              elementKind: ElementKind.ElementKindGearScore,
            },
            operationKind: OperationKind.OperationKindUpdate,
            elementKind: ElementKind.ElementKindItem,
          },
        },
        position: {
          id: 2,
          x: 1,
          y: 1,
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

  import_Update(current, update);

  const expected: Tree = {
    player: {
      1: {
        id: 1,
        items: {},
        position: {
          id: 2,
          x: 1,
          y: 1,
          operationKind: OperationKind.OperationKindUpdate,
          elementKind: ElementKind.ElementKindPosition,
        },
        operationKind: OperationKind.OperationKindUpdate,
        elementKind: ElementKind.ElementKindPlayer,
      },
    },
  };

  expect(current).toEqual(expected);
});
