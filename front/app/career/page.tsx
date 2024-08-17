import {
  Center,
  Heading,
  HStack,
  ListItem,
  UnorderedList,
  VStack,
} from "@chakra-ui/react";

export default function Career() {
  return (
    <>
      <Heading>Career</Heading>
      <Center>
        <VStack>
          <Heading>Academic background</Heading>
          <UnorderedList spacing={10}>
            <ListItem>Kobe City College of Technology, 2017~2022</ListItem>
            <ListItem>Osaka Prefecture University, 2022~2024</ListItem>
            <ListItem>Osaka Metropolitan University, 2024~</ListItem>
          </UnorderedList>

          <Heading pt={10}>Intern</Heading>
          <UnorderedList spacing={10}>
            <ListItem>sky</ListItem>
            <ListItem>Rakuten bank</ListItem>
            <ListItem>timee</ListItem>
          </UnorderedList>
        </VStack>
      </Center>
    </>
  );
}
