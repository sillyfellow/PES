import math


def largest_prime_factor(n):
    if n < 2:
        return 1
    denom = 2
    if not n % denom:
        return max(denom, largest_prime_factor(n / denom))
    sqrt = math.sqrt(n)
    denom = 3
    while True:
        if not n % denom:
            print denom
            return max(denom, largest_prime_factor(n / denom))
        denom = denom + 2
        if denom > sqrt:
            return n


print largest_prime_factor(600851475143)
