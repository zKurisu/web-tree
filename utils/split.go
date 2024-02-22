package utils

import (
	"log"
	"reflect"
	"strings"
)

// webtree add --tree=root/fir/sec,another
// Problem: "," in url
func Split2List(s string) []string {
	list := strings.Split(s, `,`)
	return RemoveEmp(list)
}

func SplitTreeLevel(s string) []string {
	levels := strings.Split(s, "/")
	return RemoveEmp(levels)
}

// func MergeList(master []string, slave []string) []string {
// 	for _, elem := range slave {
// 		master = append(master, elem)
// 	}
// 	return master
// }

func MergeList(master interface{}, slave interface{}) interface{} {
	masterType := reflect.TypeOf(master)
	slaveType := reflect.TypeOf(slave)
	if masterType.Kind() != reflect.Slice || slaveType.Kind() != reflect.Slice {
		log.Fatal("In function [MergeList], master or slave input is not a [Slice]")
	}

	masterElemType := masterType.Elem()
	slaveElemType := slaveType.Elem()

	if masterElemType != slaveElemType {
		log.Fatal("In function [MergeList], element type of master is not equal to slave's")
	}

	masterValue := reflect.ValueOf(master)
	slaveValue := reflect.ValueOf(slave)

	mergedList := reflect.MakeSlice(masterType, 0, masterValue.Len()+slaveValue.Len())

	for i := 0; i < masterValue.Len(); i++ {
		elem := masterValue.Index(i)
		mergedList = reflect.Append(mergedList, elem)
	}
	for i := 0; i < slaveValue.Len(); i++ {
		elem := slaveValue.Index(i)
		mergedList = reflect.Append(mergedList, elem)
	}

	return mergedList.Interface()
}

func RemoveEmp(list []string) []string {
	newList := []string{}
	for _, elem := range list {
		if len(elem) != 0 {
			newList = append(newList, elem)
		}
	}
	return newList
}

func List2String(list []string) string {
	var s string
	for _, elem := range list {
		s = s + elem + ","
	}
	return s
}
