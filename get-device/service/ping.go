package service

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

var count int

type Option struct {
	Address string
}

func Ping(ip string, m int) int {
	count = 0
	addr := &net.IPNet{
		IP:   net.ParseIP(ip).To4(),
		Mask: net.CIDRMask(m, 32),
	}

	out := ips(addr)

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	for _, v := range out {

		opt := Option{
			Address: v.String(),
		}
		ip, err := net.ResolveIPAddr("ip4", opt.Address)
		if err != nil {
			panic(err)
		}

		// c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
		// if err != nil {
		// 	panic(err)
		// }
		// defer c.Close()

		go try(c, ip.IP)
	}

	time.Sleep(time.Second)

	fmt.Println("現在の接続端末数：", count)
	return count

}

func try(c *icmp.PacketConn, ip net.IP) {
	now := time.Now().UnixMilli()
	result := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(result, now)

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: result,
		},
	}

	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		return
	}
	if _, err := c.WriteTo(msgBytes, &net.IPAddr{IP: ip}); err != nil {
		return
	}

	c.SetDeadline(time.Now().Add(time.Second))

	rb := make([]byte, 1500)
	n, _, _ := c.ReadFrom(rb)
	if err == nil {
		rm, err := icmp.ParseMessage(ipv4.ICMPTypeEcho.Protocol(), rb[:n])
		if err == nil && rm.Type == ipv4.ICMPTypeEchoReply {
			echo, ok := rm.Body.(*icmp.Echo)
			if ok {
				t, _ := binary.Varint(echo.Data)
				fmt.Println(ip, ":", time.Now().UnixMicro()-t*1000, " ms")
				// fmt.Printf("%d ms\n", time.Now().UnixMilli()-t)
				if time.Now().UnixMilli()-t < 1000 {
					count++
				}
			}
		}
	}
}
