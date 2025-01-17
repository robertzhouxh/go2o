/**
 * Copyright 2015 @ at3.net.
 * name : parser.go
 * author : jarryliu
 * date : 2016-11-17 15:07
 * description :
 * history :
 */
package parser

import (
	"fmt"
	"github.com/ixre/go2o/core/domain/interface/item"
	"github.com/ixre/go2o/core/domain/interface/order"
	"github.com/ixre/go2o/core/domain/interface/pro_model"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/types"
)

func ItemDto(src *item.GoodsItem) *proto.SOldItem {
	it := &proto.SOldItem{
		ItemId:       src.Id,
		ProductId:    src.ProductId,
		PromFlag:     src.PromFlag,
		CatId:        src.CategoryId,
		VendorId:     src.VendorId,
		BrandId:      src.BrandId,
		ShopId:       src.ShopId,
		ShopCatId:    src.ShopCatId,
		ExpressTid:   src.ExpressTid,
		Title:        src.Title,
		ShortTitle:   src.ShortTitle,
		Code:         src.Code,
		Image:        src.Image,
		IsPresent:    src.IsPresent,
		PriceRange:   src.PriceRange,
		StockNum:     src.StockNum,
		SaleNum:      src.SaleNum,
		SkuNum:       src.SkuNum,
		SkuId:        src.SkuId,
		Cost:         src.Cost,
		Price:        src.Price,
		RetailPrice:  src.RetailPrice,
		Weight:       src.Weight,
		Bulk:         src.Bulk,
		ShelveState:  src.ShelveState,
		ReviewState:  src.ReviewState,
		ReviewRemark: src.ReviewRemark,
		SortNum:      src.SortNum,
		CreateTime:   src.CreateTime,
		UpdateTime:   src.UpdateTime,
		PromPrice:    src.PromPrice,
	}
	if src.SkuArray != nil {
		it.SkuArray = make([]*proto.SSku, len(src.SkuArray))
		for i, v := range src.SkuArray {
			it.SkuArray[i] = SkuDto(v)
		}
	}
	return it
}

func ItemDtoV2(src *item.GoodsItem) *proto.SUnifiedViewItem {
	it := &proto.SUnifiedViewItem{
		ItemId:    src.Id,
		ProductId: src.ProductId,
		//PromFlag:     src.PromFlag,
		CategoryId: int64(src.CategoryId),
		SkuId:      src.SkuId,
		VendorId:   src.VendorId,
		BrandId:    int64(src.BrandId),
		//ShopId:       src.ShopId,
		//ShopCatId:    src.ShopCatId,
		ExpressTid: int64(src.ExpressTid),
		Title:      src.Title,
		//ShortTitle:   src.ShortTitle,
		Code:  src.Code,
		Image: src.Image,
		//SaleNum:      src.SaleNum,
		//SkuNum:       src.SkuNum,
		//SkuId:        src.SkuId,
		//Cost:         float64(src.Cost),
		Price: src.Price,
		//IsPresent:    src.IsPresent,
		PriceRange: src.PriceRange,
		StockNum:   src.StockNum,
		//RetailPrice:  float64(src.RetailPrice),
		//Weight:       src.Weight,
		//Bulk:         src.Bulk,
		ShelveState: src.ShelveState,
		ReviewState: src.ReviewState,
		//ReviewRemark: src.ReviewRemark,
		//SortNum:      src.SortNum,
		//CreateTime:   src.CreateTime,
		UpdateTime: src.UpdateTime,
		Data:       nil,
	}
	if src.SkuArray != nil {
		it.SkuArray = make([]*proto.SSku, len(src.SkuArray))
		for i, v := range src.SkuArray {
			it.SkuArray[i] = SkuDto(v)
		}
	}
	return it
}

func ItemDataDto(src *item.GoodsItem) *proto.SItemDataResponse {
	it := &proto.SItemDataResponse{
		ItemId:    src.Id,
		ProductId: src.ProductId,
		//PromFlag:     src.PromFlag,
		CategoryId: int64(src.CategoryId),
		SkuId:      src.SkuId,
		VendorId:   src.VendorId,
		BrandId:    int64(src.BrandId),
		Thumbnail:  src.Image,
		//ShopId:       src.ShopId,
		//ShopCatId:    src.ShopCatId,
		ExpressTid: int64(src.ExpressTid),
		Title:      src.Title,
		//ShortTitle:   src.ShortTitle,
		Code: src.Code,
		//SaleNum:      src.SaleNum,
		//SkuNum:       src.SkuNum,
		//SkuId:        src.SkuId,
		//Cost:         float64(src.Cost),
		Price: src.Price,
		//IsPresent:    src.IsPresent,
		PriceRange: src.PriceRange,
		StockNum:   src.StockNum,
		//RetailPrice:  float64(src.RetailPrice),
		//Weight:       src.Weight,
		//Bulk:         src.Bulk,
		//ReviewRemark: src.ReviewRemark,
		//SortNum:      src.SortNum,
		//CreateTime:   src.CreateTime,
	}
	if src.SkuArray != nil {
		it.SkuArray = make([]*proto.SSku, len(src.SkuArray))
		for i, v := range src.SkuArray {
			it.SkuArray[i] = SkuDto(v)
		}
	}
	return it
}

func SkuArrayDto(src []*item.Sku) []*proto.SSku {
	var dst = make([]*proto.SSku, len(src))
	for i, v := range src {
		dst[i] = SkuDto(v)
	}
	return dst
}

func SpecOptionsDto(list promodel.SpecList) []*proto.SSpecOption {
	arr := make([]*proto.SSpecOption, len(list))
	s := func(l *promodel.Spec) []*proto.SSpecOptionItem {
		arr := make([]*proto.SSpecOptionItem, len(l.Items))
		for i, v := range l.Items {
			arr[i] = &proto.SSpecOptionItem{
				Value: fmt.Sprintf("%d:%d", l.Id, v.Id),
				Label: v.Value,
				Color: v.Color,
			}
		}
		return arr
	}
	for i, v := range list {
		arr[i] = &proto.SSpecOption{
			Name:  v.Name,
			Items: s(v),
		}
	}
	return arr
}

func SkuDto(src *item.Sku) *proto.SSku {
	return &proto.SSku{
		SkuId:       src.Id,
		ProductId:   src.ProductId,
		ItemId:      src.ItemId,
		Title:       src.Title,
		Image:       src.Image,
		SpecData:    src.SpecData,
		SpecWord:    src.SpecWord,
		Code:        src.Code,
		RetailPrice: src.RetailPrice,
		Price:       src.Price,
		Cost:        src.Cost,
		Weight:      src.Weight,
		Bulk:        src.Bulk,
		Stock:       src.Stock,
		SaleNum:     src.SaleNum,
	}
}

func Sku(src *proto.SSku) *item.Sku {
	return &item.Sku{
		Id:          src.SkuId,
		ProductId:   src.ProductId,
		ItemId:      src.ItemId,
		Title:       src.Title,
		Image:       src.Image,
		SpecData:    src.SpecData,
		SpecWord:    src.SpecWord,
		Code:        src.Code,
		RetailPrice: src.RetailPrice,
		Price:       src.Price,
		Cost:        src.Cost,
		Weight:      src.Weight,
		Bulk:        src.Bulk,
		Stock:       src.Stock,
		SaleNum:     src.SaleNum,
	}
}

func Order(src *proto.SSingleOrder) *order.ComplexOrder {
	o := &order.ComplexOrder{
		OrderId:        src.OrderId,
		OrderType:      src.OrderType,
		OrderNo:        src.OrderNo,
		BuyerId:        src.BuyerId,
		VendorId:       src.SellerId,
		ShopId:         src.ShopId,
		Subject:        src.Subject,
		ItemAmount:     src.ItemAmount,
		DiscountAmount: src.DiscountAmount,
		ExpressFee:     src.ExpressFee,
		PackageFee:     src.PackageFee,
		FinalAmount:    src.FinalAmount,

		Consignee: &order.ComplexConsignee{
			ConsigneeName:   src.Consignee.ConsigneeName,
			ConsigneePhone:  src.Consignee.ConsigneePhone,
			ShippingAddress: src.Consignee.ShippingAddress,
		},
		BuyerComment: src.BuyerComment,
		CreateTime:   src.SubmitTime,
		State:        src.State,
		Items:        make([]*order.ComplexItem, len(src.Items)),
		Data:         src.Data,
	}
	if src.Items != nil {
		for i, v := range src.Items {
			o.Items[i] = OrderItem(v)
		}
	}
	return o
}

func OrderItem(src *proto.SOrderItem) *order.ComplexItem {
	return &order.ComplexItem{
		ID:             src.Id,
		ItemId:         src.ItemId,
		SkuId:          src.SkuId,
		SnapshotId:     src.SnapshotId,
		Quantity:       src.Quantity,
		ReturnQuantity: src.ReturnQuantity,
		Amount:         src.Amount,
		FinalAmount:    src.FinalAmount,
		IsShipped:      int32(types.ElseInt(src.IsShipped, 1, 0)),
		Data:           src.Data,
	}
}

func SubOrderItemDto(src *order.SubOrderItem) *proto.SOrderItem {
	return &proto.SOrderItem{
		Id:                   src.ID,
		SnapshotId:           src.SnapshotId,
		SkuId:                src.SkuId,
		ItemId:               src.ItemId,
		ItemTitle:            "",
		Image:                "",
		Price:                0,
		FinalPrice:           0,
		Quantity:             src.Quantity,
		ReturnQuantity:       src.ReturnQuantity,
		Amount:               src.Amount,
		FinalAmount:          src.FinalAmount,
		IsShipped:            src.IsShipped == 1,
		Data:                 nil,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
}

func SubOrderDto(src *order.NormalSubOrder) *proto.SSingleOrder {
	o := &proto.SSingleOrder{
		OrderId:       src.ID,
		ParentOrderId: src.OrderId,
		OrderNo:       src.OrderNo,
		//OrderType:            src.,
		BuyerId:        src.BuyerId,
		SellerId:       src.VendorId,
		ShopId:         src.ShopId,
		Subject:        src.Subject,
		ItemAmount:     src.ItemAmount,
		DiscountAmount: src.DiscountAmount,
		DeductAmount:   0,
		AdjustAmount:   0,
		ExpressFee:     src.ExpressFee,
		PackageFee:     src.PackageFee,
		ProcedureFee:   0,
		//TotalAmount:          src.ItemAmount,
		FinalAmount: src.FinalAmount,
		Consignee: &proto.SConsigneeInfo{
			ConsigneeName:   "",
			ConsigneePhone:  "",
			ShippingAddress: "",
		},
		BuyerComment: src.BuyerComment,
		State:        src.State,
		SubmitTime:   src.CreateTime,
		Items:        make([]*proto.SOrderItem, len(src.Items)),
	}
	for i, v := range src.Items {
		o.Items[i] = SubOrderItemDto(v)
	}
	return o
}

func OrderItemDto(src *order.ComplexItem) *proto.SOrderItem {
	return &proto.SOrderItem{
		SnapshotId:           src.SnapshotId,
		SkuId:                src.SkuId,
		ItemId:               src.ItemId,
		ItemTitle:            src.ItemTitle,
		Image:                src.MainImage,
		Price:                src.Price,
		FinalPrice:           src.FinalPrice,
		Quantity:             src.Quantity,
		ReturnQuantity:       src.ReturnQuantity,
		Amount:               src.Amount,
		FinalAmount:          src.FinalAmount,
		IsShipped:            src.IsShipped == 1,
		Data:                 src.Data,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
}

func OrderDto(src *order.ComplexOrder) *proto.SSingleOrder {
	o := &proto.SSingleOrder{
		OrderId: src.OrderId,
		//SubOrderId:      src.SubOrderId,
		OrderType:      src.OrderType,
		OrderNo:        src.OrderNo,
		BuyerId:        src.BuyerId,
		SellerId:       src.VendorId,
		ShopId:         src.ShopId,
		Subject:        src.Subject,
		ItemAmount:     src.ItemAmount,
		DiscountAmount: src.DiscountAmount,
		ExpressFee:     src.ExpressFee,
		PackageFee:     src.PackageFee,
		FinalAmount:    src.FinalAmount,
		Consignee: &proto.SConsigneeInfo{
			ConsigneeName:   src.Consignee.ConsigneeName,
			ConsigneePhone:  src.Consignee.ConsigneePhone,
			ShippingAddress: src.Consignee.ShippingAddress,
		},
		BuyerComment: src.BuyerComment,
		SubmitTime:   src.CreateTime,
		State:        src.State,
		Items:        make([]*proto.SOrderItem, len(src.Items)),
		Data:         src.Data,
	}
	if src.Items != nil {
		for i, v := range src.Items {
			o.Items[i] = OrderItemDto(v)
		}
	}
	return o
}

// PrepareOrderDto 转换为预生成订单
func PrepareOrderDto(src *order.ComplexOrder) *proto.PrepareOrderResponse {
	o := &proto.PrepareOrderResponse{
		OrderType:      src.OrderType,
		ItemAmount:     src.ItemAmount,
		DiscountAmount: src.DiscountAmount,
		ExpressFee:     src.ExpressFee,
		PackageFee:     src.PackageFee,
		FinalFee:       src.FinalAmount,
		Consignee: &proto.SConsigneeInfo{
			ConsigneeName:   src.Consignee.ConsigneeName,
			ConsigneePhone:  src.Consignee.ConsigneePhone,
			ShippingAddress: src.Consignee.ShippingAddress,
		},
		BuyerComment: src.BuyerComment,
	}
	return o
}
