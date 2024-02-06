#  https://www.codewars.com/kata/59ca8246d751df55cc00014c

def hero(bullets, dragons):
  if bullets / 2 >= dragons:
    return True
  else:
    return False
  
result = hero(5, 5)
print (result)

