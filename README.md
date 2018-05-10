## pool 提供一些常用的 pool 的封装

### BytesBufferPool

```text
BytesBufferPool 提供了类似 sync.Pool 的功能，
但是它们能永久缓存指定个数的 *bytes.Buffer，不会被 GC 回收，
这些特性在某些场合有一定的作用，比如打印日志的时候，需要一个比较大的
*bytes.Buffer，一般是 16KB 或者 32KB，这样写日志的时候就不需要 grow []byte，
也不会被 GC 回收，能够重复的使用！
```
