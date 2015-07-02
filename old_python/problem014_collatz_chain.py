

def nextnum(n):
    if n % 2:
        return 3 * n + 1
    return n / 2


dct = {}
#tab = 0


def chainlen(n):
    global dct
    #global tab
    #tab = tab + 1
    #print tab * '  ', n
    if n < 1:
        return 1 / 0
    if n < 3:
        return n
    if n not in dct:
        nn = nextnum(n)
        if nn not in dct:
            dct[nn] = chainlen(nn)
        dct[n] = 1 + dct[nn]
    #print tab * '  ', dct[n]
    #tab = tab - 1
    return dct[n]


def make_all_len(lim):
    global dct
    dct[1] = 1
    dct[2] = 2
    lex, n, mx = 0, 0, 0
    for x in range(lim, 1, -1):
        #print "--"
        if x not in dct:
            lex = chainlen(x)
        else:
            lex = dct[x]
        if lex > mx:
            n, mx = x, lex
    print mx, n

#import sys
#args = sys.argv


#make_all_len(int(args[1]))
