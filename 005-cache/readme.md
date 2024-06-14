# ref

- [GeeCache](https://geektutu.com/post/geecache.html)
- [缓存模式](https://codeahoy.com/2017/08/11/caching-strategies-and-how-to-choose-the-right-one/)

## 缓存模式 

https://codeahoy.com/2017/08/11/caching-strategies-and-how-to-choose-the-right-one/

- Cache-Aside:最常用,优先读cache,cache miss 才读数据库,适用于读多写少的场景
- Read-Through Cache:与Cache-Aside类似,不过这里cache和db的data model需保持一致,一般通过第三方中间件/工具同步数据到cache而不是应用自己处理,适用于读多写少的场景
- Write-Through Cache:先写cache,再由cache写db,保证一致性,可结合read-through使用(DynamoDB Accelerator实现原理)
- Write-Around:数据直接写db,需要读的才写cache
- Write-Back or Write-Behind: 类似write-through,不过这里写db是异步操作,减少写库的延迟