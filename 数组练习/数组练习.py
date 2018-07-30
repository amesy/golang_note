import math

def for_(num_half, res):
    for i in range(num_half):
        res.append(i)
        res.append(-i)
    return res

def array(num, res=[]):
    num_half = math.ceil(num/2)
    # num为偶数
    if num % 2 == 0:
        return for_(num_half, res)

    # num为奇数
    L = for_(num_half, res)
    return L[1:]

print(array(7))