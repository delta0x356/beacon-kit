// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package broker

import (
	"context"
	"time"
)

// Client defines a client that is registered to the broker and receives msgs
// from it.
type Client[T any] chan T

// Broker broadcasts msgs to registered clients.
type Broker[T any] struct {
	// name of the message broker.
	name string
	// clients is a map of registered clients.
	clients map[Client[T]]struct{}
	// msgs is the channel for publishing new messages.
	msgs chan T
	// timeout is the timeout for sending a msg to a client.
	timeout time.Duration
}

// New creates a new b.
func New[T any](name string) *Broker[T] {
	return &Broker[T]{
		clients: make(map[Client[T]]struct{}),
		msgs:    make(chan T, defaultBufferSize),
		timeout: defaultTimeout,
		name:    name,
	}
}

// Name returns the name of the msg broker.
func (b *Broker[T]) Name() string {
	return b.name
}

// Start starts the broker loop.
func (b *Broker[T]) Start(ctx context.Context) error {
	go b.start(ctx)
	return nil
}

// start starts the broker loop.
func (b *Broker[T]) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// close all leftover clients and break the broker loop
			for client := range b.clients {
				b.Unsubscribe(client)
			}
			return
		case msg := <-b.msgs:
			// broadcast published msg to all clients
			for client := range b.clients {
				// send msg to client (or discard msg after timeout)
				select {
				case client <- msg:
				case <-time.After(b.timeout):
				}
			}
		}
	}
}

// Publish publishes a msg to the b.
// Returns ErrTimeout on timeout.
func (b *Broker[T]) Publish(msg T) error {
	select {
	case b.msgs <- msg:
		return nil
	case <-time.After(b.timeout):
		return ErrTimeout
	}
}

// Subscribe registers a new client to the broker and returns it to the caller.
// Returns ErrTimeout on timeout.
func (b *Broker[T]) Subscribe() (Client[T], error) {
	client := make(Client[T])
	b.clients[client] = struct{}{}
	return client, nil
}

// Unsubscribe removes a client from the b.
// Returns ErrTimeout on timeout.
func (b *Broker[T]) Unsubscribe(client Client[T]) {
	// Remove the client from the broker
	delete(b.clients, client)
	// close the client channel
	close(client)
}