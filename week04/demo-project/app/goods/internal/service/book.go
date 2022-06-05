package service

import (
	"context"
	v1 "geektime/api/goods/v1"
	"geektime/app/goods/internal/biz"
)

type GoodsService struct {
	usecase *biz.GoodsUsecase
	v1.UnimplementedGoodsServiceServer
}

var _ v1.GoodsServiceServer = new(GoodsService)

func (s *GoodsService) FindGoods(ctx context.Context, in *v1.FindGoodsRequest) (*v1.GoodsReply, error) {
	res, err := s.usecase.FindOneGoods(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GoodsReply{
		Data: &v1.Goods{
			Id:   int64(res.ID),
			Name: res.Name,
			SaleInfo: &v1.SaleInfo{
				SaledAt:      res.SaleInfo.SaledAt.String(),
				CustomerId:   int64(res.SaleInfo.CustomerId),
				CustomerName: res.SaleInfo.CustomerName,
			},
		},
		Message: "Getting goods successfully",
	}, nil
}

func (s *GoodsService) SaleGoods(ctx context.Context, in *v1.SaleGoodsRequest) (*v1.GoodsReply, error) {
	res, err := s.usecase.SaleOneGoods(ctx, int(in.Id), int(in.CustomerId))
	if err != nil {
		return nil, err
	}
	return &v1.GoodsReply{
		Data: &v1.Goods{
			Id:   int64(res.ID),
			Name: res.Name,
			SaleInfo: &v1.SaleInfo{
				SaledAt:      res.SaleInfo.SaledAt.String(),
				CustomerId:   int64(res.SaleInfo.CustomerId),
				CustomerName: res.SaleInfo.CustomerName,
			},
		},
		Message: "Saling goods successfully",
	}, nil
}

func (s *GoodsService) NewGoods(ctx context.Context, in *v1.NewGoodsRequest) (*v1.GoodsReply, error) {
	res, err := s.usecase.NewGoods(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	var (
		saledAt      string
		customerId   int64
		customerName string
	)
	if res.SaleInfo != nil {
		saledAt = res.SaleInfo.SaledAt.String()
		customerId = int64(res.SaleInfo.CustomerId)
		customerName = res.SaleInfo.CustomerName
	}
	return &v1.GoodsReply{
		Data: &v1.Goods{
			Id:   int64(res.ID),
			Name: res.Name,
			SaleInfo: &v1.SaleInfo{
				SaledAt:      saledAt,
				CustomerId:   customerId,
				CustomerName: customerName,
			},
		},
		Message: "Putting a goods on the shelf successfully",
	}, nil
}

func (s *GoodsService) DeleteGoods(ctx context.Context, in *v1.DeleteGoodsRequest) (*v1.DeleteGoodsReply, error) {
	err := s.usecase.DeleteGoods(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteGoodsReply{
		Message: "Deleting a goods successfully",
	}, nil
}
