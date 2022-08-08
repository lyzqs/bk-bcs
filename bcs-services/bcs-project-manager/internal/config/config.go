/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/common/envs"
)

// EtcdConfig 依赖的 etcd 服务的配置
type EtcdConfig struct {
	EtcdEndpoints string `yaml:"endpoints" usage:"endpoints of etcd"`
	EtcdCert      string `yaml:"cert" usage:"cert file of etcd"`
	EtcdKey       string `yaml:"key" usage:"key file for etcd"`
	EtcdCa        string `yaml:"ca" usage:"ca file for etcd"`
}

// MongoConfig
type MongoConfig struct {
	Address        string `yaml:"address"`
	ConnectTimeout uint   `yaml:"connecttimeout"`
	Database       string `yaml:"database"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	MaxPoolSize    uint   `yaml:"maxpoolsize"`
	MinPoolSize    uint   `yaml:"minpoolsize"`
	Encrypted      bool   `yaml:"encrypted"`
}

// ServerConfig 服务的配置
type ServerConfig struct {
	UseLocalIP      bool   `yaml:"useLocalIP" usage:"是否使用 Local IP"`
	Address         string `yaml:"address" usage:"server address"`
	InsecureAddress string `yaml:"insecureAddress" usage:"insecurue server address"`
	Port            int    `yaml:"port" usage:"grpc port"`
	HTTPPort        int    `yaml:"httpPort" usage:"http port"`
	MetricPort      int    `yaml:"metricPort" usage:"metric port"`
	Cert            string `yaml:"cert" usage:"server cert"`
	CertPwd         string `yaml:"certPwd" usage:"server cert password"`
	Key             string `yaml:"key" usage:"server key"`
	Ca              string `yaml:"ca" usage:"server ca"`
}

// ClientConfig 客户端配置
type ClientConfig struct {
	Cert    string `yaml:"cert" usage:"client cert"`
	CertPwd string `yaml:"certPwd" usage:"client cert password"`
	Key     string `yaml:"key" usage:"client key"`
	Ca      string `yaml:"ca" usage:"client ca"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level         string `yaml:"level" usage:"log level"`
	FlushInterval int    `yaml:"flushInterval" usage:"interval of flush"`
	Path          string `yaml:"path" usage:"log file path, example: /tmp/logs"`
	Name          string `yaml:"name" usage:"log file name, example: cr.log"`
	Size          int    `yaml:"size" usage:"log file size, unit: MB"`
	Age           int    `yaml:"age" usage:"log reserved age, unit: day"`
	Backups       int    `yaml:"backups" usage:"the count of reserved log"`
	AlsoToStdErr  bool   `yaml:"alsoToStdErr" usage:"also write log to stdErr"`
}

// SwaggerConfig
type SwaggerConfig struct {
	Enable bool   `yaml:"enable" usage:"enable swagger"`
	Dir    string `yaml:"dir" usage:"swagger dir"`
}

// JwtConfig
type JWTConfig struct {
	Enable         bool   `yaml:"enable" usage:"enable jwt"`
	PublicKey      string `yaml:"publicKey" usage:"public key"`
	PublicKeyFile  string `yaml:"publicKeyFile" usage:"public key file"`
	PrivateKey     string `yaml:"privateKey" usage:"private key"`
	PrivateKeyFile string `yaml:"privateKeyFile" usage:"private key file"`
}

// AppConfig  app code and secret for requesting third service api
type AppConfig struct {
	Code   string `yaml:"code" usage:"app code"`
	Secret string `yaml:"secret" usage:"app secret"`
}

// IAMConfig iam操作需要的配置
type IAMConfig struct {
	GatewayHost string `yaml:"gatewayHost" usage:"gateway host"`
	UseGWHost   bool   `yaml:"useGWHost" usage:"use gatewayHost when true, else use iamHost and bkPaaSHost"`
	IAMHost     string `yaml:"iamHost" usage:"iam host"`
	BKPaaSHost  string `yaml:"bkPaaSHost" usage:"bk paas host"`
	Debug       bool   `yaml:"debug" usage:"debug mode"`
}

// ClientActionsConfig Client级别的访问
type ClientActionsConfig struct {
	ClientID string   `yaml:"clientID" usage:"client id"`
	Actions  []string `yaml:"actions" usage:"action name"`
}

// ClientActionExemptPermConfig 非用户态跳过指定动作的权限
type ClientActionExemptPermConfig struct {
	ClientActions []ClientActionsConfig `yaml:"clientActions" usage:"exempt perm for client id action, example: [{clientID: bcs-inner, actions: [create, view]}]"`
}

// 请求的 CMDB 服务配置
type CMDBConfig struct {
	BKSupplierAccount string `yaml:"bkSupplierAccount" usage:"developer account"`
	Host              string `yaml:"host" usage:"access cmdb api host"`
	Timeout           int    `yaml:"timeout" usage:"request cmdb api timeout"`
	Proxy             string `yaml:"proxy" usage:"proxy"`
	Debug             bool   `yaml:"debug" usage:"debug"`
}

// 请求的 bcs cc 服务配置
type BCSCCConfig struct {
	Host string `yaml:"host" usage:"access bcs cc api host"`
}

// ProjectConfig 项目的配置信息
type ProjectConfig struct {
	Etcd                   EtcdConfig                   `yaml:"etcd"`
	Mongo                  MongoConfig                  `yaml:"mongo"`
	Log                    LogConfig                    `yaml:"log"`
	Swagger                SwaggerConfig                `yaml:"swagger"`
	Server                 ServerConfig                 `yaml:"server"`
	Client                 ClientConfig                 `yaml:"client"`
	JWT                    JWTConfig                    `yaml:"jwt"`
	IAM                    IAMConfig                    `yaml:"iam"`
	ClientActionExemptPerm ClientActionExemptPermConfig `yaml:"clientActionExemptPerm"`
	CMDB                   CMDBConfig                   `yaml:"cmdb"`
	BCSCC                  BCSCCConfig                  `yaml:"bcscc"`
	App                    AppConfig                    `yaml:"app"`
}

func (conf *ProjectConfig) initServerAddress() error {
	// 若指定使用 LOCAL_IP 且环境变量中 LOCAL_IP 有值，则替换掉 Server.Address
	if conf.Server.UseLocalIP && envs.LocalIP != "" {
		conf.Server.Address = envs.LocalIP
		conf.Server.InsecureAddress = envs.LocalIP
	}
	return nil
}

// GlobalConf 项目配置信息，全局都可以使用
var GlobalConf *ProjectConfig

// LoadConfig 通过制定的path，加载对应的配置选项
func LoadConfig(filePath string) (*ProjectConfig, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	conf := &ProjectConfig{}
	if err = yaml.Unmarshal(yamlFile, conf); err != nil {
		return nil, err
	}
	// 初始化服务地址
	conf.initServerAddress()
	// 用于后续的使用
	GlobalConf = conf
	return conf, nil
}
