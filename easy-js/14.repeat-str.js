// https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

function repeatStr(n, s) {
  let result = "";
  for (let i = 0; i < n; i++) {
    result += s;
  }
  return result;
}

console.log(repeatStr(5, "hauzan"));
