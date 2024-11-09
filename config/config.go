package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"golang-question/errorx"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
	"sync"
)

type Manager[T any] interface {
	Get() T
	Update(T) errorx.Error
	OnChange(func(T)) (cancel func())
	Watch() Manager[T]     //執行Watch後，會開始監聽配置的變化，並在變化時自動更新 否則每次Get都會從數據源取得最新資料
	InitData(T) Manager[T] //如果數據源沒有資料，則使用InitData put資料
}

// LocalManager 用于本地管理配置
type LocalManager[T any] struct {
	data            T
	fileContentType string // json yaml yml todo
	mutex           sync.RWMutex
	handlers        map[int]func(T)
	handlerID       int
	watcher         *fsnotify.Watcher
	configFilePath  string
	watchOnce       sync.Once
	Logger          klog.Logger
}

// NewLocalManager 创建一个 LocalManager 实例并启动文件监听
func NewLocalManager[T any](configFilePath string) *LocalManager[T] {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	manager := &LocalManager[T]{
		handlers:       make(map[int]func(T)),
		watcher:        watcher,
		configFilePath: configFilePath,
		Logger:         klog.FromContext(context.Background()),
	}
	return manager
}

// Get 返回当前的本地配置
func (l *LocalManager[T]) Get() T {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.data
}

// Update 更新本地配置
func (l *LocalManager[T]) Update(newData T) errorx.Error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.data = newData

	// 触发回调
	for _, handler := range l.handlers {
		handler(newData)
	}
	return nil
}

// OnChange 注册本地配置变更的回调函数
func (l *LocalManager[T]) OnChange(handler func(T)) (cancel func()) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	id := l.handlerID
	l.handlers[id] = handler
	l.handlerID++

	return func() {
		l.mutex.Lock()
		defer l.mutex.Unlock()
		delete(l.handlers, id)
	}
}

// InitData 初始化本地数据
func (l *LocalManager[T]) InitData(initialData T) Manager[T] {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	// 加载文件数据
	err := l.writeDataToFile(initialData)
	if err != nil {
		l.Logger.Error(err, "failed to load data from file")
	}
	l.data = initialData
	return l
}

// writeDataToFile 保存配置内容到文件
func (l *LocalManager[T]) writeDataToFile(data T) error {

	f, err := os.OpenFile(l.configFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errorx.New(fmt.Sprintf("failed open file %s: %v", l.configFilePath, err))
	}
	defer f.Close()

	fileData, err := json.Marshal(data)
	if err != nil {
		return errorx.New(fmt.Sprintf("failed to marshal file data: %v", err))
	}

	_, err = f.Write(fileData)
	if err != nil {
		return errorx.New(fmt.Sprintf("failed write to file %s: %v", l.configFilePath, err))
	}

	return nil
}

func (l *LocalManager[T]) loadDataFromFile() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	fileData, err := ioutil.ReadFile(l.configFilePath)
	if err != nil {
		return errorx.New(fmt.Sprintf("failed to read file: %v", err))
	}

	var newData T
	err = json.Unmarshal(fileData, &newData)
	if err != nil {
		return errorx.New(fmt.Sprintf("failed to unmarshal file data: %v", err))
	}

	l.data = newData

	// 触发回调
	for _, handler := range l.handlers {
		handler(newData)
	}
	return nil
}

// watchFileChanges 监听文件变化
func (l *LocalManager[T]) watchFileChanges() {
	l.watcher.Add(l.configFilePath)
	defer l.watcher.Close()

	for {
		select {
		case event, ok := <-l.watcher.Events:
			if !ok {
				continue
			}
			if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
				// 文件内容被修改或重新创建，重新加载数据
				err := l.loadDataFromFile()
				if err != nil {
					l.Logger.Error(err, "error reloading data:")
				}
			}
		case err, ok := <-l.watcher.Errors:
			if !ok {
				continue
			}
			l.Logger.Error(err, "file watch error:")
		}
	}
}

// Watch 启动对配置变化的监听
func (l *LocalManager[T]) Watch() Manager[T] {
	l.watchOnce.Do(func() {
		go l.watchFileChanges()
	})
	return l
}

// Local 返回一个本地配置管理器
func Local[T any](filePath string) Manager[T] {
	return NewLocalManager[T](filePath)
}
