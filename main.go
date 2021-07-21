package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	ds = flag.String("ds", "", "ds server ip")
)

func main() {
	flag.Parse()
	if *ds == "" {
		log.Fatal("ds server ip is empty")
	}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range containers {
		go forward(fmt.Sprintf("%d", c.Ports[0].PublicPort), *ds)
	}
	<-ctx.Done()
}

func forward(port, dsAddr string) error {
	log.Println("forward: ", fmt.Sprintf("%s:%s", dsAddr, port))
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", port))
	if err != nil {
		log.Fatal(err.Error())
	}
	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err.Error())
	}
	ds, err := net.Dial("tcp", fmt.Sprintf("%s:%s", dsAddr, port))
	if err != nil {
		log.Fatal(err.Error())
	}
	go io.Copy(conn, ds)
	_, err = io.Copy(ds, conn)
	return err
}
