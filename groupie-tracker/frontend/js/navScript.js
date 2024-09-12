const menuToggle = document.getElementById('mobile-menu');
const navLinks = document.querySelector('.nav-links');
const body = document.body;

// Toggle menu visibility when menu button is clicked
menuToggle.addEventListener('click', (e) => {
    navLinks.classList.toggle('active');
    e.stopPropagation(); // Prevent the body click event from firing when clicking on the menu toggle
});

// Close the menu when clicking anywhere outside the menu
body.addEventListener('click', () => {
    if (navLinks.classList.contains('active')) {
        navLinks.classList.remove('active');
    }
});
