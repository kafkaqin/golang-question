package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MyLocalConfig struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func TestNewLocalManager(t *testing.T) {
	filePath := "demo.json"
	localManager := Local[MyLocalConfig](filePath)
	assert.NotNil(t, localManager)
}
func TestLocalManager(t *testing.T) {

	tests := []struct {
		name         string
		filePath     string
		initMyConfig MyLocalConfig
		newMyConfig  MyLocalConfig
	}{
		{
			name:         "local-test-key-value",
			initMyConfig: MyLocalConfig{Key: "example", Value: 1},
			filePath:     fmt.Sprintf("%d.json", time.Now().UnixNano()),
			newMyConfig:  MyLocalConfig{Key: "updated_example", Value: 2},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d: %s", i, test.name), func(t *testing.T) {
			filePath := test.filePath
			localManager := Local[MyLocalConfig](filePath)
			assert.NotNil(t, localManager)
			// 初始化数据，如果没有则设置初始值
			localManager.InitData(test.initMyConfig)
			currentConfig := localManager.Get()
			assert.Equal(t, test.initMyConfig.Key, currentConfig.Key)
			assert.Equal(t, test.initMyConfig.Value, currentConfig.Value)

			localManager.OnChange(func(cfg MyLocalConfig) {
				fmt.Printf("Config changed: %+v\n", cfg)
				//assert.Equal(t, cfg.Value, test.newMyConfig.Value)
				//assert.Equal(t, cfg.Key, test.newMyConfig.Key)
			})

			// 开始监听配置变化
			localManager.Watch()

			// 更新配置
			if err := localManager.Update(test.newMyConfig); err != nil {
				fmt.Println("Failed to update config:", err)
			}

			// 获取配置
			currentConfig = localManager.Get()
			assert.Equal(t, test.newMyConfig.Value, currentConfig.Value)
			assert.Equal(t, test.newMyConfig.Key, currentConfig.Key)
			fmt.Printf("Current config: %+v\n", currentConfig)
			time.Sleep(1 * time.Second)
			test.newMyConfig.Key = "test01"
			if err := localManager.Update(test.newMyConfig); err != nil {
				fmt.Println("Failed to update config:", err)
			}

			test.newMyConfig.Key = "test02"
			time.Sleep(1 * time.Second)
			if err := localManager.Update(test.newMyConfig); err != nil {
				fmt.Println("Failed to update config:", err)
			}

			// 等待一段时间以便观察配置的变化
			time.Sleep(100 * time.Second)
		})

	}

}
