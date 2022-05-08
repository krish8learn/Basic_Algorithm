def Swap(array, firstIndex, secondIndex):
    temp = array[firstIndex]
    array[firstIndex]=array[secondIndex]
    array[secondIndex] = temp

def MinIndexSubArray(array,startIndex):
    minIndex = startIndex
    minValue = array[startIndex]
    
    i = startIndex
    while i!=len(array):
        #print(i)
        if array[i]<minValue:
            minIndex = i
            minValue = array[i]
        i+=1
    #print(minIndex, minValue)
    return minIndex 

def SelectionSort(array):
    for i in range(len(array)):
        swapIndex = MinIndexSubArray(array,i)
        Swap(array,i,swapIndex)

array = [4, 3, 2, 1, 0, -1, -99]

SelectionSort(array)

print(array)