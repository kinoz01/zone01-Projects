function replaceSpecialChars(location) {
    // Replace hyphens with commas and underscores with spaces
    let processedLocation = location.replace(/[_]/g, ' ').replace(/[-]/g, ', ');
    // Capitalize the first letter of each word
    return processedLocation.replace(/\b\w/g, (char) => char.toUpperCase());
}

// Event listener for location links
document.querySelectorAll('.location-link').forEach(link => {
        link.addEventListener('click', function(event) {
            event.preventDefault(); // Prevent default link behavior

            // Get the location from the data attribute
            const location = this.getAttribute('data-location');
            const processedLocation = replaceSpecialChars(location); // Clean up the location string

            // Generate the Google Maps URL
            const googleMapsUrl = `https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(processedLocation)}`;

            // Open the Google Maps URL in a new tab
            window.open(googleMapsUrl, '_blank');
        });
});
