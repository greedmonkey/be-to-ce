package main

import (
	"testing"
	"fmt"
)

func TestConvertLongMonthBEtoCE(t *testing.T) {
	input := "3 มีนาคม 1992"
	output := ConvertLongMonthBEtoCE(input)
	fmt.Println(output)
}

func TestConvertShortMonthBEtoCE(t *testing.T) {
	input := "3 เม.ย. 1992"
	output := ConvertShortMonthBEtoCE(input)
	fmt.Println(output)
}

func TestConvertLongDayBEtoCE(t *testing.T) {
	input := "จันทร์ เม.ย. 1992"
	output := ConvertLongDayBEtoCE(input)
	fmt.Println(output)
}

func TestConvertShortDayBEtoCE(t *testing.T) {
	input := "อา Apr 1992"
	output := ConvertShortDayBEtoCE(input)
	fmt.Println(output)
}

func TestConvertShortDayBEtoCE2(t *testing.T) {
	input := "อ Apr 1992"
	output := ConvertShortDayBEtoCE(input)
	fmt.Println(output)
}

func TestConvertLongYearBEtoCE(t *testing.T) {
	input:="3 มกราคา 2535"
	output := ConvertLongYearBEtoCE(input)
	fmt.Println(output)
}