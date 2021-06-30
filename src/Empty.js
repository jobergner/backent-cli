import {
  NonIdealState,
} from "@blueprintjs/core";

function Empty(props) {
  return (
    <NonIdealState
      icon="time"
      title="Nothing Here Yet!"
      description={props.description}
    />
  );
}

export default Empty;
