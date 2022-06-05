package biz

import (
	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewGoodsUsecase)

func NewGoodsUsecase(
	repo GoodsRepo,
	customer CustomerClient,
) *GoodsUsecase {
	return &GoodsUsecase{
		repo:     repo,
		customer: customer,
	}
}
