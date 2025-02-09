package project

import (
	"encoding/json"

	"k8s-deploy/query"
	"k8s-deploy/svc"
	"k8s-deploy/types"

	"github.com/pkg/errors"
	"gorm.io/gen/field"
)

// Edit 编辑项目
func Edit(ctx *svc.ServiceContext, req *types.ProjectEditRequest) error {
	projectModel := query.ProjectModel

	params, _ := json.Marshal(req.Params)
	template, _ := json.Marshal(req.Template)

	var assign = []field.AssignExpr{
		projectModel.Name.Value(req.Name),
		projectModel.Desc.Value(req.Desc),
		projectModel.Git.Value(req.Git),
		projectModel.UserName.Value(req.UserName),
		projectModel.UseTag.Value(req.UseTag),
		projectModel.Params.Value(string(params)),
		projectModel.Template.Value(string(template)),
	}

	if len(req.Token) > 0 {
		assign = append(assign, projectModel.Token.Value(req.Token))
	}

	_, err := projectModel.WithContext(ctx).Where(projectModel.ID.Eq(req.ID)).UpdateColumnSimple(assign...)

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
