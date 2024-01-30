// https://www.codewars.com/kata/5a00e05cc374cb34d100000d

// const reverseSeq = (n) => {
//   const result = [];
//   for (i = n; i >= 1; i--) {
//     result.push(i);
//   }
//   return result;
// };

// const reverseSeq = (n) => {
//   return Array(n)
//     .fill()
//     .map((ell, i) => n - i);
// };

const reverseSeq = (n) => [...Array(n)].map((ell, i) => n - i);

console.log(reverseSeq(5));
