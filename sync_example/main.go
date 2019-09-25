package main

//锁的状态不是互斥的，最低三位分别表示mutexLocked, mutexWoken和mutexStarving,剩下的位置都用来表示当前有多少个goroutine等待互斥锁被释放
type Mutex struct {
	state int32
	sema  uint32
}

//互斥锁被创建出来是，所有的状态都是0，当互斥锁被锁定时mutexLocked就会被置成1，当户持所在正常状态下被唤醒时mutexWoken就会被置成1，
//mutexStarving用于表示当前的互斥锁进入了状态

//互斥锁可以同时处于两种不同的模式，也就是正常模式和饥饿模式，在正常模式下，所有锁的等待者都会按照先进先出的顺序获取锁，但是如果一个
//刚刚被唤起的 Goroutine 遇到了新的 Goroutine 进程也调用了 Lock 方法时，大概率会获取不到锁，为了减少这种情况的出现，防止 Goroutine
// 被『饿死』，一旦 Goroutine 超过 1ms 没有获取到锁，它就会将当前互斥锁切换饥饿模式。

func main() {

}
