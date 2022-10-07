import React from "react";
import Home from "./pages/Home";
import Contact from "./pages/Contact"
import About from './pages/About'
import Navbar from "./components/Navbar";

function App() {
  let pageComponent
  switch(window.location.pathname) {
    case "/":
      pageComponent = <Home/>
      break
    case "/about":
      pageComponent = <About/>
      break
    case "/contact":
      pageComponent = <Contact/>
      break
  }
  return(
    <div className="App">
      <Navbar/>
      {pageComponent}
    </div>
  );
}

export default App;
