package handler

import (
	"context"
	"fmt"

	rainbondv1alpha1 "github.com/GLYASAI/rainbond-operator/pkg/apis/rainbond/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var MonitorName = "rbd-monitor"

type monitor struct {
	component *rainbondv1alpha1.RbdComponent
	cluster   *rainbondv1alpha1.RainbondCluster
	labels    map[string]string
}

func NewMonitor(ctx context.Context, client client.Client, component *rainbondv1alpha1.RbdComponent, cluster *rainbondv1alpha1.RainbondCluster) ComponentHandler {
	return &monitor{
		component: component,
		cluster:   cluster,
		labels:    component.Labels(),
	}
}

func (m *monitor) Before() error {
	// TODO: check prerequisites
	return nil
}

func (m *monitor) Resources() []interface{} {
	return []interface{}{
		m.daemonSetForMonitor(),
	}
}

func (m *monitor) After() error {
	return nil
}

func (m *monitor) daemonSetForMonitor() interface{} {
	ds := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      MonitorName,
			Namespace: m.component.Namespace,
			Labels:    m.labels,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: m.labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   MonitorName,
					Labels: m.labels,
				},
				Spec: corev1.PodSpec{
					Tolerations: []corev1.Toleration{
						{
							Key:    "node-role.kubernetes.io/master",
							Effect: corev1.TaintEffectNoSchedule,
						},
					},
					NodeSelector: map[string]string{
						"node-role.kubernetes.io/master": "",
					},
					Containers: []corev1.Container{
						{
							Name:            MonitorName,
							Image:           m.component.Spec.Image,
							ImagePullPolicy: m.component.ImagePullPolicy(),
							Env: []corev1.EnvVar{
								{
									Name: "POD_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "status.podIP",
										},
									},
								},
							},
							Args: []string{
								"--etcd-endpoints=http://etcd0:2379", // TODO: hard code
								"--advertise-addr=$(POD_IP):9999",
								"--alertmanager-address=$(POD_IP):9093",
								"--storage.tsdb.path=/prometheusdata",
								"--storage.tsdb.no-lockfile",
								"--storage.tsdb.retention=7d",
								fmt.Sprintf("--log.level=%s", m.component.LogLevel()),
							},
						},
						// TODO: /prometheusdata
					},
				},
			},
		},
	}

	return ds
}