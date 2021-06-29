import React, { useState } from "react";
import { Card, TagInput, Button, Intent } from "@blueprintjs/core";
import evalInput from "./evalInput";
import { Popover2 } from "@blueprintjs/popover2";

const numericTypes = [
  "int8",
  "uint8",
  "int16",
  "uint16",
  "int32",
  "uint32",
  "int64",
  "uint64",
  "int",
  "uint",
  "uintptr",
  "float32",
  "float64",
  "complex64",
  "complex128",
];
const textTypes = ["string", "byte", "rune", "[]byte"];
// TODO expand list
const defualtValuePerValue = (value) => {
  if (numericTypes.includes(value)) {
    return 0;
  }
  if (textTypes.includes(value)) {
    return "";
  }
  return false;
};

function SliceInput(props) {
  const [isOpen, setOpen] = useState(false);
  const [newValue, setNewValue] = useState("");
  const { setFormContent, currentFormContent, fieldName, value } = props;
  console.log(newValue);

  const currentValues = currentFormContent[fieldName] || [];

  const clearButton = (
    <Button
      onClick={() =>
        setFormContent({
          ...currentFormContent,
          [fieldName]: [],
        })
      }
      key="clearButton"
      icon={"cross"}
      minimal={true}
    />
  );

  const popoverContent = (
    <Card elevation={2}>
      <span>Add New:</span>
      {evalInput(
        fieldName,
        value,
        (wrappedNewvalue) => {
          console.log(wrappedNewvalue);
          setNewValue(wrappedNewvalue[fieldName]);
        },
        { [fieldName]: newValue },
        true
      )}
      <Button intent={Intent.DANGER} minimal onClick={() => setOpen(false)}>
        close
      </Button>
      <Button
        intent={Intent.PRIMARY}
        minimal
        disabled={newValue === ""}
        onClick={() => {
          setOpen(false);
          setFormContent({
            ...currentFormContent,
            [fieldName]: [...currentValues, newValue || defualtValuePerValue(value)],
          });
        }}
      >
        add
      </Button>
    </Card>
  );

  const addButton = (
    <Popover2
      key="addButton"
      modifiers={{ arrow: { enabled: true } }}
      isOpen={isOpen}
      content={popoverContent}
    >
      <Button
        icon={"add"}
        minimal
        onClick={() => setOpen(true)}
        intent={Intent.PRIMARY}
      />
    </Popover2>
  );

  return (
    <div className="SliceInput">
      <TagInput
        onChange={(remainingTags) =>
          setFormContent({
            ...currentFormContent,
            [fieldName]: remainingTags,
          })
        }
        placeholder={fieldName}
        rightElement={[addButton, clearButton]}
        values={currentValues}
        inputProps={{ style: { display: "none" } }}
        tagProps={{ minimal: true }}
      />
    </div>
  );
}

export default SliceInput;
