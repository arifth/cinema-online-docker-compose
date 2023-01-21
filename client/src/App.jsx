import { Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import DetailFilm from "./pages/DetailFilm";
import DetailUser from "./pages/DetailUser";
import IncomingTrans from "./pages/IncomingTrans";
import AddFilm from "./pages/AddFilm";
import UserListFilms from "./pages/ListFilm";
import React, { useEffect, useContext } from "react";
import { UserContext } from "./context/userContext";
import { useNavigate } from "react-router-dom";
import { API, setAuthToken } from "./config/api";
import PrivateLogin from "./components/PrivateLogin";
import PrivateLoginAdmin from "./components/PrivateLoginAdmin";
// Masukkan token
if (localStorage.token) {
  setAuthToken(localStorage.token);
}

function App() {
  const navigate = useNavigate();

  const [state, dispatch] = useContext(UserContext);

  useEffect(() => {
    if (localStorage.token) {
      setAuthToken(localStorage.token);
    }
    // Redirect Auth
    if (state.isLogin === false) {
      navigate("/");
    } else {
      if (state.user.role === "ADMIN") {
        navigate("/incomingTrans");
      } else if (state.user.role === "USER") {
        navigate("/");
      }
    }
  }, [state]);

  const checkUser = async () => {
    try {
      const response = await API.get("/check-auth");

      if (response.status === 404) {
        return dispatch({
          type: "AUTH_ERROR",
        });
      }

      // Mendapatkan data user
      let payload = response.data.data;

      // Mengambil token dari localstorage
      payload.token = localStorage.token;
      // Mengirim data ke useContext
      dispatch({
        type: "LOGIN_SUCCESS",
        payload,
      });
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    if (localStorage.token) {
      checkUser();
    }
  }, []);
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/detailFilm/:id" element={<DetailFilm />} />
      <Route
        element={
          <PrivateLogin isLoggedIn={state.isLogin} role={state.user.role} />
        }
      >
        <Route path="/detailUser/:id" element={<DetailUser />} />
        <Route path={"/listFilms/:id"} element={<UserListFilms />} />
      </Route>
      <Route
        element={
          <PrivateLoginAdmin
            isLoggedIn={state.isLogin}
            role={state.user.role}
          />
        }
      >
        <Route path="/incomingTrans" element={<IncomingTrans />} />
        <Route path="/addFilm" element={<AddFilm />} />
      </Route>
    </Routes>
  );
}

export default App;
