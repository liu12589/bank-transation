# bank-transation
是个人练习项目，计划将所学到的知识以及技术栈全部融入到该项目中，目前想从三条主线出发完善该项目。
- 技术广度：涉及数据库设计、迁移、API 开发、安全、测试、部署、CI/CD、高并发、微服务、分布式等，麻雀虽小但五脏俱全。
- 技术厚度：逐步丰富功能，模块化增加各功能，例如：消息通知、排行榜、报表等功能。到达熟悉各种模块功能的实现。
- 技术深度：熟悉了整体的开发流程同时，注重积累各技术底层原理，例如：如何优雅的加锁、Paseto 实现原理、如何保证 redis 高可用等。

### 简介
- 技术栈：Go、Gin 框架、PostgreSQL、Docker、Sqlc、Paseto、Bcrypt、Viper、Gomock、Github action
- 目前已完成用户账户创建、登陆、管理模块、交易模块并进行了详尽测试，能够处理复杂场景下并发事务
- 采用 Bcrypt 对用户密码、账户信息加密处理，采用 Paseto 实现交易信息安全传输
- 使用 Gin 框架构建 RESTful HTTP API ，符合 RESTful API 接口设计标准及规范
- 使用 go mock、httpest 进行单元测试，符合测试规范。并使用 Github Action 自动化测试
- 使用 Docker 构建应用程序，并部署到阿里云服务器
- 使用 Viper 读取项目中的配置信息，降低耦合，提高程序扩展性
- 目前在学习 K8s ，下一步计划使用 ECS 创建 K8s 集群、使用 Github Action 自动构建映像并将其部署到 ECS 集群

### 相关技术总结
个人针对项目中的某些技术点做的总结
- [基于令牌的身份认证机制](http://8.142.142.69:8090/archives/%E5%9F%BA%E4%BA%8E%E4%BB%A4%E7%89%8C%E7%9A%84%E8%BA%AB%E4%BB%BD%E8%AE%A4%E8%AF%81%E6%9C%BA%E5%88%B6)
- [postgres 行锁实践分析](http://8.142.142.69:8090/archives/postgres%E8%A1%8C%E9%94%81%E5%AE%9E%E8%B7%B5%E5%88%86%E6%9E%90)
- [如何编写测试模块](http://8.142.142.69:8090/archives/%E5%A6%82%E4%BD%95%E7%BC%96%E5%86%99%E6%B5%8B%E8%AF%95%E6%A8%A1%E5%9D%97)
- [gin 中间件](http://8.142.142.69:8090/archives/gin%E4%B8%AD%E9%97%B4%E4%BB%B6)
- [gin 路由](http://8.142.142.69:8090/archives/ginlu-you)
- [如何使用 docker 部署项目](http://8.142.142.69:8090/archives/2021-12-01-21-59-55)

### 配置项目运行环境

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [DB Docs](https://dbdocs.io/docs)

    ```bash
    npm install -g dbdocs
    dbdocs login
    ```

- [DBML CLI](https://www.dbml.org/cli/#installation)

    ```bash
    npm install -g @dbml/cli
    dbml2sql --version
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

- [Gomock](https://github.com/golang/mock)

    ``` bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```

### 创建项目基础环境（具体命令可参考 MakeFile 文件）

- 创建一个 postgres 容器:

    ```bash
    make postgres
    ```

- 创建 bank-transation 数据库:

    ```bash
    make createdb
    ```

- 迁移所有的数据库表:

    ```bash
    make migrateup
    ```

### 如何生成代码（该代码本项目已经生成，因此可以不要运行）

- 使用 Sqlc 生成 SQL 增删改查语句:

    ```bash
    make sqlc
    ```

- 使用 go mock 生成测试代码:

    ```bash
    make mock
    ```


### 运行项目

- 运行 server:

    ```bash
    make server
    ```

- 运行测试:

    ```bash
    make test
    ```
