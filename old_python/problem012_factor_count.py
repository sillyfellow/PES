
from primality import miller_rabin_prime
from math import sqrt


def fcount(n, mx=500):
    if n == 1:
        return [1]
    if n == 2:
        return [1, 2]
    if n == 3:
        return [1, 3]
    if n == 4:
        return [1, 2, 2, 4]

    fc = [1, n]

    if miller_rabin_prime(n):
        return fc

    for f in range(2, int(1 + sqrt(n))):
        if not n % f:
            fc = fc + [f, n / f]
            if len(fc) > mx:
                break
    return fc


def listprod(a, b):
    c = []
    for x in a:
        for y in b:
            c = c + [x*y]
    return c


def dobiz(start, end, inter=1):
    for xx in range(start, end, inter):
        ff1, ff2 = [], []
        if miller_rabin_prime(xx):
            ff1 = fcount(xx + 1)
        elif miller_rabin_prime(xx + 1):
            ff2 = fcount(xx) + [xx + 1]
        else:
            ff1 = fcount(xx)
            ff2 = fcount(xx + 1)
        ff3 = ff1 + ff2
        ff4 = listprod(ff3, ff3)
        product = xx * (xx + 1) / 2
        rootp = sqrt(product)
        count = 0
        ff5 = []
        for yy in filter(lambda x: x <= rootp, set(ff4)):
            if not product % yy:
                #ff5 = ff5 + [yy, product / yy]
                count = count + 2
        print xx, product, count
        if count > 500:
            break



