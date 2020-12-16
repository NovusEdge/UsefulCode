
def InsertionSort( arr ):
    sortedArr = arr

    for i in range(len(arr)):
        for j in range(len(arr)):
            if sortedArr[j] > sortedArr[i]:
                sortedArr[j], sortedArr[i] = sortedArr[i], sortedArr[j]
    
    return sortedArr


if __name__ == "__main__":
    from random import randint
    from time import time

    # running test cases:
    testArr = [ randint(0, 1000) for i in range(50) ]

    print("Unsorted Array: ", testArr); start = time()
    print("Sorted Array: ", InsertionSort(testArr))
    print("Time Taken: ", time() - start)