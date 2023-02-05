package mysql

import (
	"context"
	"dragonsss.cn/lbk_user/internal/data/user"
	"dragonsss.cn/lbk_user/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}

func (m *MemberDao) SaveMember(ctx context.Context, mem *user.User) error {
	return m.conn.Session(ctx).Create(mem).Error
}

func (m *MemberDao) GetMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("account=?", account).Count(&count).Error
	return count > 0, err
}
func (m *MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("email=?", email).Count(&count).Error
	return count > 0, err
}
func (m *MemberDao) GetMemberByAccountAndEmail(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("email=? or account=?", account, account).Count(&count).Error //数据库查询
	return count > 0, err
}
func (m *MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("mobile=?", mobile).Count(&count).Error
	return count > 0, err
}
