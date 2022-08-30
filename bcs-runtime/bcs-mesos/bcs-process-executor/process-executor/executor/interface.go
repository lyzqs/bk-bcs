/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package executor

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/mesosproto/mesos"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/types"
)

// Executor xxx
type Executor interface {
	// LaunchTaskgroup xxx
	// launch taskgroup
	LaunchTaskgroup(*mesos.TaskGroupInfo)

	// Shutdown xxx
	// shut down
	Shutdown()

	// RegisterCallbackFunc xxx
	// register callback function
	RegisterCallbackFunc(types.CallbackFuncType, interface{})

	// GetExecutorStatus xxx
	// Get Executor status
	GetExecutorStatus() types.ExecutorStatus

	// ReloadTasks xxx
	// reload tasks, exec reloadCmd
	ReloadTasks() error

	// RestartTasks xxx
	// restart tasks, exec restartCmd
	RestartTasks() error

	AckTaskStatusMessage(taskId string, uuid []byte)
}
