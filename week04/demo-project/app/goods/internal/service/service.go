package service

import (
	"geektime/app/goods/internal/biz"

	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewGoodsService)

func NewGoodsService(usecase *biz.GoodsUsecase) *GoodsService {
	return &GoodsService{
		usecase: usecase,
	}
}
