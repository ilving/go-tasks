package layer1

import (
	"github.com/ilving/test/t3/dti"
)

// type ILayer2Object interface {
// 	GetIncomingData() string
// }

type ILayer2 interface {
	//L2Func() ILayer2Object
	//L2Func() string
	//L2Func() io.ReadCloser
	L2Func() dti.L2Param
}

type L1Usage struct {
	l2Link ILayer2
}

func NewLayer1(l2 ILayer2) *L1Usage {
	return &L1Usage{l2Link: l2}
}
