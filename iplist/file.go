package iplist

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

type File struct {
	ips []string
	fd *os.File
	filename string
}

func NewFile(filename string) *File {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	ret := File{
		filename:filename,
		fd:fi,
	}

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

func (s *File) GetIpSet() []string {
	return s.ips
}

func (s *File) AddIp(ip string) bool {
	wr := bufio.NewWriter(s.fd)
	wr.WriteString(ip)
	wr.Flush()
	return true
}
