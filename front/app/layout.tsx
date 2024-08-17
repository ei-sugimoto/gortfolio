import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { Providers } from "./providers";
import { Grid, GridItem } from "@chakra-ui/react";
import Header from "@/components/header";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "ei's portfolio",
  description: "ei portfolio",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={inter.className}>
        <Providers>
          <Grid gridTemplateAreas={`"header" "main"`}>
            <GridItem bg={"orange.300"} area={"header"}>
              <Header />
            </GridItem>
            <GridItem bg={"orange.100"} area={"main"} h={"100vh"}>
              {children}
            </GridItem>
          </Grid>
        </Providers>
      </body>
    </html>
  );
}
