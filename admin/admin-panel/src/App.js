import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import {

  Login,

} from "./pages";

import "./App.css";


import Dashboard from "./pages/Dashboard";

const App = () => {


  return (
    <div>
      <BrowserRouter>
        <Routes>

          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Login />} />

          {/* pages  */}
          <Route path="*" element={<Dashboard />} />
         
        </Routes>

      </BrowserRouter>
    </div>
  );
};

export default App;
