import React from 'react';
import './App.css';
import {BrowserRouter, Routes, Route} from 'react-router-dom';
import Login from './pages/login';
import Publish from './pages/publish';
import SignUp from './pages/sign-up';
import Home from './pages/home';
import Model3DDashboard from './pages/my-model3ds';
import ErrorPage from './pages/error-page';

const App = () => {
  
  return (
    <div className="App bg-black">
      <BrowserRouter>
        <Routes>
          <Route path="/home" element={<Home />} />
          <Route exact path="/my-model3ds" element={<Model3DDashboard />} />
          <Route exact path="/login" element={<Login />} />
          <Route exact path="/publish" element={<Publish />} />
          <Route exact path="/sign-up" element={<SignUp />} />
          <Route exact={true} path="*" element={<ErrorPage />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
};

export default App;
