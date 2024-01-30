// https://www.codewars.com/kata/54c27a33fb7da0db0100040e

function isSquare(n) {
  if (n < 0) {
    return false; // Bilangan negatif bukan bilangan kuadrat
  }

  let i = 0;
  while (i * i <= n) {
    if (i * i === n) {
      return true; // Menemukan bilangan bulat i yang saat dikuadratkan sama dengan n
    }
    i++;
  }

  return false; // Tidak menemukan bilangan bulat yang saat dikuadratkan sama dengan n
}

// Contoh penggunaan:
console.log(isSquare(-1)); // Output: false
console.log(isSquare(0)); // Output: true
console.log(isSquare(3)); // Output: false
console.log(isSquare(4)); // Output: true
console.log(isSquare(25)); // Output: true
