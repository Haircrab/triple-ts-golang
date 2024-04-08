import React from "react";
import { Box, Grid } from "@mui/material";
import { GameState } from "@/app/game/types";

type GameBoardProps = { gameState?: GameState };
const GameBoard: React.FC<GameBoardProps> = ({ gameState }) => {
  if (!gameState) return <></>;

  return (
    <Box sx={{ flexGrow: 1 }}>
      <Grid container spacing={16}>
        <Grid container item>
          <FormRow row={gameState.board[0]} />
        </Grid>

        <Grid container item>
          <FormRow row={gameState.board[1]} />
        </Grid>

        <Grid container item>
          <FormRow row={gameState.board[2]} />
        </Grid>

        <Grid container item></Grid>

        {/* <Grid container item> */}
        {/*   <FormRow /> */}
        {/* </Grid> */}
      </Grid>
    </Box>
  );
};
export default GameBoard;

type FormRowProps = {
  row: number[][];
};
const FormRow: React.FC<FormRowProps> = ({ row }) => {
  return (
    <React.Fragment>
      <Grid item xs={4}>
        <BoardCell
          l={row[0][0] as BoardCellIdx}
          m={row[0][1] as BoardCellIdx}
          s={row[0][2] as BoardCellIdx}
        />
      </Grid>
      <Grid item xs={4}>
        <BoardCell
          l={row[0][0] as BoardCellIdx}
          m={row[0][1] as BoardCellIdx}
          s={row[0][2] as BoardCellIdx}
        />
      </Grid>
      <Grid item xs={4}>
        <BoardCell
          l={row[0][0] as BoardCellIdx}
          m={row[0][1] as BoardCellIdx}
          s={row[0][2] as BoardCellIdx}
        />
      </Grid>
    </React.Fragment>
  );
};

type BoardCellIdx = 0 | 1 | 2 | 3 | 4;
type BoardCellProps = {
  s?: BoardCellIdx;
  m?: BoardCellIdx;
  l?: BoardCellIdx;
};

const colorMap = {
  0: "bg-zinc-300",
  1: "bg-red-300",
  2: "bg-blue-300",
  3: "bg-green-300",
  4: "bg-green-300",
};

const BoardCell: React.FC<BoardCellProps> = ({ s, m, l }) => {
  return (
    <div className="relative">
      <div
        className={`absolute bottom-0 left-0 right-0 top-0 grid place-content-center place-items-center z-10`}
      >
        <div
          className={`rounded-full w-28 h-28 ${colorMap[s ?? 0]} border-black border-2`}
        ></div>
      </div>

      <div
        className={`absolute bottom-0 left-0 right-0 top-0 grid place-content-center place-items-center z-20`}
      >
        <div
          className={`rounded-full w-20 h-20 ${colorMap[m ?? 0]} border-black border-2`}
        ></div>
      </div>

      <div
        className={`absolute bottom-0 left-0 right-0 top-0 grid place-content-center place-items-center z-30`}
      >
        <div
          className={`rounded-full w-8 h-8 ${colorMap[l ?? 0]} border-black border-2`}
        ></div>
      </div>
    </div>
  );
};
