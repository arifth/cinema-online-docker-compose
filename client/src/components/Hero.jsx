import {
  Flex,
  Image,
  Box,
  Stack,
  Heading,
  Text,
  Button,
  useDisclosure,
} from "@chakra-ui/react";
import React, { useContext } from "react";
import ModalPayment from "./ModalPayment";
import ModalLogin from "./ModalLogin";
import { UserContext } from "../context/userContext";

export default function Hero() {
  const [state, dispatch] = useContext(UserContext);
  const { isOpen, onOpen, onClose } = useDisclosure();

  const price = 15000000;

  // console.log(price);
  return (
    <>
      <Flex alignItems="center" justifyContent="center" mb={"2rem"}>
        <Box width="1200px" height="auto" style={{ position: "relative" }}>
          <Image src="./LargeImg.png" rounded="lg" objectFit="contain" />
          <Stack
            style={{
              zIndex: "2",
              top: "0",
              left: "0",
              position: "absolute",
              display: "flex",
              justifyContent: "flex-start",
            }}
            p="2rem"
            w="50%"
            ml="1rem"
          >
            <Heading fontSize="3rem">DeadPool 3</Heading>
            <Text fontSize="3rem">Action</Text>
            <Text fontSize="2rem" color={"primary"}>
              Rp, <span>{price.toLocaleString()}</span>
            </Text>
            <Text>
              Hold onto your chimichangas, folks. From the studio that brought
              you all 3 Taken films comes the block-busting,
              fourth-wall-breaking masterpiece about Marvel Comics’ sexiest
              anti-hero! Starring God’s perfect idiot Ryan Reynolds and a bunch
              of other "actors," DEADPOOL is a giddy slice of awesomeness packed
              with more twists than Deadpool’s enemies’ intestines and more
              action than prom night. Amazeballs!
            </Text>
            {!state.isLogin ? (
              <>
                <Button
                  variant="primary"
                  w="30%"
                  style={{ marginTop: "2rem" }}
                  onClick={() => onOpen()}
                >
                  Buy This
                </Button>
                <ModalLogin isOpen={isOpen} onClose={onClose} />
              </>
            ) : (
              <>
                <Button
                  variant="primary"
                  w="30%"
                  style={{ marginTop: "2rem" }}
                  onClick={() => onOpen()}
                >
                  Buy This
                </Button>
                <ModalPayment
                  isOpen={isOpen}
                  onClose={onClose}
                  title={"deadpool3"}
                  price={price}
                />
              </>
            )}
          </Stack>
        </Box>
      </Flex>
    </>
  );
}
