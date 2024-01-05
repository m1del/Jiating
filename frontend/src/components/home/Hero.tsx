import { bwLogo } from '../../assets';

function Hero() {
  return (
    <div className="flex max-w-[1350px] items-center justify-center space-x-56">
      <img src={bwLogo}></img>
      <div>
        <h2 className="border-cyan border-l-4 pl-4 leading-9">
          <span className="text-cyan text-6xl font-semibold uppercase">
            Jiating
          </span>
          <br />
          <p className="mb-10 text-4xl font-semibold uppercase">
            How we got started
          </p>
        </h2>
        <div className="w-5/6 text-lg text-gray-600">
          <p className="leading-9">
            JiaTing is an all-inclusive cultural and leadership based
            organization that aims to promote traditional/modern lion and dragon
            dance through training and performance.
          </p>
          <br />
          <p className="leading-9">
            Originating from the University of Florida's Chinese American
            Student Association (CASA), JiaTing seeks to support a greater
            understanding of Asian heritage in the Gainesville community and is
            open to all members regardless of previous experience or background.
          </p>
        </div>
      </div>
    </div>
  );
}
export default Hero;
