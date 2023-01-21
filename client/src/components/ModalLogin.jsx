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
import { useMutation } from "react-query";
import React, { useState, useContext } from "react";
import { UserContext } from "../context/userContext";
import { API } from "../config/api";
import { useNavigate } from "react-router-dom";

export default function ModalLogin({ isOpen, onClose }) {
  const [state, dispatch] = useContext(UserContext);

  const navigate = useNavigate();

  const [input, setInput] = useState({
    email: "",
    password: "",
  });
  const [succesCode, setSuccesCode] = useState(false);

  // const handleSubmit = useMutation(async (e) => {
  //   e.preventDefault();
  //   try {
  //     const config = {
  //       headers: {
  //         "Content-Type": "application/json",
  //       },
  //     };

  //     // insert user data to database
  //     const response = await API.post("/login", input);
  //     if (response.status === 200) {
  //       setSuccesCode(true);
  //     }
  //     onClose();
  //     return response;
  //   } catch (error) {
  //     console.log(error);
  //   }
  //   onClose();
  // });

  const handleSubmit = useMutation(async (e) => {
    e.preventDefault();
    try {
      const config = {
        headers: {
          "Content-Type": "application/json",
        },
      };
      // insert user data to database
      const response = await API.post("/login", input, config);
      console.log(response);

      if (response.data.code === 200 && response.data.data.role === "USER") {
        const payload = response.data.data;
        console.log(state);
        onClose();
        return dispatch({
          type: "LOGIN_SUCCESS",
          payload: payload,
        });
      } else if (
        response.data.code === 200 &&
        response.data.data.role === "ADMIN"
      ) {
        const payload = response.data.data;
        console.log(response);
        navigate("/incomingTrans");
        onClose();
        return dispatch({
          type: "LOGIN_SUCCESS",
          payload: payload,
        });
      }
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
            Login
          </ModalHeader>
          <ModalBody>
            <Flex flexDirection={"column"} gap={"1rem"} marginBottom={"1rem"}>
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
              Login
            </Button>
          </ModalFooter>
          <Text textAlign={"center"} fontSize={"1.5rem"}>
            Dont Have an Account, <span>klik Here</span>
          </Text>
        </ModalContent>
      </Modal>
    </>
  );
}
