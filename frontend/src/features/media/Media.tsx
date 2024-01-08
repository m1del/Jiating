import React, { useEffect, useState } from 'react';
import { EventNavigation, PhotoGallery, YearNavigation } from './components';

function Media() {
    const [years, setYears] = useState([]);
    const [events, setEvents] = useState([]);
    const [photos, setPhotos] = useState([]); // State to hold photo paths
    const [selectedYear, setSelectedYear] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');


    useEffect(() => {
        const fetchYears = async () => {
            setLoading(true);
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
            } finally {
                setLoading(false);
            }
        };

        fetchYears();
    }, []);

    const handleYearSelect = async (year: string) => {
        setSelectedYear(year);
        setEvents([]);
        setLoading(true);

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
            setLoading(false);
        }
    };

    const handleEventSelect = async (event: string) => {
        setLoading(true);
        try {
            const response = await fetch(`http://localhost:3000/api/get/photoshoot-photos/${selectedYear}/${event}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            console.log(data);
            setPhotos(data); // Update the photos state
        } catch (err) {
            console.error('Error fetching photos:', err);
            setError('Failed to load photos');
        } finally {
            setLoading(false);
        }
    };


    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div className="p-4">
            <YearNavigation years={years} onYearSelect={handleYearSelect} />
            <EventNavigation events={events} onEventSelect={handleEventSelect} />
            <PhotoGallery photos={photos} />
        </div>
    )
}

export default Media;
