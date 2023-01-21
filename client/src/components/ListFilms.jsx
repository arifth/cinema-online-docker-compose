import { Heading, Flex, AspectRatio, Image, Box } from "@chakra-ui/react";
import React from "react";
import { useNavigate } from "react-router-dom";
import Film from "../data/Film.json";
import { Link } from "react-router-dom";
import { API } from "../config/api";
import { useQuery } from "react-query";

export default function ListFilms() {
  const navigate = useNavigate();
  let { data: films } = useQuery("filmCache", async () => {
    let data = await API.get("/films");
    return data.data.data;
  });

  return (
    <>
      <Heading w={"100%"} mt={"4rem"} paddingLeft={"15rem"}>
        List of Films
      </Heading>
      <Flex align={"center"} justify={"center"} p="1rem" gap={"2rem"}>
        {Film.map((film) => {
          return (
            <Link key={film.id} to={`/detailFilm/${film.id}`}>
              <Box height={"250px"} width={"190px"} borderRadius={"10px"}>
                <Image
                  draggable="false"
                  src={`./${film.image}`}
                  borderRadius={"10px"}
                  height={"100%"}
                  width={"100%"}
                  objectFit={"cover"}
                />
              </Box>
            </Link>
          );
        })}
        {films?.map((film) => {
          return (
            <Link key={film?.ID} to={`/detailFilm/${film?.ID}`}>
              <Box height={"198px"} width={"143px"} borderRadius={"10px"}>
                <Image
                  borderRadius={"10px"}
                  height={"100%"}
                  width={"100%"}
                  objectFit={"cover"}
                  draggable="false"
                  src={film?.image}
                />
              </Box>
            </Link>
          );
        })}
      </Flex>
    </>
  );
}
