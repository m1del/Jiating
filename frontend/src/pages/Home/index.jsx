import React from 'react';
import { TypeAnimation } from 'react-type-animation';
import titeImage2 from '../../assets/images/dragon-dance.png';
import titleImage from '../../assets/images/line-of-lions.png';
import './index.scss';

function Home() {
  return (
    <div id='home-container'>

      <div id='text-container'>
      <h1>
        JIATING
      </h1>
      <TypeAnimation
        // Same String at the start will only be typed once, initially
        sequence={[
        'We are lions',
        1000,
        'We are dragons',
        1000,
        'We are family',
        1000,
        'We are JIATING.',
        1000,
        ]}
        speed={50} // Custom Speed from 1-99 - Default Speed: 40
        style={{ fontSize: '6rem' }}
        wrapper="span" // Animation will be rendered as a <span>
        repeat={Infinity} // Repeat this Animation Sequence infinitely
      />
      </div>
      <img id='titleImage2' src={titeImage2} alt='Dragon Dance'/>
      <img id='titleImage' src={titleImage} alt='Line of Lion Dancers'/>

      
    </div>
  )
}

export default Home