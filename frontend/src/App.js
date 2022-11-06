import './App.scss';
import Navbar from "./components/Navbar/Navbar";
import About from './pages/About/About';
import Contact from "./pages/Contact/Contact";
import Home from "./pages/Home/Home";
import Photoshoots from "./pages/Photoshoots/Photoshoots";

function App() {  
  
  return(
    <div className="App">
      <Navbar/>

      <section id='home'>
        <Home/>
      </section>

      <section id='about'>
        <About/>
      </section>

      <section id='contact'>
        <Contact/>
      </section>
      
    </div>
  );
}

export default App;
