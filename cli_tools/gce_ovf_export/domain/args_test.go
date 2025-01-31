//  Copyright 2020 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package ovfexportdomain

import (
	"testing"

	"github.com/GoogleCloudPlatform/compute-image-import/cli_tools/common/utils/daisyutils"

	"github.com/stretchr/testify/assert"
)

func TestIsInstanceExport(t *testing.T) {
	assert.True(t, GetAllInstanceExportArgs().IsInstanceExport())
	assert.False(t, GetAllMachineImageExportArgs().IsInstanceExport())
}

func TestIsMachineImageExport(t *testing.T) {
	assert.False(t, GetAllInstanceExportArgs().IsMachineImageExport())
	assert.True(t, GetAllMachineImageExportArgs().IsMachineImageExport())
}

func TestDaisyAttrs(t *testing.T) {
	params := GetAllInstanceExportArgs()
	assert.Equal(t,
		daisyutils.EnvironmentSettings{
			Project:               params.Project,
			Zone:                  params.Zone,
			GCSPath:               params.ScratchBucketGcsPath,
			OAuth:                 params.Oauth,
			Timeout:               params.Timeout.String(),
			ComputeEndpoint:       params.Ce,
			WorkflowDirectory:     params.WorkflowDir,
			DisableGCSLogs:        params.GcsLogsDisabled,
			DisableCloudLogs:      params.CloudLogsDisabled,
			DisableStdoutLogs:     params.StdoutLogsDisabled,
			NoExternalIP:          params.NoExternalIP,
			Network:               params.Network,
			Subnet:                params.Subnet,
			ComputeServiceAccount: params.ComputeServiceAccount,
			Labels:                map[string]string{},
			ExecutionID:           params.BuildID,
			Tool:                  daisyutils.Tool{HumanReadableName: "ovf export", ResourceLabelName: "gce-ovf-export"},
			DaisyLogLinePrefix:    "ovf-export",
			WorkerMachineSeries:   []string{"n2", "n1"},
		},
		params.EnvironmentSettings("ovf-export"))
}
