// src/router/Router.jsx
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Register from "../components/pages/Register";

const Router = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/register" element={<Register />} />

        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
