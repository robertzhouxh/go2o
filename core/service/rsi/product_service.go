package rsi

import (
	"go2o/core/domain/interface/pro_model"
	"go2o/core/service/thrift/idl/gen-go/define"
	"go2o/core/service/thrift/parser"
)

// 产品服务
type productService struct {
	pmRep promodel.IProModelRepo
}

func NewProService(pmRep promodel.IProModelRepo) *productService {
	return &productService{
		pmRep: pmRep,
	}
}

// Get 产品品牌
func (p *productService) GetProBrand_(id int32) *promodel.ProBrand {
	return p.pmRep.BrandService().Get(id)
}

// Save 产品品牌
func (p *productService) SaveProBrand_(v *promodel.ProBrand) (*define.Result_, error) {
	id, err := p.pmRep.BrandService().SaveBrand(v)
	r := parser.Result(err)
	r.ID = id
	return r, nil
}

// Delete 产品品牌
func (p *productService) DeleteProBrand_(id int32) (*define.Result_, error) {
	err := p.pmRep.BrandService().DeleteBrand(id)
	return parser.Result(err), nil
}
