import React from "react";
import { Navigate, Outlet } from "react-router-dom";

export default function PrivateLogin({ isLoggedIn, role }) {
  return (
    <>{!isLoggedIn && role === "USER" ? <Navigate to="/" /> : <Outlet />}</>
  );
}
