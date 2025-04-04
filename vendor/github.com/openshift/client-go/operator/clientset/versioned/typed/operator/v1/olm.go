// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// OLMsGetter has a method to return a OLMInterface.
// A group's client should implement this interface.
type OLMsGetter interface {
	OLMs() OLMInterface
}

// OLMInterface has methods to work with OLM resources.
type OLMInterface interface {
	Create(ctx context.Context, oLM *v1.OLM, opts metav1.CreateOptions) (*v1.OLM, error)
	Update(ctx context.Context, oLM *v1.OLM, opts metav1.UpdateOptions) (*v1.OLM, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, oLM *v1.OLM, opts metav1.UpdateOptions) (*v1.OLM, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.OLM, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.OLMList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OLM, err error)
	Apply(ctx context.Context, oLM *operatorv1.OLMApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OLM, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, oLM *operatorv1.OLMApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OLM, err error)
	OLMExpansion
}

// oLMs implements OLMInterface
type oLMs struct {
	*gentype.ClientWithListAndApply[*v1.OLM, *v1.OLMList, *operatorv1.OLMApplyConfiguration]
}

// newOLMs returns a OLMs
func newOLMs(c *OperatorV1Client) *oLMs {
	return &oLMs{
		gentype.NewClientWithListAndApply[*v1.OLM, *v1.OLMList, *operatorv1.OLMApplyConfiguration](
			"olms",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.OLM { return &v1.OLM{} },
			func() *v1.OLMList { return &v1.OLMList{} }),
	}
}
