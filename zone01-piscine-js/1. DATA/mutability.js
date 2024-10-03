let clone1 =  { ...person };
let clone2 =  { ...person };
  
// here its like we created a pointer to the varibale
// changing the pointer value will also change the original variable
const samePerson  = person
  
samePerson.country = 'FR'; 
samePerson.age++

console.log(person)
