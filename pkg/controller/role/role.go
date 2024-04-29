package role

import (
	"context"

	"sample-project/pkg/db"
	"sample-project/pkg/db/model"
	"sample-project/pkg/types"
)

type RoleGetter interface {
	Role() Interface
}

type Interface interface {
	Create(ctx context.Context, role *types.Role) (*types.Role, error)
	Update(ctx context.Context, roleId int64, role *types.Role) (*types.Role, error)
	Delete(ctx context.Context, roleId int64) (*types.Role, error)
	Get(ctx context.Context, roleId int64) (*types.Role, error)
	List(ctx context.Context) ([]types.Role, error)

	CreateRoleMenu(ctx context.Context, roleId int64) error
	UpdateRoleMenu(ctx context.Context, roleId int64) error
	DeleteRoleMenu(ctx context.Context, roleId int64) error
	GetRoleMenu(ctx context.Context, roleId int64) error
	ListRoleMenu(ctx context.Context) error

	RoleBindMenu(ctx context.Context, roleId int64, menus []model.Menu) error
	RoleUnBindMenu(ctx context.Context, roleId int64, menus []model.Menu) error
}

type role struct {
	factory *db.ShareDaoFactory
}

func (r *role) Create(ctx context.Context, role *types.Role) (*types.Role, error) {
	rr, err := r.factory.Role.Create(ctx, &model.Role{
		Name: role.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.Role{
		Id:   int64(rr.ID),
		Name: rr.Name,
	}, nil
}

func (r *role) Update(ctx context.Context, roleId int64, role *types.Role) (*types.Role, error) {
	return nil, nil
}

func (r *role) Delete(ctx context.Context, roleId int64) (*types.Role, error) {
	rr, err := r.factory.Role.Delete(ctx, roleId)
	if err != nil {
		return nil, err
	}

	return &types.Role{
		Id:   int64(rr.ID),
		Name: rr.Name,
	}, nil
}

func (r *role) Get(ctx context.Context, roleId int64) (*types.Role, error) {
	rr, err := r.factory.Role.Get(ctx, roleId)
	if err != nil {
		return nil, err
	}

	return &types.Role{
		Id:   int64(rr.ID),
		Name: rr.Name,
	}, nil
}

func (r *role) List(ctx context.Context) ([]types.Role, error) {
	var roles []types.Role

	rrr, err := r.factory.Role.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, rr := range rrr {
		roles = append(roles, types.Role{
			Id:   int64(rr.ID),
			Name: rr.Name,
		})
	}

	return roles, nil
}

func (r *role) CreateRoleMenu(ctx context.Context, roleId int64) error {
	return nil
}

func (r *role) UpdateRoleMenu(ctx context.Context, roleId int64) error {
	return nil
}

func (r *role) DeleteRoleMenu(ctx context.Context, roleId int64) error {
	return nil
}

func (r *role) GetRoleMenu(ctx context.Context, roleId int64) error {
	return nil
}

func (r *role) ListRoleMenu(ctx context.Context) error {
	return nil
}

func (r *role) RoleBindMenu(ctx context.Context, roleId int64, menus []model.Menu) error {
	return nil
}

func (r *role) RoleUnBindMenu(ctx context.Context, roleId int64, menus []model.Menu) error {
	return nil
}

func NewRole(f *db.ShareDaoFactory) *role {
	return &role{
		f,
	}
}
