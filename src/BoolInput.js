import { Button, ButtonGroup, Intent } from "@blueprintjs/core";

function BoolInput(props) {
  const { setFormContent, currentFormContent, fieldName } = props;
    const currentValue = currentFormContent[fieldName]
  return (
    <div className="BoolInput">
      <ButtonGroup style={{ minWidth: 120 }}>
        <Button
          onClick={() =>
            setFormContent({ ...currentFormContent, [fieldName]: true })
          }
            intent={currentValue === true ? Intent.PRIMARY : Intent.NONE}
        >
          true
        </Button>
        <Button
          onClick={() =>
            setFormContent({ ...currentFormContent, [fieldName]: false })
          }
            intent={currentValue === false ? Intent.PRIMARY : Intent.NONE}
        >
          false
        </Button>
      </ButtonGroup>
    </div>
  );
}

export default BoolInput;
