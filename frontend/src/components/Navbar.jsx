import { useRef } from 'react';
import { FaBars, FaTimes } from "react-icons/fa";
import { Link } from 'react-scroll';
import jiatingLogo from '../assets/images/white-black.png';
import "./Navbar.scss";

function Navbar() {

  const navRef = useRef();

  const showNavbar = () => {
		navRef.current.classList.toggle('responsive_nav');
  }

  return (
    <div id='parent'>
    <header>
      <div>
        <a href='/'>
        <img className='navLogo' src={jiatingLogo} alt="Navbar Logo" />
        </a>
      </div>
      
      <nav ref={navRef}>
      <ul id='nav' className="navList">
          <li className='navListItem'>
            <Link onClick={showNavbar} spy={true} smooth={true} offset={-60} duration={500} className='navLink' to='home'>
              HOME
            </Link>
          </li>
          <li className='navListItem'>
            <Link onClick={showNavbar} spy={true} smooth={true} offset={-60} duration={500} className='navLink' to='about'>
              ABOUT
            </Link>
          </li>
          <li className='navListItem'>
            <Link onClick={showNavbar} spy={true} smooth={true} offset={-60} duration={500} className='navLink' to='skills'>
              PHOTOSHOOTS
            </Link>
          </li>
          <li className='navListItem'>
            <Link onClick={showNavbar} spy={true} smooth={true} offset={-60} duration={500} className='navLink' to='contact'>
              CONTACT
            </Link>
          </li>
          <li>
            <button onClick={showNavbar} className='nav-btn nav-close-btn'>
              <FaTimes size={30}/>
            </button>
          </li>
      </ul>
      </nav>

      <button className='nav-btn' onClick={showNavbar}>
        <FaBars size={30}/>
      </button>
    </header>
    </div>
  )
}

export default Navbar