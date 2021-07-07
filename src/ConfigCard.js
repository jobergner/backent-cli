import { Divider, Icon, Intent, Card, H5 } from "@blueprintjs/core";
import ReactJson from "react-json-view";
import Empty from "./Empty";

function ConfigCard({ data }) {
  return (
    <Card elevation={0} className="card card1">
      <H5>
        <Icon
          className="HeadlineIcon"
          iconSize={17}
          icon="wrench"
          intent={Intent.PRIMARY}
        />
        Config
      </H5>
      <Divider />
      <div className="JsonWrapper">
        {data && <ReactJson collapsed src={data} />}
        {!data && (
          <Empty
            description={
              "Currently trying to retrieve the config. Is your server running?"
            }
          />
        )}
      </div>
    </Card>
  );
}

export default ConfigCard;
