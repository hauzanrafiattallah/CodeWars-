// https://www.codewars.com/kata/59ca8246d751df55cc00014c

// function hero(bullets, dragons) {
//   if (bullets >= dragons * 2) {
//     return true;
//   } else {
//     return false;
//   }
// }

// function hero(bullets, dragons) {
//   return bullets >= dragons * 2;
// }

const hero = (bullets, dragon) => bullets / 2 >= dragon;

console.log(hero(10, 5));
