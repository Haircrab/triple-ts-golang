"use client";

// @ts-ignore
import * as io from "socket.io-client";

interface ClientToServerEvents {
	msg: () => void;
}

interface ServerToClientEvents {
	connected: () => void;
	disconnected: () => void;
	msg: () => void;
	error: () => void;
	roomId: () => void;
}

// Listen for incoming messages
export default function socketClient(roomID?: string) {
	const socket = io("localhost:8080/game", {
		query: {
			roomID,
		},
	});

	return socket;
}
