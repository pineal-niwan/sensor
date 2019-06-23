package lru

import (
	"container/list"
	"sync"
	"time"
)

//LRU int64 as hash -- 最近最久没使用算法的哈希 -- int64当作key
//最近没有被使用的在容量到达限制后会被移除
type LruInt64AsKeyHash struct {
	size        int
	elementList *list.List
	elementHash map[int64]*list.Element
	sync.RWMutex
}

//线索化的元素
type _ElementInt64AsKeyInfo struct {
	key       int64       //键
	value     interface{} //值
	timestamp int64       //unix stamp -- 以毫秒计数
}

//新建，指定了hash容量大小
func NewLruInt64AsKeyHash(size int) *LruInt64AsKeyHash {
	c := &LruInt64AsKeyHash{
		size:        size,
		elementList: list.New(),
		elementHash: make(map[int64]*list.Element),
	}
	return c
}

//移除指定的节点
func (c *LruInt64AsKeyHash) removeNode(elementNode *list.Element, elementInfo *_ElementInt64AsKeyInfo) {
	if elementNode != nil {
		//从list中移除
		c.elementList.Remove(elementNode)
		//从hash中移除
		delete(c.elementHash, elementInfo.key)
	}
}

//移除末尾的元素
func (c *LruInt64AsKeyHash) removeTail() {
	elementNode := c.elementList.Back()
	element := elementNode.Value.(*_ElementInt64AsKeyInfo)
	c.removeNode(elementNode, element)
}

//删除key
func (c *LruInt64AsKeyHash) remove(key int64) (value interface{}, ok bool) {
	var elementNode *list.Element
	elementNode, ok = c.elementHash[key]
	if ok {
		element := elementNode.Value.(*_ElementInt64AsKeyInfo)
		value = element.value
		c.removeNode(elementNode, element)
	}
	return
}

//设置
func (c *LruInt64AsKeyHash) set(key int64, value interface{}) {
	c.setWithTimestamp(key, value, 0)
}

//设置 - 带时间参数 - timestamp以毫秒计
func (c *LruInt64AsKeyHash) setWithTimestamp(key int64, value interface{}, timestamp int64) {
	elementNode, ok := c.elementHash[key]

	if ok {
		//已经有此元素，只需要刷新
		//将设置的元素移动到最前面
		element := elementNode.Value.(*_ElementInt64AsKeyInfo)
		element.value = value
		element.timestamp = timestamp
		c.elementList.MoveToFront(elementNode)
		return
	}

	//没有此元素，需要加入到list和hash中，list加最前面
	elementInfo := &_ElementInt64AsKeyInfo{key, value, timestamp}
	elementNode = c.elementList.PushFront(elementInfo)
	c.elementHash[key] = elementNode

	if c.elementList.Len() > c.size {
		//超过限制了，将末尾的元素移除
		c.removeTail()
	}
}

//获取
func (c *LruInt64AsKeyHash) get(key int64) (value interface{}, ok bool) {
	var elementNode *list.Element
	elementNode, ok = c.elementHash[key]
	if ok {
		element := elementNode.Value.(*_ElementInt64AsKeyInfo)
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
func (c *LruInt64AsKeyHash) getThenRefreshTimestamp(key int64, nowTimestamp int64, newTimestamp int64) (
	value interface{}, ok bool) {

	var elementNode *list.Element
	elementNode, ok = c.elementHash[key]
	if ok {
		element := elementNode.Value.(*_ElementInt64AsKeyInfo)
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
func (c *LruInt64AsKeyHash) clear() {
	c.elementHash = make(map[int64]*list.Element)
	c.elementList.Init()
}

//清仓遍历处理
func (c *LruInt64AsKeyHash) iterateThenRemove(f func(key int64, value interface{})) {
	if f == nil {
		//没有回调函数，直接清理掉
		c.clear()
		return
	}
	for lastNode := c.elementList.Back(); lastNode != nil; lastNode = lastNode.Prev() {
		elementInfo := lastNode.Value.(*_ElementInt64AsKeyInfo)
		f(elementInfo.key, elementInfo.value)
		c.removeNode(lastNode, elementInfo)
	}
}

//公共函数-获取
func (c *LruInt64AsKeyHash) Get(key int64) (value interface{}, ok bool) {
	c.Lock()
	defer c.Unlock()
	value, ok = c.get(key)
	return
}

//公共函数-获取并刷新timeout -- timeout以ms计
func (c *LruInt64AsKeyHash) GetThenRefreshTimeout(key int64, timeout int64) (value interface{}, ok bool) {
	nowTimeStamp := time.Now().UnixNano() / int64(time.Millisecond)
	newTimeStamp := nowTimeStamp + timeout
	c.Lock()
	defer c.Unlock()
	value, ok = c.getThenRefreshTimestamp(key, nowTimeStamp, newTimeStamp)
	return
}

//公共函数-设置
func (c *LruInt64AsKeyHash) Set(key int64, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.set(key, value)
}

//公共函数-设置-带超时 -- timeout以ms计
func (c *LruInt64AsKeyHash) SetWithTimeout(key int64, value interface{}, timeout int64) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestamp += timeout
	c.Lock()
	defer c.Unlock()
	c.setWithTimestamp(key, value, timestamp)
}

//公共函数-删除
func (c *LruInt64AsKeyHash) Remove(key int64) (value interface{}, ok bool) {
	c.Lock()
	defer c.Unlock()
	value, ok = c.remove(key)
	return
}

//公共函数-移除末尾的元素--最近最久没使用
func (c *LruInt64AsKeyHash) RemoveTail() {
	c.Lock()
	defer c.Unlock()
	c.removeTail()
}

//公共函数-清除
func (c *LruInt64AsKeyHash) Clear() {
	c.Lock()
	defer c.Unlock()
	c.clear()
}

//公共函数-清仓遍历处理
func (c *LruInt64AsKeyHash) IterateThenRemove(f func(key int64, value interface{})) {
	c.Lock()
	defer c.Unlock()
	c.iterateThenRemove(f)
}

//公共函数-长度 包括 hash长度和list长度
func (c *LruInt64AsKeyHash) Len() (hashLen int, listLen int) {
	c.RLock()
	defer c.RUnlock()
	hashLen = len(c.elementHash)
	listLen = c.elementList.Len()
	return
}
