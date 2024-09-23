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

        let membersString = "";
        for (let j = 0; j < members.length; j++) {
            const memberName = members[j].textContent.toLowerCase();
            membersString += memberName + " ";

            // Add member to suggestions if it matches
            if (memberName.includes(searchValue)) {
                suggestions.push({ type: 'member', name: members[j].textContent, url: artistID });
            }
        }

        // Add artist to suggestions if it matches
        if (artistName.includes(searchValue)) {
            suggestions.push({ type: 'artist', name: artists[i].getElementsByClassName('name')[0].textContent, url: artistID });
        }

        // Add creation date with band name to suggestions if it matches
        if (creationDate.includes(searchValue)) {
            suggestions.push({ 
                type: 'creation date', 
                name: `${artists[i].getElementsByClassName('name')[0].textContent} - ${creationDate}`, 
                url: artistID 
            });
        }

        // Add first album date with band name to suggestions if it matches
        if (firstAlbumDate.includes(searchValue)) {
            suggestions.push({ 
                type: 'first album date', 
                name: `${artists[i].getElementsByClassName('name')[0].textContent} - ${firstAlbumDate}`, 
                url: artistID 
            });
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

    // Filter displayed artists based on search
    for (let i = 0; i < artists.length; i++) {
        const artistName = artists[i].getElementsByClassName('name')[0].textContent.toLowerCase();
        const members = artists[i].getElementsByClassName('member');
        const creationDate = artists[i].getElementsByClassName('creationDate')[0].textContent.toLowerCase();
        const firstAlbumDate = artists[i].getElementsByClassName('firstAlbum')[0].textContent.toLowerCase();

        let membersString = "";
        for (let j = 0; j < members.length; j++) {
            membersString += members[j].textContent.toLowerCase() + " ";
        }

        if (
            artistName.includes(searchValue) || 
            membersString.includes(searchValue) ||
            creationDate.includes(searchValue) ||
            firstAlbumDate.includes(searchValue)
        ) {
            artists[i].style.display = '';
        } else {
            artists[i].style.display = 'none';
        }
    }
}
