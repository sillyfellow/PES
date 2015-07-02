

def palindrome_number(n):
    if n < 10:
        return True
    if n < 100:
        return (n % 10) == (n / 10)
    tenmul = 10
    while tenmul < n / 10:
        tenmul = tenmul * 10
    left, right = tenmul, 10
    while right <= left:
        if (n / left) % 10 != (n % right) / (right / 10):
            return False
        left = left / 10
        right = right * 10
    return True


start = 999
end   = 100

print palindrome_number(121)

a, b, large = 0, 0, 0
for x in range(start, end, -1):
    for y in range(start, end, -1):
        if palindrome_number(x * y):
            if large < x * y:
                a, b, large = x, y, x * y

print a, b, large
