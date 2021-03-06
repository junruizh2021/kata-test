// Copyright (c) 2017 Intel Corporation
// Copyright (c) 2018 HyperHQ Inc.
//
// SPDX-License-Identifier: Apache-2.0
//

package containerdshim

import (
	"os"
	"testing"

	taskAPI "github.com/containerd/containerd/runtime/v2/task"
	"github.com/stretchr/testify/assert"

	ktu "github.com/kata-containers/kata-containers/src/runtime/pkg/katatestutils"
	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/compatoci"
	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/vcmock"
)

func TestDeleteContainerSuccessAndFail(t *testing.T) {
	assert := assert.New(t)

	sandbox := &vcmock.Sandbox{
		MockID: testSandboxID,
	}

	rootPath, bundlePath, _ := ktu.SetupOCIConfigFile(t)
	defer os.RemoveAll(rootPath)
	_, err := compatoci.ParseConfigJSON(bundlePath)
	assert.NoError(err)

	s := &service{
		id:         testSandboxID,
		sandbox:    sandbox,
		containers: make(map[string]*container),
	}

	reqCreate := &taskAPI.CreateTaskRequest{
		ID: testContainerID,
	}
	s.containers[testContainerID], err = newContainer(s, reqCreate, "", nil, true)
	assert.NoError(err)
}
