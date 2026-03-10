package resolver

import (
	"fmt"
	"net/netip"
	"os"
)

func Dns64Convert(address string, dns64Prefix string, start int) string {
	addr, err := netip.ParseAddr(address)

	if err != nil {
		return address
	}

	if addr.Is6() {
		return address
	}

	dns64PrefixAddr, err := netip.ParseAddr(dns64Prefix)
	if err != nil {
		fmt.Errorf("dns64 prefix parse error")
		os.Exit(1)
	}
	v4Slice := addr.As4()
	v6Slice := dns64PrefixAddr.As16()

	var index = start
	v6Slice[index] = v4Slice[0]

	index += 1
	v6Slice[index] = v4Slice[1]
	index += 1
	v6Slice[index] = v4Slice[2]
	index += 1
	v6Slice[index] = v4Slice[3]
	targetIp, ok := netip.AddrFromSlice(v6Slice[:])
	if !ok {
		panic(fmt.Errorf("dns64 prefix convert error"))
	}
	return targetIp.String()
}
