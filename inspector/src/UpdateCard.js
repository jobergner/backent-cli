import { Divider, Icon, Intent, Card, H5 } from "@blueprintjs/core";
import ReactJson from "react-json-view";
import Empty from "./Empty"

function UpdateCard({data}) {
  return (
    <Card elevation={0} className="card card1">
      <H5>
        <Icon
          className="HeadlineIcon"
          iconSize={17}
          icon="diagram-tree"
          intent={Intent.PRIMARY}
        />
        Latest Update
      </H5>
      <Divider />
      <div className="JsonWrapper">
      {data && <ReactJson collapsed src={data} />}
      {!data && <Empty description={"The next update the server emitts will appear here!"}/>} 
      </div>
    </Card>
  );
}

export default UpdateCard;
