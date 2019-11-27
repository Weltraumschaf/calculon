package calculon

import "net"

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

func InvertByte(input byte) byte {
    return 0xFF - input
}
