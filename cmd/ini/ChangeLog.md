# ChangeLog

> Joshua Conero
>
> 2019年6月6日 星期四



## 0.2

### 0.2.0/20190617	

- ini
  - (+) *首次进入应用时，自动创建 `global` 资源，且可通过配置文件控制*
  - (修复) *初始化应用，覆盖历史资源中的数据集*
  - (+) *oIRQueueManger 新增方法 `getCurIni` 和 `getNameList`*
    - (移除) `getCurIni` 方法，重写后台逻辑，简化代码的复杂性
    - (移除) *移除多个冗余的属性值*
  - (实现) `open`
    - (+) 添加可选参数 `--alias=<别名>`
    - (实现) *不同目录相同文件名可读取 (限制取消)、项目文件可通过 `alias` 加载*
    - (优化) *重写方法以简化文件加载逻辑*
  - (+) `about`
    - (+) 实现命令，用于打印【当前资源/指定资源】的信息
  - `get`
    - (修复) *读取参数时，资源获取失败*
  - (+) `new`
    - (+) *新增新的命令*
    - (+) *实现新增空的资源，用于创建新的资源*
  - (+) `set`
    - (+) *新增新的命令，用于设置/更新资源的值*
  - (+) `save`
    - (+) *新增命令，用于保存当前的资源。可以保存新文件或者/覆盖源文件*
  - (+) `del`
    - (+) *新增命名，用于删除键值*
  - (+) `dump`
    - (+) *新增命令，用于打印值得类型和值*

 

> **todo**

- ini
  - open 命令增强
    - 不同目录相同文件名可读取 (限制取消)
    - 项目文件可通过 `alias` 加载
    - 读取文件时，获取运行信息：*耗时，行数、文件大小、注释行等*
  - about 当前加载资源的信息展示(新增)





## 0.1

### 0.1.0/20190606

- 命令行项目程序搭建
- (+) 添加命令 `$ ini`
  - 实现的交互时的内部命令，如: *open, use, list, get , help*
- (+) *添加系统所需要的默认配置文件 `inigo.ini`*

