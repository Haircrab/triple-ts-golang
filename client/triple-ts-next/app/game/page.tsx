"use client";

import React, { FC, useEffect, useState } from "react";
import socketFactory from "./socket";
import { useRouter, useSearchParams } from "next/navigation";
import { roomIDKey } from "../utils/const";

type GamePageProps = {};

const GamePage: FC<GamePageProps> = () => {
	const router = useRouter();
	const searchParams = useSearchParams();
	const roomId = searchParams.get(roomIDKey);

	// State to store the messages
	const [messages, setMessages] = useState<string[]>([]);
	const [rid, setrid] = useState("");
	// State to store the current message
	const [currentMessage, setCurrentMessage] = useState("");

	const [socket, setSocket] = useState<any>();

	const sendMessage = () => {
		// Send the message to the server
		socket.emit("msg", currentMessage);
		// Clear the currentMessage state
		setCurrentMessage("");
	};

	useEffect(() => {
		// Create a socket connection
		const s = socketFactory(roomId ?? undefined);
		setSocket(s);

		// Listen for incoming messages
		s.on("msg", (message: any) => {
			console.log("msg received", message);

			setMessages((prevMessages) => [...prevMessages, message]);
		});
		s.on("roomID", (message: any) => {
			setrid(message);
		});
		s.on("connect", () => {
			console.log("Connected");
		});

		s.on("disconnect", () => {
			console.log("Disconnected");
			router.push("/lobby");
		});

		s.on("error", async (err: any) => {
			console.log(`error due to ${err.message}`);
		});

		// Clean up the socket connection on unmount
		return () => {
			s.disconnect();
		};
	}, [roomId, router]);

	return (
		<div>
			{messages.map((message, index) => (
				<p key={index}>{message}</p>
			))}

			<input
				type="text"
				value={currentMessage}
				onChange={(e) => setCurrentMessage(e.target.value)}
			/>

			<button onClick={sendMessage}>Send</button>

			<div>rid: {rid}</div>
		</div>
	);
};

export default GamePage;
