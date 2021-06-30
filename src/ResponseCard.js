import {Divider, Icon, Intent, Button, Card, H5, Classes } from "@blueprintjs/core";
import ReactJson from 'react-json-view'
import config from "./example.config.json"
import Empty from "./Empty"


function ResponseCard() {
  return (
      <Card elevation={0} className="card card1">
        <H5><Icon className="HeadlineIcon" iconSize={17} icon="download" intent={Intent.PRIMARY} />Latest Response</H5>
        <Divider />
      <div className="JsonWrapper">
        {/* <ReactJson collapsed src={config} /> */}
        <Empty description={"Send an Action to receive a Response."}/>
      </div>
      </Card>
  );
}

export default ResponseCard;
