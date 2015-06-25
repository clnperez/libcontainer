// +build !arm
// +build ppc64,gccgo ppc64le,gccgo, !ppc64,!ppc64le
// +build ppc64,!gc ppc64le,!gc !ppc64,!ppc64le

package netlink

func ifrDataByte(b byte) int8 {
	return int8(b)
}
