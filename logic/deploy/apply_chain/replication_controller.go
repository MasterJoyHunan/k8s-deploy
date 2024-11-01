package apply_chain

import (
	"k8s-deploy/pkg/kubectl"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

type ReplicationController struct {
	ctx       *ChainContext
	localYaml *v1.ReplicationControllerApplyConfiguration
}

func (d *ReplicationController) next(ctx *ChainContext) error {
	if !(*ctx.CdrType.APIVersion == "v1" && *ctx.CdrType.Kind == "ReplicationController") {
		return nil
	}
	d.ctx = ctx
	if err := d.parse(); err != nil {
		return err
	}
	if err := d.apply(); err != nil {
		return err
	}
	return checkAllRunning(ctx.Ctx, *d.localYaml.Namespace, d.localYaml.Spec.Template.Labels)
}

func (d *ReplicationController) parse() error {
	var applyYaml v1.ReplicationControllerApplyConfiguration
	err := yaml.Unmarshal(d.ctx.YamlByte, &applyYaml)
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
		return err
	}
	d.localYaml = &applyYaml
	return nil
}

func (d *ReplicationController) apply() error {
	_, err := kubectl.K8sClient.
		CoreV1().
		ReplicationControllers(*d.localYaml.Namespace).
		Apply(d.ctx.Ctx, d.localYaml, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		d.ctx.Ctx.Log.Errorf("%+v", errors.WithStack(err))
	}
	return err
}
