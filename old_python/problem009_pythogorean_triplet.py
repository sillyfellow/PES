

def is_pythogorean(a, b, c):
    hyp = max(a, b, c)
    return (hyp ** 2 + hyp ** 2 - a ** 2 - b ** 2 - c ** 2) == 0


def all_triplet_combi(n):
    for hyp in range(n / 2 + 1):
        for base in range(hyp):
            for alt in range(base):
                if (alt + base) > hyp:
                    if alt + base + hyp == n:
                        if is_pythogorean(alt, base, hyp):
                            print alt, base, hyp


print all_triplet_combi(1000)
