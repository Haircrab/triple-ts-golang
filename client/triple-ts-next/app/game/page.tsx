"use client";

import { useRouter, useSearchParams } from "next/navigation";
import React, { useEffect, useRef, useState } from "react";
import GameBoard from "../components/GameBoard";
import { roomIdKey } from "../utils/const";
import { c2s, s2c, socketClient as socketFactory } from "./socket";
import {
  GameState,
  RoomState,
  S2CCreateRoomOkRes,
  S2CJoinRoomOkRes,
  S2cOtherPlayerMakeMoveRes,
  S2cToggleReadyRes,
} from "./types";

type GamePageProps = {};

const GamePage: React.FC<GamePageProps> = () => {
  const router = useRouter();
  const searchParams = useSearchParams();
  const roomId = searchParams.get(roomIdKey);
  const [readyState, setReadyState] = useState<boolean[]>([
    false,
    false,
    false,
    false,
  ]);
  const [rid, setrid] = useState("");
  const [pid, setpid] = useState("");
  const [isGameStarted, setIsGameStarted] = useState(false);
  const [roomState, setRoomState] = useState<RoomState | undefined>();

  const socket = useRef<any>(null);

  const toggleIsReady = () => {
    socket.current?.emit(c2s.toggleReadyEvent);
  };
  const sendMove = () => {
    socket.current?.emit(c2s.playerMakeMoveEvent, { r: 1, c: 2, x: 2 });
  };

  useEffect(() => {
    // Create a socket connection
    socket.current = socketFactory(roomId ?? undefined);

    if (socket.current) {
      // Listen for incoming messages
      socket.current.on(s2c.createRoomOkEvent, (res: S2CCreateRoomOkRes) => {
        setrid(res.roomId);
        setpid(res.playerId);
      });
      socket.current.on(s2c.joinRoomOkEvent, (res: S2CJoinRoomOkRes) => {
        setrid(res.roomId);
        setpid(res.playerId);
        setReadyState(res.readyState);
      });
      socket.current.on(
        s2c.playerToggleReadyEvent,
        (res: S2cToggleReadyRes) => {
          setReadyState((prev) => [
            ...prev.slice(0, res.playerId),
            res.isReady,
            ...prev.slice(res.playerId + 1),
          ]);
          setIsGameStarted(res.isGameStarted);
          setRoomState(res.roomState);
        },
      );

      socket.current.on(
        s2c.otherPlayerMakeMoveEvent,
        (res: S2cOtherPlayerMakeMoveRes) => {
          setRoomState(res.roomState);
        },
      );

      // boilerplate socket events
      socket.current.on(s2c.connectedEvent, () => {
        console.log("Connected");
      });
      socket.current.on(s2c.disconnectedEvent, () => {
        console.log("Disconnected");
        router.push("/lobby");
      });
      socket.current.on(s2c.errorEvent, async (err: any) => {
        console.log(`error due to ${err.message}`);
      });
    }

    // Clean up the socket connection on unmount
    return () => {
      socket.current && socket.current.disconnect();
    };
  }, [roomId, router]);

  return (
    <div>
      <div>rid: {rid}</div>
      <div>pid: {pid}</div>
      {!isGameStarted ? (
        <button onClick={toggleIsReady}>toggleIsReady</button>
      ) : (
        "no button"
      )}

      <div>{JSON.stringify(readyState)}</div>
      <div>isGameStarted: {JSON.stringify(isGameStarted)}</div>

      {isGameStarted ? <GameBoard gameState={roomState?.gameState} /> : ""}

      <button onClick={sendMove}>sendMove</button>
    </div>
  );
};

export default GamePage;
