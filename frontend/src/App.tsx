import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import Navbar from './components/Navbar';
import { Contact, Events, Home } from './pages';
import { styles } from './styles';


function App() {
  return (
      <Router>
          <Navbar />
          <div className="content-wrapper" style={{ paddingTop: '60px' }}>
            <Routes>
              <Route path='/' element={<Home />} />
              <Route path='/events' element={<Events />} />
              <Route path='/contact' element={<Contact />} />
            </Routes>
          </div>
          
      </Router>
  )
}


export default App
