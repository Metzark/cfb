import "./globals.css";
import Header from "@/components/Header/Header";
import type { Metadata } from "next";
import { Ubuntu } from "next/font/google";

const ubuntu = Ubuntu({
  weight: ["300", "400", "500", "700"],
  style: ["normal", "italic"],
  subsets: ["latin"],
  variable: "--font-ubuntu",
});

export const metadata: Metadata = {
  title: "CFB",
  description: "CFB Predictions/Analysis Tool",
  icons: {
    icon: "/football.svg",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${ubuntu.variable}`}>
        <Header />
        {children}
      </body>
    </html>
  );
}
