// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build orchestrator

//nolint:revive // TODO(CAPP) Fix revive linter
package k8s

import (
	model "github.com/DataDog/agent-payload/v5/process"
	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/cluster/orchestrator/processors"
)

//nolint:revive // TODO(CAPP) Fix revive linter
type BaseHandlers struct{}

//nolint:revive // TODO(CAPP) Fix revive linter
func (BaseHandlers) BeforeCacheCheck(ctx *processors.ProcessorContext, resource, resourceModel interface{}) (skip bool) {
	return
}

//nolint:revive // TODO(CAPP) Fix revive linter
func (BaseHandlers) BeforeMarshalling(ctx *processors.ProcessorContext, resource, resourceModel interface{}) (skip bool) {
	return
}

//nolint:revive // TODO(CAPP) Fix revive linter
func (BaseHandlers) ScrubBeforeMarshalling(ctx *processors.ProcessorContext, resource interface{}) {}

//nolint:revive // TODO(CAPP) Fix revive linter
func (BaseHandlers) BuildManifestMessageBody(ctx *processors.ProcessorContext, resourceManifests []interface{}, groupSize int) model.MessageBody {
	return ExtractModelManifests(ctx, resourceManifests, groupSize)
}

// ExtractModelManifests creates the model manifest from the given manifests
func ExtractModelManifests(ctx *processors.ProcessorContext, resourceManifests []interface{}, groupSize int) *model.CollectorManifest {
	manifests := make([]*model.Manifest, 0, len(resourceManifests))

	for _, m := range resourceManifests {
		manifests = append(manifests, m.(*model.Manifest))
	}

	cm := &model.CollectorManifest{
		ClusterName: ctx.Cfg.KubeClusterName,
		ClusterId:   ctx.ClusterID,
		Manifests:   manifests,
		GroupId:     ctx.MsgGroupID,
		GroupSize:   int32(groupSize),
	}
	return cm
}
