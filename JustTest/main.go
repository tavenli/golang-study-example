package main

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

// 定义一个简单的 DNS 记录映射
var records = map[string][]dns.RR{
	// A 记录
	"example.com.": []dns.RR{
		&dns.A{
			Hdr: dns.RR_Header{Name: "example.com.", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 3600},
			A:   net.ParseIP("192.0.2.1"),
		},
	},
	// AAAA 记录
	"ipv6.example.com.": []dns.RR{
		&dns.AAAA{
			Hdr:  dns.RR_Header{Name: "ipv6.example.com.", Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 3600},
			AAAA: net.ParseIP("2001:db8::1"),
		},
	},
	// CNAME 记录
	"www.example.com.": []dns.RR{
		&dns.CNAME{
			Hdr:    dns.RR_Header{Name: "www.example.com.", Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 3600},
			Target: "example.com.",
		},
	},
	// TXT 记录
	"txt.example.com.": []dns.RR{
		&dns.TXT{
			Hdr: dns.RR_Header{Name: "txt.example.com.", Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 3600},
			Txt: []string{"This is a TXT record."},
		},
	},
	// MX 记录
	"mail.example.com.": []dns.RR{
		&dns.MX{
			Hdr:        dns.RR_Header{Name: "mail.example.com.", Rrtype: dns.TypeMX, Class: dns.ClassINET, Ttl: 3600},
			Preference: 10,
			Mx:         "mailserver.example.com.",
		},
	},
}

// 处理 DNS 请求
func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		for _, q := range r.Question {
			switch q.Qtype {
			case dns.TypeA, dns.TypeAAAA, dns.TypeCNAME, dns.TypeTXT, dns.TypeMX:
				if recs, ok := records[q.Name]; ok {
					for _, rec := range recs {
						if rec.Header().Rrtype == q.Qtype {
							m.Answer = append(m.Answer, rec)
						}
					}
				}
			}
		}
	}

	w.WriteMsg(m)
}

func main() {
	// 注册 DNS 请求处理函数
	dns.HandleFunc(".", handleDNSRequest)

	// 启动 UDP 服务器
	udpServer := &dns.Server{Addr: ":53", Net: "udp"}
	go func() {
		log.Printf("Starting UDP DNS server on :53")
		if err := udpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start UDP server: %v", err)
		}
	}()

	// 启动 TCP 服务器
	tcpServer := &dns.Server{Addr: ":53", Net: "tcp"}
	log.Printf("Starting TCP DNS server on :53")
	if err := tcpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}
}
