"use client";

import { Box, Flex, Link } from "@chakra-ui/react";

export default function Header() {
  return (
    <Flex justify={"space-around"} align={"center"} fontSize={"25px"}>
      <Link href="/">
        <Box
          as="button"
          bg={"gray.100"}
          px={10}
          py={1}
          borderRadius={"full"}
          m={2}
        >
          Home
        </Box>
      </Link>
      <Link href="/career">
        <Box
          as="button"
          bg={"gray.100"}
          px={10}
          py={1}
          borderRadius={"full"}
          m={2}
        >
          career
        </Box>
      </Link>
      <Link href="/article">
        <Box
          as="button"
          bg={"gray.100"}
          px={10}
          py={1}
          borderRadius={"full"}
          m={2}
        >
          article
        </Box>
      </Link>
      <Link href="/skills">
        <Box
          as="button"
          bg={"gray.100"}
          px={10}
          py={1}
          borderRadius={"full"}
          m={2}
        >
          skills
        </Box>
      </Link>
    </Flex>
  );
}
