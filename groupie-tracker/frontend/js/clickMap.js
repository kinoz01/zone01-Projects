
// Event listener for location links
document.querySelectorAll('.location-link').forEach(link => {
        link.addEventListener('click', function(event) {
            event.preventDefault(); // Prevent default link behavior

            // Get the location from the data attribute
            const location = this.getAttribute('data-location');

            // Generate the Google Maps URL
            const googleMapsUrl = `https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(location)}`;

            // Open the Google Maps URL in a new tab
            window.open(googleMapsUrl, '_blank');
        });
});
