# Calculon

[![Build Status](https://travis-ci.com/Weltraumschaf/calculon.svg?branch=master)](https://travis-ci.com/Weltraumschaf/calculon)

This is a reimplementation of the [IPv4 calculator][ipcalc] in [go-lang][go-lang] as CLI tool.

For a given IP and netmask in CIDR format (eg. `192.168.123.5/24`) it will print you some information:

```text
Address:   192.168.123.5         11000000.10101000.01111011.00000101
Netmask:   255.255.255.0 = 24    11111111.11111111.11111111.00000000
Wildcard:  0.0.0.255             00000000.00000000.00000000.11111111
=>
Network:   192.168.123.0/24      11000000.10101000.01111011.00000000
Broadcast: 192.168.123.255       11000000.10101000.01111011.11111111
HostMin:   192.168.123.1         11000000.10101000.01111011.00000001
HostMax:   192.168.123.254       11000000.10101000.01111011.11111110
Hosts/Net: 254
```

## Status of the Project

This is only for my personal training to get used to [go-lang][go-lang]. You can use this project, but you should not rely on it or any maintenance!

## Build and Run

```bash
make clean && make
./bin/calculon 192.168.123.5/24
```

## Links

- https://itnext.io/how-to-create-your-own-cli-with-golang-3c50727ac608
- [Exploring the landscape of Go testing frameworks](https://bmuschko.com/blog/go-testing-frameworks/)
    - [Testify](https://godoc.org/github.com/stretchr/testify)
- http://networkbit.ch/golang-ip-address-manipulation/
- https://github.com/golang-standards/project-layout

[ipcalc]:   http://jodies.de/ipcalc
[go-lang]:  https://golang.org/