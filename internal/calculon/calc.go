package calculon

import "net"

func DeriveBroadCast(ip net.IP, network *net.IPNet) net.IP {
	broadcast := net.IP(make([]byte, 4))
	mask := network.Mask

	for i := range ip {
		broadcast[i] = ip[i] | ^mask[i]
	}

	return broadcast
}

func DeriveWildcard(mask net.IPMask) net.IP {
	var result []byte

	for _, b := range mask {
		result = append(result, InvertByte(b))
	}

	if result != nil {
		return net.IPv4(result[0], result[1], result[2], result[3])
	}

	return nil
}

func HostMin(network *net.IPNet) net.IP {
	ip := network.IP
	b := ip[3]
	ip[3] = b + 1
	return ip
}

func HostMax(ip net.IP, network *net.IPNet) net.IP {
	broadcast := DeriveBroadCast(ip, network)
	b := broadcast[3]
	broadcast[3] = b - 1
	return broadcast
}

func InvertByte(input byte) byte {
	return 0xFF - input
}
