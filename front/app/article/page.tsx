"use client";
import { useEffect, useState } from "react";
import { ArticleResponse } from "../api/article/route";
import {
  Box,
  Center,
  Image,
  Link,
  Text,
  useDisclosure,
  VStack,
} from "@chakra-ui/react";
import Loading from "@/components/loading";

export default function Article() {
  const [articles, setArticles] = useState<ArticleResponse[]>([]);
  const [loading, setLoading] = useState(false);
  const { isOpen, onClose } = useDisclosure();

  useEffect(() => {
    setLoading(true);
    fetch("/api/article")
      .then((res) => res.json())
      .then((data) => {
        setArticles(data);
        setLoading(false);
      });
  }, []);

  useEffect(() => {
    onClose();
  }, [loading, onClose]);

  return (
    <div>
      {loading ? (
        <Loading />
      ) : (
        <VStack pt={5} gap={5}>
          {articles.map((article) => (
            <Center key={article.url}>
              <Link href={article.url}>
                <Box
                  minW={"xl"}
                  borderWidth={"1px"}
                  borderRadius={"lg"}
                  overflow={"hidden"}
                  bg={"white"}
                >
                  <Center h={"80px"}>
                    <Text fontSize={"20px"}>{article.title}</Text>
                  </Center>
                </Box>
              </Link>
            </Center>
          ))}
        </VStack>
      )}
    </div>
  );
}
