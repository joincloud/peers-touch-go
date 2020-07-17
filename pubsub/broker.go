package pubsub

import (
	"context"
	"sync"

	"github.com/ipfs/go-ipfs/core/coreapi"
)

type Broker interface {
	Pub(ctx context.Context, event Event) error
	Sub(ctx context.Context, topic string) (Subscriber, error)
	Unsub(topic string) error
	Close() error
}

type Message struct {
	Header map[string]string
	Body   []byte
}

type Subscriber interface {
	Topic() string
	Unsubscribe() error
}

type Event interface {
	Topic() string
	Message() *Message
	Ack() error
	Error() error
}

type broker struct {
	coreAPI       coreapi.CoreAPI
	subscriptions map[string]Subscriber
	muMux         sync.RWMutex
}

func (b *broker) Pub(ctx context.Context, event Event) error {
	panic("implement me")
}

func (b *broker) Sub(ctx context.Context, topic string) (Subscriber, error) {
	b.muMux.Lock()
	defer b.muMux.Unlock()

	return
}

func (b *broker) Unsub(topic string) error {
	panic("implement me")
}

func (b *broker) Close() error {
	panic("implement me")
}

func NewBroker(options ...BrokerOption) Broker {
	bo := &BrokerOptions{}
	for _, option := range options {
		option(bo)
	}

	b := &broker{
		coreAPI: bo.coreAPI,
	}

	return b
}
