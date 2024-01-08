import { useEffect, useRef, useState } from 'react';
import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import { Event, Footer, Navbar } from './components';
import { AuthProvider } from './context/AuthContext';
import AdminDashboard from './features/admin/AdminDashboard';
import ProtectedRoute from './features/authentication/components/ProtectedRoute';
import Media from './features/media/Media';
import { Contact, CreateEvent, Events, Home } from './pages';
import { styles } from './styles';

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
    <AuthProvider>
      <Router>
        <div ref={ref} className={`'h-16' absolute bg-black`} />
        {!intersecting && <div className="h-16 bg-black" />}
        <div>
          <Navbar isSticky={!intersecting} />
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/media" element={<Media />} />
            <Route path="/events" element={<Events />} />
            <Route path="/contact" element={<Contact />} />
            <Route path="/event" element={<Event />} />
            <Route
              path="/admin/eventform"
              element={
                <ProtectedRoute>
                  <CreateEvent />
                </ProtectedRoute>
              }
            />
            <Route
              path="/admin/dashboard"
              element={
                <ProtectedRoute>
                  <AdminDashboard />
                </ProtectedRoute>
              }
            />
            <Route path="/admin/get-event" element={<Event />}></Route>
          </Routes>
          <Footer />
        </div>
      </Router>
    </AuthProvider>
  );
}

export default App;
