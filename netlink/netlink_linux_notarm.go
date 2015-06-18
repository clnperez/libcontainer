// +build !arm,!ppc64,!ppc64le gccgo,ppc64 gccgo,ppc64le

package netlink

func ifrDataByte(b byte) int8 {
	return int8(b)
}
