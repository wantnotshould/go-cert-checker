// Package cmd
// Author: Perry He
// Created on: 2025-11-28 13:54:49
package cmd

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

// Execute 执行证书检查
func Execute() {
	domain := flag.String("domain", "", "请输入要检查的域名，例如 example.com")
	port := flag.Int("port", 443, "SSL 端口（默认: 443）")
	timeout := flag.Int("timeout", 5, "连接超时时间（秒）")
	flag.Parse()

	if *domain == "" {
		log.Fatal("❌ 错误：必须提供域名")
	}

	if *port < 1 || *port > 65535 {
		log.Fatalf("❌ 无效端口: %d。端口必须在 1-65535 之间", *port)
	}

	addr := fmt.Sprintf("%s:%d", *domain, *port)
	dialer := &net.Dialer{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", addr, &tls.Config{
		InsecureSkipVerify: true, // 忽略证书过期/未生效验证
	})
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			log.Fatalf("🌐 DNS 解析失败: %s", dnsErr)
		}
		var opErr *net.OpError
		if errors.As(err, &opErr) {
			log.Fatalf("🔌 连接失败: %s", opErr)
		}
		log.Fatalf("🔐 TLS 连接失败: %s", err)
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		log.Fatal("⚠️ 未找到证书")
	}

	cert := certs[0]
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now()

	fmt.Println("🌐 域名:", *domain)
	fmt.Println("🔑 证书 CN:", cert.Subject.CommonName)
	fmt.Println("📜 SANs:", cert.DNSNames)
	fmt.Println("🏛️ 签发机构:", cert.Issuer)
	fmt.Println("📅 生效时间:", cert.NotBefore.In(loc).Format("2006-01-02 15:04:05"))
	fmt.Println("📅 到期时间:", cert.NotAfter.In(loc).Format("2006-01-02 15:04:05"))
	fmt.Println("🆔 序列号:", cert.SerialNumber)

	remaining := int(time.Until(cert.NotAfter).Hours() / 24)
	fmt.Printf("⏳ 剩余天数: %d 天\n", remaining)

	// 判断证书状态
	if now.Before(cert.NotBefore) {
		fmt.Println("❌ 证书状态: 尚未生效")
	} else if now.After(cert.NotAfter) {
		fmt.Println("❌ 证书状态: 已过期")
	} else {
		fmt.Println("✅ 证书状态: 有效")
	}
}
