// https://www.codewars.com/kata/5861d28f124b35723e00005e

// const zeroFuel = (distanceToPump, mpg, fuelLeft) => {
//     return mpg * fuelLeft >= distanceToPump;

// };

const zeroFuel = (distanceToPump, mpg, fuelLeft) =>
  distanceToPump / mpg >= fuelLeft;

console.log(zeroFuel(100, 50, 2));
