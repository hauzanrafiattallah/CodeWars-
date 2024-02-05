// https://www.codewars.com/kata/51dda84f91f5b5608b0004cc/train/javascript

function solution(number) {
  const result = [0, 0, 0];
  for (let i = 1; i < number; i++) {
    if (i % 3 === 0 && i % 5 != 0) {
      result[0]++;
    } else if (i % 5 === 0 && i % 3 != 0) {
      result[1]++;
    } else if (i % 3 === 0 && i % 5 === 0) {
      result[2]++;
    }
  }
  return result;
}

console.log(solution(20));
