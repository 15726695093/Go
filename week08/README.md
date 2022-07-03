# redis benchmark practice

## 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

* 使用shell脚本执行以下redis benchmark工具，并把对应结果记录到同脚本目录小的`./q1/redis_($size)`文件中

* 测试结果：10字节到5k字节的默认100000次测试的结果看来get和set的性能没什么太大的差别，基本都是2s到2.2s之内完成.

## 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间

* 使用go-redis实例化客户端
* key使用`google/uuid`生成，而value全部使用单个字符0，测试以下数量大小的数据
  * 1w
  * 5w
  * 10w
  * 20w
  * 30w
  * 40w
  * 50w
* 每次测试完直接执行`info memory`查看当前内存信息，并将结果记录到./q2的对应文件中
* 执行flushdb命令清空缓存

* 测试结果:同等字符长度的key会随着数据量的增大，在内存中占据的空间越小