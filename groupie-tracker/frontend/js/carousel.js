
document.querySelectorAll('.carousel-section').forEach(section => {
    const carousel = section.querySelector('.carousel-container');
    const leftButton = section.querySelector('.left-button');
    const rightButton = section.querySelector('.right-button');

    leftButton.addEventListener('click', () => {
        carousel.scrollBy({
            left: -carousel.offsetWidth,
            behavior: 'smooth'
        });
    });

    rightButton.addEventListener('click', () => {
        carousel.scrollBy({
            left: carousel.offsetWidth,
            behavior: 'smooth'
        });
    });
});