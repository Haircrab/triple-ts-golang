export type Move = {
  r: number;
  c: number;
  x: number;
};
export type RoomState = {
  gameState: GameState;
  p1: Player;
  p2: Player;
  p3: Player;
  p4: Player;
  connPlayerIdx: number[];
  isGameStarted: boolean;
};
export type GameState = {
  board: Board;
  playerSeq: number[];
  nextPlayerSeqIdx: number;
  winnerIdx: number;
};
export type Player = {
  id: number;
  circles: number[];
  isReady: boolean;
};
export type Board = number[][][];

export type S2CCreateRoomOkRes = {
  roomId: string;
  playerId: string;
};
export type S2CJoinRoomOkRes = S2CCreateRoomOkRes & {
  readyState: boolean[];
  roomState: RoomState;
};

export type PlayerId = 0 | 1 | 2 | 3;
export type S2cToggleReadyRes = {
  playerId: PlayerId;
  isReady: boolean;
  isGameStarted: boolean;
  roomState: RoomState;
};

export type S2cOtherPlayerMakeMoveRes = {
  playerId: PlayerId;
  move: Move;
  roomState: RoomState;
};
