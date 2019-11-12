package calculon

type network interface {
    String() string
}

type networkConfig struct {
    ip [4]uint8
    netmask [4]uint8
}


func New() network {
    return &networkConfig{
        ip:      [4]byte{0x00, 0x00, 0x00, 0x00},
        netmask: [4]byte{0x00, 0x00, 0x00, 0x00},
    }
}

func (n networkConfig) String() string {
    panic("implement me")
}
