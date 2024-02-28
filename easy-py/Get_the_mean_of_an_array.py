import math


def get_average(marks):
    total = 0
    for i in range(len(marks)):
        total += marks[i]

    average = total / len(marks)
    rounded_average = math.floor(average)
    return rounded_average