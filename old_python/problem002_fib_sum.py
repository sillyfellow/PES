def fibsum(maxlim):
    sum = 0
    c, a, b = 2, 1, 2
    while c < maxlim:
        sum = sum + c
        a = a + b
        b = a + b
        c = a + b
        if c > maxlim:
            break
        a, b = b, c
    return sum


print fibsum(4000000)
