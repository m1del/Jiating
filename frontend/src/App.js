import React, { useEffect, useState } from "react";
import Home from "./pages/Home";
import Contact from "./pages/Contact"
import About from './pages/About'
import Navbar from "./components/Navbar";

function App() {

  //Parallax
  const [offsetY,setOffsetY] = useState(0);
  const handleScroll = () => setOffsetY(window.pageYOffset);

  useEffect(() => {
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  
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
