// Event listener for location links
document.querySelectorAll('.location-link').forEach(link => {
    link.addEventListener('click', function(event) {
        event.preventDefault(); // Prevent default link behavior

        // Get the location from the data attribute
        const location = this.getAttribute('location-str');

        // Generate the Google Maps URL with a red marker
        const googleMapsUrl = `https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(location)}`;

        // Open the Google Maps URL in a new tab with the marker
        window.open(googleMapsUrl, '_blank');
    });
});
