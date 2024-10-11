export const getArchitects = () => {
    let architects = Array.from(document.body.getElementsByTagName('a'))
    let nonArchitects = Array.from(document.body.getElementsByTagName('span'))
    return [architects, nonArchitects]
}

export const getClassical = () => {
    let architects = getArchitects()[0]
    let classical = architects.filter(elm => elm.classList.contains('classical'))
    let others = architects.filter(elm => !classical.includes(elm))

    return [classical, others]
}

export const getActive = () => {
    let classical = getClassical()[0]
    let active = classical.filter(elem => elem.classList.contains('active'))
    let others = classical.filter(elem => !active.includes(elem))

    return [active, others]
}

export const getBonannoPisano = () => {
    return [document.getElementById('BonannoPisano'), getActive()[0]]
}
