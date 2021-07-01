import {Divider, Icon, Intent, Button, Card, H5, Classes } from "@blueprintjs/core";
import ReactJson from 'react-json-view'
import Empty from "./Empty"


function ResponseCard({data}) {
  return (
      <Card elevation={0} className="card card1">
        <H5><Icon className="HeadlineIcon" iconSize={17} icon="download" intent={Intent.PRIMARY} />Latest Response</H5>
        <Divider />
      <div className="JsonWrapper">
      {data && <ReactJson collapsed src={data} />}
      {!data && <Empty description={"Send an Action to receive a Response."}/>} 
      </div>
      </Card>
  );
}

export default ResponseCard;
