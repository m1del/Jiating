import React, { useEffect, useState } from 'react';

const PhotoshootGallery = () => {

    const [photoshoots, setPhotoshoots] = useState([]);

    useEffect(() => {
        // todo: fetch photoshoots from backend
    }, []);

    return (
        <div>
            photoshoots go here.
        </div>
    )
}

export default PhotoshootGallery
