import React from "react"
import instagramSVG from '../icons/instagram.svg'
import searchSVG from '../icons/magnifying-glass-solid.svg'
import jiatingLogo from '../images/white-black.png'
import "./Navbar.css"

export default function Navbar() {
  return (
    <div className="Navbar"> 
        <div className="navLeft">

          <a href="https://www.instagram.com/jiating.lion.dragon/">
            <img className="ig-icon" src={instagramSVG}/> {/*Add ALT source for image*/}
          </a>

        </div>
        <div className="navCenter">
          <ul className="navList">
            <a className="navLink" href="/">
              <li className="navListItem">HOME</li>
            </a>
            <a className="navLink" href="/about">
              <li className="navListItem">ABOUT</li>
            </a>
            <a className="navLink" href='/contact'>
              <li className="navListItem">CONTACT</li>
            </a>
            <a className="navLink" href='/photoshoots'>
              <li className="navListItem">PHOTOSHOOTS</li>
            </a>
          </ul>
        </div>
        
        <div className="navRight">
          <img className="jiating-logo" src={jiatingLogo}/> {/*Add ALT source for image */}
          <img className="search-icon" src={searchSVG}/> {/*Add ALT source for image*/}
        </div>
    </div>
  )
}
