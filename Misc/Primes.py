from math import sqrt; from itertools import count, islice

def isPrime(n):
    '''A Method to check if a number is prime or not.'''
    return n > 1 and all(n%i for i in islice(count(2), int(sqrt(n)-1)))

# For further useful info/reference head over to : https://primes.utm.edu/prove/index.html



#Prime Sieve : Implementation of 'Sieve of Eratosthenes' in python3

def primeSieve(upperLim):
    '''Method to get a list of all prime numbers below upperLim'''
    primeList = list(range(2,upperLim))
    for i in primeList:
        j = 2
        while j < primeList[::-1][0]:
            if i*j in primeList:
                primeList.remove(i*j)
            j += 1
    return primeList


# a better approach:

def better_primeSieve(n):
    sieve = [True] * n
    for i in range(3,int(n**0.5)+1,2):
        if sieve[i]:
            sieve[i*i::2*i]=[False]*((n-i*i-1)//(2*i)+1)
    return [2] + [i for i in range(3,n,2) if sieve[i]]

