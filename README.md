# gylib
gylib是我的公司正在使用的一套分布式服务端框架，
核心部分是Library中的代码，包含了服务的核心内容，
* cache：利用redis做缓存服务器，文件，消息，用户登录session。
* logger：类似于log4j的一个文件日志系统。
* mongo：使用mongodb的数据库模块。
* param：参数解析。
* service：服务核心模块，使用grpc做远程调用，使用etcd做服务注册与发现，定义了返回值与服务。
* util：测试工具。
* uuid：使用uuid做token生成，与一些缓存中使用的临时key。
* example：两个示例项目，通过AgentServer接收http请求，通过FileServer做文件管理，文件的存储使用了mongodb gridFS。
logger, mongo, uuid都是参考了一些开源的项目，现在刚开始准备文字介绍，迟点会把原项目连接更新上来。

感谢：在QQ群里有很多高手，跟他们学到了很多新的东西。所以特别将这一套代码开源出来，不管这框架是否好，但这是自己的一份学习报告。
QQ群：
* Gopher成都（459420581）
* Leaf 游戏服务器交流群（376389675）


