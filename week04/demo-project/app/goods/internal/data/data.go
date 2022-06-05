package data

import (
	"geektime/app/goods/internal/biz"
	"geektime/app/goods/internal/conf"

	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewBD, NewGoodsRepo, NewCustomerClient)

func NewBD(conf *conf.ConfDB) fakeDB {
	db := make(fakeDB)
	err := db.Dial(
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
	)
	if err != nil {
		panic(err)
	}
	return db
}

func NewGoodsRepo(db fakeDB) biz.GoodsRepo {
	return &goodsRepo{
		db: db,
	}
}

func NewCustomerClient(conf *conf.Customer) biz.CustomerClient {
	return &CustomerClient{
		addr: conf.Addr,
	}
}
