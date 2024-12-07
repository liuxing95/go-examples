# 所以 Mutex 读写锁 只是加锁用的 而 WaitGroup 才是进程间通信的方式？

是的，你的理解是正确的。以下是对 Mutex 和 WaitGroup 的详细解释：
Mutex（互斥锁）
用途：Mutex 是用于保护共享资源的同步原语。它确保在同一时间只有一个 goroutine 可以访问被保护的资源，从而防止数据竞争。
功能：通过 Lock() 和 Unlock() 方法来加锁和解锁。只有在持有锁的情况下，其他 goroutine 才能访问被保护的资源。
适用场景：当你需要确保对共享数据的安全访问时，使用 Mutex 是合适的。
WaitGroup
用途：WaitGroup 是用于等待一组 goroutine 完成的同步原语。它允许主 goroutine 等待其他 goroutine 完成其工作。
功能：通过 Add(), Done(), 和 Wait() 方法来管理 goroutine 的计数。Add() 增加计数，Done() 减少计数，Wait() 阻塞直到计数为零。
适用场景：当你需要在主 goroutine 中等待多个 goroutine 完成时，使用 WaitGroup 是合适的。
总结
Mutex 主要用于保护共享资源，防止数据竞争。
WaitGroup 主要用于协调 goroutine 之间的执行，确保主 goroutine 等待所有相关的 goroutine 完成。
这两者可以结合使用，以确保在并发编程中既能安全地访问共享数据，又能有效地管理 goroutine 的生命周期。

# WaitGroup 跟 channel 的区别是什么

WaitGroup
用途：WaitGroup 用于等待一组 goroutine 完成。它允许主 goroutine 等待其他 goroutine 完成其工作。
功能：
使用 Add(n int) 方法增加计数，表示有 n 个 goroutine 将要运行。
使用 Done() 方法减少计数，表示一个 goroutine 完成。
使用 Wait() 方法阻塞，直到计数为零，表示所有 goroutine 都已完成。
适用场景：当你需要在主 goroutine 中等待多个 goroutine 完成时，使用 WaitGroup 是合适的。

Channel
用途：channel 用于在 goroutine 之间传递数据。它可以用来实现 goroutine 之间的通信和同步。
功能：
可以通过 make(chan Type) 创建一个 channel，Type 是要传递的数据类型。
使用 <- 操作符发送和接收数据，例如 ch <- value 发送数据，value := <-ch 接收数据。
可以使用 close(ch) 关闭 channel，表示不再发送数据。
适用场景：当你需要在 goroutine 之间传递数据或信号时，使用 channel 是合适的。

# channel 的使用场景是什么？
