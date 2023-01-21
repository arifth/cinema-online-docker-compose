import {
  Flex,
  Image,
  VStack,
  Heading,
  Button,
  HStack,
  Text,
  Box,
  Alert,
} from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import { useQuery } from "react-query";
import { API } from "../config/api";
import NavBar from "../components/NavBar";
import { Link } from "react-router-dom";

export default function DetailUser() {
  const { id } = useParams();

  let { data: profile } = useQuery("profileCache", async () => {
    let data = await API.get(`/user/${id}`);
    return data.data.data;
  });

  console.log(profile);

  return (
    <>
      <NavBar />
      <Flex
        align={"center"}
        justify={"center"}
        w="90%"
        margin={"auto"}
        padding={"3rem"}
      >
        <HStack width={"100%"}>
          <VStack>
            <Link to={"/addFilm"}>
              <Heading mb={4}>My Profile</Heading>
            </Link>

            <Box height={"300px"} width={"250px"}>
              <Image
                src="/Profile.png"
                height={"100%"}
                width={"100%"}
                objectFit={"cover"}
              />
            </Box>
          </VStack>
          <VStack alignItems={"flex-start"} style={{ marginLeft: "2rem" }}>
            <Text color="primary" fontSize={"2rem"}>
              Full Name
            </Text>
            <Text color="muted" fontSize={"2rem"}>
              {profile?.name}
            </Text>
            <Text color="primary" fontSize={"2rem"}>
              Email
            </Text>
            <Text color="muted" fontSize={"2rem"}>
              {profile?.email}
            </Text>
            <Text color="primary" fontSize={"2rem"}>
              Phone
            </Text>
            <Text color="muted" fontSize={"2rem"}>
              082222222222
            </Text>
          </VStack>
          <VStack width={"50%"} style={{ marginLeft: "10rem" }}>
            <Link to={"/incomingTrans"}>
              <Heading style={{ marginBottom: "3rem" }}>
                History Transaction
              </Heading>
            </Link>
            <Box
              width={"100%"}
              bgColor={"hsla(335,63%,49%,0.4)"}
              height={"300px"}
              padding={"3rem"}
              borderRadius={"1rem"}
            >
              <HStack>
                <VStack
                  display={"flex"}
                  alignItems={"flex-start"}
                  justifyContent={"center"}
                  width={"100%"}
                  gap={"2rem"}
                >
                  <Text fontSize={"2rem"}>Tom & jerry </Text>
                  <Text>
                    <span style={{ fontWeight: "bolder" }}>Saturday </span> 12
                    April 2021
                  </Text>
                  <Text color={"primary"}>Total : RP.150000</Text>
                </VStack>
                <VStack alignItems={"flex-end"} width={"100%"}>
                  <Alert
                    marginTop={"120px"}
                    bgColor={"green"}
                    width={"30%"}
                    style={{ borderRadius: "10px" }}
                  >
                    Finished
                  </Alert>
                </VStack>
              </HStack>
            </Box>
          </VStack>
        </HStack>
      </Flex>
    </>
  );
}
