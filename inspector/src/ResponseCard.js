import { Divider, Icon, Intent, Card, H5 } from "@blueprintjs/core";
import ReactJson from "react-json-view";
import Empty from "./Empty";

function ResponseCard({ data }) {
  return (
    <Card elevation={0} className="card Action nospacebetween">
      <div>
        <H5>
          <Icon
            className="HeadlineIcon"
            iconSize={17}
            icon="download"
            intent={Intent.PRIMARY}
          />
          Latest Response
        </H5>
        <Divider />
      </div>
      <div className="JsonWrapper">
        {data && <ReactJson collapsed src={data} />}
        {!data && (
          <div>
            <Empty
              description={
                <div style={{display: "flex", flexDirection: "column"}}>
                  <span>Send an Action to receive a Response</span>
                  <span style={{ color: "orange" }}>
                    Not all messages return a direct response though!
                  </span>
                </div>
              }
            />
          </div>
        )}
      </div>
    </Card>
  );
}

export default ResponseCard;
