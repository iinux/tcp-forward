package statistic

import "fmt"

var StaticData map[string]int

func init()  {
	StaticData = make(map[string]int)
}

func Add(ip string)  {
	_, ok := StaticData[ip]
	if !ok {
		StaticData[ip] = 1
	} else {
		StaticData[ip]++
	}
}

func Print()  {
	fmt.Println(StaticData)
}

func Get() map[string]int {
	return StaticData
}
