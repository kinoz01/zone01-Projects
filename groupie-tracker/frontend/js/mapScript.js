function initMap() {
    // Define a default location (e.g., center of the map)
    var defaultLocation = {lat: 20.6843, lng: -88.5678};  // This can be any default point
    var map = new google.maps.Map(document.getElementById('map'), {
        zoom: 2,
        center: defaultLocation,  // Start with the default location
        gestureHandling: 'greedy'
    });

    // Geocoder to convert address names to lat/lng
    var geocoder = new google.maps.Geocoder();
    
    // Add markers for each location in the window.locations array
    window.locations.forEach(function(location) {
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
