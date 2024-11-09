package config

import (
	"context"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang-question/errorx"
	"k8s.io/klog/v2"
	"sync"
	"time"
)

// EtcdManager 用于通过 Etcd 管理配置
type EtcdManager[T any] struct {
	endpoints       []string
	client          *clientv3.Client
	key             string
	data            T
	mutex           sync.RWMutex
	handlers        map[int]func(T)
	handlerID       int
	watchOnce       sync.Once
	cancelFunc      context.CancelFunc
	Logger          klog.Logger
	fileContentType string // json yaml yml todo
}

// NewEtcdManager 创建新的 EtcdManager
func NewEtcdManager[T any](client *clientv3.Client, key string, endpoints []string) *EtcdManager[T] {
	return &EtcdManager[T]{
		client:    client,
		key:       key,
		handlers:  make(map[int]func(T)),
		endpoints: endpoints,
		Logger:    klog.FromContext(context.Background()),
	}
}

// Get 返回当前的配置
func (e *EtcdManager[T]) Get() T {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.data
}

// Update 更新配置到 Etcd
func (e *EtcdManager[T]) Update(newData T) errorx.Error {
	dataBytes, err := json.Marshal(newData)
	if err != nil {
		return errorx.New("failed to marshal data")
	}

	_, err = e.client.Put(context.Background(), e.key, string(dataBytes))
	if err != nil {
		return errorx.New("failed to update config in etcd")
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.data = newData
	return nil
}

// OnChange 注册配置变更的回调函数
func (e *EtcdManager[T]) OnChange(handler func(T)) (cancel func()) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	id := e.handlerID
	e.handlers[id] = handler
	e.handlerID++

	return func() {
		e.mutex.Lock()
		defer e.mutex.Unlock()
		delete(e.handlers, id)
	}
}

// Watch 启动对配置变化的监听
func (e *EtcdManager[T]) Watch() Manager[T] {
	e.watchOnce.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		e.cancelFunc = cancel

		go func() {
			rch := e.client.Watch(ctx, e.key)
			for wresp := range rch {
				for _, ev := range wresp.Events {
					if ev.Type == clientv3.EventTypePut {
						var newData T
						if err := json.Unmarshal(ev.Kv.Value, &newData); err == nil {
							e.mutex.Lock()
							e.data = newData
							e.mutex.Unlock()

							e.mutex.RLock()
							for _, handler := range e.handlers {
								handler(newData)
							}
							e.mutex.RUnlock()
						}
					}
				}
			}
		}()
	})
	return e
}

// InitData 如果 Etcd 中没有数据，初始化数据
func (e *EtcdManager[T]) InitData(initialData T) Manager[T] {
	resp, err := e.client.Get(context.Background(), e.key)
	if err == nil && resp.Count == 0 {
		err := e.Update(initialData)
		if err != nil {
			e.Logger.Error(err, "InitData: Update to etcd failed")
		}
	}
	return e
}

// Etcd 返回一个 Etcd 配置管理器
func Etcd[T any](endpoints []string, configFileKey string) Manager[T] {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints, // 配置Etcd的地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic("failed to connect to etcd")
	}
	return NewEtcdManager[T](client, configFileKey, endpoints)
}
