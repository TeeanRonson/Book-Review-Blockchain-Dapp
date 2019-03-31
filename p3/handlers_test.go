package p3

import "testing"

func TestConvertStringToInt32(t *testing.T) {

    input := "12"
    var expect int32
    expect = 12
    result := convertToInt32(input)

    if result != expect {
        t.Fail()
    }
}