package apply_chain

import (
	deploylog "k8s-deploy/logic/deploy_log"
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

type ConfigMap struct {
	ctx       *ChainContext
	localYaml *v1.ConfigMapApplyConfiguration
}

func (d *ConfigMap) next(ctx *ChainContext) error {
	if *ctx.CdrType.Kind != "ConfigMap" {
		return nil
	}
	d.ctx = ctx

	err := d.parse()
	if err != nil {
		return err
	}
	return d.applyWarp()
}

func (d *ConfigMap) parse() error {
	var applyYaml v1.ConfigMapApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *ConfigMap) applyWarp() error {
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 0, "部署 ConfigMap \n"+string(d.ctx.YamlByte))
	err := d.apply()
	if err != nil {
		deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 3, "部署 ConfigMap 失败, 失败原因："+err.Error())
		return err
	}
	deploylog.RecordLog(d.ctx.Ctx, d.ctx.ID, 1, "部署 ConfigMap 成功")
	return nil
}

func (d *ConfigMap) apply() error {
	_, err := kubectl.K8sClient.
		CoreV1().
		ConfigMaps(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch", DryRun: d.ctx.DryRun})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
