function filterEntries(obj, func) {
    return Object.fromEntries(Object.entries(obj).filter(func))
}

function mapEntries(obj, func) {
    return Object.fromEntries(Object.entries(obj).map(func))
}

function reduceEntries(obj, ...args) {
    return Object.entries(obj).map(...args)
}

function totalCalories(obj) {
    let res = 0
    Object.keys(obj).forEach(key => {
        let amount = obj[key]
        res += nutritionDB[key]['calories'] * (amount / 100)
    })
    return Number(res.toFixed(1))
}

function lowCarbs(obj) {
    let res = {}
    Object.keys(obj).forEach(key => {
        let carbsAmount = (obj[key] / 100) * Number(nutritionDB[key]['carbs'])
        if (carbsAmount < 50) res[key] = obj[key]
    })
    return res
}

function cartTotal(obj) {
    let res = {}
    Object.keys(obj).forEach(key => {
        res[key] = nutritionDB[key]
    })
    Object.keys(res).forEach(key => {
        let amount = Number(obj[key])
        // range on subkeys
        Object.keys(nutritionDB[key]).forEach(subKey => {
            res[key][subKey] = Number((nutritionDB[key][subKey] * (amount / 100)).toFixed(2))
        })

    })
    return res
}

/************** Another ways using reduce, filter and map ***********/
// using reduce
const totalCal = obj => {
    return Number(Object.entries(obj).reduce((acc, keyVal) =>
        acc + nutritionDB[keyVal[0]]['calories'] * (keyVal[1] / 100), 0
    ).toFixed(1))
}

const lowCarb = obj => {
    return Object.fromEntries(Object.entries(obj).filter(keyVal =>
        nutritionDB[keyVal[0]]['carbs'] * (keyVal[1]) / 100 < 50
    ))
}

const cartTot = obj => {
    return Object.fromEntries(Object.entries(obj).map(([food, amount]) => {
        const macroNutrients = Object.fromEntries(
            Object.entries(nutritionDB[food]).map(([macro, perCent]) => {
                return [macro, Number((amount * (perCent / 100)).toFixed(2))]
            }))
        return [food, macroNutrients]
    }))
}

// const nutritionDB = {
//     tomato: { calories: 18, protein: 0.9, carbs: 3.9, sugar: 2.6, fiber: 1.2, fat: 0.2 },
//     vinegar: { calories: 20, protein: 0.04, carbs: 0.6, sugar: 0.4, fiber: 0, fat: 0 },
//     oil: { calories: 48, protein: 0, carbs: 0, sugar: 123, fiber: 0, fat: 151 },
//     onion: { calories: 0, protein: 1, carbs: 9, sugar: 0, fiber: 0, fat: 0 },
//     garlic: { calories: 149, protein: 6.4, carbs: 33, sugar: 1, fiber: 2.1, fat: 0.5 },
//     paprika: { calories: 282, protein: 14.14, carbs: 53.99, sugar: 1, fiber: 0, fat: 12.89 },
//     sugar: { calories: 387, protein: 0, carbs: 100, sugar: 100, fiber: 0, fat: 0 },
//     orange: { calories: 49, protein: 0.9, carbs: 13, sugar: 9, fiber: 0.2, fat: 0.1 },
// }

// const groceriesCart = { orange: 500, oil: 20, sugar: 480 }
// console.log('Total calories:')
// console.log(totalCalories(groceriesCart))
// console.log('Items with low carbs:')
// console.log(lowCarb(groceriesCart))
// console.log('Total cart nutritional facts:')
// console.log(cartTot(groceriesCart))
