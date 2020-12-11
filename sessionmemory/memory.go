package sessionmemory

import (
	"container/list"
	"sync"
	"time"

	seesioncookie "github.com/DesistDaydream/GoWeb/sessioncookie"
)

var pder = &FromMemory{list: list.New()}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	//注册  memory 调用的时候一定有一致
	seesioncookie.Register("sessionmemory", pder)
}

// SessionStore session实现
type SessionStore struct {
	sid              string                      //session id 唯一标示
	LastAccessedTime time.Time                   //最后访问时间
	value            map[interface{}]interface{} //session 里面存储的值
}

// Set 设置
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

// Get 获取session
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

// Delete 删除
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

// SessionID 标识符
func (st *SessionStore) SessionID() string {
	return st.sid
}

// FromMemory session来自内存实现
type FromMemory struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做 gc
}

// SessionInit 初始化
func (frommemory *FromMemory) SessionInit(sid string) (seesioncookie.Session, error) {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, LastAccessedTime: time.Now(), value: v}
	element := frommemory.list.PushBack(newsess)
	frommemory.sessions[sid] = element
	return newsess, nil
}

// SessionRead 读取
func (frommemory *FromMemory) SessionRead(sid string) (seesioncookie.Session, error) {
	if element, ok := frommemory.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := frommemory.SessionInit(sid)
		return sess, err
	}
}

// SessionDestroy 摧毁
func (frommemory *FromMemory) SessionDestroy(sid string) error {
	if element, ok := frommemory.sessions[sid]; ok {
		delete(frommemory.sessions, sid)
		frommemory.list.Remove(element)
		return nil
	}
	return nil
}

// SessionGC 垃圾收集
func (frommemory *FromMemory) SessionGC(maxLifeTime int64) {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	for {
		element := frommemory.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).LastAccessedTime.Unix() + maxLifeTime) <
			time.Now().Unix() {
			frommemory.list.Remove(element)
			delete(frommemory.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

// SessionUpdate 更新
func (frommemory *FromMemory) SessionUpdate(sid string) error {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	if element, ok := frommemory.sessions[sid]; ok {
		element.Value.(*SessionStore).LastAccessedTime = time.Now()
		frommemory.list.MoveToFront(element)
		return nil
	}
	return nil
}
