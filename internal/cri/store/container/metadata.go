/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package container

import (
	"encoding/json"
	"fmt"

	"github.com/containerd/log"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"
)

// NOTE(random-liu):
// 1) Metadata is immutable after created.
// 2) Metadata is checkpointed as containerd container label.

// metadataVersion is current version of container metadata.
const metadataVersion = "v1"

// versionedMetadata is the internal versioned container metadata.
type versionedMetadata struct {
	// Version indicates the version of the versioned container metadata.
	Version string
	// Metadata's type is metadataInternal. If not there will be a recursive call in MarshalJSON.
	Metadata metadataInternal
}

// metadataInternal is for internal use.
type metadataInternal struct {
	ID           string
	Name         string
	SandboxID    string
	Config       json.RawMessage
	ImageRef     string
	LogPath      string
	StopSignal   string
	ProcessLabel string
}

// Metadata is the unversioned container metadata.
type Metadata struct {
	// ID is the container id.
	ID string
	// Name is the container name.
	Name string
	// SandboxID is the sandbox id the container belongs to.
	SandboxID string
	// Config is the CRI container config.
	// NOTE(random-liu): Resource limits are updatable, the source
	// of truth for resource limits are in containerd.
	Config *runtime.ContainerConfig
	// ImageRef is the reference of image used by the container.
	ImageRef string
	// LogPath is the container log path.
	LogPath string
	// StopSignal is the system call signal that will be sent to the container to exit.
	// TODO(random-liu): Add integration test for stop signal.
	StopSignal string
	// ProcessLabel is the SELinux process label for the container
	ProcessLabel string
}

// MarshalJSON encodes Metadata into bytes in json format.
func (c *Metadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(&versionedMetadata{
		Version:  metadataVersion,
		Metadata: metadataToInternal(*c),
	})
}

// UnmarshalJSON decodes Metadata from bytes.
func (c *Metadata) UnmarshalJSON(data []byte) error {
	versioned := &versionedMetadata{}
	if err := json.Unmarshal(data, versioned); err != nil {
		return err
	}
	// Handle old version after upgrade.
	switch versioned.Version {
	case metadataVersion:
		*c = internalToMetadata(versioned.Metadata)
		return nil
	}
	return fmt.Errorf("unsupported version: %q", versioned.Version)
}

func internalToMetadata(c metadataInternal) Metadata {
	config := &runtime.ContainerConfig{}
	if err := json.Unmarshal(c.Config, config); err != nil {
		log.L.WithError(err).WithField("container", c.ID).Error("Failed to unmarshal container metadata ContainerConfig")
	}
	return Metadata{
		ID:           c.ID,
		Name:         c.Name,
		SandboxID:    c.SandboxID,
		Config:       config,
		ImageRef:     c.ImageRef,
		LogPath:      c.LogPath,
		StopSignal:   c.StopSignal,
		ProcessLabel: c.ProcessLabel,
	}
}

func metadataToInternal(c Metadata) metadataInternal {
	config, err := json.Marshal(c.Config)
	if err != nil {
		log.L.WithError(err).WithField("container", c.ID).Error("Failed to marshal container metadata ContainerConfig")
	}
	return metadataInternal{
		ID:           c.ID,
		Name:         c.Name,
		SandboxID:    c.SandboxID,
		Config:       config,
		ImageRef:     c.ImageRef,
		LogPath:      c.LogPath,
		StopSignal:   c.StopSignal,
		ProcessLabel: c.ProcessLabel,
	}
}
