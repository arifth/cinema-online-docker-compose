import React from "react";
import NavBar from "../components/NavBar";
import Hero from "../components/Hero";
import ListFilms from "../components/ListFilms";

export default function Home() {
  return (
    <>
      <NavBar />
      <Hero />
      <ListFilms />
    </>
  );
}
