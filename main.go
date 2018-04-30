package main

import (
	"strconv"
	"strings"
	"regexp"
)

var thaiLongMonthNames = []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}
var thaiShortMonthNames = []string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."}
var thaiLongDayNames = []string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัสบดี", "ศุกร์", "เสาร์"}
var thaiShortDayNames = []string{"อา", "จ", "พฤ", "ศ", "ส", "อ", "พ"}
var longDayNames = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var shortDayNames = []string{"Sun", "Mon", "Thu", "Fri", "Sat", "Tue", "Wed"}
var shortMonthNames = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
var longMonthNames = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func ConvertLongMonthBEtoCE(inputDate string) string {
	for i, month := range thaiLongMonthNames {
		inputDate = strings.Replace(inputDate, month, longMonthNames[i], -1)
	}
	return inputDate
}

func ConvertShortMonthBEtoCE(inputDate string) string {
	for i, month := range thaiShortMonthNames {
		inputDate = strings.Replace(inputDate, month, shortMonthNames[i], -1)
	}
	return inputDate
}

func ConvertLongDayBEtoCE(inputDate string) string {
	for i, month := range thaiLongDayNames {
		inputDate = strings.Replace(inputDate, month, longDayNames[i], -1)
	}
	return inputDate
}

func ConvertShortDayBEtoCE(inputDate string) string {
	for i, month := range thaiShortDayNames {
		inputDate = strings.Replace(inputDate, month, shortDayNames[i], -1)
	}
	return inputDate
}

func ConvertLongYearBEtoCE(inputDate string) string {
	if rex, err := regexp.Compile(`(\d\d\d\d)`); err == nil {
		yearBE := rex.FindString(inputDate)
		yearBEint, _ := strconv.Atoi(yearBE)
		yearBEint = yearBEint - 543
		inputDate = strings.Replace(inputDate, yearBE, strconv.Itoa(yearBEint), 1)
	}
	return inputDate
}
