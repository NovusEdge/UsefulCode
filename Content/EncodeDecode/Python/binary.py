from functools import reduce

def stringToBin( data ):
    if type(data) == int:
        return bin( data )[2:]
    elif type(data) == str:
        return
    elif type(data) in (list, tuple):
        return map(lambda a: bin(x)[2:], data)
    else:
        return




