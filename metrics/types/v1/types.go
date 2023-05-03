//go:build linux

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

package v1

import (
	cgroupv1 "github.com/containerd/cgroups/v3/cgroup1/stats"
)

type (
	// Metrics alias
	Metrics = cgroupv1.Metrics
	// BlkIOEntry alias
	BlkIOEntry = cgroupv1.BlkIOEntry
	// MemoryStat alias
	MemoryStat = cgroupv1.MemoryStat
	// CPUStat alias
	CPUStat = cgroupv1.CPUStat
	// CPUUsage alias
	CPUUsage = cgroupv1.CPUUsage
	// BlkIOStat alias
	BlkIOStat = cgroupv1.BlkIOStat
	// PidsStat alias
	PidsStat = cgroupv1.PidsStat
	// RdmaStat alias
	RdmaStat = cgroupv1.RdmaStat
	// RdmaEntry alias
	RdmaEntry = cgroupv1.RdmaEntry
	// HugetlbStat alias
	HugetlbStat = cgroupv1.HugetlbStat
)
