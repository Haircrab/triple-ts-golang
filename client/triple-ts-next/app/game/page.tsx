"use client";

import React, { FC, useEffect, useState } from "react";
import { socketClient as socketFactory, c2s, s2c } from "./socket";
import { useRouter, useSearchParams } from "next/navigation";
import { roomIdKey } from "../utils/const";
import { S2CCreateRoomOkRes, S2CJoinRoomOkRes, S2cToggleReadyRes } from "./types";

type GamePageProps = {};

const GamePage: FC<GamePageProps> = () => {
  const router = useRouter();
  const searchParams = useSearchParams();
  const roomId = searchParams.get(roomIdKey);
  const [readyState, setReadyState] = useState<boolean[]>([false, false, false, false]);
  const [rid, setrid] = useState("");
  const [pid, setpid] = useState("");
  const [isGameStarted, setIsGameStarted] = useState(false);

  const [socket, setSocket] = useState<any>();

  const toggleIsReady = () => {
    socket.emit(c2s.toggleReadyEvent);
  };

  useEffect(() => {
    // Create a socket connection
    const s = socketFactory(roomId ?? undefined);
    setSocket(s);

    // Listen for incoming messages
    s.on(s2c.createRoomOkEvent, (res: S2CCreateRoomOkRes) => {
      setrid(res.roomId);
      setpid(res.playerId);
    });
    s.on(s2c.joinRoomOkEvent, (res: S2CJoinRoomOkRes) => {
      console.log("isGameStarted", isGameStarted)
      setrid(res.roomId);
      setpid(res.playerId);
      setReadyState(res.readyState)
    });
    s.on(s2c.playerToggleReadyEvent, (res: S2cToggleReadyRes) => {
      console.log("playerToggleReadyEvent", res)
      setReadyState((prev) => [...prev.slice(0, res.playerId), res.isReady, ...prev.slice(res.playerId + 1)])
      setIsGameStarted(res.isGameStarted)
    })

    // boilerplate socket events
    s.on(s2c.connectedEvent, () => {
      console.log("Connected");
    });
    s.on(s2c.disconnectedEvent, () => {
      console.log("Disconnected");
      router.push("/lobby");
    });
    s.on(s2c.errorEvent, async (err: any) => {
      console.log(`error due to ${err.message}`);
    });

    // Clean up the socket connection on unmount
    return () => {
      s.disconnect();
    };
  }, [roomId, router])

  return (
    <div>
      <div>rid: {rid}</div>
      <div>pid: {pid}</div>
      <div>{JSON.stringify(readyState)}</div>
      {!isGameStarted ? <button onClick={toggleIsReady}>toggleIsReady</button> : "no button"}
      <div>isGameStarted: {JSON.stringify(isGameStarted)}</div>
    </div>
  );
};

export default GamePage;
