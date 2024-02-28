function getAverage(marks){
    let result = 0
    for (let i = 0 ; i < marks.length; i++){
        result += marks[i]
    }
    return Math.floor(result / marks.length)
}