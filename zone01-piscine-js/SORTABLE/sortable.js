let heroes = [];
let filteredHeroes = [];
let currentPage = 1;
let pageSize = 20; // Default page size
let sortOrder = { column: 'name', ascending: true };

// Fetch data on page load
window.onload = () => {
  fetch('https://rawcdn.githack.com/akabab/superhero-api/0.2.0/api/all.json')
    .then(response => response.json())
    .then(data => {
      heroes = data;
      filteredHeroes = data;

      // Set page size to default (20) and select it in the dropdown
      pageSize = 20;
      document.getElementById('page-size').value = pageSize;

      renderTable(); // Initial table render
    });
};

// Render table rows based on the current page and page size
function renderTable() {
  const tableBody = document.getElementById('hero-body');
  tableBody.innerHTML = ''; // Clear previous content

  const startIndex = (currentPage - 1) * pageSize;
  const endIndex =
    pageSize === 'all' ? filteredHeroes.length : Math.min(startIndex + pageSize, filteredHeroes.length);

  const currentHeroes = filteredHeroes.slice(startIndex, endIndex);

  // Using a document fragment for better performance during large DOM updates
  const fragment = document.createDocumentFragment();

  currentHeroes.forEach(hero => {
    const row = document.createElement('tr');

    // Prepare height and weight values with both units
    const heightImperial = displayValue(hero.appearance.height[0]);
    const heightMetric = displayValue(hero.appearance.height[1]);
    const weightImperial = displayValue(hero.appearance.weight[0]);
    const weightMetric = displayValue(hero.appearance.weight[1]);

    // Combine both units, each on a separate line
    const heightDisplay = `${heightImperial}<br>${heightMetric}`;
    const weightDisplay = `${weightImperial}<br>${weightMetric}`;

    row.innerHTML = `
      <td>${displayValue(hero.name)}</td>
      <td>${displayValue(hero.biography.fullName)}</td>
      <td>${displayValue(hero.appearance.race)}</td>
      <td>${displayValue(hero.appearance.gender)}</td>
      <td>${heightDisplay}</td>
      <td>${weightDisplay}</td>
      <td>${displayValue(hero.biography.placeOfBirth)}</td>
      <td>${displayValue(hero.biography.alignment)}</td>
      <td>${displayValue(hero.powerstats.intelligence)}</td>
      <td>${displayValue(hero.powerstats.strength)}</td>
      <td>${displayValue(hero.powerstats.speed)}</td>
      <td>${displayValue(hero.powerstats.durability)}</td>
      <td>${displayValue(hero.powerstats.power)}</td>
      <td>${displayValue(hero.powerstats.combat)}</td>
      <td><img src="${hero.images.xs}" alt="${hero.name}"></td>
    `;
    fragment.appendChild(row);
  });

  tableBody.appendChild(fragment); // Append all rows at once for better performance
  renderPagination();
}

// Helper function to display values, keeping 'null' as 'null' and handling other missing values
function displayValue(value) {
  if (value === null || value === 'null') return 'null';
  if (value === undefined || value === '' || value === '-' || value === '0') return '-';
  return value;
}

// Pagination
function renderPagination() {
  const pagination = document.getElementById('pagination');
  pagination.innerHTML = ''; // Clear previous pagination

  const totalPages =
    pageSize === 'all' ? 1 : Math.ceil(filteredHeroes.length / pageSize);

  for (let i = 1; i <= totalPages; i++) {
    const button = document.createElement('button');
    button.textContent = i;
    button.disabled = i === currentPage;
    button.onclick = () => {
      currentPage = i;
      renderTable();
    };
    pagination.appendChild(button);
  }
}

// Change page size and reset the current page
function changePageSize() {
  const selector = document.getElementById('page-size');
  pageSize = selector.value === 'all' ? filteredHeroes.length : parseInt(selector.value);
  currentPage = 1; // Reset to the first page whenever the page size changes
  renderTable();
}

// Live search by name without debounce
function filterData() {
  const query = document.getElementById('search').value.toLowerCase();
  filteredHeroes = heroes.filter(hero => hero.name.toLowerCase().includes(query));
  currentPage = 1; // Reset to the first page whenever the filter changes
  renderTable();
}

// Sorting function with missing values sorted last
function sortTable(column) {
  const isAscending = sortOrder.column === column ? !sortOrder.ascending : true;
  sortOrder = { column, ascending: isAscending };

  filteredHeroes.sort((a, b) => {
    const valA = getColumnValue(a, column);
    const valB = getColumnValue(b, column);

    // Handle missing values (move them to the end)
    const missingValues = ['N/A', 'null', null, '-', undefined];
    const isValAMissing = missingValues.includes(valA);
    const isValBMissing = missingValues.includes(valB);

    // Always move missing values to the end
    if (isValAMissing && isValBMissing) return 0;
    if (isValAMissing) return 1; // Missing values come last
    if (isValBMissing) return -1;

    // Special case for sorting by weight and height
    if (column === 'weight') {
      return compareWeight(valA, valB, isAscending);
    }

    if (column === 'height') {
      return compareHeight(valA, valB, isAscending);
    }

    // Regular sorting for other columns
    if (valA === valB) return 0;

    // Determine if values are numbers
    const valANumber = typeof valA === 'number';
    const valBNumber = typeof valB === 'number';

    if (valANumber && valBNumber) {
      // Both are numbers
      return isAscending ? valA - valB : valB - valA;
    } else {
      // Compare as strings
      if (isAscending) {
        return valA > valB ? 1 : -1;
      } else {
        return valA < valB ? 1 : -1;
      }
    }
  });

  renderTable();
}

// Custom function to compare heights (converted to centimeters)
function compareHeight(heightA, heightB, isAscending) {
  const parsedHeightA = parseHeight(heightA);
  const parsedHeightB = parseHeight(heightB);

  // Handle missing or invalid heights
  if (parsedHeightA.value === null && parsedHeightB.value === null) return 0;
  if (parsedHeightA.value === null) return 1; // Missing values come last
  if (parsedHeightB.value === null) return -1;

  // Sort by numerical value
  if (parsedHeightA.value === parsedHeightB.value) return 0;
  return isAscending
    ? parsedHeightA.value - parsedHeightB.value
    : parsedHeightB.value - parsedHeightA.value;
}

// Custom function to compare weights (converted to kilograms)
function compareWeight(weightA, weightB, isAscending) {
  const parsedWeightA = parseWeight(weightA);
  const parsedWeightB = parseWeight(weightB);

  // Handle missing or invalid weights
  if (parsedWeightA.value === null && parsedWeightB.value === null) return 0;
  if (parsedWeightA.value === null) return 1; // Missing values come last
  if (parsedWeightB.value === null) return -1;

  // Sort by numerical value
  if (parsedWeightA.value === parsedWeightB.value) return 0;
  return isAscending
    ? parsedWeightA.value - parsedWeightB.value
    : parsedWeightB.value - parsedWeightA.value;
}

// Helper function to parse weight (kg and tons)
function parseWeight(weightStr) {
  if (!weightStr || weightStr === 'null' || weightStr === '-' || weightStr === '0 kg') return { value: null, unit: null };

  weightStr = weightStr.trim();

  // Check for tons
  if (weightStr.includes(' ton')) {
    const valueInTons = parseFloat(weightStr.replace(/,/g, '').replace(' tons', '').trim());
    if (isNaN(valueInTons) || valueInTons === 0) return { value: null, unit: null }; // Treat 0 tons as missing
    return { value: valueInTons * 1000, unit: 'kg' }; // Convert tons to kg
  }

  // Check for kilograms
  if (weightStr.includes(' kg')) {
    const valueInKg = parseFloat(weightStr.replace(/,/g, '').replace(' kg', '').trim());
    if (isNaN(valueInKg) || valueInKg === 0) return { value: null, unit: null }; // Treat 0 kg as missing
    return { value: valueInKg, unit: 'kg' };
  }

  // Check for pounds (if needed)
  if (weightStr.includes(' lb')) {
    const valueInLb = parseFloat(weightStr.replace(/,/g, '').replace(' lb', '').trim());
    if (isNaN(valueInLb) || valueInLb === 0) return { value: null, unit: null }; // Treat 0 lb as missing
    const kg = valueInLb * 0.453592; // Convert lb to kg
    return { value: kg, unit: 'kg' };
  }

  return { value: null, unit: null }; // Unable to parse
}

// Helper function to parse height and convert to centimeters
function parseHeight(heightStr) {
  if (!heightStr || heightStr === 'null' || heightStr === '-' || heightStr === '0 cm') return { value: null, unit: null };

  heightStr = heightStr.trim();

  // Check for meters
  if (heightStr.includes(' m')) {
    const valueInMeters = parseFloat(heightStr.replace(' m', '').trim());
    if (isNaN(valueInMeters) || valueInMeters === 0) return { value: null, unit: null }; // Treat 0 m as missing
    return { value: valueInMeters * 100, unit: 'cm' }; // Convert to centimeters
  }

  // Check for centimeters
  if (heightStr.includes(' cm')) {
    const valueInCm = parseFloat(heightStr.replace(' cm', '').trim());
    if (isNaN(valueInCm) || valueInCm === 0) return { value: null, unit: null }; // Treat 0 cm as missing
    return { value: valueInCm, unit: 'cm' };
  }

  // Check for feet and inches (e.g., 6'11")
  if (heightStr.includes("'")) {
    const regex = /(\d+)'(\d+)"/;
    const match = heightStr.match(regex);
    if (match) {
      const feet = parseInt(match[1], 10);
      const inches = parseInt(match[2], 10);
      const totalInches = feet * 12 + inches;
      const cm = totalInches * 2.54;
      if (cm === 0) return { value: null, unit: null }; // Treat 0 cm as missing
      return { value: cm, unit: 'cm' };
    }
  }

  return { value: null, unit: null }; // Unable to parse
}

// Helper to get value from hero object for sorting
function getColumnValue(hero, column) {
  switch (column) {
    case 'name':
      return hero.name;
    case 'fullName':
      const fullName = hero.biography.fullName;
      // Treat '-' as missing value
      if (fullName === '-' || fullName === undefined || fullName === null || fullName === '') {
        return 'N/A';
      } else {
        return fullName;
      }
    case 'race':
      return hero.appearance.race !== undefined ? hero.appearance.race : 'N/A';
    case 'gender':
      return hero.appearance.gender !== undefined ? hero.appearance.gender : '-'; // Treat "-" as missing for sorting
    case 'height':
      // Use combined height (both units) for sorting
      // We'll use the metric value for parsing
      const heightMetric = hero.appearance.height[1];
      return heightMetric || 'N/A';
    case 'weight':
      // Use combined weight (both units) for sorting
      // We'll use the metric value for parsing
      const weightMetric = hero.appearance.weight[1];
      return weightMetric || 'N/A';
    case 'placeOfBirth':
      return hero.biography.placeOfBirth !== undefined ? hero.biography.placeOfBirth : 'N/A';
    case 'alignment':
      return hero.biography.alignment !== undefined ? hero.biography.alignment : 'N/A';
    case 'intelligence':
    case 'strength':
    case 'speed':
    case 'durability':
    case 'power':
    case 'combat':
      const statValue = parseInt(hero.powerstats[column]);
      return !isNaN(statValue) ? statValue : hero.powerstats[column] || 'N/A';
    default:
      return 'N/A';
  }
}
