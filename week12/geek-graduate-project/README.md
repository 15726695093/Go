# Geekband Graduate Project



## 目录

| 内容                                                         |
| ------------------------------------------------------------ |
| [毕业项目文档](https://github.com/15726695093/go/tree/main/week12/geek-graduate-project/README.md)，其实就是当前文件... |
| [毕业总结](https://github.com/15726695093/go/tree/main/week12/geek-graduate-project/SUMMARY.md) |





## 题目

对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：

- 微服务架构（BFF、Service、Admin、Job、Task 分模块）

  具体体现在[服务分层](#服务分层)部分

- API 设计（包括 API 定义、错误码规范、Error 的使用）

  具体体现在[接口与错误定义](#接口与错误定义)部分

- gRPC 的使用

  具体体现在[业务代码结构](#业务代码结构)部分

- Go 项目工程化（项目结构、DI、代码分层、ORM 框架）

  具体体现在[业务代码结构](#业务代码结构)部分

- 并发的使用（errgroup 的并行链路请求）

  具体体现在[admin](#admin)部分

- 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）

  只实现了MQ中间件的使用....具体体现在[打卡记录服务](#打卡记录服务(record-service)), [worktime-job](#worktime-job)两部分

- 缓存的使用优化（一致性处理、Pipeline 优化）

  具体体现在[worktime-service](#worktime-service)的查询用户工时部分

## 项目设定

实现一个简单的打卡引用，包含以下内容

* 管理端
  * 查看指定用户的指定日期的打卡时长
* 一般用户
  * 注册与登录
  * 上班/下班打卡



## 框架使用

微服务框架

* kratos version v2.1.4

数据库部分

| 数据库    | ORM                 |
| --------- | ------------------- |
| mysql 5.7 | entgo.io/ent v0.9.1 |

中间件部分

| 数据库       | SDK                       |
| ------------ | ------------------------- |
| redis 5.0.7  | go-redis/redis/v8 v8.11.4 |
| rabbitmq 3.8 | streadway/amqp v1.0.0     |



## 项目结构

### 服务分层

针对[项目设定](#项目设定)，对项目进行以下分层

BFF

| 名称              | 描述                                        | 代码位置                                                     |
| ----------------- | ------------------------------------------- | ------------------------------------------------------------ |
| clockin-interface | 普通用户BFF接口，提供http与grpc协议的接口   | 业务逻辑: /app/clockin/interface<br>接口定义: /api/clockin/interface/v1 |
| clockin-admin     | 管理员用户BFF接口，提供http与grpc协议的接口 | 业务逻辑: /app/clockin/admin<br/>接口定义: /api/clockin/admin/v1 |



内部微服务

| 名称             | 描述                                                         | 代码位置                                                     |
| ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| user-service     | 用户服务，管理用户信息，提供grpc协议的接口                   | 业务逻辑: /app/user/service<br/>接口定义: /api/user/v1       |
| record-service   | 打卡记录服务，记录用户每次打开的详情，提供grpc协议的接口，提供grpc协议的接口 | 业务逻辑: /app/record/service<br/>接口定义: /api/record/v1   |
| worktime-service | 工作时长服务，记录用户每天打卡计算得出的工作时长，提供grpc协议的接口 | 业务逻辑: /app/worktime/service<br/>接口定义: /api/worktime/service/v1 |
| worktime-job     | 用户下班打卡后触发的当天的工作时长计算的任务                 | 业务逻辑: /app/worktime/job<br/>接口定义: /api/worktime/job/v1 |



### 服务结构

采取[kratos-layout](https://github.com/go-kratos/kratos-layout)推荐形式，结合[beer-shop](https://github.com/go-kratos/beer-shop)项目的实践形式



#### 接口与错误定义

使用proto buffer来定义接口与接口返回的错误，放置于`/api/<服务名称>/<服务类型>/<接口版本>`路径下，分别使用`*.proto`来定义接口信息和`*_error.proto`来定义错误信息

在定义好接口与错误信息后，使用`kratos proto client <protobuf文件路径>`来对proto buffer进行go文件转译

实际目录接口举例：

```bash
api/worktime/service/
└── v1
    ├── worktime_error_errors.pb.go
    ├── worktime_error.pb.go
    ├── worktime_error.proto
    ├── worktime_grpc.pb.go
    ├── worktime.pb.go
    └── worktime.proto
```



#### 业务代码结构

根据[kratos-layout](https://github.com/go-kratos/kratos-layout)与[beer-shop](https://github.com/go-kratos/beer-shop)项目的最佳实践，每个服务会按以下分层

* cmd

  项目启动文件main.go

  构建依赖关系的wire.go，使用[google/wire](https://github.com/google/wire)项目来实现项目启动时的依赖注入的生命周期管理

* configs

  采取yaml文件定义配置信息，如服务启动端口，数据库连接信息等

* internal/biz

  定于以下内容：

  1. 领域对象
  2. 服务用例
  3. 与外部数据源进行交互的组件的接口

* internal/conf

  使用proto buffer定义需要读取的配置信息的格式

* internal/data

  基于internal/biz定义的与外部数据源进行交互的组件的接口，实现具体的数据交互的逻辑

* internal/service

  基于1. internal/service中定义的服务用例， 2.[api](#接口与错误定义)中protobuf编译后的go文件接口，实现以下内容：

  * 处理与传递外部请求所发过来数据
  * 调用biz层中的服务用例来响应请求，并返回请求结果

* internal/server

  * 定义http与grpc服务对象的创建方法



#### 接口测试

每个服务的接口测试都定义在`test/<服务名>/<服务类型>/*_test.go`中，如：

```go
package service

import (
	recordv1 "clock-in/api/record/v1"
	"context"
	"testing"

	"google.golang.org/grpc"
)

var ctx = context.TODO()
var client = newUserClient("127.0.0.1:9002")

func TestClockInOnWork(t *testing.T) {
	_, err := client.ClockInOnWork(ctx, &recordv1.ClockInOnWorkRequest{
		User: 13,
	})
	if err != nil {
		t.Error(err.Error())
	}
}

func newUserClient(addr string) recordv1.RecordServiceClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err.Error())
	}
	return recordv1.NewRecordServiceClient(conn)
}
```





### 服务描述



#### 用户服务(user-service)

主要实现以下服务接口

* 通过ID批量查询用户信息 - GetUserById

  从数据库中根据提交的id列表查询相印的用户信息

* 通过用户名获取该用户信息 - GetUserByName

  从数据库中根据提交的名字(精准匹配)获取用户信息，主要是用户登录时使用

* 根据用户名关键词搜索用户 - SearchUserByName

  从数据库中根据提交的名字(模糊匹配)获取用户信息

* 持久化一个用户

  创建/更新用户信息到数据库中，用户注册时使用

* 移除一个用户

  标记数据库中的用户为删除状态(软删除)，管理员移除用户时使用



#### 打卡记录服务(record-service)

主要实现以下服务接口

* 用户上班打卡

  写入一条上班类型的打卡记录到数据库中

* 用户下班打卡

  * 写入一条下班类型的打卡记录到数据库中
  * 触发mq，作为消息队列生产者，给worktime-job服务发送一条计算用户当天工时计算的异步任务

  rabbitmq生产者实现代码如下，代码位置是[/app/record/service/internal/data/record.go](https://github.com/AkiOuma/geek-graduate-project/blob/main/app/record/service/internal/data/record.go)

  ```go
  func (r *recordRepo) CaculateWorkTime(ctx context.Context, user int64, day int64) error {
  	rows, err := r.data.db.Record.Query().
  		Where(
  			record.UserEQ(user),
  			record.DayEQ(day),
  		).
  		All(ctx)
  	if err != nil {
  		return err
  	}
  	message, err := buildMessage(rows)
  	if err != nil {
  		return err
  	}
  	return r.data.mq.Publish(
  		"",         // exchange
  		"worktime", // routing key
  		false,      // mandatory
  		false,      // immediate
  		amqp.Publishing{
  			ContentType: "text/plain",
  			Body:        message,
  		})
  }
  ```

  



#### 工时服务



##### worktime-job

实现消息队列的消费者，监听record-service发送的消息，并且根据消息内容，调用worktime-service的计算用户工时服务，计算用户当天的工作时长

rabbitmq的消费者队列实现如下，代码位置是[app/worktime/job/cmd/job/main.go](https://github.com/AkiOuma/geek-graduate-project/blob/main/app/worktime/job/cmd/job/main.go)

```go
func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	uc, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	consumer := mq.NewConsumer(bc.Data)
	go func() {
		for m := range consumer {
			err := uc.CaculateWorktime(context.Background(), string(m.Body))
			if err != nil {
				stdlog.Printf("caculate worktime failed: reason:%v", err.Error())
			}
		}
	}()
	<-forever
}
```







##### worktime-service

主要实现以下服务接口

* 查询用户工时

  从数据库查询用户每天工作时长，主要用于管理端的服务

  当成功从数据库从查询到指定用户的指定时间的工时时，会向缓存中写入一条一样的数据(使用的数据格式为Hash，每个用户按id持有一个集合，每天的日期与实践作为hash map储存到集合中)。查询时会优先到缓存中查询用户有无指定日期的工时记录，若有，则从缓存中直接取出使用。若无，则回到数据库中进行查询，并将查询结果写入缓存的相应位置

  缓存部分实现的代码如下，代码位置是[/app/worktime/service/internal/data/worktime.go](https://github.com/AkiOuma/geek-graduate-project/blob/main/app/worktime/service/internal/data/worktime.go)

  ```go
  func (r *workTimeRepo) GetUserWorkTime(ctx context.Context, user int64, day []int64) ([]*biz.WorkTime, error) {
  	reply, err := r.getUserWorkTimeFromCache(ctx, user, day)
  	if err != nil {
  		reply, err = r.getUserWorkTimeFromDB(ctx, user, day)
  		if err != nil {
  			return nil, err
  		}
  		return reply, r.storeUserWorkTime2Cache(ctx, user, reply)
  	}
  	return reply, nil
  }
  
  func (r *workTimeRepo) getUserWorkTimeFromDB(ctx context.Context, user int64, day []int64) ([]*biz.WorkTime, error) {
  	rows, err := r.data.db.Worktime.Query().
  		Where(
  			worktime.UserEQ(user),
  			worktime.DayIn(day...),
  		).
  		All(ctx)
  	if err != nil {
  		return nil, err
  	}
  	return bulk2zBizWorkTime(rows), nil
  }
  
  func (r *workTimeRepo) getUserWorkTimeFromCache(ctx context.Context, user int64, day []int64) ([]*biz.WorkTime, error) {
  	val, err := r.data.rc.HGetAll(ctx, strconv.Itoa(int(user))).Result()
  	if err != nil {
  		return nil, err
  	}
  	worktime := make([]*biz.WorkTime, 0)
  	for _, v := range day {
  		s, ok := val[strconv.Itoa(int(v))]
  		if !ok {
  			return nil, ErrWorkTimeNotFoundInCache
  		}
  		minute, err := strconv.ParseInt(s, 10, 64)
  		if err != nil {
  			return nil, err
  		}
  		worktime = append(worktime, &biz.WorkTime{
  			Day:    v,
  			Minute: minute,
  		})
  	}
  	return worktime, nil
  }
  
  func (r *workTimeRepo) storeUserWorkTime2Cache(ctx context.Context, user int64, record []*biz.WorkTime) error {
  	data := make(map[string]interface{})
  	for _, v := range record {
  		data[strconv.Itoa(int(v.Day))] = v.Minute
  	}
  	return r.data.rc.HSet(ctx, strconv.Itoa(int(user)), data).Err()
  }
  ```

  

* 计算用户工时

  根据worktime-job提交的用户id与用户上下班的时间信息，计算该用户的工作时长



#### BFF

##### interface

主要实现以下服务接口(http)

* 用户登录

  提交用户名和密码到user-service中进行查询，若用户名匹配成功且密码正确，会根据用户id和用户名生成一个token(这里偷懒简单使用了base64.....)并返回给客户端，其余接口需要在http请求头中获取Authorization字段中的token，进行解密后，将用户id写入context中才可进入业务逻辑，否则将中断请求

* 用户注册

  调用user-service的持久化用户服务

* 用户上班打卡

  调用record-service的用户上班打卡服务

* 用户下班打卡

  调用record-service的用户下班打卡服务



##### admin

主要实现以下服务接口(http)

* 查询指定用户指定日期的打卡记录

  这里采取了并行请求的模式，每次请求的内容为单个用户的多个日期的工时记录，使用errgroup来统一管理错误信息，实现如下（代码位置是[/app/clockin/admin/internal/biz/clockin.go](https://github.com/AkiOuma/geek-graduate-project/blob/main/app/clockin/admin/internal/biz/clockin.go)）：

```go
func (uc *ClockinUsecase) GetWorkTime(ctx context.Context, user []int64, day []int64) ([]*UserWorkTime, error) {
	worktime := make([]*UserWorkTime, len(user))
	g, ctx := errgroup.WithContext(ctx)
	for k, v := range user {
		index := k
		userId := v
		g.Go(func() error {
			reply, err := uc.repo.GetUserWorkTime(ctx, userId, day)
			if err != nil {
				return err
			}
			worktime[index] = &UserWorkTime{
				User:     userId,
				Worktime: reply,
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return worktime, nil
}
```





## 遗留问题

由于这次大作业和公司年前要完成的新坑的时间冲突了....所以完成的比较仓促，还有许多问题待解决....



1. user和record服务在接口定义忘了应该还有用来区分类别的service层了.....
2. 原本是想用kafka作为消息中间件的，但是不知道为啥突然docker的kafka还是zookeeper除了点状况搞到用不了，重新拉起部署似乎还是没解决....就先用rabbitmq顶上mq部分的中间件了QAQ
3. worktime-job中的格式没有严格安装beer-shop中的job模式来写...因为rabbitmq的sdk的使用形式的原因，并且时间有点紧没来得及研究怎么封装一下变成最佳实践中的模式.....
4. worktime-service中的缓存还有进一步优化的空间，当要查询的数据在缓存和数据库中都没有时，每次查询的时候还是会两个库都要经历查询，针对这个问题还需要进行后序优化
5. 用户登录时返回的令牌没有使用jwt而是简单使用了base64加密，后续需要改为jwt保证安全性
6. 验证用户token的时候，突然没想明白中断请求直接返回信息的正确姿势是什么(像gin可以直接在它定义的context中直接调用abort方法来直接中断请求....)....需要后面再熟悉一下kratos的中间件设计来完善
7. 对metadata的时候还有一些疑问，因此没有直接用上，直接把校验后的用户id直接用最原始的方法塞入上下文对象中了....之后需要改造这部分的内容...
8. 本次作业没有考虑到logger的一些具体的实现和熔断限流等问题....


