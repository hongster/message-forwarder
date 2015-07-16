# message-forwarder
Get task from message queue and forward to an URL.

## Setup

Copy [message-forwarder.cfg](https://github.com/hongster/message-forwarder/blob/master/message-forwarder.cfg) to */etc/*, and edit accordingly.

## Linux Setup

Optional instruction for running message-forwarder as [systemd](https://en.wikipedia.org/wiki/Systemd) service.

Generate executable binary

```bash
cd $GOPATH
go install github.com/hongster/message-forwarder
sudo cp bin/message-forwarder /usr/bin/
```

Copy [doc/message-forwarder.service](https://github.com/hongster/message-forwarder/blob/master/doc/message-forwarder.service) to */etc/systemd/system/*. Then run the following commands to enable and start service.

```bash
sudo systemctl enable message-forwarder
sudo systemctl start message-forwarder
```

## Dependencies

Automatically resolved using `go get github.com/hongster/message-forwarder`.

* [streadway/amqp](https://github.com/streadway/amqp)
* [robfig/config](https://github.com/robfig/config)
