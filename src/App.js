import React from "react"
import "./App.css";
import Actions from "./Actions"
import ResponseCard from "./ResponseCard"
import InitialStateCard from "./InitialStateCard"
import UpdateCard from "./UpdateCard"
import AppBar from "./AppBar"
import { useEffect, useState, useRef } from "react";

class App extends React.Component {

  state = {
    ws: null,
    receivedData: {},
    socketStatus: "closed",
  }

  setSocketStatus = (newStatus) => {
    this.setState({socketStatus: newStatus })
  }

  setReceivedData = (key, data) => {
    this.setState({receivedData: {...this.state.receivedData, [key]: data }})
  }
  
  componentDidMount() {
        const ws = new WebSocket("ws://localhost:8080/ws");
        ws.open = () => this.setSocketStatus("open");
        ws.onclose = () => this.setSocketStatus("closed");

        this.setState({ws: ws})


        ws.onmessage = e => {
            const message = JSON.parse(e.data);
            if (message.kind === "currentState") {
              this.setReceivedData("initialState",JSON.parse(message.content))
            } else if (message.kind === "update") {
              this.setReceivedData("update",JSON.parse(message.content))
            } else {
              this.setReceivedData("latestResponse",JSON.parse(message.content))
            }
        };
  }

  render() {
  return (
    <>
      <div className="bp3-dark">
        <AppBar />
      </div>
      <div className="App bp3-dark">
        <InitialStateCard data={this.state.receivedData.initialState}  />
        <UpdateCard data={this.state.receivedData.update}/>
        <ResponseCard data={this.state.receivedData.latestResponse} />
        <Actions ws={this.state.ws} />
      </div>
    </>
  );
  }
}

export default App;
