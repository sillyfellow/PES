import primality


def sumprimes(maxlim):
    primes = [2]
    for end in range(3, maxlim, 2):
        newp = end
        for p in primes:
            if not end % p:
                newp = 0
                break
        if newp:
            primes.append(newp)
            #if not len(primes) % 10000:
                #print ts, primes[-1], end
                #ts = ts + 10
    return sum(primes)


def primsum(s, e):
    sum = 0
    for xxx in range(s, e):
        if primality.miller_rabin_prime(xxx):
            sum = sum + xxx
    return sum

TOP=2000000

#print sumprimes(TOP)
print primsum(0, TOP)
