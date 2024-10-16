const baseUrl = 'https://cdn.jsdelivr.net/gh/akabab/superhero-api@0.3.0/api/all.json';
let heroes = [];

let currentPage = 1;
let pageSize = 20; // Set default page size to 20
let sortColumn = null; // The currently sorted column
let sortOrder = 'asc'; // 'asc' or 'desc'
let searchQuery = '';

// Set the page size dropdown to 20
document.getElementById('page-size').value = pageSize;

function loadData(mydata) {
    heroes = mydata;
    updateDisplay();
}

const fetchData = () => {
    fetch(baseUrl)
        .then((response) => response.json())
        .then(loadData)
        .catch((error) => {
            console.error('Error fetching data:', error);
        });
};

const displayHeroes = (heroes) => {
    const tableBody = document.getElementById('hero-data');
    tableBody.innerHTML = ''
    heroes.forEach(hero => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td><img src="${hero.images.xs}" alt="${hero.name}"></td>
            <td>${hero.name}</td>
            <td>${hero.biography.fullName}</td>
            <td>intelligence: ${hero.powerstats.intelligence}, strength: ${hero.powerstats.strength}, speed: ${hero.powerstats.speed}, durability: ${hero.powerstats.durability}, power: ${hero.powerstats.power}, combat: ${hero.powerstats.combat}</td>
            <td>${hero.appearance.race}</td>
            <td>${hero.appearance.gender}</td>
            <td>${hero.appearance.height[0]}<br>${hero.appearance.height[1]}</td>
            <td>${hero.appearance.weight[0]}<br>${hero.appearance.weight[1]}</td>
            <td>${hero.biography.placeOfBirth}</td>
            <td>${hero.biography.alignment}</td>
        `;
        tableBody.appendChild(row);
    });
};

fetchData();

document.getElementById('search').addEventListener('input', function () {
    searchQuery = this.value.toLowerCase();
    currentPage = 1;
    updateDisplay();
});

document.getElementById('page-size').addEventListener('change', function () {
    const query = this.value;
    if (query !== '') {
        if (query === 'all') {
            pageSize = heroes.length;
        } else {
            pageSize = parseInt(query);
        }
        currentPage = 1;
        updateDisplay();
    }
});

const paginateData = (data, page, size) => {
    const start = (page - 1) * size;
    const end = start + size;
    return data.slice(start, end);
};

const getFilteredAndSortedHeroes = () => {
    let filteredHeroes = heroes.filter(hero => hero.name.toLowerCase().includes(searchQuery));
    if (sortColumn) {
        filteredHeroes.sort((a, b) => {
            let valA = getValueByField(a, sortColumn);
            let valB = getValueByField(b, sortColumn);

            let missingA = isMissingValue(valA);
            let missingB = isMissingValue(valB);

            if (missingA && missingB) return 0;
            if (missingA) return 1;
            if (missingB) return -1;

            if (typeof valA === 'string') valA = valA.toLowerCase();
            if (typeof valB === 'string') valB = valB.toLowerCase();

            if (valA < valB) return sortOrder === 'asc' ? -1 : 1;
            if (valA > valB) return sortOrder === 'asc' ? 1 : -1;
            return 0;
        });
    }
    return filteredHeroes;
};

function isMissingValue(val) {
    return val === null || val === undefined || val === '' || val === '-' || val === 'N/A' || val === 0;
}

function parseWeightHeight(valueStr, WeightOrHeight) {
    if (!valueStr || valueStr === '-' || valueStr === '0' || valueStr === 'N/A') {
        return null
    }
    valueStr = valueStr.trim().toLowerCase().replace(/,/g, '')

    if (WeightOrHeight === 'weight' && (valueStr.includes('tons') || valueStr.includes('ton'))) {
        let value = parseFloat(valueStr)
        return isNaN(value) ? null : value * 1000
    }
    if (WeightOrHeight === 'height' && valueStr.includes('meter')) {
        let value = parseFloat(valueStr)
        return isNaN(value) ? null : value * 100
    }
    let value = parseFloat(valueStr)
    return isNaN(value) ? null : value
}

function getValueByField(hero, field) {
    switch (field) {
        case 'name':
            return hero.name || '';
        case 'fullName':
            return hero.biography.fullName || '';
        case 'powerstatsSum':
            let sum = 0;
            for (let stat in hero.powerstats) {
                let statValue = hero.powerstats[stat];
                let value = parseInt(statValue);
                if (isNaN(value) || value === 0) {
                    return null;
                } else {
                    sum += value;
                }
            }
            return sum;
        case 'race':
            return hero.appearance.race || '';
        case 'gender':
            return hero.appearance.gender || '';
        case 'height':
            let heightCmStr = hero.appearance.height[1];
            let heightCm = parseWeightHeight(heightCmStr, 'height');
            return heightCm;
        case 'weight':
            let weightKgStr = hero.appearance.weight[1];
            let weightKg = parseWeightHeight(weightKgStr, 'weight');
            return weightKg;
        case 'placeOfBirth':
            return hero.biography.placeOfBirth || '';
        case 'alignment':
            return hero.biography.alignment || '';
        default:
            return '';
    }
}

function updateDisplay() {
    const filteredAndSortedHeroes = getFilteredAndSortedHeroes();
    const paginatedHeroes = paginateData(filteredAndSortedHeroes, currentPage, pageSize);
    displayHeroes(paginatedHeroes);
    renderPagination(filteredAndSortedHeroes);
    updateSortIcons();
}

const renderPagination = (heroesList) => {
    const paginationContainer = document.getElementById('pagination');
    paginationContainer.innerHTML = '';
    const totalPages = Math.ceil(heroesList.length / pageSize);
    for (let i = 1; i <= totalPages; i++) {
        const link = document.createElement('a');
        link.innerText = i;
        if (i === currentPage) {
            link.classList.add('active');
        }
        
        link.addEventListener('click', function () {
            currentPage = i;
            const allLinks = paginationContainer.querySelectorAll('a');
            allLinks.forEach(link => link.classList.remove('active'));
            this.classList.add('active');
            updateDisplay();
        });
        paginationContainer.appendChild(link);
    }
};

function updateSortIcons() {
    const headerButtons = document.querySelectorAll('table th button');
    headerButtons.forEach(button => {
        const dataField = button.getAttribute('data-field');
        const th = button.parentNode;
        th.classList.remove('asc', 'desc');
        if (dataField === sortColumn) {
            th.classList.add(sortOrder);
        }
    });
}

// Add event listeners to header buttons for sorting
const headerButtons = document.querySelectorAll('table th button');
headerButtons.forEach(button => {
    button.addEventListener('click', function () {
        const dataField = this.getAttribute('data-field');
        if (sortColumn === dataField) {
            sortOrder = sortOrder === 'asc' ? 'desc' : 'asc';
        } else {
            sortColumn = dataField;
            sortOrder = 'asc';
        }
        currentPage = 1;
        updateDisplay();
    });
});
