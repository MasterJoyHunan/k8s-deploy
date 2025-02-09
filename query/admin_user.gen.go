// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"k8s-deploy/model"
	"k8s-deploy/types"
)

func newAdminUserModel(db *gorm.DB, opts ...gen.DOOption) adminUserModel {
	_adminUserModel := adminUserModel{}

	_adminUserModel.adminUserModelDo.UseDB(db, opts...)
	_adminUserModel.adminUserModelDo.UseModel(&model.AdminUserModel{})

	tableName := _adminUserModel.adminUserModelDo.TableName()
	_adminUserModel.ALL = field.NewAsterisk(tableName)
	_adminUserModel.ID = field.NewInt(tableName, "id")
	_adminUserModel.Username = field.NewString(tableName, "username")
	_adminUserModel.Realname = field.NewString(tableName, "realname")
	_adminUserModel.Password = field.NewString(tableName, "password")
	_adminUserModel.Phone = field.NewString(tableName, "phone")
	_adminUserModel.RoleID = field.NewInt(tableName, "role_id")
	_adminUserModel.Status = field.NewInt(tableName, "status")
	_adminUserModel.CreateTime = field.NewTime(tableName, "create_time")
	_adminUserModel.UpdateTime = field.NewTime(tableName, "update_time")
	_adminUserModel.DeleteTime = field.NewField(tableName, "delete_time")

	_adminUserModel.fillFieldMap()

	return _adminUserModel
}

// adminUserModel 后台用户
type adminUserModel struct {
	adminUserModelDo adminUserModelDo

	ALL        field.Asterisk
	ID         field.Int
	Username   field.String // 用户名
	Realname   field.String // 真实姓名
	Password   field.String // 密码
	Phone      field.String // 手机号
	RoleID     field.Int    // 角色ID
	Status     field.Int    // 状态 0:未启用 1:正常
	CreateTime field.Time
	UpdateTime field.Time
	DeleteTime field.Field

	fieldMap map[string]field.Expr
}

func (a adminUserModel) Table(newTableName string) *adminUserModel {
	a.adminUserModelDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminUserModel) As(alias string) *adminUserModel {
	a.adminUserModelDo.DO = *(a.adminUserModelDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminUserModel) updateTableName(table string) *adminUserModel {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt(table, "id")
	a.Username = field.NewString(table, "username")
	a.Realname = field.NewString(table, "realname")
	a.Password = field.NewString(table, "password")
	a.Phone = field.NewString(table, "phone")
	a.RoleID = field.NewInt(table, "role_id")
	a.Status = field.NewInt(table, "status")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.DeleteTime = field.NewField(table, "delete_time")

	a.fillFieldMap()

	return a
}

func (a *adminUserModel) WithContext(ctx context.Context) IAdminUserModelDo {
	return a.adminUserModelDo.WithContext(ctx)
}

func (a adminUserModel) TableName() string { return a.adminUserModelDo.TableName() }

func (a adminUserModel) Alias() string { return a.adminUserModelDo.Alias() }

func (a adminUserModel) Columns(cols ...field.Expr) gen.Columns {
	return a.adminUserModelDo.Columns(cols...)
}

func (a *adminUserModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminUserModel) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["username"] = a.Username
	a.fieldMap["realname"] = a.Realname
	a.fieldMap["password"] = a.Password
	a.fieldMap["phone"] = a.Phone
	a.fieldMap["role_id"] = a.RoleID
	a.fieldMap["status"] = a.Status
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["delete_time"] = a.DeleteTime
}

func (a adminUserModel) clone(db *gorm.DB) adminUserModel {
	a.adminUserModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminUserModel) replaceDB(db *gorm.DB) adminUserModel {
	a.adminUserModelDo.ReplaceDB(db)
	return a
}

type adminUserModelDo struct{ gen.DO }

type IAdminUserModelDo interface {
	gen.SubQuery
	Debug() IAdminUserModelDo
	WithContext(ctx context.Context) IAdminUserModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminUserModelDo
	WriteDB() IAdminUserModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminUserModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminUserModelDo
	Not(conds ...gen.Condition) IAdminUserModelDo
	Or(conds ...gen.Condition) IAdminUserModelDo
	Select(conds ...field.Expr) IAdminUserModelDo
	Where(conds ...gen.Condition) IAdminUserModelDo
	Order(conds ...field.Expr) IAdminUserModelDo
	Distinct(cols ...field.Expr) IAdminUserModelDo
	Omit(cols ...field.Expr) IAdminUserModelDo
	Join(table schema.Tabler, on ...field.Expr) IAdminUserModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminUserModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminUserModelDo
	Group(cols ...field.Expr) IAdminUserModelDo
	Having(conds ...gen.Condition) IAdminUserModelDo
	Limit(limit int) IAdminUserModelDo
	Offset(offset int) IAdminUserModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminUserModelDo
	Unscoped() IAdminUserModelDo
	Create(values ...*model.AdminUserModel) error
	CreateInBatches(values []*model.AdminUserModel, batchSize int) error
	Save(values ...*model.AdminUserModel) error
	First() (*model.AdminUserModel, error)
	Take() (*model.AdminUserModel, error)
	Last() (*model.AdminUserModel, error)
	Find() ([]*model.AdminUserModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminUserModel, err error)
	FindInBatches(result *[]*model.AdminUserModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminUserModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminUserModelDo
	Assign(attrs ...field.AssignExpr) IAdminUserModelDo
	Joins(fields ...field.RelationField) IAdminUserModelDo
	Preload(fields ...field.RelationField) IAdminUserModelDo
	FirstOrInit() (*model.AdminUserModel, error)
	FirstOrCreate() (*model.AdminUserModel, error)
	FindByPage(offset int, limit int) (result []*model.AdminUserModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminUserModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetUserInfo(id int) (result types.UserInfoResponse)
}

// GetUserInfo 获取用户基本信息
//
// SELECT
//
//	      u.id,
//	      u.username,
//	      u.realname,
//	      u.phone,
//	      r.NAME AS rolename,
//	      ( SELECT GROUP_CONCAT(`key`) FROM admin_auth WHERE JSON_CONTAINS( r.auth, JSON_ARRAY(id))) as authkeys
//	  FROM
//	      admin_user u
//	      LEFT JOIN admin_role r ON u.role_id = r.id
//	  WHERE
//			u.id = @id
func (a adminUserModelDo) GetUserInfo(id int) (result types.UserInfoResponse) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT u.id, u.username, u.realname, u.phone, r.NAME AS rolename, ( SELECT GROUP_CONCAT(`key`) FROM admin_auth WHERE JSON_CONTAINS( r.auth, JSON_ARRAY(id))) as authkeys FROM admin_user u LEFT JOIN admin_role r ON u.role_id = r.id WHERE u.id = ? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	_ = executeSQL

	return
}

func (a adminUserModelDo) Debug() IAdminUserModelDo {
	return a.withDO(a.DO.Debug())
}

func (a adminUserModelDo) WithContext(ctx context.Context) IAdminUserModelDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminUserModelDo) ReadDB() IAdminUserModelDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminUserModelDo) WriteDB() IAdminUserModelDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminUserModelDo) Session(config *gorm.Session) IAdminUserModelDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminUserModelDo) Clauses(conds ...clause.Expression) IAdminUserModelDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminUserModelDo) Returning(value interface{}, columns ...string) IAdminUserModelDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminUserModelDo) Not(conds ...gen.Condition) IAdminUserModelDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminUserModelDo) Or(conds ...gen.Condition) IAdminUserModelDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminUserModelDo) Select(conds ...field.Expr) IAdminUserModelDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminUserModelDo) Where(conds ...gen.Condition) IAdminUserModelDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminUserModelDo) Order(conds ...field.Expr) IAdminUserModelDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminUserModelDo) Distinct(cols ...field.Expr) IAdminUserModelDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminUserModelDo) Omit(cols ...field.Expr) IAdminUserModelDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminUserModelDo) Join(table schema.Tabler, on ...field.Expr) IAdminUserModelDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminUserModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminUserModelDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminUserModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminUserModelDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminUserModelDo) Group(cols ...field.Expr) IAdminUserModelDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminUserModelDo) Having(conds ...gen.Condition) IAdminUserModelDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminUserModelDo) Limit(limit int) IAdminUserModelDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminUserModelDo) Offset(offset int) IAdminUserModelDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminUserModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminUserModelDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminUserModelDo) Unscoped() IAdminUserModelDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminUserModelDo) Create(values ...*model.AdminUserModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminUserModelDo) CreateInBatches(values []*model.AdminUserModel, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminUserModelDo) Save(values ...*model.AdminUserModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminUserModelDo) First() (*model.AdminUserModel, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserModel), nil
	}
}

func (a adminUserModelDo) Take() (*model.AdminUserModel, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserModel), nil
	}
}

func (a adminUserModelDo) Last() (*model.AdminUserModel, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserModel), nil
	}
}

func (a adminUserModelDo) Find() ([]*model.AdminUserModel, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminUserModel), err
}

func (a adminUserModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminUserModel, err error) {
	buf := make([]*model.AdminUserModel, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminUserModelDo) FindInBatches(result *[]*model.AdminUserModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminUserModelDo) Attrs(attrs ...field.AssignExpr) IAdminUserModelDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminUserModelDo) Assign(attrs ...field.AssignExpr) IAdminUserModelDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminUserModelDo) Joins(fields ...field.RelationField) IAdminUserModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminUserModelDo) Preload(fields ...field.RelationField) IAdminUserModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminUserModelDo) FirstOrInit() (*model.AdminUserModel, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserModel), nil
	}
}

func (a adminUserModelDo) FirstOrCreate() (*model.AdminUserModel, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminUserModel), nil
	}
}

func (a adminUserModelDo) FindByPage(offset int, limit int) (result []*model.AdminUserModel, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminUserModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminUserModelDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminUserModelDo) Delete(models ...*model.AdminUserModel) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminUserModelDo) withDO(do gen.Dao) *adminUserModelDo {
	a.DO = *do.(*gen.DO)
	return a
}
