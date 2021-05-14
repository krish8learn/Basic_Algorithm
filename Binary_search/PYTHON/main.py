def binarySearch(arr, min, max, targetvalue):
    if max >= min:
        mid = min + (max-min)//2
        if arr[mid] == targetvalue:
            return mid
        elif arr[mid]>targetvalue:
            max = mid -1
            return binarySearch(arr,min, max, targetvalue) 
        elif arr[mid]<targetvalue:
            min = mid + 1
            return binarySearch(arr, min, max, targetvalue)    
    else:
        return -1

arr = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97]

result = binarySearch(arr,0,len(arr)-1, 3)

print(result)