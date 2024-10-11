import { styles } from './pimp-my-style.data.js'

let click = 0

export const pimp = () => {
    let button = document.querySelector('button')
    if (button.classList.contains('unpimp')) {
        click--
        if (click == 0) {
            button.classList.remove(styles[click], 'unpimp')
        } else {
            button.classList.remove(styles[click])
        }
    } else {
        if (click == 14) {
            button.classList.add(styles[click], 'unpimp')
        } else {
            button.classList.add(styles[click])
        }
        click++
    }
}