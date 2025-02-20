package containermanger

import (
	"context"

	"github.com/docker/docker/client"
)

type Manager struct {
	ctx    context.Context
	client *client.Client
}

type DeploymentMode string

const (
	DeploymentModeReplicated DeploymentMode = "replicated"
	DeploymentModeGlobal     DeploymentMode = "global"
)

type Service struct {
	Name                 string            `json:"name"`
	Image                string            `json:"image"`
	Command              []string          `json:"command,omitempty"`
	Env                  map[string]string `json:"env,omitempty"`
	Capabilities         []string          `json:"capabilities,omitempty"`
	Sysctls              map[string]string `json:"sysctl,omitempty"`
	VolumeMounts         []VolumeMount     `json:"volumemounts,omitempty"`
	VolumeBinds          []VolumeBind      `json:"volumebinds,omitempty"`
	Networks             []string          `json:"networks,omitempty"`
	DeploymentMode       DeploymentMode    `json:"deploymentmode"`
	Replicas             uint64            `json:"replicas"`
	PlacementConstraints []string          `json:"placementconstraints,omitempty"`
	ReservedResource     Resource          `json:"reserved_resource,omitempty"`
	ResourceLimit        Resource          `json:"resource_limit,omitempty"`
}

type ServiceRealtimeInfo struct {
	Name              string                     `json:"name"`
	PlacementInfos    []ServiceTaskPlacementInfo `json:"placementinfos"`
	DesiredReplicas   int                        `json:"desiredreplicas"`
	RunningReplicas   int                        `json:"runningreplicas"`
	ReplicatedService bool                       `json:"replicatedservice"`
}

type ServiceTaskPlacementInfo struct {
	NodeID          string `json:"nodeid"`
	NodeName        string `json:"nodename"`
	IsManagerNode   bool   `json:"ismanagernode"`
	RunningReplicas int    `json:"runningreplicas"`
}

type VolumeMount struct {
	Source   string `json:"source"`
	Target   string `json:"target"`
	ReadOnly bool   `json:"readonly"`
}

type VolumeBind struct {
	Source string `json:"source"`
	Target string `json:"target"`
	// No need to specify readonly for volume bind
	// VolumeBind is for special internal use, and need readwrite access all the time
}

type Resource struct {
	MemoryMB int `json:"memory_mb,omitempty"`
}
