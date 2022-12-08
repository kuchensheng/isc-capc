# infrastructure 基础构建层
基础构建层存储了基础模型信息以及基础的仓库信息\
包结构如下
```text
├─common 存储了通用的异常信息、全局变量等
├─connetor 主要用于与数据库建立联系
├─model 存储了模型信息和仓库信息
│  ├─api api信息的模型信息和仓库信息
│  └─category 分组信息的模型和仓库信息
└─vo 值对象，用于对信息进行检索
    ├─api api的值对象
    └─category 分组信息的值对象
```
