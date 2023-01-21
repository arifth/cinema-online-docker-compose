import { ChakraProvider } from "@chakra-ui/react";
import React from "react";
import * as ReactDOM from "react-dom/client";
import { MyNewTheme } from "./ExtendTheme";
import App from "./App";
import { BrowserRouter as Router } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "react-query";
import { UserContextProvider } from "./context/userContext";

const rootElement = document.getElementById("root");

const client = new QueryClient();

ReactDOM.createRoot(rootElement).render(
  <React.StrictMode>
    <ChakraProvider resetCSS theme={MyNewTheme}>
      <UserContextProvider>
        <Router>
          <QueryClientProvider client={client}>
            <App />
          </QueryClientProvider>
        </Router>
      </UserContextProvider>
    </ChakraProvider>
  </React.StrictMode>
);
