def BubbleSort( arr ):
    """ BubbleSort sorts an array using the bubble sort algorithm """

    sortedArr = arr

    for _ in arr:
        for i in range(len(arr)-1):
            if arr[i] > arr[i+1]:
                arr[i], arr[i+1] = arr[i+1], arr[i]

    return sortedArr


if __name__ == "__main__":
    from random import randint
    from time import time

    # running test cases:
    testArr = [ randint(0, 1000) for i in range(50) ]

    print("Unsorted Array: ", testArr); start = time()
    print("Sorted Array: ", BubbleSort(testArr))
    print("Time Taken: ", time() - start)