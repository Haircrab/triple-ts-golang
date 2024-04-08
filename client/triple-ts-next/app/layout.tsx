import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
	title: "Triple tic-tac-toe",
	description: "Otrio!!!!",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="en" suppressHydrationWarning>
			<body className="w-full h-full bg-slate-500">{children}</body>
		</html>
	);
}
