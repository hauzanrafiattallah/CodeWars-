# https://www.codewars.com/kata/55685cd7ad70877c23000102/train/python

def make_negative( number ):
    if number > 0 :
        return number * -1
    else:
        return number
    
result = make_negative(-1)
print(result)