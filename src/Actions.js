import "./Actions.css";
import Input from "./Input"
import { Button, Card, H5, Classes, Intent } from "@blueprintjs/core";
import config from "./example.config.json"

function Actions() {

    return (
        <>
            {Object.entries(config.actions).map(([key, value], i) => {
                console.log(key, value)
                return <Card elevation={0} className="card" key={key}>
                    <H5>Action: {key}</H5>
                    <Input action={value} />
                    <Button intent={Intent.PRIMARY} text="Send" className={Classes.BUTTON} />
                </Card>
            })}

        </>
    );
}

export default Actions;


// <Card elevation={0} className="card ActionMessage">
// <H5>Define Action</H5>
// <Tabs
//     animate
//     id="TabsExample"
//     vertical
// >
//     {Object.entries(config.actions).map(([key, value], i) => {
//         console.log(key, value)
//         return <Tab id={key} key={key} title={key} panel={<div><Divider vertical />{key}</div>} />
//     })}
// </Tabs>
// </Card>
// <Card elevation={0} className="card ActionResponse">
// <H5>Action Response</H5>
// <p>
//     User interfaces that enable people to interact smoothly with data, ask
//     better questions, and make better decisions.
// </p>
// <Button text="Explore products" className={Classes.BUTTON} />
// </Card>