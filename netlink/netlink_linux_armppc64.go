// +build arm ppc64,!gccgo ppc64le,!gccgo

package netlink

func ifrDataByte(b byte) uint8 {
	return uint8(b)
}
