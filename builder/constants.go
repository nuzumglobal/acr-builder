// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package builder

const (
	// NoBaseImageSpecifierLatest is the empty base image
	// Note that :latest is not valid in the FROM clause, but we're
	// always appending :latest to tags during processing.
	NoBaseImageSpecifierLatest = "scratch:latest"

	// DockerHubRegistry is the docker hub registry
	DockerHubRegistry = "registry.hub.docker.com"

	containerWorkspaceDir = "/workspace"
	rmContainer           = true

	// homeVol is the volume to manage $HOME
	homeVol = "home"

	// homeWorkDir is the working directory to start at in $HOME
	homeWorkDir = "/acb/home"
)