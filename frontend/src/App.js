import './App.scss';
import Navbar from "./components/Navbar";
import About from './pages/About';
import Contact from "./pages/Contact";
import Home from "./pages/Home";
import Photoshoots from "./pages/Photoshoots";

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
