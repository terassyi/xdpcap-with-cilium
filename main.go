package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf"
	"github.com/vishvananda/netlink"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang XdpcapProg ./bpf/prog.c -- -I./bpf/header

func main() {
	var iface string
	flag.StringVar(&iface, "iface", "", "interface attached xdp program.")
	flag.Parse()

	link, err := netlink.LinkByName(iface)
	if err != nil {
		panic(err)
	}
	collect := &Collect{}
	spec, err := LoadXdpcapProg()
	if err != nil {
		panic(err)
	}
	if err := spec.LoadAndAssign(collect, nil); err != nil {
		panic(err)
	}

	if err := netlink.LinkSetXdpFd(link, collect.XdpProg.FD()); err != nil {
		panic(err)
	}
	tmpDir := "/sys/fs/bpf/xdpcap"
	if err != nil {
		panic(err)
	}
	if err := collect.XdpcapHook.Pin(tmpDir); err != nil {
		panic(err)
	}

	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Println("Xdpcap program running...")
	fmt.Println("Press CTRL+C to exit.")
	for {
		select {
		case <-ctrlC:
			fmt.Println("detaching xdp program...")
			if err := exit(link, tmpDir); err != nil {
				panic(err)
			}
			return
		}
	}
}


func exit(link netlink.Link, dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return netlink.LinkSetXdpFd(link, -1)
}

type Collect struct {
	XdpProg *ebpf.Program `ebpf:"prog"`
	XdpcapHook *ebpf.Map `ebpf:"xdpcap_hook"`
}
