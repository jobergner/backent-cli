import "./App.css";
import {Icon, Intent, Button, Card, H5, Classes } from "@blueprintjs/core";
import Actions from "./Actions"
import ResponseCard from "./ResponseCard"
import InitialStateCard from "./InitialStateCard"
import UpdateCard from "./UpdateCard"
import AppBar from "./AppBar"

function App() {
  return (
    <>
      <div className="bp3-dark">
        <AppBar />
      </div>
      <div className="App bp3-dark">
        <InitialStateCard />
        <UpdateCard />
        <ResponseCard />
        <Actions />
      </div>
    </>
  );
}

export default App;
