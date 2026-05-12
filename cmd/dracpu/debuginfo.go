/*
Copyright The Kubernetes Authors.

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

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kubernetes-sigs/dra-driver-cpu/pkg/cpuinfo"
)

type debugInfo struct {
	Topology *cpuinfo.CPUTopology `json:"topology"`
	CPUs     []cpuinfo.CPUInfo    `json:"cpus"`
}

// runDebugInfo collects CPU topology data and prints it as JSON to stdout.
func runDebugInfo() error {
	sys := cpuinfo.NewSystemCPUInfo()

	topology, err := sys.GetCPUTopology()
	if err != nil {
		return fmt.Errorf("failed to get CPU topology: %w", err)
	}

	cpus, err := sys.GetCPUInfos()
	if err != nil {
		return fmt.Errorf("failed to get CPU infos: %w", err)
	}

	info := debugInfo{
		Topology: topology,
		CPUs:     cpus,
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(info); err != nil {
		return fmt.Errorf("failed to encode output: %w", err)
	}

	return nil
}
