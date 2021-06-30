import { InputGroup } from "@blueprintjs/core";

function TextInput(props) {
  const { setFormContent, currentFormContent, fieldName } = props;
  return (
    <>
      <InputGroup
        className="TextInput"
        onChange={(e) => {
          setFormContent({
            ...currentFormContent,
            [fieldName]: e.target.value,
          });
        }}
        placeholder={fieldName}
      />
    </>
  );
}

export default TextInput;
