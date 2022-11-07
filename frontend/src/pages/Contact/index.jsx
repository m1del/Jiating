import React from 'react';
import { FaFacebook, FaInstagram } from "react-icons/fa";
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
        <div id='media-icons'>
          <FaFacebook/>
          <FaInstagram/>
        </div>
      </div>
    </div>
  )
}

export default Contact