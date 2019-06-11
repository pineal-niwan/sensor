package hash

import (
	"container/list"
	"sync"
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
	key   string
	value interface{}
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
func (c *LruHash) removeNode(elementNode *list.Element) {
	if elementNode != nil {
		//从list中移除
		c.elementList.Remove(elementNode)
		elementInfo := elementNode.Value.(*_ElementInfo)
		//从hash中移除
		delete(c.elementHash, elementInfo.key)
	}
}

//移除末尾的元素
func (c *LruHash) removeTail() {
	elementNode := c.elementList.Back()
	c.removeNode(elementNode)
}

//删除key
func (c *LruHash) remove(key string) (value interface{}, ok bool) {
	var elementNode *list.Element
	elementNode, ok = c.elementHash[key]
	if ok {
		value = elementNode.Value.(*_ElementInfo).value
		c.removeNode(elementNode)
	}
	return
}

//设置
func (c *LruHash) set(key string, value interface{}) {
	elementNode, ok := c.elementHash[key]
	if ok {
		//已经有此元素，只需要刷新
		//将设置的元素移动到最前面
		elementNode.Value.(*_ElementInfo).value = value
		c.elementList.MoveToFront(elementNode)
		return
	}

	//没有此元素，需要加入到list和hash中，list加最前面
	elementInfo := &_ElementInfo{key, value}
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
		value = elementNode.Value.(*_ElementInfo).value
		//如果元素存在，需要将其移动到list头部
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
		c.removeNode(lastNode)
	}
}

//公共函数-获取
func (c *LruHash) Get(key string) (value interface{}, ok bool) {
	c.Lock()
	value, ok = c.get(key)
	c.Unlock()
	return
}

//公共函数-设置
func (c *LruHash) Set(key string, value interface{}) {
	c.Lock()
	c.set(key, value)
	c.Unlock()
}

//公共函数-删除
func (c *LruHash) Remove(key string) (value interface{}, ok bool) {
	c.Lock()
	value, ok = c.remove(key)
	c.Unlock()
	return
}

//公共函数-移除末尾的元素--最近最久没使用
func (c *LruHash) RemoveTail() {
	c.Lock()
	c.removeTail()
	c.Unlock()
}

//公共函数-清除
func (c *LruHash) Clear() {
	c.Lock()
	c.clear()
	c.Unlock()
}

//公共函数-清仓遍历处理
func (c *LruHash) IterateThenRemove(f func(key string, value interface{})) {
	c.Lock()
	c.iterateThenRemove(f)
	c.Unlock()
}

//公共函数-长度
func (c *LruHash) Len() int {
	c.RLock()
	l := c.elementList.Len()
	c.RUnlock()
	return l
}

//公共函数-hash长度
func (c *LruHash) HashLen() int {
	c.RLock()
	l := len(c.elementHash)
	c.RUnlock()
	return l
}
