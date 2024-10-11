export const build = (num) => {

    let id = 1
    const create = () => {
        let block = document.createElement('div') // live only in memory
        block.id = `brick-${id}`

        if ((id + 1) % 3 === 0) {
            block.dataset.foundation = 'true'
        }
        document.body.append(block) // actually append the div to the html

        if (id === num) {
            clearInterval(interval) // Stops the interval when the number is reached
        }
        id++
    }
    let interval = setInterval(create, 100) // Calls create every 100ms
}

export const repair = (...ids) => {
    ids.forEach(id => {
        let element = document.getElementById(id)
        if (element && element.hasAttribute('data-foundation')) {
            element.dataset.repaired = 'in progress'
        } else if (element) {
            element.dataset.repaired = 'true'
        }
    })
}

export const destroy = () => {
    let lastBlock = Array.from(document.body.getElementsByTagName('div')).reverse()[0]
    if (lastBlock) lastBlock.remove()
}