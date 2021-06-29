import React, { useState } from "react";

import "./Actions.css";
import Input from "./Input";
import {
  Icon,
  Button,
  Card,
  H5,
  Classes,
  Intent,
  Divider,
} from "@blueprintjs/core";
import config from "./example.config.json";

function Action(props) {
  const [formContent, setFormContent] = useState({});
  const { keyName, value } = props;
  return (
    <Card elevation={0} className="card Action">
      <>
        <div className="ActionUpperSection">
          <H5>
            <Icon
              className="HeadlineIcon"
              iconSize={17}
              icon="send-to"
              intent={Intent.PRIMARY}
            />
            {keyName}
          </H5>
          <Divider />
          <div className="InputsWrapper">
            <Input currentFormContent={formContent} setFormContent={setFormContent} action={value} />
          </div>
        </div>
        <div className="ActionLower">
          <Divider />
          <div className="ActionSendButtonWrapper">
            <Button
              className="CardButton"
              intent={Intent.PRIMARY}
              rightIcon="inbox"
              disabled
              minimal
              text="View Response"
            />
            <Button
              intent={Intent.PRIMARY}
              rightIcon="send-message"
              text="Send"
              className={Classes.BUTTON}
                onClick={() => console.log(formContent)}
            />
          </div>
        </div>
      </>
    </Card>
  );
}

function Actions() {
  return (
    <>
      {Object.entries(config.actions).map(([keyName, value]) => {
        return <Action key={keyName} keyName={keyName} value={value} />;
      })}
    </>
  );
}

export default Actions;
