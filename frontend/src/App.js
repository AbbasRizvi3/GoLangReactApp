import './App.css';
import Home from './Components/Home';
import Login from './Components/Login';
import Signup from './Components/Signup';
import { Routes, Route } from "react-router-dom";

function App() {
  return (
    <>
     <Routes>
         <Route path="/" element={<Home/>} />
        <Route path="/login" element={<Login/>} />
        <Route path="/signup" element={<Signup/>} />
      </Routes>
       </>
  );
}

export default App;
