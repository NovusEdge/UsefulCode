def CountingSort(arr, place):
    temp = arr
    size = len(temp)
    res = [0] * len(temp)
    count = [0] * 10

    for i in range(0, len(temp)):
        index = arr[i] // place
        count[index % 10] += 1

    for i in range(1, 10):
        count[i] += count[i - 1]

    i = size - 1
    while i >= 0:
        index = arr[i] // place
        res[count[index % 10] - 1] = arr[i]
        count[index % 10] -= 1
        i -= 1


    for i in range(0, size):
        temp[i] = res[i]

    return temp


def RadixSort( arr ):
    sortedArr = arr
    max_element = max( sortedArr )

    place = 1
    while max_element // place > 0:
        CountingSort( sortedArr , place )
        place *= 10

    return sortedArr

if __name__ == "__main__":
    from random import randint
    from time import time

    # running test cases:
    testArr = [ randint(0, 1000) for i in range(50) ]

    print("Unsorted Array: ", testArr, "\n"); start = time()
    print("Sorted Array: ", RadixSort(testArr), "\n")
    print("Time Taken: ", time() - start)
