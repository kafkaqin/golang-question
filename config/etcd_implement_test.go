package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MyConfig struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func TestNewEtcdManager(t *testing.T) {
	// 使用 Etcd 管理配置
	endpoints := []string{"localhost:2379"}
	configFileKey := "demo"
	etcdManager := Etcd[MyConfig](endpoints, configFileKey)
	defer etcdManager.(*EtcdManager[MyConfig]).client.Close()
	defer etcdManager.(*EtcdManager[MyConfig]).cancelFunc()
	assert.NotNil(t, etcdManager)
}
func TestEtcdManager(t *testing.T) {

	tests := []struct {
		name          string
		configFileKey string
		endpoints     []string
		initMyConfig  MyConfig
		newMyConfig   MyConfig
	}{
		{
			name:          "etcd-test-key-value",
			initMyConfig:  MyConfig{Key: "example", Value: 1},
			configFileKey: fmt.Sprintf("%d", time.Now().UnixNano()),
			newMyConfig:   MyConfig{Key: "updated_example", Value: 2},
			endpoints:     []string{"localhost:2379"},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d: %s", i, test.name), func(t *testing.T) {
			// 使用 Etcd 管理配置
			endpoints := test.endpoints
			configFileKey := test.configFileKey
			etcdManager := Etcd[MyConfig](endpoints, configFileKey)
			defer etcdManager.(*EtcdManager[MyConfig]).client.Close()
			assert.NotNil(t, etcdManager)
			// 初始化数据，如果没有则设置初始值
			etcdManager.InitData(test.initMyConfig)
			currentConfig := etcdManager.Get()
			assert.Equal(t, test.initMyConfig.Key, currentConfig.Key)
			assert.Equal(t, test.initMyConfig.Value, currentConfig.Value)
			// 注册配置变更监听
			etcdManager.OnChange(func(cfg MyConfig) {
				fmt.Printf("Config changed: %+v\n", cfg)
				assert.Equal(t, cfg.Value, test.newMyConfig.Value)
				assert.Equal(t, cfg.Key, test.newMyConfig.Key)
			})

			// 开始监听配置变化
			etcdManager.Watch()

			// 更新配置
			if err := etcdManager.Update(test.newMyConfig); err != nil {
				fmt.Println("Failed to update config:", err)
			}

			// 获取配置
			currentConfig = etcdManager.Get()
			assert.Equal(t, test.newMyConfig.Value, currentConfig.Value)
			assert.Equal(t, test.newMyConfig.Key, currentConfig.Key)
			fmt.Printf("Current config: %+v\n", currentConfig)

			// 等待一段时间以便观察配置的变化
			time.Sleep(1 * time.Second)
			// 更新配置
			test.newMyConfig.Value = 9
			if err := etcdManager.Update(test.newMyConfig); err != nil {
				fmt.Println("Failed to update config:", err)
			}
			// 更新配置
			time.Sleep(1 * time.Second)
			test.newMyConfig.Value = 10
			if err := etcdManager.Update(test.newMyConfig); err != nil {
				fmt.Println("Failed to update config:", err)
			}
			time.Sleep(1 * time.Second)
		})

	}

}
