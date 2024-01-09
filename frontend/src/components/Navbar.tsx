import React, { useState } from 'react';
import { IoClose, IoMenu } from 'react-icons/io5';
import { Link, useNavigate } from 'react-router-dom';
import { logo } from '../assets';
import { navLinks } from '../constants';
import { styles } from '../styles';

function Navbar({ isSticky }) {
  const [active, setActive] = React.useState('');
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
        document
          .getElementById(link.id)
          ?.scrollIntoView({ behavior: 'smooth' });
      }
    }
  };

  return (
    <nav
      className={`${styles.paddingX} items-center ${
        isSticky ? 'bg-black-rgba fixed' : 'relative bg-black'
      } top-0 z-20
    flex h-16 w-full py-5 backdrop-blur-sm`}
    >
      <div
        className="mx-auto flex w-full max-w-[1350px] 
      items-center justify-between"
      >
        <Link
          to="/"
          className="flex items-center gap-2"
          onClick={() => {
            setActive('');
            window.scroll(0, 0);
          }}
        >
          <img
            src={logo}
            alt="logo"
            draggable="false"
            className="h-9 w-9 object-contain"
          />
          <p className="flex cursor-pointer text-3xl font-bold text-white">
            JIATING
          </p>
        </Link>

        <ul className="hidden list-none flex-row gap-10 md:flex">
          {navLinks.map((link) => (
            <li
              key={link.title}
              className={`${
                active === link.title ? 'text-cyan' : 'text-white'
              } hover:text-cyan cursor-pointer text-[16px] font-medium uppercase transition-colors`}
            >
              <a onClick={() => handleNavClick(link)}>{link.title}</a>
            </li>
          ))}
        </ul>

        {/*Mobile nav*/}
        <div className="flex flex-1 items-center justify-end md:hidden">
          {toggle ? (
            <IoClose
              className="h-8 w-8 cursor-pointer text-white"
              onClick={() => setToggle(!toggle)}
            />
          ) : (
            <IoMenu
              className="h-8 w-8 cursor-pointer text-white"
              onClick={() => setToggle(!toggle)}
            />
          )}
          <div
            className={`${!toggle ? 'hidden' : 'flex'} 
              absolute right-0 top-20 z-10 mx-4 my-2 min-w-[140px]
              rounded-xl p-6`}
          >
            <ul
              className="flex list-none flex-col items-start
              justify-end gap-4"
            >
              {navLinks.map((link) => (
                <li
                  key={link.title}
                  className={`${
                    active === link.title ? 'text-white' : 'text-secondary'
                  } cursor-pointer text-[16px] font-medium`}
                  onClick={() => {
                    setToggle(!toggle);
                    setActive(link.title);
                  }}
                >
                  <a onClick={() => handleNavClick(link)}>{link.title}</a>
                </li>
              ))}
            </ul>
          </div>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
