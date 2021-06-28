import "./App.css";
import { Button, Card, H5, Classes } from "@blueprintjs/core";
import Actions from "./Actions"

function App() {
  return (
    <div className="App bp3-dark">
      <Card elevation={0} className="card card1">
        <H5>hello</H5>
        <p>
          User interfaces that enable people to interact smoothly with data, ask
          better questions, and make better decisions.
        </p>
        <Button text="Explore products" className={Classes.BUTTON} />
      </Card>
      <Actions />
    </div>
  );
}

export default App;
