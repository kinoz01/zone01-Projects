// filter.js

// Initialize the filter toggle button event listener
document.getElementById('filterToggleBtn').addEventListener('click', toggleFilter);

window.onload = function() {
    generateLocationOptions();
    applyFilters(); // Apply filters on page load if necessary
};

function generateLocationOptions() {
    const artistList = document.getElementById('artistList');
    const artists = artistList.getElementsByTagName('li');
    const locationSet = new Set();
    
    for (let i = 0; i < artists.length; i++) {
        const locations = artists[i].getElementsByClassName('location');
        for (let k = 0; k < locations.length; k++) {
            const location = locations[k].textContent.trim();
            if (location) {
                // Split location into parts and add each part to the set
                const locationParts = location.split(',').map(part => part.trim().toLowerCase());
                // Add full location
                locationSet.add(location.toLowerCase());
                // Add individual parts
                locationParts.forEach(part => locationSet.add(part));
            }
        }
    }
    
    const locationSelect = document.getElementById('locationSelect');
    
    // Convert set to array and sort
    const sortedLocations = Array.from(locationSet).sort();
    
    sortedLocations.forEach(location => {
        const option = document.createElement('option');
        option.value = location;
        option.textContent = location;
        locationSelect.appendChild(option);
    });
}

function parseDateString(dateString) {
    // Expected format: "DD-MM-YYYY"
    const parts = dateString.split('-');
    if (parts.length !== 3) return null;
    const day = parseInt(parts[0], 10);
    const month = parseInt(parts[1], 10) - 1; // Months are zero-based in JS Date
    const year = parseInt(parts[2], 10);
    return new Date(year, month, day);
}

function applyFilters() {
    const artistList = document.getElementById('artistList');
    const artists = artistList.getElementsByTagName('li');

    // Get filter values
    const creationDateFrom = document.getElementById('creationDateFrom').value;
    const creationDateTo = document.getElementById('creationDateTo').value;
    const firstAlbumFrom = document.getElementById('firstAlbumFrom').value;
    const firstAlbumTo = document.getElementById('firstAlbumTo').value;
    
    const selectedLocation = document.getElementById('locationSelect').value.toLowerCase();

    // Get selected number of members
    const membersCheckboxes = document.querySelectorAll('.membersCheckbox:checked');
    const selectedMembers = Array.from(membersCheckboxes).map(cb => parseInt(cb.value));

    // Convert filter dates to Date objects
    const firstAlbumFromDate = firstAlbumFrom ? new Date(firstAlbumFrom) : null;
    const firstAlbumToDate = firstAlbumTo ? new Date(firstAlbumTo) : null;

    // Loop through artists and apply filters
    for (let i = 0; i < artists.length; i++) {
        const artist = artists[i];
        const members = artist.getElementsByClassName('member');
        const creationDate = artist.getElementsByClassName('creationDate')[0].textContent;
        const firstAlbumDate = artist.getElementsByClassName('firstAlbum')[0].textContent;
        const locations = artist.getElementsByClassName('location');

        // Convert dates and numbers
        const creationYear = parseInt(creationDate);
        const firstAlbum = parseDateString(firstAlbumDate);
        const numberOfMembers = members.length;

        // Initialize display flag
        let displayArtist = true;

        // Apply creation date filter
        if (creationDateFrom && creationYear < parseInt(creationDateFrom)) {
            displayArtist = false;
        }
        if (creationDateTo && creationYear > parseInt(creationDateTo)) {
            displayArtist = false;
        }

        // Apply first album date filter
        if (firstAlbumFromDate && firstAlbum && firstAlbum < firstAlbumFromDate) {
            displayArtist = false;
        }
        if (firstAlbumToDate && firstAlbum && firstAlbum > firstAlbumToDate) {
            displayArtist = false;
        }

        // Apply number of members filter
        if (selectedMembers.length > 0 && !selectedMembers.includes(numberOfMembers)) {
            displayArtist = false;
        }

        // Apply locations filter
        if (selectedLocation && selectedLocation !== 'all') {
            let locationMatch = false;
            for (let k = 0; k < locations.length; k++) {
                const locationText = locations[k].textContent.toLowerCase();
                if (locationText.includes(selectedLocation)) {
                    locationMatch = true;
                    break;
                } else {
                    // Split location into parts and check each part
                    const locationParts = locationText.split(',').map(part => part.trim());
                    for (let part of locationParts) {
                        if (part.includes(selectedLocation)) {
                            locationMatch = true;
                            break;
                        }
                    }
                    if (locationMatch) break;
                }
            }
            if (!locationMatch) {
                displayArtist = false;
            }
        }

        // Set data attribute instead of changing display directly
        if (displayArtist) {
            artist.dataset.filtersMatch = 'true';
        } else {
            artist.dataset.filtersMatch = 'false';
        }
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

function toggleFilter() {
    const filterContainer = document.getElementById('filterContainer');
    filterContainer.classList.toggle('show');
}

function clearFilters() {
    // Clear creation date inputs
    document.getElementById('creationDateFrom').value = '';
    document.getElementById('creationDateTo').value = '';

    // Clear first album date inputs
    document.getElementById('firstAlbumFrom').value = '';
    document.getElementById('firstAlbumTo').value = '';

    // Uncheck all members checkboxes
    const membersCheckboxes = document.querySelectorAll('.membersCheckbox');
    membersCheckboxes.forEach(checkbox => {
        checkbox.checked = false;
    });

    // Reset location select to 'all'
    document.getElementById('locationSelect').value = 'all';

    // Re-apply filters to show all artists
    applyFilters();
}
