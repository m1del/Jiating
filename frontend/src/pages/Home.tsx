import { Button } from '../components';
import Hero from '../components/home/Hero';
import HomeHeading from '../components/home/HomeHeading';
import { loginGoogleUser } from '../services/authService';

function Home() {
  return (
    <div className="mx-auto max-w-[1920px]">
      <div className="flex flex-col items-center">
        <HomeHeading />

        <Button
          buttonText="Login with Google"
          onClick={() => loginGoogleUser()}
        />

        <Hero />
      </div>
    </div>
  );
}

export default Home;
