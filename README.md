# gylib
gylib是我的公司正在使用的一套分布式服务端框架，
核心部分是Library中的代码，包含了服务的核心内容，
gycache：利用redis做缓存服务器，文件，消息，用户登录session。
gylogger：类似于log4j的一个文件日志系统。
gymongo：使用mongodb的数据库模块
gyparam：参数解析
gyservice：服务核心模块，使用grpc做远程调用，使用etcd做服务注册与发现，定义了返回值与服务。
gyutil：测试工具
gyuuid：使用uuid做token生成，与一些缓存中使用的临时key

gylogger, gymongo, gyuuid都是参考了一些开源的项目，现在刚开始准备文字介绍，迟点会把原项目连接更新上来。

感谢
QQ群：Gopher成都（459420581）， Leaf 游戏服务器交流群（376389675）
在里面很多高手，跟他们学到了很多新的东西。所以特别将这一套代码开源出来，不管是不是够好，就当是自己的一份学习报告。
