import { Divider, Icon, Intent, Card, H5 } from "@blueprintjs/core";
import ReactJson from "react-json-view";
import Empty from "./Empty";

function MessageCard({ data }) {
  return (
    <Card elevation={0} className="card Action nospacebetween">
      <div>
        <H5>
          <Icon
            className="HeadlineIcon"
            iconSize={17}
            icon="send-to"
            intent={Intent.PRIMARY}
          />
          Latest Sent Message
        </H5>
        <Divider />
      </div>
      <div className="JsonWrapper">
        {data && <ReactJson collapsed src={data} />}
        {!data && (
          <div>
            <Empty description={"Send an action and it's content will be displayed here!"} />
          </div>
        )}
      </div>
    </Card>
  );
}

export default MessageCard;
