package util

import (
	"testing"
	"fmt"
)

func TestCheckWin(t *testing.T) {
	for _, tiles := range []string{
		"123456789m 12344s",     // [44s 123m 456m 789m 123s][一气通贯]
		"11122345678999s",       // [22s 111s 999s 345s 678s][九莲宝灯]
		"111234678m 11122z",     // [22z 111m 111z 234m 678m]
		"22334455m 234s 234p",   // [22m 345m 345m 234p 234s][一杯口], [55m 234m 234m 234p 234s][一杯口]
		"111222333m 234s 11z",   // [11z 111m 222m 333m 234s], [11z 123m 123m 123m 234s][一杯口]
		"112233m 112233p 11z",   // [11z 123m 123m 123p 123p][两杯口]   不是七对子，且不算一杯口
		"11223344556677z",       // [七对子]
		"119m 19p 19s 1234567z", // 国士无双自行判断
		"11m 345p",
		"1122m",
		"11m 112233p",
		"11m 123456789p", // [11m 123p 456p 789p][一气通贯]
		"11m 111p 111s",
		"111m 11p 111s",
		"111m 111p 11s",
	} {
		fmt.Print(tiles + " = ")
		tiles34 := MustStrToTiles34(tiles)
		results := DivideTiles34(tiles34)
		if len(results) == 0 {
			fmt.Println("[国士/未和牌]")
			continue
		}
		for i, result := range results {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(result)
		}
		fmt.Println()
	}
}

func BenchmarkCheckWin(b *testing.B) {
	// without divide result: 92.7 ns/op
	// with divide result: 235 ns/op
	tiles34 := MustStrToTiles34("123456789m 12344s")
	for i := 0; i < b.N; i++ {
		DivideTiles34(tiles34)
	}
}
