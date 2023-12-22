import { FaDiscord, FaFacebookF, FaInstagram, FaTiktok, FaYoutube } from "react-icons/fa";
import { MdOutlineEmail, MdOutlineLocationOn } from "react-icons/md";
import { teamLogoBW } from '../assets';
import { styles } from '../styles';


const Footer = () => (
  <div className={`${styles.paddingX} bg-black
  items-center border-t w- full py-5 px-5`}>
    <div className='mx-auto flex w-full max-w-[1350px] items-center
    justify-between'>
        <div className="logo-section">
        <img src={teamLogoBW} alt="Logo" className="h-24 w-24"/>
        </div>

        <div className="text-white contact-info flex flex-col items-center">
        <div className="email-section mb-2 flex">
            <MdOutlineEmail className="text-white text-2xl mr-2"/>
            <span>jiating.lion.dragon@gmail.com</span>
        </div>
        <div className="address-section flex">
            <MdOutlineLocationOn className="text-white text-2xl mr-2"/>
            <span>University of Florida, Gainesville, FL</span>
        </div>
        </div>

        <div className="social-media flex items-center">
        <a href="https://tiktok.com/yourprofile" className="mr-2" target="_blank" rel="noopener noreferrer">
            <FaTiktok className="text-white text-2xl"/>
        </a>
        <a href="https://youtube.com/yourchannel" className="mr-2" target="_blank" rel="noopener noreferrer">
            <FaYoutube className="text-white text-2xl"/>
        </a>
        <a href="https://facebook.com/yourpage" className="mr-2" target="_blank" rel="noopener noreferrer">
            <FaFacebookF className="text-white text-2xl"/>
        </a>
        <a href="https://instagram.com/yourhandle" className="mr-2" target="_blank" rel="noopener noreferrer">
            <FaInstagram className="text-white text-2xl"/>
        </a>
        <a href="https://discord.com/yourserver" target="_blank" rel="noopener noreferrer">
            <FaDiscord className="text-white text-2xl"/>
        </a>
        </div>
    </div>

    
  </div>
);

export default Footer;
