from math import sqrt; from itertools import count, islice

def isPrime(n):
    '''A Method to check if a number is prime or not.'''
    return n > 1 and all(n%i for i in islice(count(2), int(sqrt(n)-1)))

# For further useful info/reference head over to : https://primes.utm.edu/prove/index.html
