import { Divider, Icon, Intent, Card, H5 } from "@blueprintjs/core";
import ReactJson from "react-json-view";
import Empty from "./Empty";

function CurrentStateCard({ data }) {
  return (
    <Card elevation={0} className="card card1">
      <H5>
        <Icon
          className="HeadlineIcon"
          iconSize={17}
          icon="diagram-tree"
          intent={Intent.PRIMARY}
        />
        Current State
      </H5>
      <Divider />
      <div className="JsonWrapper">
        {data && <ReactJson collapsed src={data} />}
        {!data && (
          <Empty
            description={
              "As soon as you connect to the server the current server state will appear here!"
            }
          />
        )}
      </div>
    </Card>
  );
}

export default CurrentStateCard;
