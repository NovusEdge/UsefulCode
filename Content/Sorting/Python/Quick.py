def QuickSort(inputArray):
    '''Method to recursively perform quick-sort on an array'''
    if len(inputArray) <= 1:
        return inputArray

    else:

        begin = QuickSort([i for i in inputArray[1:] if i < inputArray[0]])
        pivot = [inputArray[0]]
        end = QuickSort([i for i in inputArray[1:] if i >= inputArray[0]])

        return begin + pivot + end


if __name__ == "__main__":
    from random import randint
    from time import time

    # running test cases:
    testArr = [ randint(0, 1000) for i in range(50) ]

    print("Unsorted Array: ", testArr, "\n"); start = time()
    print("Sorted Array: ", QuickSort(testArr), "\n")
    print("Time Taken: ", time() - start)
