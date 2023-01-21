import { extendTheme } from "@chakra-ui/react";
import { ButtonStyles as Button } from "../extendstyles/ButtonStyles";

export const MyNewTheme = extendTheme({
  styles: {
    global: {
      body: {
        bg: "hsla(0, 0%, 0%, 0.9)",
        color: "white",
      },
    },
  },
  colors: {
    primary: "hsla(335,63%,49%,1)",
    secondary: "#153e75",
    muted: "hsla(0,0%,38%,1)",
  },
  components: {
    Button,
  },
});
