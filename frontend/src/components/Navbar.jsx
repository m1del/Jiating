import React from "react"
import "./Navbar.css"
import instagramSVG from '../icons/instagram.svg'
import searchSVG from '../icons/magnifying-glass-solid.svg'
import jiatingLogo from '../images/white-black.png'

export default function Navbar() {
  return (
    <div className="Navbar"> 
        <div className="navLeft">

          <img className="ig-icon" src={instagramSVG}/> {/*Add ALT source for image*/}

        </div>
        <div className="navCenter">
          <ul className="navList">
            <li className="navListItem">HOME</li>
            <li className="navListItem">ABOUT</li>
            <li className="navListItem">CONTACT</li>
          </ul>
        </div>
        
        <div className="navRight">
          <img className="jiating-logo" src={jiatingLogo}/> {/*Add ALT source for image */}
          <img className="search-icon" src={searchSVG}/> {/*Add ALT source for image*/}
        </div>
    </div>
  )
}
