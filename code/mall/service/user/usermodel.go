package user

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UserModel interface {
		Insert(data *User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		FindOneByMobile(mobile string) (*User, error)
		Update(data *User) error
		Delete(id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`     // 用户姓名
		Gender     int64     `db:"gender"`   // 用户性别
		Mobile     string    `db:"mobile"`   // 用户电话
		Password   string    `db:"password"` // 用户密码
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Insert(data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Gender, data.Mobile, data.Password)
	return ret, err
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByMobile(mobile string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRow(&resp, query, mobile)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(data *User) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Gender, data.Mobile, data.Password, data.Id)
	return err
}

func (m *defaultUserModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}