package role

import (
	"strconv"

	"k8s-deploy/config"
	"k8s-deploy/model"
	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
)

// Del 删除角色
func Del(ctx *svc.ServiceContext, req *types.PathID) error {
	if req.ID == config.SuperAdminRoleID {
		return errors.New("无法修改超级管理员角色")
	}

	userModel := query.AdminUserModel
	user, _ := userModel.WithContext(ctx).Where(userModel.RoleID.Eq(req.ID)).First()
	if user != nil {
		return errors.New("当前角色正在使用,无法删除")
	}

	defer model.Enforcer.LoadPolicy()

	err := query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.AdminRoleModel.WithContext(ctx).Where(tx.AdminRoleModel.ID.Eq(req.ID)).Delete()
		if err != nil {
			return err
		}

		_, err = tx.AdminCasbinRuleModel.WithContext(ctx).
			Where(
				tx.AdminCasbinRuleModel.Ptype.Eq("p"),
				tx.AdminCasbinRuleModel.V0.Eq("role:"+strconv.Itoa(req.ID)),
			).
			Delete()
		return err
	})
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}
