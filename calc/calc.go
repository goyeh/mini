package calc

import (
    "bytes"
    "fmt"
    "github.com/goyeh/mini/logic"
    "strconv"
)

const (
    resultFormat     = "%d.%s"
    errorResult      = "ERROR"
    leftParenthesis  = "("
    rightParenthesis = ")"
)

// ----------------------------------------------------
//              _       _  ______ _       _     _
//             (_)     | | |  _  (_)     (_)   | |
//   _ __  _ __ _ _ __ | |_| | | |___   ___  __| | ___
//  | '_ \| '__| | '_ \| __| | | | \ \ / / |/ _` |/ _ \
//  | |_) | |  | | | | | |_| |/ /| |\ V /| | (_| |  __/
//  | .__/|_|  |_|_| |_|\__|___/ |_| \_/ |_|\__,_|\___|
//  | |
//  |_|
// Divide divides two integers & Return the fraction in parenthesis as a string.
func PrintDivide(a int, b int) ( printDivide string) {
    if b == 0 {
        printDivide = errorResult
    }
    value := a / b
    remainder := a % b
    printDivide = fmt.Sprintf(resultFormat, value, Divide(remainder, b))
    return
}

// -------------------------------
//       _ _       _     _
//      | (_)     (_)   | |
//    __| |___   ___  __| | ___
//   / _` | \ \ / / |/ _` |/ _ \
//  | (_| | |\ V /| | (_| |  __/
//   \__,_|_| \_/ |_|\__,_|\___|
//  Core divide function
func Divide(a int, b int) (Divide string) {
    remainderIndexMap := make(map[int]int)
    values := make([]int, 0)

    value := 0
    remainder := a
    for !logic.Exists(remainderIndexMap, remainder) {
        remainderIndexMap[remainder] = len(values)

        remainder *= 10
        value = remainder / b
        remainder = remainder % b
        values = append(values, value)
    }

    index, _ := remainderIndexMap[remainder]

    var buffer bytes.Buffer
    for i := 0; i < index; i++ {
        buffer.WriteString(strconv.Itoa(values[i]))
    }

    buffer.WriteString(leftParenthesis)

    for i := index; i < len(values); i++ {
        buffer.WriteString(strconv.Itoa(values[i]))
    }

    buffer.WriteString(rightParenthesis)
    Divide = buffer.String()
    return
}
