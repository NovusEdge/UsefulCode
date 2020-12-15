def getDigitalRoot(n = None):
    '''A Method to get the digital root of a number i.e. the sum of all its digits.'''
    if n is None:
        return None
    digitList = [int(i) for i in str(n)] 
    return sum(digitList)
