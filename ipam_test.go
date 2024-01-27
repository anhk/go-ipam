package ipam

import (
	"fmt"
	"testing"
)

func TestIPAM(t *testing.T) {
	ipam, _ := NewIPAM("192.168.5.0/24")
	fmt.Println(ipam.Alloc())
	a, _ := ipam.Alloc()
	fmt.Println(ipam.Alloc())
	ipam.Release(a)
	fmt.Println(ipam.Alloc())
}
