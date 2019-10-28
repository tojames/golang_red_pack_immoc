# resk

#### 介绍

#### 

慕课网《3小时极简春节抢红包之Go的实战源代码》《GO从0到1实战微服务版抢红包系统》

课程地址1《3小时极简春节抢红包之Go的实战》：https://www.imooc.com/learn/1101

课程地址2《GO从0到1实战微服务版抢红包系统》：https://coding.imooc.com/class/345.html

resk: red envelope seckill（second kill） 红包秒杀系统

分支说明，分支名称由“v+章号+节号1_节号N""组成，比如：

- v3-4 表示第三章第四节
- v4-10_12  表示第四章第10节至第12节

| 章编号 | 节编号 | 课程标题                                         | 分支     |
| ------ | ------ | ------------------------------------------------ | -------- |
| 3      | 4      | 初始源代码： 3-4 Go module模块化管理代码依赖     | v3-4     |
| 4      | 1      | 在Golang中如何设计枚举？                         | v4-1     |
| 4      | 2      | 在Golang中如何使用Json序列化：从标准库和jsoniter | v4-2     |
| 4      | 3~4    | 基础资源-配置设计-配置starter编码实践            | v4-3_4   |
| 4      | 5      | 基础资源-配置设计-启动管理器编码实践             | v4-5     |
| 4      | 6~7    | 基础设施资源-mysql starter编码实践               | v4-6_7   |
| 4      | 8~9    | 基础设施资源-log starter编码实践                 | v4-8_9   |
| 4      | 10~12  | 基础设施资源-验证器 starter编码实践              | v4-10_12 |
| 4      | 13~15  | 基础设施资源-web框架 starter编码实践             | v4-13_15 |
| 5      | 0      | 在第4章的基础上重构了代码，以便更好的运行        | v5-0     |
| 5      | 3      | 资金账户模块服务接口设计和定义                   | v5-3     |
| 5      | 4~8    | 资金账户模块-数据库访问层的定义和编码实践        | v5-4_8   |
| 5      | 9~13   | 资金账户-业务领域层的定义和编码实践              | v5-9_13  |
| 5      | 14-17  | 资金账户模块-应用服务层实现编码实践              | v5-14_21 |
| 5      | 18-21  | 资金账户模块-Web接口的定义和编码实践             | v5-14_21 |
| 6    | 4     | 红包模块-服务接口设计和定义编码实践                   | v6-4     |
| 6    | 5~7   | 红包模块-数据访问层的定义和编码实践                   | v6-5_7   |
| 6    | 8~10  | 发红包业务领域层的编码实践                            | v6-8_10  |
| 6    | 11~15 | 红包模块-发红包web和goRpc用户接口层编码实践           | v6-11_15 |
| 6    | 17~21 | 红包模块-抢红包业务领域层和应用服务层编码实践         | v6-17_21 |
| 6    | 22    | 红包模块-抢红包web和GoRPC用户接口的轻松适配和编码实践 | v6-22    |
| 6    | 23~28 | 红包模块-过期红包                                     | v6-23_28 |
| 6    | 29~31 | 红包模块-mobile H5 UI                                     | v6.29_31 |
|7|1~6|集成测试来保证代码健壮和稳定运行|v7|
|8|1~6|Golang项目打包运维部署|v8|
|9|1~7|红包微服务化架构概述和实践|v9-1_7|
|9|8~10|红包微服务化服务拆分实践：拆分后包含所有微服务代码的分支|v9-8_10|
|9|10|红包微服务化服务拆分实践：拆分后清理只包含resk微服务代码的分支，其他微服务查看account/infra/resk-ui仓库的v9-10分支|v9-10|
|10|1~9|微服务http服务发现客户端实践|v10-1_9|
|10|10|基于服务发现负载均衡的goRpc客户端编码实践：备注resk/infra|v10-10|
|11|-|微服务下分布式配置管理实践|v11|
|12|-|微服务部署管理实践|v12|
|13|-|红包系统性能分析实战演示|v13|





## 备注：

- 红包微服务化服务拆分实践后提交了2个分支，分别是：
- 分支：v9-8_10，拆分后包含所有微服务代码的分支
- 分支v9-10：拆分后清理只包含resk微服务代码的分支，其他微服务查看account/infra/resk-ui仓库的v9-10分支
- infra公共基础组件仓库地址： https://git.imooc.com/coding-345/infra.git
- account资金账户仓库地址：https://git.imooc.com/coding-345/account.git
- resk-ui红包mobile ui仓库地址：https://git.imooc.com/coding-345/resk-ui.git