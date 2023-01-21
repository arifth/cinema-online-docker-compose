import NavBar from "../components/NavBar";
import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Text,
  TableContainer,
  Heading,
  VStack,
  Popover,
  PopoverTrigger,
  PopoverContent,
  PopoverBody,
  PopoverArrow,
  PopoverCloseButton,
  Button,
  Box,
  HStack,
  Divider,
} from "@chakra-ui/react";
import { useQuery } from "react-query";
import { API } from "../config/api";

export default function IncomingTrans() {
  let { data: books } = useQuery("profileCache", async () => {
    let data = await API.get("/books");
    console.log(data.data.data);
    return data.data.data;
  });

  return (
    <>
      <NavBar />
      <VStack width={"80%"} margin={"auto"}>
        <Heading alignSelf={"start"}>Incoming Transaction</Heading>
        <TableContainer margin={"auto"} width={"100%"} marginTop={"5rem"}>
          <Table variant="striped" colorScheme="grey">
            <Thead>
              <Tr>
                <Th>Number</Th>
                <Th>Users</Th>
                <Th>Bukti Transfer</Th>
                <Th>Film</Th>
                <Th>Number Account</Th>
                <Th>Status Payment</Th>
                <Th>Action</Th>
              </Tr>
            </Thead>
            <Tbody>
              {books?.map((book) => {
                return (
                  <Tr>
                    <Td>{book?.ID}</Td>
                    <Td>{book?.user.full_name}</Td>
                    <Td width={"20px"} overflow={"scroll"}>
                      <a href={book?.transfer_proof}></a>
                    </Td>
                    <Td>{book?.ID}</Td>
                    <Td>{book?.account_number}</Td>
                    <Td>{book?.status}</Td>
                    <Td>
                      <Popover>
                        <PopoverTrigger>
                          <Button variant={"primary"}>Trigger</Button>
                        </PopoverTrigger>
                        <PopoverContent width={"200px"}>
                          <PopoverArrow />
                          <PopoverCloseButton />
                          <PopoverBody>
                            <Box marginBottom={"1rem"} width={"100%"}>
                              <HStack
                                width={"100%"}
                                justifyContent={"flex-start"}
                                marginBottom={"1rem"}
                              >
                                <Text fontSize={"1.5rem"} color={"green"}>
                                  Approve
                                </Text>
                              </HStack>
                              <Divider />
                              <HStack
                                width={"100%"}
                                justifyContent={"flex-start"}
                              >
                                <Text fontSize={"1.5rem"} color={"red"}>
                                  Cancel
                                </Text>
                              </HStack>
                            </Box>
                          </PopoverBody>
                        </PopoverContent>
                      </Popover>
                    </Td>
                  </Tr>
                );
              })}
            </Tbody>
          </Table>
        </TableContainer>
      </VStack>
    </>
  );
}
