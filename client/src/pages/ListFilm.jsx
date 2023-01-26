import React, { useContext } from "react";
import NavBar from "../components/NavBar";
import ListFilms from "../components/ListFilms";
import { useQuery } from "react-query";
import { API } from "../config/api";
import { UserContext } from "../context/userContext";
import { Heading } from "@chakra-ui/react";
import { useParams } from "react-router-dom";

export default function UserListFilms() {
  const [state, dispatch] = useContext(UserContext);

  let { data: books } = useQuery("profileCache", async () => {
    let data = await API.get("/books");
    return data.data.data;
  });

  let { id } = useParams();

  // const filtered = books?.filter((book) => {
  //   if (book?.user_id === id) {
  //     return trans;
  //   }
  // });

  return (
    <>
      <NavBar />
      <Heading textAlign="center" fontSize="6xl">
        My List Film
      </Heading>
      <ListFilms />
    </>
  );
}
