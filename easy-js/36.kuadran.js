// https://www.codewars.com/kata/643af0fa9fa6c406b47c5399/train/javascript

function quadrant(x, y) {
  if (x > 0 && y > 0) {
    return 1;
  } else if (x < 0 && y > 0) {
    return 2;
  } else if (x < 0 && y < 0) {
    return 3;
  } else {
    return 4;
  }
}

console.log(quadrant(1, -2));
