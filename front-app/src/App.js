import './App.css';
import React, { lazy, Suspense } from 'react';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'

const Navigation = lazy(()=>import('./Navigation/Navigation'));
const Registration = lazy(()=>import('./Registration/Registration'));
const Login = lazy(()=>import('./Login/Login'));
const Chat = lazy(()=>import('./Chat/Chat'));

function App() {
  return (
    <>
    <Navigation />
    <Router>
      <Suspense fallback="Loading...">
      <Routes>
        <Route path="/registration" exact element={<Registration/>}></Route>
        <Route path="/login" exact element={<Login/>}></Route>
        <Route path="/chat" exact element={<Chat/>}></Route>
      </Routes> 
    </Suspense>
    </Router> 
    </>
  );
}

export default App;
