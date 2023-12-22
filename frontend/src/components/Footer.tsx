import { FaDiscord, FaFacebookF, FaInstagram, FaTiktok, FaYoutube } from "react-icons/fa";
import { MdOutlineEmail, MdOutlineLocationOn } from "react-icons/md";
import { teamLogoBW } from '../assets';
import GoogleLoginButton from "../features/authentication/components/GoogleLoginButton";
import { styles } from '../styles';

const Footer = () => (
  <div className={`${styles.paddingX} bg-black border-t w-full py-5 px-5`}>

    <div className='mx-auto flex flex-col sm:flex-row w-full max-w-[1350px] items-center justify-between'>
      <div className="logo-section mb-4 sm:mb-0">
        <img src={teamLogoBW} alt="Logo" className="hidden sm:block sm:h-24 sm:w-24"/>
      </div>

      <div className="text-white contact-info flex flex-col items-start mb-4 sm:mb-0">
        <div className="email-section mb-2 flex">
          <MdOutlineEmail className="text-white text-2xl mr-2"/>
          <span>jiating.lion.dragon@gmail.com</span>
        </div>
        <div className="address-section flex">
          <MdOutlineLocationOn className="text-white text-2xl mr-2"/>
          <span>University of Florida, Gainesville, FL</span>
        </div>
      </div>

      <GoogleLoginButton/>

      <div className="social-media flex">
        <a href="https://www.tiktok.com/@jiatingliondragon" className="mr-2" target="_blank" rel="noopener noreferrer">
          <FaTiktok className="text-white text-2xl"/>
        </a>
        <a href="https://www.youtube.com/channel/UCH-SNRsw6u9_549hUCxM9Zw" className="mr-2" target="_blank" rel="noopener noreferrer">
          <FaYoutube className="text-white text-2xl"/>
        </a>
        <a href="https://www.facebook.com/jiatingliondragon" className="mr-2" target="_blank" rel="noopener noreferrer">
          <FaFacebookF className="text-white text-2xl"/>
        </a>
        <a href="https://www.instagram.com/jiating.lion.dragon/" className="mr-2" target="_blank" rel="noopener noreferrer">
          <FaInstagram className="text-white text-2xl"/>
        </a>
        <a href="https://discord.gg/AHnNzWZTax" target="_blank" rel="noopener noreferrer">
          <FaDiscord className="text-white text-2xl"/>
        </a>
      </div>
    </div>
  </div>
);

export default Footer;
