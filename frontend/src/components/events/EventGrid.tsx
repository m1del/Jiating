import EventTile from './EventTile';

function EventGrid() {
  return (
    <div className="my-48 grid h-[56.25rem] w-[75rem] grid-cols-4 grid-rows-3 gap-1.5">
      <EventTile
        size="large"
        gridPos="row-start-1 row-end-3 col-start-1 col-end-3"
      />
      <EventTile
        size="long"
        gridPos="row-start-1 row-end-2 col-start-3 col-end-5"
      />
      <EventTile
        size="medium"
        gridPos="row-start-2 row-end-3 col-start-3 col-end-4"
      />
      <EventTile
        size="medium"
        gridPos="row-start-2 row-end-3 col-start-4 col-end-5"
      />
      <EventTile
        size="medium"
        gridPos="row-start-3 row-end-4 col-start-1 col-end-2"
      />
      <EventTile
        size="medium"
        gridPos="row-start-3 row-end-4 col-start-2 col-end-3"
      />
      <EventTile
        size="long"
        gridPos="row-start-3 row-end-4 col-start-3 col-end-5"
      />
    </div>
  );
}

export default EventGrid;
