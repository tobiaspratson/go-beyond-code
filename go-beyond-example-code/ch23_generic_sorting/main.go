package main

import "fmt"

// Constraint for types that support ordering
type Ordered interface {
    ~int | ~float64 | ~string
}

// Generic bubble sort - O(n²) time complexity
func BubbleSort[T Ordered](slice []T) []T {
    result := make([]T, len(slice))
    copy(result, slice)
    
    for i := 0; i < len(result); i++ {
        for j := 0; j < len(result)-1-i; j++ {
            if result[j] > result[j+1] {
                result[j], result[j+1] = result[j+1], result[j]
            }
        }
    }
    
    return result
}

// Generic quick sort - O(n log n) average time complexity
func QuickSort[T Ordered](slice []T) []T {
    if len(slice) <= 1 {
        return slice
    }
    
    pivot := slice[0]
    var left, right []T
    
    for _, v := range slice[1:] {
        if v < pivot {
            left = append(left, v)
        } else {
            right = append(right, v)
        }
    }
    
    left = QuickSort(left)
    right = QuickSort(right)
    
    return append(append(left, pivot), right...)
}

// Generic merge sort - O(n log n) time complexity
func MergeSort[T Ordered](slice []T) []T {
    if len(slice) <= 1 {
        return slice
    }
    
    mid := len(slice) / 2
    left := MergeSort(slice[:mid])
    right := MergeSort(slice[mid:])
    
    return merge(left, right)
}

// Helper function for merge sort
func merge[T Ordered](left, right []T) []T {
    result := make([]T, 0, len(left)+len(right))
    i, j := 0, 0
    
    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    
    return result
}

// Generic binary search - O(log n) time complexity
func BinarySearch[T Ordered](slice []T, target T) int {
    left, right := 0, len(slice)-1
    
    for left <= right {
        mid := (left + right) / 2
        if slice[mid] == target {
            return mid
        } else if slice[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}

// Generic insertion sort - O(n²) time complexity
func InsertionSort[T Ordered](slice []T) []T {
    result := make([]T, len(slice))
    copy(result, slice)
    
    for i := 1; i < len(result); i++ {
        key := result[i]
        j := i - 1
        
        for j >= 0 && result[j] > key {
            result[j+1] = result[j]
            j--
        }
        result[j+1] = key
    }
    
    return result
}

func main() {
    // Test sorting with integers
    intSlice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
    fmt.Printf("Original: %v\n", intSlice)
    
    bubbleSorted := BubbleSort(intSlice)
    fmt.Printf("Bubble sort: %v\n", bubbleSorted)
    
    quickSorted := QuickSort(intSlice)
    fmt.Printf("Quick sort: %v\n", quickSorted)
    
    mergeSorted := MergeSort(intSlice)
    fmt.Printf("Merge sort: %v\n", mergeSorted)
    
    insertionSorted := InsertionSort(intSlice)
    fmt.Printf("Insertion sort: %v\n", insertionSorted)
    
    // Test binary search
    index := BinarySearch(quickSorted, 8)
    fmt.Printf("Binary search for 8: %d\n", index)
    
    index = BinarySearch(quickSorted, 6)
    fmt.Printf("Binary search for 6: %d\n", index)
    
    // Test sorting with strings
    stringSlice := []string{"banana", "apple", "cherry", "date", "elderberry"}
    fmt.Printf("Original: %v\n", stringSlice)
    
    sorted := QuickSort(stringSlice)
    fmt.Printf("Sorted: %v\n", sorted)
    
    // Test binary search with strings
    index = BinarySearch(sorted, "cherry")
    fmt.Printf("Binary search for 'cherry': %d\n", index)
    
    index = BinarySearch(sorted, "grape")
    fmt.Printf("Binary search for 'grape': %d\n", index)
    
    // Test with floats
    floatSlice := []float64{3.14, 2.71, 1.41, 0.57, 1.73}
    fmt.Printf("Original floats: %v\n", floatSlice)
    
    floatSorted := QuickSort(floatSlice)
    fmt.Printf("Sorted floats: %v\n", floatSorted)
}