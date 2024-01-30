// https://www.codewars.com/kata/57a429e253ba3381850000fb

function bmi(weight, height) {
  let restbmi = weight / (height * height);
  if (restbmi <= 18.5) {
    return "Underweight";
  } else if (restbmi <= 25.0) {
    return "Normal";
  } else if (restbmi <= 30.0) {
    return "Overweight";
  } else {
    return "Obese";
  }
}

console.log(bmi(80, 1.8));
