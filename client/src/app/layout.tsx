import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "Photos",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <div
          style={{ filter: "drop-shadow(5px 5px 42px #000000)" }}
          className="fixed left-12 top-12"
        >
          <div className="text-4xl font-bold">Photos</div>
          <div className="text-base font-bold">2,375 Items</div>
        </div>

        {children}
      </body>
    </html>
  );
}
