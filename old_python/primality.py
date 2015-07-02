
from random import randint


def powmod(a, b, m):
    if m == 0:
        return -1
    if a == 0:
        return 0
    if a == 1:
        return 1
    if b == 0:
        return 1
    if b == 1:
        return a % m
    if b == 2:
        return (a * a) % m
    if b % 2:
        return (a * powmod(a, b - 1, m)) % m
    else:
        interim = powmod(a, b / 2, m)
        return (interim * interim) % m


def fermatcomposite(n):
    return powmod(randint(1, n - 1), n - 1, n) != 1


def fermatprime(n, testtimes=7):
    if n == 1:
        return False
    for t in range(testtimes):
        if fermatcomposite(n):
            return False
    return True


def largest_odd_div(d):
    s = 0
    while not d % 2:
        s = s + 1
        d = d / 2
    return (s, d)


def miller_rabin_composite(n, testtimes=7):
    if n < 7:
        return (n == 1) or (n == 4) or (n == 6)
    s, d = largest_odd_div(n - 1)
    for t in range(testtimes):
        a = randint(2, n - 2)
        x = powmod(a, d, n)
        if (x == 1) or (x == n - 1):
            continue
        reloop = False
        for q in range(s - 1):
            x = powmod(x, 2, n)
            if x == 1:
                return True
            if x == n - 1:
                reloop = True
                break
        if not reloop:
            return True
    return False


def miller_rabin_prime(n, testtimes=3):
    if n == 1:
        return False
    for t in range(testtimes):
        if miller_rabin_composite(n, testtimes):
            return False
    return True
