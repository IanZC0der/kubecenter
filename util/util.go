package util

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"time"
)

type ListItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func MapConverter(data []*ListItem) map[string]string {
	dataMap := make(map[string]string)
	for _, item := range data {
		dataMap[item.Key] = item.Value
	}
	return dataMap
}

func ListConverter(data map[string]string) []*ListItem {
	l := make([]*ListItem, 0)
	for k, v := range data {
		l = append(l, &ListItem{Key: k, Value: v})
	}
	return l
}

func ListConverterFromByte(data map[string][]byte) []*ListItem {
	l := make([]*ListItem, 0)
	for k, v := range data {
		l = append(l, &ListItem{Key: k, Value: string(v)})
	}
	return l
}

func FormatTime(creationTime int64) string {
	// compare the creation time with the current time, return a string denoting how much time has passed since creation of the cluster
	if creationTime == 0 {
		return "unknown"
	}
	timeNow := time.Now()
	timeStamp := time.Unix(creationTime, 0)
	days := int(timeNow.Sub(timeStamp).Hours() / 24)

	years := days / 365
	modDays := days % 365

	months := days / 30
	modDays = days % 30

	result := ""

	if years > 0 {
		result += fmt.Sprintf("%d year(s) ", years)
	}
	if months > 0 {
		result += fmt.Sprintf("%d month(s) ", months)
	}

	if modDays > 0 {
		result += fmt.Sprintf("%d day(s) ", modDays)
	}

	return result
}

func GenerateRGB(str string) string {
	val := HashString(str)
	r, g, b := HashToRBB(val)
	return strconv.Itoa(r) + "," + strconv.Itoa(g) + "," + strconv.Itoa(b)
}

func HashString(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

func HashToRBB(val uint32) (r, g, b int) {
	r = int(val & 0xFF) //last 8 bit as red
	g = int(val>>8) & 0xFF
	b = int(val>>16) & 0xFF
	return
}
