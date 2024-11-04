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
 */

// Package options xxx
package options

import (
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/conf"

	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/pkg/common"
)

// Config defines the config
type Config struct {
	conf.FileConfig
	conf.LogConfig

	ListenAddr string `json:"listenAddr" value:"0.0.0.0" usage:"proxy server listen addr"`
	ListenPort int    `json:"listenPort" value:"8080" usage:"proxy server listen port"`

	TlsCert string `json:"tlsCert" value:"/etc/webhook/certs/cert.pem" usage:"webhook server cert"`
	TlsKey  string `json:"tlsKey" value:"/etc/webhook/certs/key.pem" usage:"webhook server key"`

	ArgoService string `json:"argoService" value:"" usage:"the service address of argo"`
	ArgoUser    string `json:"argoUser" value:"" usage:"the user of argo"`
	ArgoPass    string `json:"argoPass" value:"" usage:"the password of argo"`
	ArgoRepoUrl string `json:"argoRepoUrl" value:"" usage:"the url of argo repo-server"`

	DBConfig common.DBConfig `json:"dbConfig,omitempty"`

	RecoverProjects string `json:"recoverProjects" value:"" usage:""`
	AdminNamespace  string `json:"adminNamespace" value:"" usage:""`

	InterceptSyncProjectsStr string   `json:"interceptSyncProjects" value:"" usage:"the projects need intercept sync"`
	InterceptSyncProjects    []string `json:"-" value:"" usage:""`

	PublicProjectsStr string   `json:"publicProjects,omitempty"`
	PublicProjects    []string `json:"-"`
}

// Parse parse the options
func Parse() *Config {
	cfg := new(Config)
	conf.Parse(cfg)
	cfg.PublicProjects = strings.Split(cfg.PublicProjectsStr, ",")
	cfg.InterceptSyncProjects = strings.Split(cfg.InterceptSyncProjectsStr, ",")
	return cfg
}
