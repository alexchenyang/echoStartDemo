# echo简单服务Demo


# 一、项目目录结构
- controller 控制层
- logic 逻辑层， 实际底层逻辑代码 
  - entity: 接口出入参文件
- model 与数据库交互层

- conf 存储配置文件
- config 解析配置文件代码
- common 通用包
  - constants  常量定义
  - handlers   全局句柄定义&初始化区域，如mysql全局句柄 & 初始化
  - util       功能包定义
