import {
  Box,
  Flex,
  Text,
  Button,
  Stack,
  useColorModeValue,
  useDisclosure,
  Image,
  Avatar,
  Popover,
  PopoverTrigger,
  PopoverContent,
  HStack,
  VStack,
  PopoverBody,
  Divider,
} from "@chakra-ui/react";
import ModalLogin from "./ModalLogin";
import ModalRegister from "./ModalRegister";
import React, { useState, useContext } from "react";
import { Link, useNavigate } from "react-router-dom";
import { UserContext } from "../context/userContext";
import User from "../assets/userIcon.svg";
import Film from "../assets/film.svg";
import Logout from "../assets/logout.svg";

export default function NavBar() {
  const Navigate = useNavigate();
  const [state, dispatch] = useContext(UserContext);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [opened, setOPened] = useState(false);
  const handleOpen = () => {
    setOPened(true);
  };
  const closed = () => {
    setOPened(false);
  };

  const handleLogout = () => {
    dispatch({ type: "LOGOUT" });
  };

  return (
    <Box>
      <ModalLogin isOpen={isOpen} onClose={onClose} />
      <ModalRegister isOpen={opened} onClose={closed} />

      <Flex
        bg="transparent"
        color={useColorModeValue("gray.600", "white")}
        minH={"60px"}
        py="1.5rem"
        px="3rem"
        align={"center"}
      >
        <Flex flex={{ base: 1 }} justify={{ base: "center", md: "start" }}>
          <a href="/">
            <Image src="/Icon.svg" />
          </a>
        </Flex>

        <Stack
          flex={{ base: 1, md: 0 }}
          justify={"flex-end"}
          direction={"row"}
          spacing={6}
        >
          {!state.isLogin ? (
            <>
              <Button variant="secondary" onClick={onOpen}>
                Login
              </Button>
              <Button variant="primary" onClick={handleOpen}>
                Register
              </Button>
            </>
          ) : (
            <>
              <Popover>
                <PopoverTrigger>
                  <Avatar />
                </PopoverTrigger>
                <PopoverContent style={{ border: "none" }}>
                  <PopoverBody
                    bgColor={"blackAlpha.800"}
                    style={{ border: "none" }}
                    padding={"1rem"}
                  >
                    {state.user.role === "USER" ? (
                      <VStack>
                        <Box marginBottom={"1rem"} width={"100%"}>
                          <Link to={`/detailUser/${state?.user.id_user}`}>
                            <HStack
                              width={"100%"}
                              justifyContent={"flex-start"}
                            >
                              <Image src={User} alt="" marginRight="2rem" />
                              <Text fontSize={"1.5rem"} color={"white"}>
                                Profile
                              </Text>
                            </HStack>
                          </Link>
                        </Box>

                        <Box marginBottom={"1rem"} width={"100%"}>
                          <Link to={`/listFilms/${state?.user.id_user}`}>
                            <HStack
                              width={"100%"}
                              justifyContent={"flex-start"}
                            >
                              <Image src={Film} alt="" marginRight="2rem" />
                              <Text fontSize={"1.5rem"} color={"white"}>
                                List Film
                              </Text>
                            </HStack>
                          </Link>
                        </Box>
                        <Divider />
                        <HStack width={"100%"} justifyContent={"flex-start"}>
                          <Box
                            marginBottom={"1rem"}
                            width={"100%"}
                            display={"flex"}
                            justifyContent={"space-between"}
                          >
                            <Button
                              bgColor={"blackAlpha.900"}
                              onClick={() => handleLogout()}
                              width={"100%"}
                            >
                              <Image src={User} alt="" marginRight="2rem" />
                              <Text fontSize={"1.5rem"} color={"white"}>
                                Logout
                              </Text>
                            </Button>
                          </Box>
                        </HStack>
                      </VStack>
                    ) : (
                      <VStack>
                        <Box marginBottom={"1rem"}>
                          <Link to={`/addFilm`}>
                            <HStack
                              width={"100%"}
                              justifyContent={"flex-start"}
                            >
                              <Image src={Film} alt="" marginRight="2rem" />
                              <Text fontSize={"1.5rem"} color={"white"}>
                                Add Film
                              </Text>
                            </HStack>
                          </Link>
                        </Box>
                        <HStack width={"100%"} justifyContent={"flex-start"}>
                          <Button
                            bgColor={"blackAlpha.900"}
                            onClick={() => handleLogout()}
                            width={"100%"}
                          >
                            <Image src={User} alt="" marginRight="2rem" />
                            <Text fontSize={"1.5rem"} color={"white"}>
                              Logout
                            </Text>
                          </Button>
                        </HStack>
                      </VStack>
                    )}
                  </PopoverBody>
                </PopoverContent>
              </Popover>
            </>
          )}
        </Stack>
      </Flex>
    </Box>
  );
}
