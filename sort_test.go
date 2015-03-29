package main

import (
    "testing"
    "reflect"
)

func Test(t *testing.T) {
    var tests = []struct{
        arr, want []int
    }{
        {
            []int{2, 6, 345, 54, 23, 34, 3, 3, 434, 432, 1},
            []int{1, 2, 3, 3, 6, 23, 34, 54, 345, 432, 434},
        },
    }

    for _, c := range tests {
        got := Mergesort(c.arr)
        if !reflect.DeepEqual(got, c.want) {
            t.Errorf("Mergesort(%v) == %v, wanted %v", c.arr, got, c.want)
        }
    }
}
