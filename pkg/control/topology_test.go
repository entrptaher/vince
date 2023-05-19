package control

import (
	"bytes"
	"os"
	"testing"

	"github.com/gernest/vince/pkg/apis/vince/v1alpha1"
	vince_informers "github.com/gernest/vince/pkg/gen/client/vince/informers/externalversions"
	"github.com/gernest/vince/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
)

func mock(t *testing.T, path string) *k8s.Mock {
	m, err := k8s.NewMock(path)
	if err != nil {
		t.Fatalf("failed to initialize mock client %v", err)
	}
	return m
}

func build(clients k8s.Client) *Topology {
	vince := vince_informers.NewSharedInformerFactory(clients.Vince(), k8s.ResyncPeriod)
	k8s := informers.NewSharedInformerFactory(clients.Kube(), k8s.ResyncPeriod)
	return &Topology{
		vinceLister:   vince.Staples().V1alpha1().Vinces().Lister(),
		siteLister:    vince.Staples().V1alpha1().Sites().Lister(),
		podLister:     k8s.Core().V1().Pods().Lister(),
		secretsLister: k8s.Core().V1().Secrets().Lister(),
	}
}

func TestFirstApply(t *testing.T) {
	// Running apply to Vince crd for the first time.
	clients := mock(t, "topology/first_apply.yml")
	topo := build(clients)
	_, err := topo.Build(&k8s.ResourceFilter{}, "vince", func(s string) {})
	if err != nil {
		t.Error(err)
	}
	var b bytes.Buffer
	err = k8s.Encode(&b,
		&v1alpha1.Vince{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Vince",
				APIVersion: "staples/v1alpha1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "topology",
				Namespace: "ns0",
				Labels:    baseLabels(),
			},
			Spec: v1alpha1.VinceSpec{
				Volume: v1.PersistentVolumeClaimSpec{
					Resources: v1.ResourceRequirements{
						Requests: v1.ResourceList{
							v1.ResourceStorage: resource.MustParse("1Gi"),
						},
					},
				},
				Container: v1.Container{},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("./testdata/topology/first_apply.yml", b.Bytes(), 0600)
}