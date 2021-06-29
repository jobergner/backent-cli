
import * as React from "react";

import { Icon, Button, H5, Intent,  NumericInput, PanelStack2  } from "@blueprintjs/core";

const Panel1 = props => {
    const [counter, setCounter] = React.useState(0);
    const shouldOpenPanelType2 = counter % 2 === 0;

    const openNewPanel = () => {
        if (shouldOpenPanelType2) {
            props.openPanel({
                props: { counter },
                renderPanel: Panel2,
    title: <H5><Icon className="HeadlineIcon" iconSize={17} icon="inbox" intent={Intent.PRIMARY} /><span className="ResponsePanelTitle">Panel2</span></H5>,
            });
        } else {
            props.openPanel({
                props: { intent: counter % 3 === 0 ? Intent.SUCCESS : Intent.WARNING },
                renderPanel: Panel3,
    title: <H5><Icon className="HeadlineIcon" iconSize={17} icon="inbox" intent={Intent.PRIMARY} /><span className="ResponsePanelTitle">Panel3</span></H5>,
            });
        }
    };

    return (
        <div className="docs-panel-stack-contents-example">
            <Button
                intent={Intent.PRIMARY}
                onClick={openNewPanel}
                text={`Open panel type ${shouldOpenPanelType2 ? 2 : 3}`}
            />
            <NumericInput value={counter} stepSize={1} onValueChange={setCounter} />
        </div>
    );
};


const Panel2= props => {
    const openNewPanel = () => {
        props.openPanel({
            props: {},
            renderPanel: Panel1,
    title: <H5><Icon className="HeadlineIcon" iconSize={17} icon="inbox" intent={Intent.PRIMARY} /><span className="ResponsePanelTitle">Panel3</span></H5>,
        });
    };

    return (
        <div className="docs-panel-stack-contents-example">
            <H5>Parent counter was {props.counter}</H5>
            <Button intent={Intent.PRIMARY} onClick={openNewPanel} text="Open panel type 1" />
        </div>
    );
};

const Panel3 = props => {
    const openNewPanel = () => {
        props.openPanel({
            props: {},
            renderPanel: Panel1,
    title: <H5><Icon className="HeadlineIcon" iconSize={17} icon="inbox" intent={Intent.PRIMARY} /><span className="ResponsePanelTitle">Panel1</span></H5>,
        });
    };

    return (
        <div className="docs-panel-stack-contents-example">
            <Button intent={props.intent} onClick={openNewPanel} text="Open panel type 1" />
        </div>
    );
};

const RegularPanel = props => {
    const openNewPanel = () => {
        props.openPanel({
            props: {panelNumber: props.panelNumber + 1},
            renderPanel: RegularPanel,
            title: <H5><Icon className="HeadlineIcon" iconSize={17} icon="inbox" intent={Intent.PRIMARY} /><span className="ResponsePanelTitle">{props.panelNumber + 1}</span></H5>,
        });
    };

    return (
        <div className="docs-panel-stack-contents-example">
            Hello I'm panel number {props.panelNumber + 1}
            <Button intent={props.intent} onClick={openNewPanel} text="Open panel type 1" />
        </div>
    );
};

const initialPanel =  {
    props: {
        panelNumber: 1,
    },
    renderPanel: RegularPanel,
    title: <H5><Icon className="HeadlineIcon" iconSize={17} icon="inbox" intent={Intent.PRIMARY} /><span className="ResponsePanelTitle">1</span></H5>,
};

const ResponsePanel = () => {
    const [currentPanels, setCurrentPanelStack] = React.useState([initialPanel]);

    console.log(currentPanels[currentPanels.length -1])
    const addToPanelStack = React.useCallback(
        (newPanel) => setCurrentPanelStack(stack => [newPanel, ...stack]),
        [],
    );

    const removeFromPanelStack = React.useCallback(() => setCurrentPanelStack(stack => stack.slice(1)), []);

    return (
        <>
        <PanelStack2
            className="ResponsePanel"
            initialPanel={initialPanel}
            onOpen={addToPanelStack}
            onClose={removeFromPanelStack}
            showPanelHeader
            renderActivePanelOnly
        />
        </>
    );
};

export default ResponsePanel
