// Code generated by goctl. DO NOT EDIT.
package routes

import (
	adminAuth "k8s-deploy/routes/admin/auth"
	adminBase "k8s-deploy/routes/admin/base"
	adminLogin "k8s-deploy/routes/admin/login"
	adminRole "k8s-deploy/routes/admin/role"
	adminUser "k8s-deploy/routes/admin/user"
	common "k8s-deploy/routes/common"
	deploy "k8s-deploy/routes/deploy"
	namespace "k8s-deploy/routes/namespace"
	project "k8s-deploy/routes/project"
	run "k8s-deploy/routes/run"
	template "k8s-deploy/routes/template"
	websocket "k8s-deploy/routes/websocket"

	"github.com/gin-gonic/gin"
)

func Setup(e *gin.Engine) {
	common.RegisterCommonRoute(e)
	adminLogin.RegisterAdminLoginRoute(e)
	adminUser.RegisterAdminUserRoute(e)
	adminRole.RegisterAdminRoleRoute(e)
	adminAuth.RegisterAdminAuthRoute(e)
	adminBase.RegisterAdminBaseRoute(e)
	namespace.RegisterNamespaceRoute(e)
	project.RegisterProjectRoute(e)
	template.RegisterTemplateRoute(e)
	deploy.RegisterDeployRoute(e)
	websocket.RegisterWebsocketRoute(e)
	run.RegisterRunRoute(e)
}
