import Navbar from "./components/Navbar";
import About from './pages/About';
import Contact from "./pages/Contact";
import Home from "./pages/Home";
import Photoshoots from "./pages/Photoshoots";

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
    case "/photoshoots":
      pageComponent = <Photoshoots/>
      break
    default:
      pageComponent = <Home/>
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
