package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func Probe_Demo1_main() {
	//https://github.com/shirou/gopsutil
	//https://pkg.go.dev/github.com/shirou/gopsutil/v3

	//系统探针

	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	fmt.Println("==========================")
	fmt.Println("host.Info")
	hostInfo, err := host.Info()
	//hostInfo.Hostname
	//hostInfo.Platform
	//hostInfo.OS
	fmt.Println(hostInfo, err)
	fmt.Println("==========================")

	fmt.Println("cpu.Info")
	cpuInfo, err := cpu.Info()
	fmt.Println(cpuInfo, err)
	fmt.Println("==========================")

	fmt.Println("net.Info")
	//网卡信息
	netInfos, err := net.Interfaces()
	fmt.Println(netInfos, err)
	fmt.Println("==========================")

	fmt.Println("disk.Info")
	diskInfos, err := disk.Partitions(true)
	fmt.Println(diskInfos, err)
	fmt.Println("==========================")

}
