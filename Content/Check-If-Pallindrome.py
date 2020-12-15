def isPallindrome(n):
    '''returns true if entered number/string is a pallindrome'''
    
    return str(n) == str(n)[::-1]
 
# We can also implement this with a lambda method : 
isPall = lambda n : str(n) == str(n)[::-1]
