// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
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
	sysApiFieldNames          = builder.RawFieldNames(&SysApi{})
	sysApiRows                = strings.Join(sysApiFieldNames, ",")
	sysApiRowsExpectAutoSet   = strings.Join(stringx.Remove(sysApiFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), ",")
	sysApiRowsWithPlaceHolder = strings.Join(stringx.Remove(sysApiFieldNames, "`id`", "`updatedTime`", "`deletedTime`", "`createdTime`"), "=?,") + "=?"
)

type (
	sysApiModel interface {
		Insert(ctx context.Context, data *SysApi) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysApi, error)
		FindOneByRoute(ctx context.Context, route string) (*SysApi, error)
		Update(ctx context.Context, data *SysApi) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysApiModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysApi struct {
		Id           int64        `db:"id"`           // 编号
		Route        string       `db:"route"`        // 路由
		Method       string       `db:"method"`       // 请求方式
		Name         string       `db:"name"`         // 请求名称
		BusinessType int64        `db:"businessType"` // 业务类型（1新增 2修改 3删除 4查询 5其它）
		Group        string       `db:"group"`        // 接口组
		Desc         string       `db:"desc"`         // 备注
		CreatedTime  time.Time    `db:"createdTime"`  // 创建时间
		UpdatedTime  time.Time    `db:"updatedTime"`  // 更新时间
		DeletedTime  sql.NullTime `db:"deletedTime"`
	}
)

func newSysApiModel(conn sqlx.SqlConn) *defaultSysApiModel {
	return &defaultSysApiModel{
		conn:  conn,
		table: "`sys_api`",
	}
}

func (m *defaultSysApiModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysApiModel) FindOne(ctx context.Context, id int64) (*SysApi, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysApiRows, m.table)
	var resp SysApi
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysApiModel) FindOneByRoute(ctx context.Context, route string) (*SysApi, error) {
	var resp SysApi
	query := fmt.Sprintf("select %s from %s where `route` = ? limit 1", sysApiRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, route)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysApiModel) Insert(ctx context.Context, data *SysApi) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysApiRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Route, data.Method, data.Name, data.BusinessType, data.Group, data.Desc)
	return ret, err
}

func (m *defaultSysApiModel) Update(ctx context.Context, newData *SysApi) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysApiRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Route, newData.Method, newData.Name, newData.BusinessType, newData.Group, newData.Desc, newData.Id)
	return err
}

func (m *defaultSysApiModel) tableName() string {
	return m.table
}
