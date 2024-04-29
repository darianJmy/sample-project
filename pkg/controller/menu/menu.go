package menu

import (
	"context"

	"sample-project/pkg/db"
	"sample-project/pkg/db/model"
	"sample-project/pkg/types"
)

type MenuGetter interface {
	Menu() Interface
}

type Interface interface {
	Create(ctx context.Context, menu *types.Menu) (*types.Menu, error)
	Update(ctx context.Context, menuId int64, menu *types.Menu) (*types.Menu, error)
	Delete(ctx context.Context, menuId int64) (*types.Menu, error)
	Get(ctx context.Context, menuId int64) (*types.Menu, error)
	List(ctx context.Context) ([]types.Menu, error)
}

type menu struct {
	factory *db.ShareDaoFactory
}

func (m *menu) Create(ctx context.Context, menu *types.Menu) (*types.Menu, error) {
	mm, err := m.factory.Menu.Create(ctx, &model.Menu{
		URL:  menu.URL,
		Name: menu.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.Menu{
		Id:   int64(mm.ID),
		URL:  mm.URL,
		Name: mm.Name,
	}, nil
}

func (m *menu) Update(ctx context.Context, menuId int64, menu *types.Menu) (*types.Menu, error) {
	return nil, nil
}

func (m *menu) Delete(ctx context.Context, menuId int64) (*types.Menu, error) {
	mm, err := m.factory.Menu.Delete(ctx, menuId)
	if err != nil {
		return nil, err
	}

	return &types.Menu{
		Id:   int64(mm.ID),
		URL:  mm.URL,
		Name: mm.Name,
	}, nil
}

func (m *menu) Get(ctx context.Context, menuId int64) (*types.Menu, error) {
	mm, err := m.factory.Menu.Get(ctx, menuId)
	if err != nil {
		return nil, err
	}

	return &types.Menu{
		Id:   int64(mm.ID),
		URL:  mm.URL,
		Name: mm.Name,
	}, nil
}

func (m *menu) List(ctx context.Context) ([]types.Menu, error) {
	var menus []types.Menu

	mmm, err := m.factory.Menu.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, mm := range mmm {
		menus = append(menus, types.Menu{
			Id:   int64(mm.ID),
			URL:  mm.URL,
			Name: mm.Name,
		})
	}

	return menus, nil
}

func NewMenu(f *db.ShareDaoFactory) *menu {
	return &menu{f}
}
