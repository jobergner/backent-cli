import { NumericInput } from "@blueprintjs/core";

function NumberInput(props) {
  const { setFormContent, currentFormContent, fieldName } = props;
  return (
    <>
      <NumericInput
        className="NumberInput"
        onValueChange={(value) => {
          setFormContent({ ...currentFormContent, [fieldName]: value });
        }}
        placeholder={fieldName}
      />
    </>
  );
}

export default NumberInput;
