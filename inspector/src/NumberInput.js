import { NumericInput } from "@blueprintjs/core";

function NumberInput(props) {
  const { setFormContent, currentFormContent, fieldName } = props;
  return (
    <>
      <NumericInput
        className="NumberInput"
        placeholder={currentFormContent[fieldName]}
        onValueChange={(value) => {
          setFormContent({ ...currentFormContent, [fieldName]: value });
        }}
      />
    </>
  );
}

export default NumberInput;
