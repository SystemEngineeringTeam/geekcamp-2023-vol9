/*
arpscan.go

Copyright (c) 2012 Google, Inc. All rights reserved.
Copyright (c) 2009-2011 Andreas Krennmair. All rights reserved.

Released under the "BSD 3-Clause "New" or "Revised" License".
*/

package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	i, _ := strconv.Atoi(os.Getenv("INTERFACE_INDEX"))
	m, _ := strconv.Atoi(os.Getenv("NETWORK_MASK"))
	ha, _ := net.ParseMAC(os.Getenv("INTERFACE_HARDWAREADDR"))
	iface := net.Interface{
		Index:        i,
		MTU:          1500,
		Name:         os.Getenv("INTERFACE_NAME"),
		HardwareAddr: ha,
	}
	ip := net.ParseIP(os.Getenv("NETWORK_ADDR"))
	ip4 := ip.To4()
	mask := net.CIDRMask(m, 32)
	if err := scan(&iface, ip4, mask); err != nil {
		log.Printf("interface %v: %v", iface.Name, err)
	}
}

// scan scans an individual interface's local network for machines using ARP requests/replies.
//
// scan loops forever, sending packets out regularly.  It returns an error if
// it's ever unable to write a packet.
// ARPを使ってネットワーク上のデバイスをスキャンする
func scan(iface *net.Interface, ip net.IP, mask net.IPMask) error {
	addr := &net.IPNet{
		IP:   ip,
		Mask: mask,
	}
	log.Printf("Using network range %v for interface %v", addr, iface.Name)

	// Open up a pcap handle for packet reads/writes.
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	// Start up a goroutine to read in packet data.
	stop := make(chan struct{})
	go readARP(handle, iface, stop)
	defer close(stop)
	// Write our scan packets out to the handle.
	writeARP(handle, iface, addr)
	time.Sleep(1 * time.Second)
	return nil
}

// readARP watches a handle for incoming ARP responses we might care about, and prints them.
//
// readARP loops until 'stop' is closed.
// ネットワーク上で発生するARPレスポンスを受信し、その応答を解析する
func readARP(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-stop:
			return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {
				// This is a packet I sent.
				continue
			}
			// Note:  we might get some packets here that aren't responses to ones we've sent,
			// if for example someone else sends US an ARP request.  Doesn't much matter, though...
			// all information is good information :)
			if arp.SourceProtAddress[len(arp.SourceProtAddress)-1] == 1 {
				continue
			}
			log.Printf("IP %v is at %v", net.IP(arp.SourceProtAddress), net.HardwareAddr(arp.SourceHwAddress))
		}
	}
}

// writeARP writes an ARP request for each address on our local network to the
// pcap handle.
// ARPリクエストパケットをネットワーク内の全てのデバイスに送信する
func writeARP(handle *pcap.Handle, iface *net.Interface, addr *net.IPNet) error {
	// Set up all the layers' fields we can.
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(addr.IP),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}
	// Set up buffer and options for serialization.
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	// Send one packet for every address.
	for _, ip := range ips(addr) {
		arp.DstProtAddress = []byte(ip)
		gopacket.SerializeLayers(buf, opts, &eth, &arp)
		if err := handle.WritePacketData(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

// ips is a simple and not very good method for getting all IPv4 addresses from a
// net.IPNet.  It returns all IPs it can over the channel it sends back, closing
// the channel when done.
// 指定されたネットワークのサブネット内に存在する全てのIPv4アドレスを生成する
func ips(n *net.IPNet) (out []net.IP) {
	num := binary.BigEndian.Uint32([]byte(n.IP))
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	network := num & mask
	broadcast := network | ^mask
	for network++; network < broadcast; network++ {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], network)
		out = append(out, net.IP(buf[:]))
	}
	return
}
