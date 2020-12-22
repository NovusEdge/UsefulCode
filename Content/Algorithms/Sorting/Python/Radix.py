def CountingSort( arr, sigPos=1 ):
    pass

def RadixSort( arr ):
    sortedArr = arr
    sigDigit = lambda num, b: num % 10**b
    flag = len(str(max(arr)))

    for i in range(len(arr)):
        pass


if __name__ == "__main__":
    from random import randint
    from time import time

    # running test cases:
    testArr = [ randint(0, 1000) for i in range(50) ]

    print("Unsorted Array: ", testArr, "\n"); start = time()
    print("Sorted Array: ", RadixSort(testArr), "\n")
    print("Time Taken: ", time() - start)