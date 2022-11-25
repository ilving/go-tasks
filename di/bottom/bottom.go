package layer2

import "github.com/ilving/test/di/dti"

type L2Param struct{}

func (l2p *L2Param) GetIncomingData() string { return "test" }

type L2 struct{}

// func (l2 *L2) L2Func() *L2Param { return &L2Param{} }
// func (l2 *L2) L2Func() *L2Param { return &L2Param{} }
// func (l2 *L2) L2Func() string { return "&L2Param{}" }
// func (l2 *L2) L2Func() io.ReadCloser { return nil }
func (l2 *L2) L2Func() dti.L2Param { return dti.L2Param(&L2Param{}) }

func NewL2() *L2 {
	return &L2{}
}
