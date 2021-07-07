import React from "react";
import "./App.css";
import Actions from "./Actions";
import axios from "axios";
import ResponseCard from "./ResponseCard";
import MessageCard from "./MessageCard";
import InitialStateCard from "./InitialStateCard";
import ConfigCard from "./ConfigCard";
import UpdateCard from "./UpdateCard";
import AppBar from "./AppBar";
import { useEffect, useState, useRef } from "react";

class App extends React.Component {
  state = {
    ws: null,
    receivedData: {},
    socketStatus: "closed",
    sentData: null,
    configData: null,
  };

  setSocketStatus = (newStatus) => {
    this.setState({ socketStatus: newStatus });
  };

  setReceivedData = (key, data) => {
    this.setState({
      receivedData: { ...this.state.receivedData, [key]: data },
    });
  };

  setSentData = (newData) => {
    this.setState({
      sentData: { ...newData },
    });
  };

  setConfigData = (newData) => {
    this.setState({
      configData: { ...newData },
    });
  };

  componentDidMount() {
    var params = new URLSearchParams(window.location.search);
    let port = params.get("port")
    if (!port) {
      window.location = window.location + "?port=8080"  
    }
    const ws = new WebSocket("ws://localhost:"+ port +"/ws");
    ws.open = () => this.setSocketStatus("open");
    ws.onclose = () => this.setSocketStatus("closed");

    this.setState({ ws: ws });

    ws.onmessage = (e) => {
      const message = JSON.parse(e.data);
      if (message.kind === "currentState") {
        this.setReceivedData("initialState", JSON.parse(message.content));
      } else if (message.kind === "update") {
        this.setReceivedData("update", JSON.parse(message.content));
      } else {
        this.setReceivedData("latestResponse", JSON.parse(message.content));
      }
    };

    axios.get("http://localhost:"+port+"/inspect").then((res) => {
      this.setConfigData(res.data)
    });
  }

  render() {
    return (
      <>
        <div className="bp3-dark">
          <AppBar />
        </div>
        <div className="App bp3-dark">
          <ConfigCard data={this.state.configData}/>
          <InitialStateCard data={this.state.receivedData.initialState} />
          <UpdateCard data={this.state.receivedData.update} />
          <ResponseCard data={this.state.receivedData.latestResponse} />
          <MessageCard data={this.state.sentData} />
          <Actions
            config={this.state.configData}
            setSentData={this.setSentData.bind(this)}
            ws={this.state.ws}
          />
        </div>
      </>
    );
  }
}

export default App;
