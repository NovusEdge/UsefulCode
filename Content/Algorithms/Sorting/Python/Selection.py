
def SelectionSort( arr ):
    sortedArr = arr

    for i in range(len(arr)):
        flag = i

        for j in range(i+1, len(arr)):
            if sortedArr[flag] > sortedArr[j]:
                flag = j

        sortedArr[flag], sortedArr[i] = sortedArr[i], sortedArr[flag]
    
    return sortedArr


if __name__ == "__main__":
    from random import randint
    from time import time

    # running test cases:
    testArr = [ randint(0, 1000) for i in range(50) ]

    print("Unsorted Array: ", testArr); start = time()
    print("Sorted Array: ", SelectionSort(testArr))
    print("Time Taken: ", time() - start)