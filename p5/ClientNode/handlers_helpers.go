package ClientNode

import (
    "fmt"
    "strconv"
)

func ConvertToFloat32(value string) float32 {

    i, err := strconv.ParseFloat(value, 10)
    if err != nil {
        fmt.Println("Unable to convert string to Float32")
        panic(err)
    }
    return float32(i)
}

