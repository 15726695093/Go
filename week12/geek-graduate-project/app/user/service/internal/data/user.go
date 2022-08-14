package data

import (
	"clock-in/app/user/service/internal/biz"
	"clock-in/app/user/service/internal/data/ent"
	"clock-in/app/user/service/internal/data/ent/user"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.UserRepo = (*userRepo)(nil)

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetUserById(ctx context.Context, id []int64) ([]*biz.User, error) {
	rows, err := r.data.db.User.Query().
		Where(
			user.IDIn(int642Int(id)...),
			user.DeletedAtIsNil(),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return bulk2BizUser(rows), nil
}

func (r *userRepo) SearchUserByName(ctx context.Context, name string) ([]*biz.User, error) {
	rows, err := r.data.db.User.Query().
		Where(
			user.NameContains(name),
			user.DeletedAtIsNil(),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return bulk2BizUser(rows), nil
}

func (r *userRepo) SaveUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	if user.Id == 0 {
		return r.createUser(ctx, user)
	}
	return r.updateUser(ctx, user)
}

func (r *userRepo) GetUserByName(ctx context.Context, name string) (*biz.User, error) {
	user, err := r.data.db.User.Query().
		Where(
			user.NameEQ(name),
			user.DeletedAtIsNil(),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return toBizUser(user), nil
}

func (r *userRepo) createUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	newUser, err := r.data.db.User.Create().
		SetName(user.Name).
		SetPassword(user.Password).
		SetPhone(user.Phone).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toBizUser(newUser), nil
}

func (r *userRepo) updateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	updateUser, err := r.data.db.User.UpdateOneID(int(user.Id)).
		SetName(user.Name).
		SetPassword(user.Password).
		SetPhone(user.Phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toBizUser(updateUser), nil
}

func (r *userRepo) RemoveUser(ctx context.Context, id []int64) error {
	_, err := r.data.db.User.Update().
		Where(
			user.IDIn(int642Int(id)...),
		).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return err
}

func toBizUser(source *ent.User) *biz.User {
	return &biz.User{
		Id:       int64(source.ID),
		Name:     source.Name,
		Password: source.Password,
		Phone:    source.Phone,
	}
}

func bulk2BizUser(source []*ent.User) []*biz.User {
	user := make([]*biz.User, 0, len(source))
	for _, v := range source {
		user = append(user, toBizUser(v))
	}
	return user
}

func int642Int(source []int64) []int {
	intSli := make([]int, 0, len(source))
	for _, v := range source {
		intSli = append(intSli, int(v))
	}
	return intSli
}
