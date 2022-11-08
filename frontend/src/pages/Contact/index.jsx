import React from 'react';
import { FaFacebook, FaInstagram } from "react-icons/fa";
import { HiOutlineMail } from "react-icons/hi";
import './index.scss';

function Contact() {
  return (
    <div id='contact-container'>

      <div id='text-container'>
        <h1>
          CONTACT
        </h1>
        <h4>
          Connect with us on our social medias!
        </h4>
        <div id='social-media'>
          <a href='https://www.instagram.com/jiating.lion.dragon/' target="_blank" rel="noopener noreferrer"> 
            <FaFacebook className='media-icon' size={60}/>
          </a>
          <a href='https://www.facebook.com/groups/192755019429528' target="_blank" rel="noopener noreferrer"> 
            <FaInstagram className='media-icon' size={60}/>
          </a>
          <HiOutlineMail className='media-icon' size={60}
          onClick={() => window.location 
          = 'mailto:jiating.lion.dragon@gmail.com'}/>
          <button>
            add linktree
          </button>
        </div>
      </div>
    </div>
  )
}

export default Contact