import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import Navbar from './components/Navbar';
import { Contact, Events, Home } from './pages';
import { styles } from './styles';
import { useEffect, useState, useRef } from 'react';

function App() {
  const [intersecting, setIntersecting] = useState(true);
  const ref = useRef(null);

  useEffect(() => {
    const observer = new IntersectionObserver(([entry]) => {
      setIntersecting(entry.isIntersecting);
    });
    if (ref.current) observer.observe(ref.current);
  }, []);

  return (
    <Router>
      <div ref={ref} className={`'h-16' absolute bg-black`} />
      {!intersecting && <div className="h-16 bg-black" />}
      <div>
        <Navbar isSticky={!intersecting} />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/events" element={<Events />} />
          <Route path="/contact" element={<Contact />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
