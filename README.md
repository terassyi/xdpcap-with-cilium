# xdpcap with cilium

This code is an example of xdpcap using [cilium/ebpf](https://github.com/cilium/ebpf).
[cloudflare/xdpcap](https://github.com/cloudflare/xdpcap) is a tcpdump like tool for XDP.
It can capture packets and action code of XDP.
Please see each repositories for detail.

## Usage
First, you must mount bpffs and install xdpcap.
To install xdpcap, Please see [cloudflare/xdpcap](https://github.com/cloudflare/xdpcap).
Please run this command.
**We cannot run this in network namespaces.**
```shell
$ sudo mount -t bpf none /sys/fs/bpf
```

After that, we can built and run program.
```shell
$ make
$ sudo ./xdpcap-with-cilium -iface <interface>
```

To capture to a pcap file and display captured packets, you can run commands below.
```shell
$ sudo xdpcap /sys/fs/bpf/xdpcap <pcap file> "filter rules"
```

```shell
$ sudo xdpcap /sys/fs/bpf/xdpcap - "filter rules" | sudo tcpdump -r -
```
