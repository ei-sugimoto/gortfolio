"use client";

import { Flex, Link } from "@chakra-ui/react";

export default function Header() {
  return (
    <Flex justify={"space-around"} align={"center"}>
      <Link href="/">Home</Link>
      <Link href="/career">career</Link>
      <Link href="/article">article</Link>
      <Link href="/skills">skills</Link>
    </Flex>
  );
}
