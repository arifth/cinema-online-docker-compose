import React, { useState } from "react";
import {
  Flex,
  Image,
  VStack,
  Heading,
  Button,
  HStack,
  Text,
  Box,
  useDisclosure,
} from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import NavBar from "../components/NavBar";
import Film from "../data/Film.json";
import ModalPayment from "../components/ModalPayment";

export default function DetailFilm() {
  const [opened, setOpened] = useState(false);
  const handleOpen = () => {
    setOpened(true);
  };
  const handleClose = () => {
    setOpened(false);
  };
  const { id } = useParams();
  console.log(id);

  return (
    <>
      <NavBar />

      <Flex align={"center"} justify={"center"} w="90%" margin={"auto"}>
        {Film.map((film) => {
          if (film.id == id) {
            return (
              <>
                <ModalPayment
                  isOpen={opened}
                  onOpen={handleOpen}
                  onClose={handleClose}
                  title={film.title}
                  total={film.price}
                />
                <Image
                  key={film.id}
                  src={`/${film.image}`}
                  height={"500px"}
                  draggable="false"
                />
                <VStack p={"3rem"}>
                  <HStack
                    ml={"2rem"}
                    mb={"2rem"}
                    display={"flex"}
                    alignContent={"center"}
                    justifyContent={"space-around"}
                    w={"100%"}
                  >
                    <Heading>{film.title}</Heading>
                    <Button variant={"primary"} onClick={handleOpen}>
                      Buy Now
                    </Button>
                  </HStack>
                  <Box width={"85%"} height={"500px"}>
                    <iframe
                      height={"100%"}
                      width={"100%"}
                      title="naruto"
                      src={`https://www.youtube.com/embed/${film.trailer}`}
                      allowFullScreen
                    />
                  </Box>
                  <VStack style={{ marginTop: "3rem" }}>
                    <Text
                      style={{ marginBottom: "2rem" }}
                      w={"100%"}
                      color={"primary"}
                      fontWeight={"black"}
                      fontSize={"2rem"}
                    >
                      Rp, {film.price}
                    </Text>
                    <Text w={"100%"} fontWeight={"black"} fontSize={"2rem"}>
                      {film.genre}
                    </Text>
                    <Text style={{ marginTop: "3rem" }}>
                      {film.description}
                    </Text>
                  </VStack>
                </VStack>
              </>
            );
          }
        })}
      </Flex>
    </>
  );
}
