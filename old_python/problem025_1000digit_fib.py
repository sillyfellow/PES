#!/usr/bin/python


def fibn():
    a, b, c = 1, 1, 0
    while True:
        c = a
        a = b
        b = b + c
        yield c

count = 0
for x in fibn():
    count = count + 1
    if x > (10 ** 999):
        print count
        print x
        break
    if x == 144:
        print count
