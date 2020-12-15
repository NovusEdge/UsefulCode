def quick_Sort(inputArray):
    '''Method to recursively perform quick-sort on an array'''
    if len(inputArray) <= 1:
        return inputArray
    
    else:
        
        begin = quick_Sort([i for i in inputArray[1:] if i < inputArray[0]])
        pivot = [inputArray[0]]
        end = quick_Sort([i for i in inputArray[1:] if i >= inputArray[0]])
        
        return begin + pivot + end
