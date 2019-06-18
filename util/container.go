package util

import (
	"github.com/iinux/tcp-forward/iplist"
	"fmt"
)

const SYS_WHITELIST = "sys_whitelist"

type Container struct {
	Data map[string]interface{}
}

func (c *Container) Init() {
	c.Data = make(map[string]interface{})
}

func (c *Container) GetWhitelist() *iplist.WhiteList {
	fmt.Println(c.Data)
	return c.Data[SYS_WHITELIST].(*iplist.WhiteList)
}

func (c *Container) Put(key string, data interface{}) {
	c.Data[key] = data
	fmt.Println(c.Data)
}
