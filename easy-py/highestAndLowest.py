def high_and_low(numbers):
    num_list = [int(num) for num in numbers.split()]
    highest = max(num_list)
    lowest = min(num_list)
    result = f"{highest} {lowest}"
    return result

print(high_and_low("1 2 3 4 6 8 10"))