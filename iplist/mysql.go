package iplist

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

type Mysql struct {
	ips []string
}

func NewMysql(filename string) *File {
	ret := File{}
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		ip := string(a)
		fmt.Println(ip)
		ret.ips = append(ret.ips, ip)
	}
	return &ret
}

func (s *Mysql) GetIpSet() []string {
	return s.ips
}
