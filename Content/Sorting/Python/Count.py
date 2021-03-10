
def CountingSort( arr ):
    countArr = [0]*(max(arr)+1)
    for i in arr:
        countArr[i] += 1

    resArr = []

    for i in range(len(countArr)):
        resArr.extend( [i]*countArr[i] )

    return resArr

if __name__ == "__main__":
    from random import randint
    from time import time

    # running test case:
    testArr = [ randint(0, 1000) for i in range(50) ]

    print("Unsorted Array: ", testArr, "\n"); start = time()
    print("Sorted Array: ", CountingSort(testArr), "\n")
    print("Time Taken: ", time() - start)