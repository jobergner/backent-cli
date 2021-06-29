import "./Actions.css";
import TextInput from "./TextInput";
import BoolInput from "./BoolInput";
import NumberInput from "./NumberInput";
import SliceInput from "./SliceInput";
import { Label } from "@blueprintjs/core";

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

function Input(props) {
  const { setFormContent, currentFormContent, action } = props;
  return (
    <>
      {Object.entries(action).map(([key, value]) => {
        const label = (
          <Label className="InputLabel">
            {key}: <span className="bp3-text-muted">({value})</span>
          </Label>
        );

        if (value.startsWith("[]")) {
          return (
            <div className="InputField" key={key}>
              {label}
              <SliceInput
                fieldName={key}
                currentFormContent={currentFormContent}
                setFormContent={setFormContent}
              />
            </div>
          );
        }
        if (textTypes.includes(value)) {
          return (
            <div className="InputField" key={key}>
              {label}
              <TextInput
                fieldName={key}
                currentFormContent={currentFormContent}
                setFormContent={setFormContent}
              />
            </div>
          );
        }
        if (numericTypes.includes(value) || value.endsWith("ID")) {
          return (
            <div className="InputField" key={key}>
              {label}
              <NumberInput
                fieldName={key}
                currentFormContent={currentFormContent}
                setFormContent={setFormContent}
              />
            </div>
          );
        }
        if (value === "bool") {
          return (
            <div className="InputField" key={key}>
              {label}
              <BoolInput
                fieldName={key}
                currentFormContent={currentFormContent}
                setFormContent={setFormContent}
              />
            </div>
          );
        }
        console.log(value);
        return <div />;
      })}
    </>
  );
}

export default Input;
