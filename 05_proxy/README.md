代理模式在不改变原始类代码的情况下，通过引入代理类来给原始类附加功能。


监控、统计、鉴权、限流、事务、幂等、日志


# 静态代理





# 动态代理

略过些检查之类的代码，例如 User  是否真正实现了 IUser  这种情况。

- 读取文件, 获取文件的 ast 语法树
- 通过 NewCommentMap 构建 node 和 comment 的关系
- 通过 comment 是否包含 @proxy 接口名  的接口，判断该节点是否需要生成代理类
- 通过 Lookup 方法找到接口
- 循环获取接口的每个方法的，方法名、参数、返回值信息
- 将方法信息，包名、需要代理类名传递给构建好的模板文件，生成代理类
- 最后用 format 包的方法格式化源代码



