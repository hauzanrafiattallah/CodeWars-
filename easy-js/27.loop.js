// https://www.codewars.com/kata/55ecd718f46fba02e5000029

function between(a, b) {
  let hasil = [];
  while (a <= b) {
    hasil.push(a);
    a++;
  }
  return hasil;
}

console.log(between(2, 8));
