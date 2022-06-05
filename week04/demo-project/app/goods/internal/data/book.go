package data

import (
	"errors"
	"geektime/app/goods/internal/biz"
	"sync"
	"time"
)

var errNotFound = errors.New("error: Not found")
var errExist = errors.New("error: record existed")
var errNotExist = errors.New("error: record not existed")
var errDialFailed = errors.New("error: dail failed")

type row struct {
	id         int
	name       string
	customerId int
	saledAt    time.Time
}

type fakeDB map[int]*row

type goodsRepo struct {
	db fakeDB
	sync.RWMutex
}

// simulate connect to db
func (d *fakeDB) Dial(host, port, user, password string) error {
	time.Sleep(time.Millisecond)
	if host == "localhost" && port == "3306" && user == "root" && password == "root" {
		return nil
	}
	return errDialFailed
}

var _ biz.GoodsRepo = new(goodsRepo)

func (r *goodsRepo) FindGoodsByID(id int) (*biz.Goods, error) {
	r.RLock()
	defer r.RUnlock()
	goods, ok := r.db[id]
	if !ok {
		return nil, errNotFound
	}
	return &biz.Goods{
		ID:   goods.id,
		Name: goods.name,
		SaleInfo: &biz.SaleInfo{
			SaledAt:    goods.saledAt,
			CustomerId: goods.customerId,
			// TODO: customer's name
		},
	}, nil
}

func (r *goodsRepo) SaveGoods(goods *biz.Goods) (*biz.Goods, error) {
	r.Lock()
	defer r.Unlock()
	if goods.ID == 0 {
		id := len(r.db) + 1
		goods.ID = id
		return r.createGoods(goods)
	}
	return r.updateGoods(goods)
}

func (r *goodsRepo) createGoods(goods *biz.Goods) (*biz.Goods, error) {
	if _, ok := r.db[goods.ID]; ok {
		return nil, errExist
	}
	r.db[goods.ID] = &row{
		id:   goods.ID,
		name: goods.Name,
	}
	return goods, nil
}

func (r *goodsRepo) updateGoods(goods *biz.Goods) (*biz.Goods, error) {
	if _, ok := r.db[goods.ID]; !ok {
		return nil, errNotExist
	}
	r.db[goods.ID] = &row{
		id:         goods.ID,
		name:       goods.Name,
		saledAt:    goods.SaleInfo.SaledAt,
		customerId: goods.SaleInfo.CustomerId,
	}
	return goods, nil
}

func (r *goodsRepo) DeleteGoods(id int) error {
	_, ok := r.db[id]
	if !ok {
		return errNotFound
	}
	delete(r.db, id)
	return nil
}
