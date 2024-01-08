const PhotoGallery = ({ photos }) => {
    return (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 mt-4">
            {photos.map((photo, index) => (
                <div key={index} className="overflow-hidden rounded-lg shadow-lg">
                    <img src={photo} alt={`Photo ${index + 1}`} className="object-cover w-full h-full transition-transform duration-300 ease-in-out hover:scale-105" />
                </div>
            ))}
        </div>
    );
};

export default PhotoGallery;
