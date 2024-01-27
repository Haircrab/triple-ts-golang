import React, { FC } from "react";
import GamePage from "./page";

type LayoutProps = {
	children?: React.ReactNode;
};

const Layout: FC<LayoutProps> = () => {
	return (
		<>
			<div>layout</div>
			<GamePage />
		</>
	);
};

export default Layout;
