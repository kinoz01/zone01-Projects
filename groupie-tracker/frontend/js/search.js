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
        const artistName = artists[i].getElementsByClassName('name')[0].textContent.toLowerCase();
        const members = artists[i].getElementsByClassName('member');
        const artistID = artists[i].getElementsByTagName('a')[0].getAttribute('href'); // Get artist URL

        const creationDate = artists[i].getElementsByClassName('creationDate')[0].textContent.toLowerCase();
        const firstAlbumDate = artists[i].getElementsByClassName('firstAlbum')[0].textContent.toLowerCase();
        const locations = artists[i].getElementsByClassName('location');

        // Check members
        let memberMatchFound = false; // Track if a member's name starts with the search value
        for (let j = 0; j < members.length; j++) {
            const memberName = members[j].textContent.toLowerCase();

            // Add member to suggestions if it matches (includes)
            if (memberName.includes(searchValue)) {
                suggestions.push({ type: 'member', name: `${artistName} - ${memberName}`, url: artistID });
            }

            // If member's name starts with the search value, mark match found
            if (memberName.startsWith(searchValue)) {
                memberMatchFound = true;
            }
        }

        // Add artist to suggestions if it matches (includes)
        if (artistName.includes(searchValue)) {
            suggestions.push({ type: 'artist', name: artistName, url: artistID });
        }

        // Add creation date with band name to suggestions if it matches (includes)
        if (creationDate.includes(searchValue)) {
            suggestions.push({ 
                type: 'creation date', 
                name: `${artistName} - ${creationDate}`, 
                url: artistID 
            });
        }

        // Add first album date with band name to suggestions if it matches (includes)
        if (firstAlbumDate.includes(searchValue)) {
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
            if (location.includes(searchValue)) {
                suggestions.push({
                    type: 'location',
                    name: `${artistName} - ${location}`,
                    url: artistID
                });
            }
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

    // Filter displayed artists
    for (let i = 0; i < artists.length; i++) {
        const artistName = artists[i].getElementsByClassName('name')[0].textContent.toLowerCase();
        const members = artists[i].getElementsByClassName('member');
        const creationDate = artists[i].getElementsByClassName('creationDate')[0].textContent.toLowerCase();
        const firstAlbumDate = artists[i].getElementsByClassName('firstAlbum')[0].textContent.toLowerCase();
        const locations = artists[i].getElementsByClassName('location');

        let memberStartsWithSearch = false;
        let locationStartsWithSearch = false;

        // Check members for starting match
        for (let j = 0; j < members.length; j++) {
            const memberName = members[j].textContent.toLowerCase();
            if (memberName.startsWith(searchValue)) {
                memberStartsWithSearch = true;
            }
        }

        // Check locations for starting match
        for (let k = 0; k < locations.length; k++) {
            const location = locations[k].textContent.toLowerCase();
            if (location.startsWith(searchValue)) {
                locationStartsWithSearch = true;
            }
        }

        // Show or hide artist based on matches
        if (
            artistName.startsWith(searchValue) || // Starts with search value
            memberStartsWithSearch || // Any member starts with search value
            creationDate.startsWith(searchValue) ||
            firstAlbumDate.startsWith(searchValue) ||
            locationStartsWithSearch // Any location starts with search value
        ) {
            artists[i].style.display = ''; // Show artist if condition matches
        } else {
            artists[i].style.display = 'none'; // Hide artist if no match
        }
    }
}
