package calculon

import (
	"encoding/binary"
	"net"
)

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
	min := ip2int(network.IP)
	min = min + 1
	return int2ip(min)
}

func HostMax(ip net.IP, network *net.IPNet) net.IP {
	max := ip2int(DeriveBroadCast(ip, network))
	max = max - 1
	return int2ip(max)
}

func HostsPerNetwork(ip net.IP, network *net.IPNet) uint32 {
	hostMin := ip2int(HostMin(network))
	hostMax := ip2int(HostMax(ip, network))

	return hostMax - hostMin + 1
}

func InvertByte(input byte) byte {
	return 0xFF - input
}


// https://gist.github.com/ammario/649d4c0da650162efd404af23e25b86b
func ip2int(input net.IP) uint32 {
	if len(input) == 16 {
		return binary.BigEndian.Uint32(input[12:16])
	}

	return binary.BigEndian.Uint32(input)
}

// https://gist.github.com/ammario/649d4c0da650162efd404af23e25b86b
func int2ip(input uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, input)
	return ip
}
