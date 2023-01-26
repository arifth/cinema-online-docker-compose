import {
  Modal,
  Button,
  ModalOverlay,
  ModalBody,
  Center,
  ModalContent,
  ModalHeader,
  Input,
  Flex,
  Text,
  HStack,
  Divider,
  Spinner,
} from "@chakra-ui/react";
import React, { useState, useEffect, useContext, useRef } from "react";
import Proof from "../assets/proof.svg";
import { useMutation } from "react-query";
import { API } from "../config/api";
import { UserContext } from "../context/userContext";

export default function ModalPayment({ isOpen, onClose, title, price }) {
  const [input, setInput] = useState({
    accNumber: "",
    file: "",
    title: title || "",
    price: price || "",
  });

  const inputFile = useRef(null);

  const [state, dispatch] = useContext(UserContext);

  const handleSubmit = useMutation(async () => {
    try {
      const config = {
        method: "POST",
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "multipart/form-data",
        },
      };
      setInput({
        ...input,
        title: title,
      });

      const date = new Date().toString();
      console.log(date);
      let data = new FormData();
      data.append("price", input.price);
      data.append("image", input.file);
      data.append("account_number", input.accNumber);
      data.append("order_date", date);
      data.append("film_id", "1");
      data.append("user_id", state.user.id_user);

      const response = await API.post("/book", data, config);

      const token = response.data.data.token;
      console.log(token);

      window.snap.pay(token, {
        onSuccess: function (result) {
          console.log(result);
        },
        onPending: function (result) {
          console.log(result);
        },
        onError: function (result) {
          navigate("/payment");
        },
      });
      return response.data;
    } catch (error) {
      console.log(error);
    }
  });

  useEffect(() => {
    //change this to the script source you want to load, for example this is snap.js sandbox env
    const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
    //change this according to your client-key
    const myMidtransClientKey = "SB-Mid-client-X1qte313qM27Xpbq";

    let scriptTag = document.createElement("script");
    scriptTag.src = midtransScriptUrl;
    // optional if you want to set script attribute
    scriptTag.setAttribute("data-client-key", myMidtransClientKey);
    document.body.appendChild(scriptTag);
    return () => {
      document.body.removeChild(scriptTag);
    };
  }, []);

  return (
    <>
      <Modal isOpen={isOpen} onClose={onClose} isCentered>
        <ModalOverlay backdropFilter="blur(10px) hue-rotate(0deg)" />
        <ModalContent
          bgColor={"hsla(0, 0%, 0%, 0.7)"}
          padding={"1rem"}
          width={"800px"}
        >
          <ModalHeader fontSize={"1.5rem"}>
            Cinema <span style={{ color: "hsla(335,63%,49%,1)" }}>Online</span>{" "}
            : 0981312323
          </ModalHeader>
          <ModalBody>
            <Flex flexDirection={"column"} marginBottom={"1rem"} gap={"10px"}>
              <Text fontSize={"2.5rem"}> {title}</Text>
              <Divider />
              <Text fontSize={"2rem"}>
                Total :
                <span style={{ color: "#CD2E71" }}>
                  Rp{price?.toLocaleString()}
                </span>
              </Text>
              <Input
                bgColor={"#343434"}
                placeholder="Account Number"
                type="text"
                value={input.accNumber}
                onChange={(e) =>
                  setInput({ ...input, accNumber: e.target.value })
                }
              />
              <HStack>
                <Input
                  style={{ display: "none" }}
                  type={"file"}
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
                <Text width={"60%"} color="muted" fontSize={"1rem"}>
                  *transfers can be made to holyways accounts
                </Text>
              </HStack>
              <Button
                bgColor={"primary"}
                onClick={() => handleSubmit.mutate()}
                width={"100%"}
              >
                {handleSubmit.isLoading ? (
                  <>
                    <Center>
                      <p> Loading cuy </p>
                      <Spinner
                        marginLeft={"1rem"}
                        thickness="10px"
                        speed="0.65s"
                        emptyColor="gray.800"
                        color="grey.500"
                        size="l"
                      />
                    </Center>
                  </>
                ) : (
                  <p> Pay</p>
                )}
              </Button>
            </Flex>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
}
