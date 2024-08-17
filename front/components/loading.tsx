import { Spinner, Box } from "@chakra-ui/react";

export default function Loading() {
  return (
    <Box
      display="flex"
      alignItems="center"
      justifyContent="center"
      height="100vh"
      width="100vw"
      position="fixed"
      top="0"
      left="0"
      bg="white"
      zIndex={99999}
      opacity={0.5}
    >
      <Spinner size="xl" color="orange" speed="0.8s" thickness="5px" />
    </Box>
  );
}
