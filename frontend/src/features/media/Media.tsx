import { useEffect, useState } from 'react';
import { EventNavigation, Loader, PhotoGallery, YearNavigation } from './components';

function Media() {
    const [years, setYears] = useState([]);
    const [events, setEvents] = useState([]);
    const [photos, setPhotos] = useState([]); // State to hold photo paths
    const [selectedYear, setSelectedYear] = useState('');
    const [selectedEvent, setSelectedEvent] = useState('');
    const [error, setError] = useState('');

    // handle loading for events and photos separately
    const [loadingEvents, setLoadingEvents] = useState(false);
    const [loadingPhotos, setLoadingPhotos] = useState(false);


    useEffect(() => {
        const fetchYears = async () => {
            try {
                const response = await fetch('http://localhost:3000/api/get/photoshoot-years');
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const data = await response.json();
                setYears(data);
            } catch (error) {
                console.error('Error fetching years:', error);
                setError('Failed to load years');
            }
        };

        fetchYears();
    }, []);

    const handleYearSelect = async (year: string) => {
        setSelectedYear(year);
        setEvents([]);
        setLoadingEvents(true);

        try {
            const response = await fetch(`http://localhost:3000/api/get/photoshoot-events/${year}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setEvents(data);
        } catch (err) {
            console.error('Error fetching events:', err);
            setError('Failed to load events');
        } finally {
            setLoadingEvents(false);
        }
    };

    const handleEventSelect = async (event: string) => {
        setLoadingPhotos(true);
        try {
            const response = await fetch(`http://localhost:3000/api/get/photoshoot-photos/${selectedYear}/${event}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setSelectedEvent(event)
            setPhotos(data);
        } catch (err) {
            console.error('Error fetching photos:', err);
            setError('Failed to load photos');
        } finally {
            setLoadingPhotos(false);
        }
    };

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div className="min-h-screen container mx-auto p-4">
            <div className="bg-gray-700 text-white p-6 rounded-md shadow-md">
                <h1 className="text-4xl font-bold mb-2">Media</h1>
                <p className="text-xl">Explore our collection of photos and events</p>
            </div>

            <YearNavigation years={years} selectedYear={selectedYear} onYearSelect={handleYearSelect} />

            {loadingEvents ? <Loader /> : (
                <EventNavigation events={events} selectedEvent={selectedEvent} onEventSelect={handleEventSelect} />
            )}            

             {loadingPhotos ? <Loader /> : (
                <PhotoGallery photos={photos} />
            )}
            
        </div>
    )
}

export default Media;
