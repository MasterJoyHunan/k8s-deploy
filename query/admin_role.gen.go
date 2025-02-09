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
)

func newAdminRoleModel(db *gorm.DB, opts ...gen.DOOption) adminRoleModel {
	_adminRoleModel := adminRoleModel{}

	_adminRoleModel.adminRoleModelDo.UseDB(db, opts...)
	_adminRoleModel.adminRoleModelDo.UseModel(&model.AdminRoleModel{})

	tableName := _adminRoleModel.adminRoleModelDo.TableName()
	_adminRoleModel.ALL = field.NewAsterisk(tableName)
	_adminRoleModel.ID = field.NewInt(tableName, "id")
	_adminRoleModel.Name = field.NewString(tableName, "name")
	_adminRoleModel.Auth = field.NewString(tableName, "auth")
	_adminRoleModel.CreateTime = field.NewTime(tableName, "create_time")
	_adminRoleModel.UpdateTime = field.NewTime(tableName, "update_time")
	_adminRoleModel.DeleteTime = field.NewField(tableName, "delete_time")

	_adminRoleModel.fillFieldMap()

	return _adminRoleModel
}

// adminRoleModel 角色
type adminRoleModel struct {
	adminRoleModelDo adminRoleModelDo

	ALL        field.Asterisk
	ID         field.Int
	Name       field.String // 角色名
	Auth       field.String // 权限ID
	CreateTime field.Time
	UpdateTime field.Time
	DeleteTime field.Field

	fieldMap map[string]field.Expr
}

func (a adminRoleModel) Table(newTableName string) *adminRoleModel {
	a.adminRoleModelDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminRoleModel) As(alias string) *adminRoleModel {
	a.adminRoleModelDo.DO = *(a.adminRoleModelDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminRoleModel) updateTableName(table string) *adminRoleModel {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt(table, "id")
	a.Name = field.NewString(table, "name")
	a.Auth = field.NewString(table, "auth")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.DeleteTime = field.NewField(table, "delete_time")

	a.fillFieldMap()

	return a
}

func (a *adminRoleModel) WithContext(ctx context.Context) IAdminRoleModelDo {
	return a.adminRoleModelDo.WithContext(ctx)
}

func (a adminRoleModel) TableName() string { return a.adminRoleModelDo.TableName() }

func (a adminRoleModel) Alias() string { return a.adminRoleModelDo.Alias() }

func (a adminRoleModel) Columns(cols ...field.Expr) gen.Columns {
	return a.adminRoleModelDo.Columns(cols...)
}

func (a *adminRoleModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminRoleModel) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["auth"] = a.Auth
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["delete_time"] = a.DeleteTime
}

func (a adminRoleModel) clone(db *gorm.DB) adminRoleModel {
	a.adminRoleModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminRoleModel) replaceDB(db *gorm.DB) adminRoleModel {
	a.adminRoleModelDo.ReplaceDB(db)
	return a
}

type adminRoleModelDo struct{ gen.DO }

type IAdminRoleModelDo interface {
	gen.SubQuery
	Debug() IAdminRoleModelDo
	WithContext(ctx context.Context) IAdminRoleModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminRoleModelDo
	WriteDB() IAdminRoleModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminRoleModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminRoleModelDo
	Not(conds ...gen.Condition) IAdminRoleModelDo
	Or(conds ...gen.Condition) IAdminRoleModelDo
	Select(conds ...field.Expr) IAdminRoleModelDo
	Where(conds ...gen.Condition) IAdminRoleModelDo
	Order(conds ...field.Expr) IAdminRoleModelDo
	Distinct(cols ...field.Expr) IAdminRoleModelDo
	Omit(cols ...field.Expr) IAdminRoleModelDo
	Join(table schema.Tabler, on ...field.Expr) IAdminRoleModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminRoleModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminRoleModelDo
	Group(cols ...field.Expr) IAdminRoleModelDo
	Having(conds ...gen.Condition) IAdminRoleModelDo
	Limit(limit int) IAdminRoleModelDo
	Offset(offset int) IAdminRoleModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminRoleModelDo
	Unscoped() IAdminRoleModelDo
	Create(values ...*model.AdminRoleModel) error
	CreateInBatches(values []*model.AdminRoleModel, batchSize int) error
	Save(values ...*model.AdminRoleModel) error
	First() (*model.AdminRoleModel, error)
	Take() (*model.AdminRoleModel, error)
	Last() (*model.AdminRoleModel, error)
	Find() ([]*model.AdminRoleModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminRoleModel, err error)
	FindInBatches(result *[]*model.AdminRoleModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminRoleModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminRoleModelDo
	Assign(attrs ...field.AssignExpr) IAdminRoleModelDo
	Joins(fields ...field.RelationField) IAdminRoleModelDo
	Preload(fields ...field.RelationField) IAdminRoleModelDo
	FirstOrInit() (*model.AdminRoleModel, error)
	FirstOrCreate() (*model.AdminRoleModel, error)
	FindByPage(offset int, limit int) (result []*model.AdminRoleModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminRoleModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRoleByAuthId(authId int) (result *model.AdminRoleModel)
}

// GetRoleByAuthId 根据 auth_id 获取角色
//
// where(json_contains( auth, json_array( @authId )) )
func (a adminRoleModelDo) GetRoleByAuthId(authId int) (result *model.AdminRoleModel) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, authId)
	generateSQL.WriteString("json_contains( auth, json_array( ? )) ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	_ = executeSQL

	return
}

func (a adminRoleModelDo) Debug() IAdminRoleModelDo {
	return a.withDO(a.DO.Debug())
}

func (a adminRoleModelDo) WithContext(ctx context.Context) IAdminRoleModelDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminRoleModelDo) ReadDB() IAdminRoleModelDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminRoleModelDo) WriteDB() IAdminRoleModelDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminRoleModelDo) Session(config *gorm.Session) IAdminRoleModelDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminRoleModelDo) Clauses(conds ...clause.Expression) IAdminRoleModelDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminRoleModelDo) Returning(value interface{}, columns ...string) IAdminRoleModelDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminRoleModelDo) Not(conds ...gen.Condition) IAdminRoleModelDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminRoleModelDo) Or(conds ...gen.Condition) IAdminRoleModelDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminRoleModelDo) Select(conds ...field.Expr) IAdminRoleModelDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminRoleModelDo) Where(conds ...gen.Condition) IAdminRoleModelDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminRoleModelDo) Order(conds ...field.Expr) IAdminRoleModelDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminRoleModelDo) Distinct(cols ...field.Expr) IAdminRoleModelDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminRoleModelDo) Omit(cols ...field.Expr) IAdminRoleModelDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminRoleModelDo) Join(table schema.Tabler, on ...field.Expr) IAdminRoleModelDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminRoleModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminRoleModelDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminRoleModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminRoleModelDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminRoleModelDo) Group(cols ...field.Expr) IAdminRoleModelDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminRoleModelDo) Having(conds ...gen.Condition) IAdminRoleModelDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminRoleModelDo) Limit(limit int) IAdminRoleModelDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminRoleModelDo) Offset(offset int) IAdminRoleModelDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminRoleModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminRoleModelDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminRoleModelDo) Unscoped() IAdminRoleModelDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminRoleModelDo) Create(values ...*model.AdminRoleModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminRoleModelDo) CreateInBatches(values []*model.AdminRoleModel, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminRoleModelDo) Save(values ...*model.AdminRoleModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminRoleModelDo) First() (*model.AdminRoleModel, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleModel), nil
	}
}

func (a adminRoleModelDo) Take() (*model.AdminRoleModel, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleModel), nil
	}
}

func (a adminRoleModelDo) Last() (*model.AdminRoleModel, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleModel), nil
	}
}

func (a adminRoleModelDo) Find() ([]*model.AdminRoleModel, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminRoleModel), err
}

func (a adminRoleModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminRoleModel, err error) {
	buf := make([]*model.AdminRoleModel, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminRoleModelDo) FindInBatches(result *[]*model.AdminRoleModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminRoleModelDo) Attrs(attrs ...field.AssignExpr) IAdminRoleModelDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminRoleModelDo) Assign(attrs ...field.AssignExpr) IAdminRoleModelDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminRoleModelDo) Joins(fields ...field.RelationField) IAdminRoleModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminRoleModelDo) Preload(fields ...field.RelationField) IAdminRoleModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminRoleModelDo) FirstOrInit() (*model.AdminRoleModel, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleModel), nil
	}
}

func (a adminRoleModelDo) FirstOrCreate() (*model.AdminRoleModel, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleModel), nil
	}
}

func (a adminRoleModelDo) FindByPage(offset int, limit int) (result []*model.AdminRoleModel, count int64, err error) {
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

func (a adminRoleModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminRoleModelDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminRoleModelDo) Delete(models ...*model.AdminRoleModel) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminRoleModelDo) withDO(do gen.Dao) *adminRoleModelDo {
	a.DO = *do.(*gen.DO)
	return a
}
