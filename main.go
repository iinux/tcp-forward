package main

import (
	"io"
	"log"
	"net"
	"fmt"

	"github.com/iinux/tcp-forward/server"
	"github.com/iinux/tcp-forward/conf"
	"github.com/iinux/tcp-forward/statistic"
	"strings"
	"github.com/iinux/tcp-forward/iplist"
	"github.com/iinux/tcp-forward/util"
)

func main() {
	conf.Container.Init()
	wl := iplist.NewWhiteList(iplist.NewFile(conf.WhitelistFile))
	conf.Container.Put(util.SYS_WHITELIST, wl)

	go server.HttpServer()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":"+conf.ListenPort)
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		remoteAddr := client.RemoteAddr().String()
		remoteHost := strings.Split(remoteAddr, ":")[0]

		if !wl.Allow(remoteHost) {
			fmt.Println(remoteAddr, " be denied")
			client.Close()
			continue
		}

		fmt.Println(remoteAddr, "be allowed")
		statistic.Add(remoteHost)
		if err != nil {
			log.Panic(err)
		}

		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	serverConn, err := net.Dial("tcp", net.JoinHostPort(conf.RemoteHost, conf.RemotePort))
	if err != nil {
		log.Println(err)
		return
	}
	defer serverConn.Close()

	go io.Copy(serverConn, client)
	io.Copy(client, serverConn)
}
