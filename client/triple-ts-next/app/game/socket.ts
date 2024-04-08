"use client";

// @ts-ignore
import io from "socket.io-client";

export const c2s = {
  sendMsgEvent: "sendMsg",
  toggleReadyEvent: "toggleReady",
  playerMakeMoveEvent: "playerMakeMove",
};

export const s2c = {
  connectedEvent: "connect",
  disconnectedEvent: "disconnect",
  errorEvent: "error",

  receiveMsgEvent: "receiveMsg",
  createRoomOkEvent: "createRoomOk",
  joinRoomOkEvent: "joinRoomOk",
  playerToggleReadyEvent: "playerToggleReady",
  otherPlayerMakeMoveEvent: "otherPlayerMakeMove",
};

// Listen for incoming messages
export function socketClient(roomId?: string) {
  const socket = io("localhost:8080/game", {
    query: {
      roomId: roomId,
    },
  });

  return socket;
}
