function EventTile({ size, gridPos }) {
  const imgPath = '../lion_1.png';
  const convertPath = `\'${imgPath}\'`;
  let spefClass = '';

  if (size == 'large') {
    spefClass = `row-span-2 col-span-2 ${gridPos} hover:-translate-y-4 duration-500`;
  } else if (size == 'long') {
    spefClass = `col-span-2 ${gridPos}`;
  } else {
    spefClass = `relative group ${gridPos}`;
  }
  const eventClass = `overflow-hidden rounded-2xl h-full w-full bg-[url('../lion_1.png')] bg-contain bg-no-repeat ${spefClass} text-white hover:cursor-pointer`;

  if (size == 'large') {
    return (
      <div className={eventClass}>
        <div className="flex h-full flex-col justify-end">
          <div className="bg-black-rgba px-4 py-6">
            <h2 className="pb-5 text-4xl">CNY 2023</h2>
            <p className="pb-2 text-lg">
              Lorem ipsum dolor, sit amet consectetur adipisicing elit.
              Consequatur eaque accusantium tempore deleniti odit cum obcaecati.
            </p>
            <p>Written by: Michael Shi</p>
            <p>05/15/2002</p>
          </div>
        </div>
      </div>
    );
  } else if (size == 'medium') {
    return (
      <div className={eventClass}>
        <div className="absolute bottom-0 left-0 right-0 h-0 w-full overflow-hidden bg-black-rgba transition-all duration-500 ease-in group-hover:h-full">
          <div className="flex h-full flex-col items-center justify-center px-5">
            <h2 className="pb-4 text-2xl">CNY 2023</h2>
            <p className="pb-3 text-center text-lg">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus
              deleniti, quo eius similique ipsa ab labore.
            </p>
            <p>Written by: Michael Shi</p>
            <p>05/15/2002</p>
          </div>
        </div>
      </div>
    );
  }
  return <div className={eventClass}></div>;
}

export default EventTile;

// header, imgPath, author, desc, size
