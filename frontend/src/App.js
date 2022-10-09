import React, { useEffect, useState } from "react";
import Home from "./pages/Home";
import Contact from "./pages/Contact"
import About from './pages/About'
import Navbar from "./components/Navbar";

function App() {  
  //Accessing different pages
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
      <section className="page-content">
        {pageComponent}
      </section>
    </div>
  );
}

export default App;
