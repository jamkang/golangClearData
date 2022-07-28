# 一、amqp基本使用和原理
>go get -v github.com/streadway/amqp 安裝好amqp包

## 1、生产者的流程
1. 连接Connection `connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")`
   - 参数形式：name:pwd@hostname:port/
2. 创建Channel `channel, err := connection.Channel()`
3. 创建或连接一个交换器 `err = channel.ExchangeDeclare("e1", "direct", true, false, false, true, nil)`
   - name 交换机名称
   - kind 交换机类型
   - durable 持久化标识
   - autoDelete 是否自动删除
   - internal 是否是内置交换机
   - noWait 是否等待服务器确认
   - args 其它配置
4. 创建或连接一个队列 `q, err := channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, nil)`
   - name 队列名称
   - durable 持久化
   - autoDelete 自动删除
   - 自动删除的前提是:至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。
   - exclusive 排他
   - 排他队列只对首次创建它的连接可见，排他队列是基于连接（Connection可见的，并且该连接内的所有信道（Channel都可以访问这个排他队列,在这个连接断开之后，该队列自动删除，由此可见这个队列可以说是绑到连接上的，对同一服务器的其他连接不可见
   - noWait 是否等待服务器确认
   - args Table
5. 交换器绑定队列 `err = channel.QueueBind(name, key BindingKey, exchange, noWait, nil)`
   - name 队列名称
   - key BindingKey 根据交换机类型来设定
   - exchange 交换机名称
   - noWait 是否等待服务器确认
   - args Table
6. 投递消息</br>`
   err = channel.Publish("e1", "q1Key", true, false, amqp.Publishing{
   Timestamp:   time.Now(),
   DeliveryMode: amqp.Persistent, //Msg set as persistent
   ContentType: "text/plain",
   Body:        []byte("Hello Golang and AMQP(Rabbitmq)!"),
   })`
    - exchange 交换器名称
    - key RouterKey 根据交换机类型来设定
    - mandatory 是否为无法路由的消息进行返回处理,设置为 true 表示将消息返回到生产者，否则直接丢弃消息
    - immediate 是否对路由到无消费者队列的消息进行返回处理 RabbitMQ 3.0 废弃
    - msg 消息体
7. 关闭 Channel
8. 关闭 Connection

## 2、消费者模式
### 2.1 拉模式
由消费者主动拉取信息来消费，每次只消费一条消息，同样也需要进行 ack 确认消费。</br>
`channel.Get(queue string, autoAck bool)`
### 2.2 推模式
推模式是通过持续订阅的方式来消费信息， Consume 将信道( Channel )设置为接收模式，
直到取消队列的订阅为止。在接收模式期间， RabbitMQ 会不断地推送消息给消费者。
推送消息的个数还是会受到 channel.Qos 的限制。</br>
`deliveries, err := channel.Consume("q1", "any", false, false, false, true, nil)`
- queue 队列名称 
- consumer 消息者名称 
- autoAck 是否确认消费,设为 true 则消费者一接收到就从 queue 中去除了,设为 false 则消费者在处理完消息后，调用 msg.Ack(false) 后消息才从 queue 中去除
- exclusive 排他 
- noLocal  设置为 true 则表示不能将同一个 Connection 中生产者发送的消息传送给这个 Connection 中的消费者
- noWait bool 
- args Table

# 二、redis的发布订阅