import {
  Center,
  Container,
  Heading,
  Image,
  Link,
  Spacer,
  Stack,
  Text,
} from "@chakra-ui/react";

export default function Home() {
  return (
    <>
      <Container>
        <Stack gap={5}>
          <Heading> Hi, I am ei sugimoto</Heading>
          <Text fontSize={"2xl"}>I am a software engineer</Text>
          <Spacer h={"100px"} />
          <Center>
            <Image
              src="icons.jpeg"
              alt=""
              borderRadius={"full"}
              boxSize={"200px"}
            />
          </Center>
          <Stack direction={"row"}>
            <Link href="https://github.com/ei-sugimoto">
              <Image src="github.svg" alt="github" />
            </Link>
            <Link href="https://qiita.com/ei_sugimoto">
              <Image
                src="qiita.png"
                alt="qiita"
                borderRadius={"full"}
                boxSize={"30px"}
              />
            </Link>
          </Stack>
        </Stack>
      </Container>
    </>
  );
}
