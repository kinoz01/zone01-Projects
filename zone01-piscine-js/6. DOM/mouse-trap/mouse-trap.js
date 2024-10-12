let outside = true
let box

const createCircle = () => {
    document.addEventListener('click', event => {
        createCircleElement(event);
        outside = true;
    })
}

const moveCircle = () => {
    document.addEventListener('mousemove', e => {
        document.querySelectorAll('.removeCircle').forEach(elem => {
            elem.remove()
        })

        let circle = createCircleElement(e, 'removeCircle');
        const boxRect = box.getBoundingClientRect();

        if ((e.clientX >= boxRect.left + 25 && e.clientX <= boxRect.right - 25) && (e.clientY >= boxRect.top + 25 && e.clientY <= boxRect.bottom - 25)) {
            outside = false
        }
        if (!outside) {
            if (e.clientX - 25 < boxRect.left) {
                circle.style.left = boxRect.left + "px"             
            }
            if (e.clientX + 25 > boxRect.right) {
                circle.style.left = boxRect.right - 50 + "px"
            }
            if (e.clientY - 25 < boxRect.top) {
                circle.style.top = boxRect.top + "px"
            }
            if (e.clientY + 25 > boxRect.bottom) {
                circle.style.top = boxRect.bottom - 50 + "px"
            }
            document.querySelector(".circle").style.background = 'var(--purple)'
        }
    })
}

const setBox = () => {
    box = document.createElement("div")
    box.className = "box"
    document.body.appendChild(box)
}


const createCircleElement = (event, extraClass = '') => {

    let circle = document.createElement('div')
    circle.className = 'circle'
    if (extraClass) circle.classList.add(extraClass)

    if (outside) circle.style.background = 'white'
    else circle.style.background = 'var(--purple)' 
     
    circle.style.left = event.clientX - 25 + 'px'
    circle.style.top = event.clientY - 25 + 'px'
    document.body.appendChild(circle)
    return circle;
}

export { createCircle, moveCircle, setBox }