package repo

import (
	"context"
	"dragonsss.cn/lbk_user/internal/data/user"
	"dragonsss.cn/lbk_user/internal/database"
)

type UserRepo interface {
	SaveMember(conn database.DbConn, ctx context.Context, mem *user.User) error
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByAccountAndEmail(ctx context.Context, account string) (bool, error)
	GetMemberByName(ctx context.Context, name string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	FindMember(ctx context.Context, account string, pwd string) (mem *user.User, err error)
}

type MemberRepo interface {
	SaveMember(ctx context.Context, member *user.User) error
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
}
