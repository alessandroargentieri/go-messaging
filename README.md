# go-messaging
Comparison among famous message broker or chaching tools

## Redis

Here are showed three examples of Redis usage:
- as a cache
- as an event notifier
- as a stream writer

### Redis as a cache

In the example, we use Redis to save in memory data, to recover, delete or expire it.
All the components of a system can share info using this powerful tool. It can also be persisted with opportune conf.

### Redis as an event notifier

Redis can be also used to publish events to which many component can act as subscribers. As soon as they subscribe, they start to get the notifications from Redis. Redis PUSHES the events to all the interested components.
No queue, no buffer, no backpressure is implemented.

### Redis as a stream writer

In the third example, Redis acts as a stream writer. It publish information as stream to which many listener can attach.
The listener which reads the information, consumes it from the stream, so all other listeners are prevented to get the same information. This system is good to divide tasks among more components.

## RabbitMQ

RabbitMQ is a famous message broker which implements both publish/subscribe and queue mechanisms.
In the example, a publisher creates an *exchange* which is a topic on which to publish some notifications. Other consumer, just create a queue, attach it to the exchange and read from the queue.
This way the information is replicated to all the queues and each consumer can access to a replica of the same information.
