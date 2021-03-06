package common

const (
	// DefaultBootHelmfileRepository default git repo for boot with helmfile
	DefaultBootHelmfileRepository = "https://github.com/jenkins-x/jenkins-x-boot-helmfile-config.git"

	// DefaultVersionsRef default version stream ref
	DefaultVersionsRef = "master"

	// DefaultVersionsURL default version stream url
	DefaultVersionsURL = "https://github.com/jenkins-x/jenkins-x-versions.git"

	// HelmfileBuildPackName the build pack name for helm 3 / helmfile style environments
	HelmfileBuildPackName = "environment-helmfile"

	// PipelineActivitiesYAMLFile the name of the YAML file to help migrate PipelineActivity resources to a new cluster
	PipelineActivitiesYAMLFile = "pipelineActivities.yaml"
)
