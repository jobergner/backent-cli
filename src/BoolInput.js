import { Button, ButtonGroup, Intent } from "@blueprintjs/core";

function BoolInput() {
    return (<div className="BoolInput">
        <ButtonGroup style={{ minWidth: 120 }}>
            <Button intent={Intent.PRIMARY} >true</Button>
            <Button>false</Button>
        </ButtonGroup>
    </div>);
}

export default BoolInput;

