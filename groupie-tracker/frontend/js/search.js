
let currentFocus = -1; // To track the current suggestion focus

function filterArtists() {
    const searchValue = document.getElementById('searchBar').value.toLowerCase();
    const artistList = document.getElementById('artistList');
    const suggestionsList = document.getElementById('suggestionsList');
    const artists = artistList.getElementsByTagName('li');

    suggestionsList.innerHTML = ""; // Clear suggestions
    currentFocus = -1; // Reset focus

    // Collect and display suggestions
    let suggestions = [];

    for (let i = 0; i < artists.length; i++) {
        const artist = artists[i];
        const artistName = artist.getElementsByClassName('name')[0].textContent.toLowerCase();
        const members = artist.getElementsByClassName('member');
        const artistID = artist.getElementsByTagName('a')[0].getAttribute('href'); // Get artist URL

        const creationDate = artist.getElementsByClassName('creationDate')[0].textContent.toLowerCase();
        const firstAlbumDate = artist.getElementsByClassName('firstAlbum')[0].textContent.toLowerCase();
        const locations = artist.getElementsByClassName('location');

        // Check members
        for (let j = 0; j < members.length; j++) {
            const memberName = members[j].textContent.toLowerCase();

            // Add member to suggestions if it matches (includes)
            if (memberName.includes(searchValue) && searchValue !== "") {
                suggestions.push({ type: 'member', name: `${artistName} - ${memberName}`, url: artistID });
            }
        }

        // Add artist to suggestions if it matches (includes)
        if (artistName.includes(searchValue) && searchValue !== "") {
            suggestions.push({ type: 'artist', name: artistName, url: artistID });
        }

        // Add creation date with band name to suggestions if it matches (includes)
        if (creationDate.includes(searchValue) && searchValue !== "") {
            suggestions.push({ 
                type: 'creation date', 
                name: `${artistName} - ${creationDate}`, 
                url: artistID 
            });
        }

        // Add first album date with band name to suggestions if it matches (includes)
        if (firstAlbumDate.includes(searchValue) && searchValue !== "") {
            suggestions.push({ 
                type: 'first album date', 
                name: `${artistName} - ${firstAlbumDate}`, 
                url: artistID 
            });
        }

        // Check each location
        for (let k = 0; k < locations.length; k++) {
            const location = locations[k].textContent.toLowerCase();

            // Add location to suggestions if it matches (includes)
            if (location.includes(searchValue) && searchValue !== "") {
                suggestions.push({
                    type: 'location',
                    name: `${artistName} - ${location}`,
                    url: artistID
                });
            }
        }

        // Determine if artist matches the search value
        let memberStartsWithSearch = false;
        let locationStartsWithSearch = false;

        // Check members for starting match
        for (let j = 0; j < members.length; j++) {
            const memberName = members[j].textContent.toLowerCase();
            if (memberName.startsWith(searchValue)) {
                memberStartsWithSearch = true;
                break;
            }
        }

        // Check locations for starting match
        for (let k = 0; k < locations.length; k++) {
            const location = locations[k].textContent.toLowerCase();
            if (location.startsWith(searchValue)) {
                locationStartsWithSearch = true;
                break;
            }
        }

        // Set data attribute instead of changing display directly
        if (
            artistName.startsWith(searchValue) || // Starts with search value
            memberStartsWithSearch || // Any member starts with search value
            creationDate.startsWith(searchValue) ||
            firstAlbumDate.startsWith(searchValue) ||
            locationStartsWithSearch // Any location starts with search value
        ) {
            artist.dataset.searchMatch = 'true';
        } else {
            artist.dataset.searchMatch = 'false';
        }
    }

    // Render suggestions
    if (searchValue) {
        suggestions.forEach(suggestion => {
            const li = document.createElement('li');
            li.textContent = `${suggestion.name} (${suggestion.type})`;
            li.setAttribute('data-url', suggestion.url); // Store URL in a data attribute
            li.addEventListener('click', function() {
                window.location.href = suggestion.url; // Navigate to the URL when clicked
            });
            suggestionsList.appendChild(li);
        });
    }

    // Update visibility based on both search and filters
    updateArtistVisibility();
}

// Function to update artist visibility based on search and filter matches
function updateArtistVisibility() {
    const artists = document.getElementById('artistList').getElementsByTagName('li');

    for (let i = 0; i < artists.length; i++) {
        const artist = artists[i];
        const searchMatch = artist.dataset.searchMatch !== 'false'; // Default to true if undefined
        const filtersMatch = artist.dataset.filtersMatch !== 'false'; // Default to true if undefined

        if (searchMatch && filtersMatch) {
            artist.style.display = ''; // Show artist
        } else {
            artist.style.display = 'none'; // Hide artist
        }
    }
}

// Event listener to hide suggestions when clicking outside of the search bar
document.addEventListener('click', function (e) {
    const searchBar = document.getElementById('searchBar');
    const suggestionsList = document.getElementById('suggestionsList');
    
    // Hide suggestions if clicked outside of search bar and suggestions list
    if (e.target !== searchBar && e.target.parentNode !== suggestionsList) {
        suggestionsList.innerHTML = ""; // Clear the suggestions list
    }
});
