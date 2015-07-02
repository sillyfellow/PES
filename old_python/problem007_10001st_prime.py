
def any_one_divides(n, lp):
    for p in lp:
        if n % p == 0: return True
    return False


def all_n_primes(n):
    primes = []
    end = 1
    while True:
        end = end + 1
        remainder = 0
        if any_one_divides(end, primes):
            continue
        primes = primes + [end]
        if len(primes) >= n:
            break
    return primes


print all_n_primes(10001)
