package arr
// Set of array functions to enhance array operations to more functionality.


import (
    "math/rand"
)

// shuffle array using a random source
func Shuffle(array []interface{}, source rand.Source) {
    random := rand.New(source)
    for i := len(array) - 1; i > 0; i-- {
        j := random.Intn(i + 1)
        array[i], array[j] = array[j], array[i]
    }
}

// bubble sort the array Low to High
func BubbleSort(items []int) {
    var (
        n = len(items)
        sorted = false
    )
    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
                items[i+1], items[i] = items[i], items[i+1]
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
    }
}


func LargestValue(items []int) ( Largest int ){
    for loop := 1; loop < len(items); loop++ {
        if Largest < items[loop] {Largest = items[loop]}
    }
    return
}

func LargestElement(items []int) ( Element int ){
    var Largest = 0
    for loop := 1; loop < len(items); loop++ {
        if Largest < items[loop] {Largest = items[loop]; Element = loop}
    }
    return
}