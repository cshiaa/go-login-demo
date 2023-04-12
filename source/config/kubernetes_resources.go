package config

import (
	appsv1 "k8s.io/api/apps/v1"
)


type ResourcesMetadata struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
}

type ResourcesDeployment struct {
	ResourcesMetadata 				`json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Status appsv1.DeploymentStatus 	`json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}