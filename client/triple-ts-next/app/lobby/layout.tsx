import React, { FC } from "react";
import LobbyPage from "./page";

type LayoutProps = { children?: React.ReactNode };

const Layout: FC<LayoutProps> = () => {
	return (
		<>
			<LobbyPage />
		</>
	);
};

export default Layout;
