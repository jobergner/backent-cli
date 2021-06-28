import * as React from "react";
import { TagInput, Button, Intent } from "@blueprintjs/core";

class SliceInput extends React.Component {
    state = {
        addOnBlur: false,
        addOnPaste: true,
        disabled: false,
        fill: false,
        intent: "none",
        large: false,
        leftIcon: true,
        tagIntents: false,
        tagMinimal: false,
    };

    render() {
        const clearButton = (
            <Button
                key="clearButton"
                icon={"cross"}
                minimal={true}
                onClick={this.handleClear}
            />
        );
        const addButton = <Button
            icon={"add"}
            key="addButton"
            minimal
            onClick={this.handleClear}
            intent={Intent.PRIMARY}
        />

        return <div className="SliceInput">
            <TagInput
                onChange={this.handleChange}
                placeholder="Separate values with commas..."
                rightElement={[addButton, clearButton]}
                values={[<div>hello</div>]}
                inputProps={{ style: { display: "none" } }}
                tagProps={{ minimal: true }}
            />
        </div>
    }

}

export default SliceInput;

