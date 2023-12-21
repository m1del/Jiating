import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { close, logo, menu } from '../assets';
import { navLinks } from '../constants';
import { styles } from '../styles';

function Navbar() {
  const [active, setActive] = React.useState("");
  const [toggle, setToggle] = useState(false);

  const navigate = useNavigate();

  const handleNavClick = (link) => {
    setActive(link.title);
    setToggle(false); // for mobile nav, ensures the menu closes on click

    if (link.type === 'page') {
        navigate(link.path);
    } else { 
      // handle section links
      const targetUrl = `/#${link.id}`;
      
      if (window.location.pathname !== '/') {
          // navigate to home with hash
          navigate(targetUrl, { replace: true });
      } else {
          // if already on home, just update hash and scroll
          window.history.pushState({}, '', targetUrl);
          document.getElementById(link.id)?.scrollIntoView({ behavior: 'smooth' });
      }
    }
  };

  return (
    <nav className={`${styles.paddingX} w-full flex items-center py-5
    fixed top-0 z-20`}>
      <div className='w-full flex justify-between items-center 
      max-w-7xl mx-auto'>
        <Link
          to='/'
          className='flex items-center gap-2'
          onClick={() => {
            setActive("");
            window.scroll(0,0);
          }}
        >
          <img src={logo} alt='logo' draggable='false'
          className='w-9 h-9 object-contain'/>
          <p className='text-red-500 text-[24px] font-bold cursor-pointer flex'>
            JIATING
          </p>
        </Link>

        <ul className='list-none hidden sm:flex flex-row gap-10'>
          {navLinks.map((link) => (
            <li key={link.title}
              className={`${
                active === link.title ?
                "text-white" : "text-secondary"
              } font-medium cursor-pointer text-[16px]`}
            >
              <a onClick={() => handleNavClick(link)}>
                {link.title}
              </a>
            </li>
          ))}
        </ul>

        {/*Mobile nav*/}
        <div className='sm:hidden flex flex-1 justify-end items-center'>
          <img src={toggle ? close : menu} 
            alt='menu' className='w-[28px] h-[28px] cursor-pointer'
            onClick={() => setToggle(!toggle)}/>
            <div className={`${!toggle ? 'hidden' : 'flex'} 
              p-6 absolute top-20 right-0 mx-4 my-2 min-w-[140px]
              z-10 rounded-xl`}>
              <ul className='list-none flex justify-end items-start
              flex-col gap-4'>
                {navLinks.map((link) => (
                  <li
                    key={link.title}
                    className={`${
                      active === link.title ? "text-white" : "text-secondary"
                    } font-medium cursor-pointer text-[16px]`}
                    onClick={() => {
                      setToggle(!toggle);
                      setActive(link.title);
                    }}
                  >
                    <a onClick={() => handleNavClick(link)}>
                      {link.title}
                    </a>
                  </li>
                ))}
              </ul>
          </div>
        </div>
        
      </div>
    </nav>
  )
}

export default Navbar;
