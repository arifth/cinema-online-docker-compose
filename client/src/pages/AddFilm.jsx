import NavBar from "../components/NavBar";
import React, { useState, useRef } from "react";
import {
  FormControl,
  InputGroup,
  InputRightAddon,
  Input,
  Heading,
  VStack,
  HStack,
  Textarea,
  Button,
  Select,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { API } from "../config/api";
import { useQuery } from "react-query";
import Proof from "../assets/proof.svg";

export default function AddFilm() {
  const [input, setInput] = useState({
    file: "",
    title: "",
    category: "",
    price: "",
    link_film: "",
    description: "",
    thumbnnail: "",
  });

  const inputFile = useRef(null);

  let { data: category } = useQuery("categoryCache", async () => {
    const response = await API.get("/categories");
    return response.data.data;
  });

  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    try {
      e.preventDefault();
      const data = new FormData();

      data.append("image", input.file);
      data.append("title", input.title);
      data.append("category_id", input.category);
      data.append("price", input.price);
      data.append("film_url", input.link_film);
      data.append("description", input.description);
      data.append("thumbnail", input.thumbnnail);

      console.log(data);

      const post = await API.post("/film", data);
      navigate("/");
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <>
      <NavBar />
      <VStack width={"80%"} margin={"auto"}>
        <FormControl
          margin={"auto"}
          padding={"3rem"}
          display={"flex"}
          gap={"2rem"}
          flexDirection={"column"}
        >
          <Heading alignSelf={"flex-start"}>Add Movie</Heading>
          <HStack>
            <Input
              bgColor={"#343434"}
              width={"80%"}
              placeholder="Title of Movies"
              type="text"
              value={input.title}
              onChange={(e) => setInput({ ...input, title: e.target.value })}
            />
            <InputGroup width={"20%"}>
              <Input
                style={{ display: "none" }}
                type="file"
                name="file"
                ref={inputFile}
                onChange={(e) =>
                  setInput({ ...input, file: e.target.files[0] })
                }
              />
              <Button
                onClick={() => inputFile.current.click()}
                bgColor={"primary"}
                width={"100%"}
                padding={"1.5rem"}
              >
                Attach Payment
                <img src={Proof} alt="proof" style={{ marginLeft: "1rem" }} />
              </Button>
            </InputGroup>
          </HStack>
          <Select
            placeholder="Select Category"
            onChange={(e) => setInput({ ...input, category: e.target.value })}
          >
            {category?.map((cat) => {
              console.log(cat?.ID);
              return <option value={cat?.ID}>{cat?.name}</option>;
            })}
          </Select>
          <Input
            bgColor={"#343434"}
            placeholder="price"
            type="text"
            value={input.price}
            onChange={(e) => setInput({ ...input, price: e.target.value })}
          />
          <Input
            bgColor={"#343434"}
            placeholder="Link Film"
            type="text"
            value={input.link_film}
            onChange={(e) => setInput({ ...input, link_film: e.target.value })}
          />
          <Textarea
            bgColor={"#343434"}
            placeholder="description"
            type="text"
            value={input.description}
            onChange={(e) =>
              setInput({ ...input, description: e.target.value })
            }
          />
          <Button
            bgColor={"primary"}
            alignSelf={"flex-end"}
            onClick={handleSubmit}
          >
            Add Film
          </Button>
        </FormControl>
      </VStack>
    </>
  );
}
