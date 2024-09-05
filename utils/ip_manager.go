package utils

import (
	"google.golang.org/grpc"
	"net"
)

func getLocalIP(conn *grpc.ClientConn, serviceName string) (string, error) {
	// Extract the underlying connection from grpc.ClientConn
	addr := conn.Target()

	// Create a temporary dial to get the local address and port
	tempConn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer func(tempConn net.Conn) {
		err := tempConn.Close()
		if err != nil {
			panic(err)
		}
	}(tempConn)

	// Retrieve the local port
	localAddr := tempConn.LocalAddr().(*net.TCPAddr)
	//fmt.Println(serviceName, "local IP:", localAddr.IP.String())
	return localAddr.IP.String(), nil
}

func getLocalPort(conn *grpc.ClientConn, serviceName string) (uint32, error) {
	// Extract the underlying connection from grpc.ClientConn
	addr := conn.Target()

	// Create a temporary dial to get the local address and port
	tempConn, err := net.Dial("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func(tempConn net.Conn) {
		err := tempConn.Close()
		if err != nil {
			panic(err)
		}
	}(tempConn)

	// Retrieve the local port
	localAddr := tempConn.LocalAddr().(*net.TCPAddr)
	//fmt.Println(serviceName, "local port:", localAddr.Port)
	return uint32(localAddr.Port), nil
}

func GetLocalIPPort(conn *grpc.ClientConn, serviceName string) (string, uint32, error) {
	ip, err := getLocalIP(conn, serviceName)
	if err != nil {
		return "", 0, err
	}

	port, err := getLocalPort(conn, serviceName)
	if err != nil {
		return "", 0, err
	}

	return ip, port, nil
}

func GetIpFromListener(lis net.Listener) (string, error) {
	// Extract the IP address of the listener
	localAddr := lis.Addr().(*net.TCPAddr)
	localIp := localAddr.IP.String()

	// Check if the IP is an unspecified address (e.g., "::" or "0.0.0.0")
	if localAddr.IP.IsUnspecified() {
		// Get all the system's IP addresses
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				localIp = ipNet.IP.String()
				break
			}
		}
	}

	return localIp, nil
}
