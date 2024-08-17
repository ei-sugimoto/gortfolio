import { Heading, HStack, VStack } from "@chakra-ui/react";
import Image from "next/image";

export default function Skills() {
  return (
    <>
      <Heading>Skills</Heading>
      <VStack>
        <Heading>language</Heading>
        <HStack w={"100%"} justify={"space-evenly"}>
          <Image
            width={60}
            height={60}
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg"
            alt="go"
          />

          <Image
            width={60}
            height={60}
            alt="ts"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/typescript/typescript-original.svg"
          />

          <Image
            width={60}
            height={60}
            alt="ruby"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/ruby/ruby-original.svg"
          />

          <Image
            alt="C++"
            width={60}
            height={60}
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/cplusplus/cplusplus-plain.svg"
          />
        </HStack>
        <Heading>framework</Heading>
        <HStack w={"100%"} justify={"space-evenly"}>
          <Image
            width={60}
            height={60}
            alt="rails"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/rails/rails-plain.svg"
          />

          <Image
            width={60}
            height={60}
            alt="react"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/react/react-original-wordmark.svg"
          />

          <Image
            width={60}
            height={60}
            alt="next.js"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/nextjs/nextjs-original.svg"
          />

          <Image
            alt="nest"
            width={60}
            height={60}
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/nestjs/nestjs-original.svg"
          />

          <Image
            width={60}
            height={60}
            alt="vue"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/vuejs/vuejs-original.svg"
          />
        </HStack>
        <Heading>Others</Heading>
        <HStack w={"100%"} justify={"space-evenly"}>
          <Image
            width={60}
            height={60}
            alt="docker"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original-wordmark.svg"
          />

          <Image
            width={60}
            height={60}
            alt="mysql"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/mysql/mysql-original-wordmark.svg"
          />

          <Image
            width={60}
            height={60}
            alt="postgresql"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original-wordmark.svg"
          />

          <Image
            width={60}
            height={60}
            alt="prisma"
            src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/prisma/prisma-original.svg"
          />
        </HStack>
      </VStack>
    </>
  );
}
