def bubble_sort(arr:list):
    ln = len(arr)
    for i in range(ln):
        isSwapped = False
        for j in range(0, ln - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                isSwapped = True
        if not isSwapped:
            break

a = [4, 3, 2, 1]

print('Array before sorting', a)

bubble_sort(a)

print('Array after sorting', a)