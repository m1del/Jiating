import React from 'react'
import lionsG14 from '../images/background1.jpg'
import liongif1 from '../images/lion-gif1.gif'
import liongif2 from '../images/lion-gif2.gif'
import lionjpg from '../images/lion-jpg.jpg'
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
        >
        </ParallaxLayer>

        <ParallaxLayer 
        offset={.9}
        speed={.7}
        factor={1}>
          <img className='gif1' src={liongif1}/>
        </ParallaxLayer>

        <ParallaxLayer 
        offset={.9}
        speed={.7}
        factor={1}>
          <img className='gif2' src={liongif2}/>
        </ParallaxLayer>

        <ParallaxLayer
        offset={0}
        speed={0.4}
        sticky={{start:0.1, end: 0.9}}>
          <div className='jiating'>
            <h1 className='title'>JIATING</h1>
            <p1 className='subtitle'>Lion & Dragon</p1>
          </div>
        </ParallaxLayer>

        

      </Parallax>
    </div>
  )
}

export default Home