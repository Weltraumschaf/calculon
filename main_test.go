package main

import "testing"

func TestExtractIpAddress(t *testing.T) {
    ip := ExtractIpAddress("192.168.123.0/24")
    expected := "192.168.123.0"

    if ip != expected {
        t.Errorf("IP was incorrect, got: %s, want: %s.", ip, expected)
    }
}

func TestExtractIpAddress_emptySubnet(t *testing.T) {
    ip := ExtractIpAddress("192.168.123.0/")
    expected := "192.168.123.0"

    if ip != expected {
        t.Errorf("IP was incorrect, got: %s, want: %s.", ip, expected)
    }
}

func TestExtractIpAddress_noSubnet(t *testing.T) {
    ip := ExtractIpAddress("192.168.123.0")
    expected := "192.168.123.0"

    if ip != expected {
        t.Errorf("IP was incorrect, got: %s, want: %s.", ip, expected)
    }
}

func TestExtractIpAddress_emptyIp(t *testing.T) {
    ip := ExtractIpAddress("/24")
    expected := ""

    if ip != expected {
        t.Errorf("IP was incorrect, got: %s, want: %s.", ip, expected)
    }
}

func TestExtractIpAddress_empty(t *testing.T) {
    ip := ExtractIpAddress("")
    expected := ""

    if ip != expected {
        t.Errorf("IP was incorrect, got: %s, want: %s.", ip, expected)
    }
}

func TestExtractNetMask(t *testing.T) {
    subnet := ExtractNetMask("192.168.123.0/24")
    expected := "24"

    if subnet != expected {
        t.Errorf("Subnet was incorrect, got: %s, want: %s.", subnet, expected)
    }
}

func TestExtractNetMask_emptySubnet(t *testing.T) {
    subnet := ExtractNetMask("192.168.123.0/")
    expected := ""

    if subnet != expected {
        t.Errorf("Subnet was incorrect, got: %s, want: %s.", subnet, expected)
    }
}

func TestExtractNetMask_noSubnet(t *testing.T) {
    subnet := ExtractNetMask("192.168.123.0")
    expected := ""

    if subnet != expected {
        t.Errorf("Subnet was incorrect, got: %s, want: %s.", subnet, expected)
    }
}

func TestExtractNetMask_emptyIp(t *testing.T) {
    subnet := ExtractNetMask("/24")
    expected := "24"

    if subnet != expected {
        t.Errorf("Subnet was incorrect, got: %s, want: %s.", subnet, expected)
    }
}

func TestExtractNetMask_empty(t *testing.T) {
    subnet := ExtractNetMask("")
    expected := ""

    if subnet != expected {
        t.Errorf("Subnet was incorrect, got: %s, want: %s.", subnet, expected)
    }
}