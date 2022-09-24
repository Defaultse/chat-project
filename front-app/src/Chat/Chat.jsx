import React from 'react'
import {useEffect, useState} from 'react'
import * as WebSocket from "websocket"

export default function Chat() { 
    // const [msgArr, setMsgArr] = useState([])

    // useEffect(() => {
    //   console.log("in it")
    //   const socket = new WebSocket.w3cwebsocket('ws://localhost:8080/ws');
  
    //   socket.onopen = function () {
    //     socket.send("helloheee!")
    //     socket.onmessage = (msg) => {
    //       msgArr.push(msgArr)
    //     };
    //   };
    // }, []);
    
  return (
    <div>Chat</div>
  )
}
