import {
  FaDiscord,
  FaFacebookF,
  FaInstagram,
  FaTiktok,
  FaYoutube,
} from 'react-icons/fa';
import { MdOutlineEmail, MdOutlineLocationOn } from 'react-icons/md';
import { teamLogoBW } from '../assets';
import { styles } from '../styles';

const Footer = () => (
  <div className={`${styles.paddingX} w-full border-t bg-black px-5 py-5`}>
    <div className="mx-auto flex w-full max-w-[1350px] flex-col items-center justify-between sm:flex-row">
      <div className="logo-section mb-4 sm:mb-0">
        <a href="http://localhost:3000/auth/google">
          <img
            src={teamLogoBW}
            alt="Logo"
            className="hidden sm:block sm:h-24 sm:w-24"
          />
        </a>
      </div>

      <div className="contact-info mb-4 flex flex-col items-start text-white sm:mb-0">
        <div className="email-section mb-2 flex">
          <MdOutlineEmail className="mr-2 text-2xl text-white" />
          <span>jiating.lion.dragon@gmail.com</span>
        </div>
        <div className="address-section flex">
          <MdOutlineLocationOn className="mr-2 text-2xl text-white" />
          <span>University of Florida, Gainesville, FL</span>
        </div>
      </div>

      <div className="social-media flex">
        <a
          href="https://www.tiktok.com/@jiatingliondragon"
          className="mr-2"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaTiktok className="text-2xl text-white" />
        </a>
        <a
          href="https://www.youtube.com/channel/UCH-SNRsw6u9_549hUCxM9Zw"
          className="mr-2"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaYoutube className="text-2xl text-white" />
        </a>
        <a
          href="https://www.facebook.com/jiatingliondragon"
          className="mr-2"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaFacebookF className="text-2xl text-white" />
        </a>
        <a
          href="https://www.instagram.com/jiating.lion.dragon/"
          className="mr-2"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaInstagram className="text-2xl text-white" />
        </a>
        <a
          href="https://discord.gg/AHnNzWZTax"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaDiscord className="text-2xl text-white" />
        </a>
      </div>
    </div>
  </div>
);

export default Footer;
