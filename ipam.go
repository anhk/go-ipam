package ipam

import (
	"errors"
	"math/big"
	"net"

	"github.com/RoaringBitmap/roaring"
)

type IPAM struct {
	m *roaring.Bitmap

	base *big.Int
	max  uint32
}

func NewIPAM(cidr string) (*IPAM, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	ones, bits := ipNet.Mask.Size()
	ipam := &IPAM{m: roaring.NewBitmap()}
	ipam.base = big.NewInt(0).SetBytes(ipNet.IP)
	ipam.max = 1 << (bits - ones)
	return ipam, nil
}

func (ipam *IPAM) Alloc() (net.IP, error) {
	for i := uint32(2); i < ipam.max; i++ {
		if ok := ipam.m.CheckedAdd(i); ok {
			return big.NewInt(0).Add(ipam.base, big.NewInt(int64(i))).Bytes(), nil
		}
	}

	return nil, errors.New("unavailable")
}

func (ipam *IPAM) Release(ip net.IP) {
	offset := big.NewInt(0).Sub(big.NewInt(0).SetBytes(ip), ipam.base).Int64()
	ipam.m.CheckedRemove(uint32(offset))
}
