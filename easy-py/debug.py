def how_many_dalmatians(number):
    dogs =  ["Hardly any", "More than a handful!", "Woah that's a lot of dogs!", "101 DALMATIONS!!!"]
    respond = []
    if number <= 10 :
        respond = dogs[0]
    elif  number <= 50 :
        respond =  dogs[1]
    elif number == 101 :
        respond = dogs[3]
    else:
        respond = dogs[2]

    return respond

print(how_many_dalmatians(12))