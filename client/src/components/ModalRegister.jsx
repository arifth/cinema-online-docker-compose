import {
  Modal,
  Button,
  ModalOverlay,
  ModalBody,
  ModalFooter,
  ModalContent,
  ModalHeader,
  Input,
  Flex,
  Text,
  Fade,
  Box,
} from "@chakra-ui/react";
import React, { useState } from "react";
import { useMutation } from "react-query";
import { API } from "../config/api";

export default function ModalRegister({ isOpen, onClose }) {
  const [input, setInput] = useState({
    name: "",
    email: "",
    password: "",
  });

  const [succesCode, setSuccesCode] = useState(false);

  const handleSubmit = useMutation(async (e) => {
    e.preventDefault();
    // NOTES new implementation using useMutation
    try {
      const config = {
        headers: {
          "Content-Type": "application/json",
        },
      };

      // insert user data to database
      const response = await API.post("/register", input);
      if (response.status === 200) {
        setSuccesCode(true);
        setTimeout(onClose(), 2000);
      }

      return response;
    } catch (error) {
      console.log(error);
    }
  });

  return (
    <>
      <Modal isOpen={isOpen} onClose={onClose} isCentered>
        <ModalOverlay backdropFilter="blur(10px) hue-rotate(0deg)" />
        <ModalContent bgColor={"hsla(0, 0%, 0%, 0.7)"} padding={"2rem"}>
          <ModalHeader color={"primary"} fontSize={"3rem"}>
            Register
          </ModalHeader>
          <ModalBody>
            <Flex flexDirection={"column"} gap={"1rem"} marginBottom={"1rem"}>
              <Input
                bgColor={"#343434"}
                placeholder="Name"
                type="text"
                value={input.name}
                onChange={(e) => setInput({ ...input, name: e.target.value })}
              />
              <Input
                bgColor={"#343434"}
                placeholder="Email"
                type="text"
                value={input.email}
                onChange={(e) => setInput({ ...input, email: e.target.value })}
              />
              <Input
                bgColor={"#343434"}
                placeholder="Password"
                type="password"
                value={input.password}
                onChange={(e) =>
                  setInput({ ...input, password: e.target.value })
                }
              />
            </Flex>
          </ModalBody>

          <ModalFooter>
            <Button
              bgColor={"primary"}
              onClick={(e) => handleSubmit.mutate(e)}
              width={"100%"}
            >
              Register
            </Button>
          </ModalFooter>

          {succesCode && (
            <Fade in={isOpen}>
              <Text textAlign={"center"} color="red.300">
                Berhasil Mendaftar
              </Text>
            </Fade>
          )}
          <Text textAlign={"center"} fontSize={"1.5rem"}>
            Already Have Account, <span>klik Here</span>
          </Text>
        </ModalContent>
      </Modal>
    </>
  );
}
