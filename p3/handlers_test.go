package p3

import (
    "testing"
)

func TestConvertStringToInt32(t *testing.T) {

    input := "12"
    var expect int32
    expect = 12
    result := ConvertToInt32(input)

    if result != expect {
        t.Fail()
    }
}

func TestProofOfWork(t *testing.T) {

    parentHash := "12312h12j3j12b3k12k3238424"
    mptRootHash := "123jn3446h3j12b3k12k323842"
   if ProofOfWorkTest(mptRootHash, parentHash) != true {
       t.Fail()
   }

}