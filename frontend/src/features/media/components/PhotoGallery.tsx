import { useEffect, useState } from 'react';

const PhotoModal = ({ photo, onClose }) => {
  useEffect(() => {
    const handleKeyUp = (e) => {
      if (e.key === 'Escape') {
        onClose();
      }
    };

    window.addEventListener('keyup', handleKeyUp);

    return () => {
      window.removeEventListener('keyup', handleKeyUp);
    };
  }, [onClose]);

  return (
    <div className="z-10 fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center p-4">
      <div className="relative max-w-full max-h-full overflow-auto">
        <button 
          onClick={onClose} 
          className="absolute top-2 right-2 bg-black bg-opacity-50 text-white text-5xl 
          font-bold hover:bg-opacity-70 rounded-full shadow-lg h-10 w-10 flex items-center justify-center"        >
          &times;
        </button>
        <img src={photo} alt="Enlarged view" className="max-h-[80vh] w-auto" />
      </div>
    </div>
  );
};

const PhotoGallery = ({ photos }) => {
    const [selectedPhoto, setSelectedPhoto] = useState(null);

    const handlePhotoClick = (photo, isLoaded) => {
        if (isLoaded) {
            setSelectedPhoto(photo);
        }
    };

    const handleCloseModal = () => {
        setSelectedPhoto(null);
    };

    return (
        <div>
            {selectedPhoto && <PhotoModal photo={selectedPhoto} onClose={handleCloseModal} />}
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 mt-4">
                {photos.map((photo, index) => (
                    <div key={index} className="overflow-hidden rounded-lg shadow-lg cursor-pointer">
                        <img 
                            src={photo} 
                            alt={`Photo ${index + 1}`} 
                            className="object-cover w-full h-full transition-transform duration-300 ease-in-out hover:scale-110 hover:shadow-xl" 
                            onClick={(e) => handlePhotoClick(photo, e.target.complete)}
                            loading="lazy"
                        />
                    </div>
                ))}
            </div>
        </div>
    );
};

export default PhotoGallery;
