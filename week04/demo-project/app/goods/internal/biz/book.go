package biz

import (
	"context"
	"errors"
	"time"
)

type SaleInfo struct {
	SaledAt      time.Time
	CustomerId   int
	CustomerName string
}

type Goods struct {
	ID       int
	Name     string
	SaleInfo *SaleInfo
}

var ErrGoodsNotFound = errors.New("error: goods is not found")
var ErrGoodsSaled = errors.New("error: this goods has been saled")

type GoodsRepo interface {
	FindGoodsByID(id int) (*Goods, error)
	SaveGoods(goods *Goods) (*Goods, error)
	DeleteGoods(id int) error
}

type GoodsUsecase struct {
	repo     GoodsRepo
	customer CustomerClient
}

// FindOneGoods 按id查找商品
func (uc *GoodsUsecase) FindOneGoods(ctx context.Context, id int) (*Goods, error) {
	goods, err := uc.repo.FindGoodsByID(id)
	if err != nil {
		return nil, ErrGoodsNotFound
	}
	if !goods.SaleInfo.SaledAt.IsZero() || goods.SaleInfo.CustomerId > 0 {
		customer, err := uc.customer.FindCustomer(ctx, goods.SaleInfo.CustomerId)
		if err != nil {
			return goods, err
		}
		goods.SaleInfo.CustomerName = customer.Name
	}
	return goods, nil
}

// SaleOneGoods 出售一本商品
func (uc *GoodsUsecase) SaleOneGoods(ctx context.Context, id, customerId int) (*Goods, error) {
	goods, err := uc.FindOneGoods(ctx, id)
	if err != nil {
		return nil, err
	}
	customer, err := uc.customer.FindCustomer(ctx, customerId)
	if err != nil {
		return nil, err
	}
	if !goods.SaleInfo.SaledAt.IsZero() {
		return nil, ErrGoodsSaled
	}
	goods.SaleInfo.CustomerId = customer.Id
	goods.SaleInfo.CustomerName = customer.Name
	goods.SaleInfo.SaledAt = time.Now()
	return uc.repo.SaveGoods(goods)
}

// NewGoods 上架商品
func (uc *GoodsUsecase) NewGoods(ctx context.Context, name string) (*Goods, error) {
	newGoods := &Goods{
		Name: name,
	}
	return uc.repo.SaveGoods(newGoods)
}

// DeleteGoods 删除商品
func (uc *GoodsUsecase) DeleteGoods(ctx context.Context, id int) error {
	err := uc.repo.DeleteGoods(id)
	if err != nil {
		return ErrGoodsNotFound
	}
	return nil
}
