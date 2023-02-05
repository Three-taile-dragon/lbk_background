package mysql

import (
	"context"
	"dragonsss.cn/lbk_user/internal/data/user"
	"dragonsss.cn/lbk_user/internal/database"
	"dragonsss.cn/lbk_user/internal/database/gorms"
	"gorm.io/gorm"
)

type UserDao struct {
	conn *gorms.GormConn
}

func NewUserDao() *UserDao {
	return &UserDao{
		conn: gorms.New(),
	}
}

func (m *UserDao) SaveMember(conn database.DbConn, ctx context.Context, mem *user.User) error {
	m.conn = conn.(*gorms.GormConn) //使用事务操作
	return m.conn.Tx(ctx).Create(mem).Error
}

func (m *UserDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("email=?", email).Count(&count).Error //数据库查询
	return count > 0, err
}

func (m *UserDao) GetMemberByAccountAndEmail(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("email=? or account=?", account, account).Count(&count).Error //数据库查询
	return count > 0, err
}

func (m *UserDao) GetMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("account=?", account).Count(&count).Error //数据库查询
	return count > 0, err
}

func (m *UserDao) GetMemberByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("name=?", name).Count(&count).Error //数据库查询
	return count > 0, err
}

func (m *UserDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&user.User{}).Where("mobile=?", mobile).Count(&count).Error //数据库查询
	return count > 0, err
}
func (m *UserDao) FindMember(ctx context.Context, account string, pwd string) (*user.User, error) {
	var mem *user.User
	err := m.conn.Session(ctx).Where("account=? and password=?", account, pwd).First(&mem).Error
	if err == gorm.ErrRecordNotFound {
		//未查询到对应的信息
		return nil, nil
	}
	return mem, err
}
