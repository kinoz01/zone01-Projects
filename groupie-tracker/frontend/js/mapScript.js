function initMap() {
    // Define a default location (e.g., center of the map)
    var defaultLocation = {lat: 40, lng: -30};  // This can be any default point
    var map = new google.maps.Map(document.getElementById('map'), {
        zoom: 3,
        center: defaultLocation,  // Start with the default location
        gestureHandling: 'greedy' // zoom and pan using both scroll gestures
    });

    // Geocoder to convert address names to lat/lng
    var geocoder = new google.maps.Geocoder();
    
    // Add markers for each location in the window.locations array
    locations.forEach(function(location) {
        geocoder.geocode({'address': location}, function(results, status) {
            if (status === 'OK') {
                var marker = new google.maps.Marker({
                    map: map,
                    position: results[0].geometry.location,
                    title: location,
                    icon: {
                        url: "http://maps.google.com/mapfiles/ms/icons/red-dot.png"
                    }
                });
            } else {
                console.error('Geocode was not successful for the following reason: ' + status);
            }
        });
    });
}
