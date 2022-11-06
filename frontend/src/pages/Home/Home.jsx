import { Parallax, ParallaxLayer } from '@react-spring/parallax'
import React from 'react'
import left from '../../assets/images/ab1.jpg'
import right from '../../assets/images/ab2.jpg'
import lionsG14 from '../../assets/images/background1.jpg'
import errbody from '../../assets/images/errbody.jpg'
import bg2 from '../../assets/images/lamp-bg.jpg'
import liongif1 from '../../assets/images/lion-gif1.gif'
import liongif2 from '../../assets/images/lion-gif2.gif'
import logo from '../../assets/images/sketchy.png'
import './Home.scss'

function Home() {
  return (
    <div id='home-container'>

      <h1>
        HELLO FROM HOME
      </h1>
      {/* <Parallax pages={3}>

        <ParallaxLayer
        offset={0}
        speed={1}
        factor={1}
        style={{
          backgroundImage: `url(${lionsG14})`,
          backgroundSize: 'cover',
          
        }}
        />

        <ParallaxLayer
        offset={1.1}
        speed={1}
        factor={1}
        style={{
          backgroundImage: `url(${bg2})`,
          backgroundSize: 'cover',
        }}/>

        <ParallaxLayer
        offset={.8}
        speed={1}
        factor={1}
        style={{
          backgroundImage: `url(${logo})`,
          backgroundSize: 'cover',
        }}/>


        <ParallaxLayer 
        offset={.95}
        speed={.5}
        factor={1}>
          <img className='gif1' src={liongif1} alt="lion dancer gif"/>
        </ParallaxLayer>

        <ParallaxLayer 
        offset={.95}
        speed={.5}
        factor={1}>
          <img className='gif2' src={liongif2} alt="lion dancer gif"/>
        </ParallaxLayer>

        <ParallaxLayer
        offset={0}
        speed={0.4}
        sticky={{start:0.1, end: 0.5}}>
          <div className='jiating'>
            <h1 className='title'>JIATING</h1>
            <p1 className='subtitle'>Lion & Dragon</p1>
          </div>
        </ParallaxLayer>

        <ParallaxLayer
        offset={1}
        speed={0.6}
        factor={1}>
          <div className='jiating-desc'>
            <p className='bozo-descriptionslmao'>
              Jiating is a non-profit organization based in Gainesville, Florida. Consisting of university students and young professionals, we strive to keep the traditional art of lion dance thriving in the melting pot
              of the United States. Join us as in our expeditions to find the hidden dragon (in your butt).
            </p>
            <div className='row'>
              <img className='abt-img' src={errbody} alt="jiating group pic" width="33%" height="100%"/>
              <img className='abt-img' src={left} alt="jiating group pic" width="33%" height="100%"/>
              <img className='abt-img' src={right} alt="jiating group pic"width="33%" height="100%"/>
            </div>
          </div>
        </ParallaxLayer>

        

      </Parallax> */}
    </div>
  )
}

export default Home