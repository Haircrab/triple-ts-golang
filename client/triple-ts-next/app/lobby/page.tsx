"use client";

import React, { FC, useState } from "react";
import { useRouter } from "next/navigation";
import { roomIdKey } from "../utils/const";

type LobbyPageProps = { children?: React.ReactNode };

const LobbyPage: FC<LobbyPageProps> = () => {
	const router = useRouter();

	const [roomId, setRoomId] = useState<string>("");

	return (
		<>
			<div className="flex items-center">
				<div className="flex items-center">
					<div>Room ID: </div>
					<input
						className="border-orange-300 border-2"
						value={roomId}
						onChange={(e) => setRoomId(e.target.value)}
					></input>
				</div>

				<button disabled={!roomId} onClick={() => router.push(`/game?${roomIdKey}=${roomId}`)}>
					Join Room
				</button>
			</div>

			<div>
				<button onClick={() => router.push("/game")}>New Room</button>
			</div>
		</>
	);
};

export default LobbyPage;
