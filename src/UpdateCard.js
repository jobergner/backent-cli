import { Divider, Icon, Intent, Card, H5 } from "@blueprintjs/core";
import ReactJson from "react-json-view";
import config from "./example.config.json";

function UpdateCard() {
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
        <ReactJson collapsed src={config} />
      </div>
    </Card>
  );
}

export default UpdateCard;
