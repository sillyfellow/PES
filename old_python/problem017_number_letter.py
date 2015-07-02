
hundreds = ['one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']
tens     = ['twenty', 'thirty', 'forty', 'fifty', 'sixty', 'seventy', 'eighty', 'ninety']
specs    = ['ten', 'eleven', 'twelve', 'thirteen', 'fourteen', 'fifteen', 'sixteen', 'seventeen', 'eighteen', 'nineteen']
units    = ['one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']


def inside100():
    for u in units:
        yield u + " "
    for s in specs:
        yield s + " "
    for t in tens:
        yield t + " "
        for u in units:
            yield t + " " + u + " "


def all100s():
    for n in inside100():
        yield n
    for h in hundreds:
        yield h + " hundred "
        for n in inside100():
            yield h + " hundred and " + n + " "
    yield "one thousand"


#>>>
#>>> for nn in pnl.all100s():
    #...     sss = sss + nn
    #...
    #>>> sss=''
    #>>> for nn in pnl.all100s():
        #...     sss = sss + nn
        #...
        #>>> len(''.join(sss.split()))
        #21124

