package repo

import (
	"context"
	"dragonsss.cn/lbk_project/internal/data/menu"
)

type MenuRepo interface {
	FindMenus(ctx context.Context) ([]*menu.ProjectMenu, error)
}
