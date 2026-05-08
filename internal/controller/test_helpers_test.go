// Copyright Contributors to the KubeOpenCode project

//go:build !integration

package controller

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// hasEnvVar checks if an env var list contains a variable with the given name and value.
func hasEnvVar(envs []corev1.EnvVar, name, value string) bool {
	for _, e := range envs {
		if e.Name == name && e.Value == value {
			return true
		}
	}
	return false
}

// hasVolumeMount checks if a volume mount list contains a mount with the given name and mountPath.
func hasVolumeMount(mounts []corev1.VolumeMount, name, mountPath string) bool {
	for _, m := range mounts {
		if m.Name == name && m.MountPath == mountPath {
			return true
		}
	}
	return false
}

// findPodInitContainer returns the init container with the given name from a Pod, or nil.
func findPodInitContainer(pod *corev1.Pod, name string) *corev1.Container {
	for i := range pod.Spec.InitContainers {
		if pod.Spec.InitContainers[i].Name == name {
			return &pod.Spec.InitContainers[i]
		}
	}
	return nil
}

// findDeploymentInitContainer returns the init container with the given name from a Deployment, or nil.
func findDeploymentInitContainer(deployment *appsv1.Deployment, name string) *corev1.Container {
	for i := range deployment.Spec.Template.Spec.InitContainers {
		if deployment.Spec.Template.Spec.InitContainers[i].Name == name {
			return &deployment.Spec.Template.Spec.InitContainers[i]
		}
	}
	return nil
}
