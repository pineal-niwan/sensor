package lru

import (
	"container/list"
	"sync"
	"time"
)

//LRU hash -- 最近最久没使用算法的哈希
//最近没有被使用的在容量到达限制后会被移除
type LruHash struct {
	size        int
	elementList *list.List
	elementHash map[string]*list.Element
	sync.RWMutex
}

//线索化的元素
type _ElementInfo struct {
	key       string      //键
	value     interface{} //值
	timestamp int64       //unix stamp -- 以毫秒计数
}

//新建，指定了hash容量大小
func NewLruHash(size int) *LruHash {
	c := &LruHash{
		size:        size,
		elementList: list.New(),
		elementHash: make(map[string]*list.Element),
	}
	return c
}

//移除指定的节点
func (c *LruHash) removeNode(elementNode *list.Element, elementInfo *_ElementInfo) {
	if elementNode != nil {
		//从list中移除
		c.elementList.Remove(elementNode)
		//从hash中移除
		delete(c.elementHash, elementInfo.key)
	}
}

//移除末尾的元素
func (c *LruHash) removeTail() {
	elementNode := c.elementList.Back()
	element := elementNode.Value.(*_ElementInfo)
	c.removeNode(elementNode, element)
}

//删除key
func (c *LruHash) remove(key string) (value interface{}, ok bool) {
	var elementNode *list.Element
	elementNode, ok = c.elementHash[key]
	if ok {
		element := elementNode.Value.(*_ElementInfo)
		value = element.value
		c.removeNode(elementNode, element)
	}
	return
}

//设置
func (c *LruHash) set(key string, value interface{}) {
	c.setWithTimestamp(key, value, 0)
}

//设置 - 带时间参数 - timestamp以毫秒计
func (c *LruHash) setWithTimestamp(key string, value interface{}, timestamp int64) {
	elementNode, ok := c.elementHash[key]

	if ok {
		//已经有此元素，只需要刷新
		//将设置的元素移动到最前面
		element := elementNode.Value.(*_ElementInfo)
		element.value = value
		element.timestamp = timestamp
		c.elementList.MoveToFront(elementNode)
		return
	}

	//没有此元素，需要加入到list和hash中，list加最前面
	elementInfo := &_ElementInfo{key, value, timestamp}
	elementNode = c.elementList.PushFront(elementInfo)
	c.elementHash[key] = elementNode

	if c.elementList.Len() > c.size {
		//超过限制了，将末尾的元素移除
		c.removeTail()
	}
}

//获取
func (c *LruHash) get(key string) (value interface{}, ok bool) {
	var elementNode *list.Element
	elementNode, ok = c.elementHash[key]
	if ok {
		element := elementNode.Value.(*_ElementInfo)
		timestamp := element.timestamp
		value = element.value
		if timestamp > 0 {
			//需要检查timeout
			nowMilUnixStamp := time.Now().UnixNano() / int64(time.Millisecond)
			if timestamp < nowMilUnixStamp {
				//已经超时
				c.removeNode(elementNode, element)
				ok = false
				return
			}
		}
		//将其移动到list头部
		c.elementList.MoveToFront(elementNode)
	}
	return
}

//获取并刷新timestamp - timestamp以毫秒计
func (c *LruHash) getThenRefreshTimestamp(key string, nowTimestamp int64, newTimestamp int64) (
	value interface{}, ok bool) {

	var elementNode *list.Element
	elementNode, ok = c.elementHash[key]
	if ok {
		element := elementNode.Value.(*_ElementInfo)
		timestamp := element.timestamp
		value = element.value
		if timestamp > 0 {
			//需要检查timeout
			if timestamp < nowTimestamp {
				//已经超时
				c.removeNode(elementNode, element)
				ok = false
				return
			}
		}
		element.timestamp = newTimestamp
		//将其移动到list头部
		c.elementList.MoveToFront(elementNode)
	}
	return
}

//清除
func (c *LruHash) clear() {
	c.elementHash = make(map[string]*list.Element)
	c.elementList.Init()
}

//清仓遍历处理
func (c *LruHash) iterateThenRemove(f func(key string, value interface{})) {
	if f == nil {
		//没有回调函数，直接清理掉
		c.clear()
		return
	}
	for lastNode := c.elementList.Back(); lastNode != nil; lastNode = lastNode.Prev() {
		elementInfo := lastNode.Value.(*_ElementInfo)
		f(elementInfo.key, elementInfo.value)
		c.removeNode(lastNode, elementInfo)
	}
}

//公共函数-获取
func (c *LruHash) Get(key string) (value interface{}, ok bool) {
	c.Lock()
	defer c.Unlock()
	value, ok = c.get(key)
	return
}

//公共函数-获取并刷新timeout -- timeout以ms计
func (c *LruHash) GetThenRefreshTimeout(key string, timeout int64) (value interface{}, ok bool) {
	nowTimeStamp := time.Now().UnixNano() / int64(time.Millisecond)
	newTimeStamp := nowTimeStamp + timeout
	c.Lock()
	defer c.Unlock()
	value, ok = c.getThenRefreshTimestamp(key, nowTimeStamp, newTimeStamp)
	return
}

//公共函数-设置
func (c *LruHash) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.set(key, value)
}

//公共函数-设置-带超时 -- timeout以ms计
func (c *LruHash) SetWithTimeout(key string, value interface{}, timeout int64) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestamp += timeout
	c.Lock()
	defer c.Unlock()
	c.setWithTimestamp(key, value, timestamp)
}

//公共函数-删除
func (c *LruHash) Remove(key string) (value interface{}, ok bool) {
	c.Lock()
	defer c.Unlock()
	value, ok = c.remove(key)
	return
}

//公共函数-移除末尾的元素--最近最久没使用
func (c *LruHash) RemoveTail() {
	c.Lock()
	defer c.Unlock()
	c.removeTail()
}

//公共函数-清除
func (c *LruHash) Clear() {
	c.Lock()
	defer c.Unlock()
	c.clear()
}

//公共函数-清仓遍历处理
func (c *LruHash) IterateThenRemove(f func(key string, value interface{})) {
	c.Lock()
	defer c.Unlock()
	c.iterateThenRemove(f)
}

//公共函数-长度 包括 hash长度和list长度
func (c *LruHash) Len() (hashLen int, listLen int) {
	c.RLock()
	defer c.RUnlock()
	hashLen = len(c.elementHash)
	listLen = c.elementList.Len()
	return
}
