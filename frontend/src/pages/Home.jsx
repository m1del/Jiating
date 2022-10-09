import React from 'react'
import lionsG14 from '../images/background1.jpg'
import liongif1 from '../images/lion-gif1.gif'
import liongif2 from '../images/lion-gif2.gif'
import errbody from '../images/errbody.jpg'
import left from '../images/ab1.jpg'
import right from '../images/ab2.jpg'
import bg2 from '../images/lamp-bg.jpg'
import logo from '..//images/sketchy.png'
import {Parallax, ParallaxLayer} from '@react-spring/parallax'
import "./styles.css"

function Home() {
  return (
    <div className='Home'>
      <Parallax pages={4}>

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
          <img className='gif1' src={liongif1}/>
        </ParallaxLayer>

        <ParallaxLayer 
        offset={.95}
        speed={.5}
        factor={1}>
          <img className='gif2' src={liongif2}/>
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
              <img className='abt-img' src={errbody}/>
              <img className='abt-img' src={left}/>
              <img className='abt-img' src={right}/>
            </div>
          </div>
        </ParallaxLayer>

        

      </Parallax>
    </div>
  )
}

export default Home