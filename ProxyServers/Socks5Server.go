package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

var (
	NO_AUTH       = []byte{0x05, 0x00}
	USERPASS_AUTH = []byte{0x05, 0x02}

	AUTH_SUCCESS = []byte{0x05, 0x00}
	AUTH_FAILED  = []byte{0x05, 0x01}

	CONNECT_SUCCESS = []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

//实现一个 Socks5 的代理服务器，可以通过该服务访问网络
//实现 Socks5 代理协议的服务端
//Socks4 只支持TCP，Socks5 支持TCP和UDP
/*
支持的验证方式：
X'00’ NO AUTHENTICATION REQUIRED（不需要验证）
X'01’ GSSAPI
X'02’ USERNAME/PASSWORD（用户名密码）
X'03’ to X'7F’ IANA ASSIGNED
X'80’ to X’FE’ RESERVED FOR PRIVATE METHODS
X’FF’ NO ACCEPTABLE METHODS（都不支持，没法连接了
*/
type Socks5Server struct {
	ServerListenAddr string
	serverListener   net.Listener
	AuthType         int
	UserName         string
	PassWord         string
}

func NewSocks5Server(serverListenAddr string, authType int, userName, passWord string) *Socks5Server {
	return &Socks5Server{
		ServerListenAddr: serverListenAddr,
		AuthType:         authType,
		UserName:         userName,
		PassWord:         passWord,
	}
}

func (_self *Socks5Server) Start() {
	var err error
	_self.serverListener, err = net.Listen("tcp", _self.ServerListenAddr)
	if err != nil {
		fmt.Println("Socks Listen err:", err)
		return
	}

	for {
		fmt.Println("Ready to Accept ...")
		clientConn, err := _self.serverListener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			break
		}

		//go _self.printHandler(clientConn)
		go _self.proxyHandler(clientConn)

	}

}

func (_self *Socks5Server) printHandler(conn net.Conn) {
	//打印测试
	for {
		dataBuf := make([]byte, 10)
		n, err := conn.Read(dataBuf)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " Read error: ", err)
			return
		}

		fmt.Println("接收到请求数据 ", conn.RemoteAddr().String(), "长度：", n, "dataBuf：", dataBuf)

	}
}

func (_self *Socks5Server) proxyHandler(conn net.Conn) {

	//协议文档：https://www.ietf.org/rfc/rfc1928.txt

	//VER：1个byte，代表 Socks 的版本，Socks5 默认为0x05
	//NMETHODS：1个byte，表示字段METHODS的长度
	//METHODS：目前是支持6种验证方式

	/*
			请求的报文：
		   +----+----------+----------+
		   |VER | NMETHODS | METHODS  |
		   +----+----------+----------+
		   | 1  |    1     | 1 to 255 |
		   +----+----------+----------+

		  o  X'00' NO AUTHENTICATION REQUIRED
		  o  X'01' GSSAPI
		  o  X'02' USERNAME/PASSWORD
		  o  X'03' to X'7F' IANA ASSIGNED
		  o  X'80' to X'FE' RESERVED FOR PRIVATE METHODS
		  o  X'FF' NO ACCEPTABLE METHODS
	*/
	//先读前面2个byte
	headBuf := make([]byte, 2)
	_, err := conn.Read(headBuf)

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), " Read error: ", err)
		return
	}

	//VER
	if headBuf[0] != 0x05 {
		fmt.Println("只支持Socks5代理")
		conn.Close()
		return
	}

	//NMETHODS
	nMethods := headBuf[1]
	//METHODS
	methods := make([]byte, nMethods)
	if n, err := conn.Read(methods); n != int(nMethods) || err != nil {
		fmt.Println("Get methods error", err)
		conn.Close()
		return
	}

	/*
		返回报文
		 +----+--------+
		 |VER | METHOD |
		 +----+--------+
		 | 1  |   1    |
		 +----+--------+
	*/
	if _self.AuthType == 0 {
		//回复2个byte，表示 VER=0x05 METHOD=0x00
		conn.Write(NO_AUTH)
	} else {
		//告诉客户端，需要用户名密码验证
		conn.Write(USERPASS_AUTH)

		//接收用户名密码信息 https://www.ietf.org/rfc/rfc1929.txt
		/*
			+----+------+----------+------+----------+
			|VER | ULEN |  UNAME   | PLEN |  PASSWD  |
			+----+------+----------+------+----------+
			| 1  |  1   | 1 to 255 |  1   | 1 to 255 |
			+----+------+----------+------+----------+
			注意：这里的VER 不是指socks版本
			The VER field contains the current version of the subnegotiation, which is X'01'.
		*/

		//ver := uint8(authBuf1[0])
		//uLen := uint8(authBuf1[1])
		//fmt.Println("Auth ver：", ver, "uLen：", uLen, "authBuf1：", authBuf1)

		authBuf := make([]byte, 520)
		n, err := conn.Read(authBuf)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " Read error: ", err)
			return
		}

		fmt.Println("读取到验证信息：", n, authBuf)

		ver := uint8(authBuf[0])
		uLen := uint8(authBuf[1])

		fmt.Println("Auth ver：", ver, "uLen：", uLen, "authBuf：", authBuf)

		uname := string(authBuf[2 : 2+uLen])
		pLen := uint8(authBuf[2+uLen])
		pass := string(authBuf[2+1+uLen : 2+1+uLen+pLen])

		//fmt.Println("Auth UserName：", uname, "PassWord：", pass, authBuf2[2:2+uLen], authBuf2[2+1+uLen:2+1+uLen+pLen])
		//校验用户名和密码
		if _self.UserName == uname && _self.PassWord == pass {
			fmt.Println("AUTH_SUCCESS")
			conn.Write(AUTH_SUCCESS)
		} else {
			//用户名密码不对，验证失败关闭连接
			conn.Write(AUTH_FAILED)
			conn.Close()
			fmt.Println("AUTH_FAILED，Closed conn")
			return
		}

	}

	/*
			解析代理要请求的连接信息
			+----+-----+-------+------+----------+----------+
			|VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
			+----+-----+-------+------+----------+----------+
			| 1  |  1  | X'00' |  1   | Variable |    2     |
			+----+-----+-------+------+----------+----------+

		  o  VER    protocol version: X'05'
		  o  CMD
			 o  CONNECT X'01'
			 o  BIND X'02'
			 o  UDP ASSOCIATE X'03'
		  o  RSV    RESERVED
		  o  ATYP   address type of following address
			 o  IP V4 address: X'01'
			 o  DOMAINNAME: X'03'
			 o  IP V6 address: X'04'
		  o  DST.ADDR       desired destination address
		  o  DST.PORT desired destination port in network octet order
	*/

	b := make([]byte, 32)
	n, err := conn.Read(b)
	fmt.Println(conn.RemoteAddr().String(), "len：", n, "reqBuf b：", b)

	//接收到HTTP请求，测试通过
	//127.0.0.1:7319 len： 17 reqBuf b： [5 1 0 3 10 104 105 99 111 100 101 46 116 111 112 1 187]

	//接收到该指令，是UDP请求
	//127.0.0.1:7203 len： 10 reqBuf b： [5 3 0 1 0 0 0 0 255 111]

	var host string
	switch b[3] {
	case 0x01: //IP V4
		host = net.IPv4(b[4], b[5], b[6], b[7]).String()
	case 0x03: //domain
		host = string(b[5 : n-2]) //b[4] length of domain
	case 0x04: //IP V6
		host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
	default:
		return
	}
	port := strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))
	targetAddr := net.JoinHostPort(host, port)
	fmt.Println("req to targetAddr：", targetAddr)

	network := "tcp"
	if b[1] == 0x03 {
		network = "udp"
	}
	//reqServer, err := net.Dial(network, targetAddr)
	reqServer, err := net.DialTimeout(network, targetAddr, 30*time.Second)
	if err != nil {
		fmt.Println("reqServer err:", err)
		return
	}
	conn.Write(CONNECT_SUCCESS)

	go func() {
		_, err := StdIOCopy(reqServer, conn)
		if err != nil {
			//logger.Errors("数据转发到目标端口异常：", err)
			conn.Close()
		}
	}()

	go func() {
		_, err := StdIOCopy(conn, reqServer)
		if err != nil {
			//logger.Errors("返回响应数据异常：", err)
			conn.Close()
		}
	}()

}

func (_self *Socks5Server) Stop() {

	_self.serverListener.Close()
	_self.serverListener = nil
}

func StdIOCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	//标准库的io复制
	return io.Copy(dst, src)
}
