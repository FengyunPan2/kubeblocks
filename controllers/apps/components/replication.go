/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package components

import (
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	"github.com/apecloud/kubeblocks/internal/controller/component"
	"github.com/apecloud/kubeblocks/internal/controller/graph"
	ictrltypes "github.com/apecloud/kubeblocks/internal/controller/types"
	intctrlutil "github.com/apecloud/kubeblocks/internal/controllerutil"
)

func newReplicationComponent(cli client.Client,
	recorder record.EventRecorder,
	cluster *appsv1alpha1.Cluster,
	clusterVersion *appsv1alpha1.ClusterVersion,
	synthesizedComponent *component.SynthesizedComponent,
	dag *graph.DAG) *replicationComponent {
	comp := &replicationComponent{
		statefulComponentBase: statefulComponentBase{
			componentBase: componentBase{
				Client:         cli,
				Recorder:       recorder,
				Cluster:        cluster,
				ClusterVersion: clusterVersion,
				Component:      synthesizedComponent,
				ComponentSet: &replicationSet{
					stateful: stateful{
						componentSetBase: componentSetBase{
							Cli:                  cli,
							Cluster:              cluster,
							SynthesizedComponent: synthesizedComponent,
							ComponentSpec:        nil,
							ComponentDef:         nil,
						},
					},
				},
				Dag:            dag,
				WorkloadVertex: nil,
			},
		},
	}
	return comp
}

type replicationComponent struct {
	statefulComponentBase
}

var _ Component = &replicationComponent{}

func (c *replicationComponent) newBuilder(reqCtx intctrlutil.RequestCtx, cli client.Client,
	action *ictrltypes.LifecycleAction) componentWorkloadBuilder {
	builder := &replicationComponentWorkloadBuilder{
		componentWorkloadBuilderBase: componentWorkloadBuilderBase{
			ReqCtx:        reqCtx,
			Client:        cli,
			Comp:          c,
			DefaultAction: action,
			Error:         nil,
			EnvConfig:     nil,
			Workload:      nil,
		},
	}
	builder.ConcreteBuilder = builder
	return builder
}

func (c *replicationComponent) GetWorkloadType() appsv1alpha1.WorkloadType {
	return appsv1alpha1.Replication
}

func (c *replicationComponent) GetBuiltObjects(reqCtx intctrlutil.RequestCtx, cli client.Client) ([]client.Object, error) {
	return c.statefulComponentBase.GetBuiltObjects(c.newBuilder(reqCtx, cli, ictrltypes.ActionCreatePtr()))
}

func (c *replicationComponent) Create(reqCtx intctrlutil.RequestCtx, cli client.Client) error {
	return c.statefulComponentBase.Create(reqCtx, cli, c.newBuilder(reqCtx, cli, ictrltypes.ActionCreatePtr()))
}

func (c *replicationComponent) Update(reqCtx intctrlutil.RequestCtx, cli client.Client) error {
	return c.statefulComponentBase.Update(reqCtx, cli, c.newBuilder(reqCtx, cli, nil))
}

func (c *replicationComponent) Status(reqCtx intctrlutil.RequestCtx, cli client.Client) error {
	return c.statefulComponentBase.Status(reqCtx, cli, c.newBuilder(reqCtx, cli, ictrltypes.ActionNoopPtr()))
}
