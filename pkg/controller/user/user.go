package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sample-project/cmd/app/config"
	"sample-project/pkg/db"
	"sample-project/pkg/db/model"
	"sample-project/pkg/types"
)

type UserGetter interface {
	User() Interface
}

type Interface interface {
	Create(ctx context.Context, user *types.User) (*types.User, error)
	Update(ctx context.Context, userId int64, user *types.User) (*types.User, error)
	Delete(ctx context.Context, userId int64) (*types.User, error)
	Get(ctx context.Context, userId int64) (*types.User, error)
	List(ctx context.Context) ([]types.User, error)

	Login(ctx context.Context, user *types.User) (string, error)
	ChangePassword(ctx context.Context, userId int64, pwd *types.Password) (*types.User, error)
	ResetPassword(ctx context.Context, userId int64) (*types.User, error)

	CreateUserRole(ctx context.Context, roleId int64) error
	UpdateUserRole(ctx context.Context, roleId int64) error
	DeleteUserRole(ctx context.Context, roleId int64) error
	GetUserRole(ctx context.Context, roleId int64) error
	ListUserRole(ctx context.Context, roleId int64) error

	UserBindRole(ctx context.Context, roleId int64, menus []model.Menu) error
	UserUnBindRole(ctx context.Context, roleId int64, menus []model.Menu) error
}

type user struct {
	cc      config.Config
	factory *db.ShareDaoFactory
}

func (u *user) Create(ctx context.Context, user *types.User) (*types.User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	uu, err := u.factory.User.Create(ctx, &model.User{
		Name:        user.Name,
		Password:    string(encryptedPassword),
		Status:      1,
		Description: user.Description,
	})
	if err != nil {
		return nil, err
	}

	return &types.User{
		Id:          int64(uu.ID),
		Name:        uu.Name,
		Description: uu.Description,
	}, nil
}

func (u *user) Update(ctx context.Context, userId int64, user *types.User) (*types.User, error) {
	uu, err := u.factory.User.Update(ctx, userId, map[string]interface{}{"description": user.Description})
	if err != nil {
		return nil, err
	}

	return &types.User{
		Id:          int64(uu.ID),
		Name:        uu.Name,
		Description: uu.Description,
	}, nil
}

func (u *user) Delete(ctx context.Context, userId int64) (*types.User, error) {
	uu, err := u.factory.User.Delete(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &types.User{
		Id:          int64(uu.ID),
		Name:        uu.Name,
		Description: uu.Description,
	}, nil
}

func (u *user) Get(ctx context.Context, userId int64) (*types.User, error) {
	uu, err := u.factory.User.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &types.User{
		Id:          int64(uu.ID),
		Name:        uu.Name,
		Description: uu.Description,
	}, nil

}

func (u *user) List(ctx context.Context) ([]types.User, error) {
	var users []types.User

	uuu, err := u.factory.User.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, uu := range uuu {
		users = append(users, types.User{
			Id:       int64(uu.ID),
			Name:     uu.Name,
			Password: uu.Password,
		})
	}

	return users, nil
}

func (u *user) Login(ctx context.Context, user *types.User) (string, error) {
	uu, err := u.factory.User.FirstByWhere(ctx, "name = ?", user.Name)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(uu.Password), []byte(user.Password)); err != nil {
		return "", errors.New("the password is incorrect")
	}

	return "1111111111111", nil
}

func (u *user) ChangePassword(ctx context.Context, userId int64, pwd *types.Password) (*types.User, error) {
	uu, err := u.factory.User.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(uu.Password), []byte(pwd.OriginPassword)); err != nil {
		return nil, errors.New("the password is incorrect")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd.CurrentPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	uuu, err := u.factory.User.Update(ctx, userId, map[string]interface{}{"password": encryptedPassword})
	if err != nil {
		return nil, err
	}

	return &types.User{
		Id:          int64(uuu.ID),
		Name:        uuu.Name,
		Description: uuu.Description,
	}, nil
}

func (u *user) ResetPassword(ctx context.Context, userId int64) (*types.User, error) {
	_, err := u.factory.User.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	uu, err := u.factory.User.Update(ctx, userId, map[string]interface{}{"password": encryptedPassword})
	if err != nil {
		return nil, err
	}

	return &types.User{
		Id:          int64(uu.ID),
		Name:        uu.Name,
		Description: uu.Description,
	}, nil
}

func (u *user) CreateUserRole(ctx context.Context, roleId int64) error {

	return nil
}

func (u *user) UpdateUserRole(ctx context.Context, roleId int64) error {
	return nil
}

func (u *user) DeleteUserRole(ctx context.Context, roleId int64) error {

	return nil
}

func (u *user) GetUserRole(ctx context.Context, roleId int64) error {

	return nil
}

func (u *user) ListUserRole(ctx context.Context, roleId int64) error {

	return nil
}

func (u *user) UserBindRole(ctx context.Context, roleId int64, menus []model.Menu) error {
	return nil
}

func (u *user) UserUnBindRole(ctx context.Context, roleId int64, menus []model.Menu) error {
	return nil
}

func NewUser(cfg config.Config, f *db.ShareDaoFactory) *user {
	return &user{
		cc:      cfg,
		factory: f,
	}
}
