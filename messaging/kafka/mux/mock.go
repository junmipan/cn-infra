// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mux

import (
	"github.com/Shopify/sarama/mocks"
	"github.com/ligato/cn-infra/logging/logroot"
	"github.com/ligato/cn-infra/messaging/kafka/client"
	"testing"
)

func getMockConsumerFactory(t *testing.T) ConsumerFactory {
	return func(topics []string, name string) (*client.Consumer, error) {
		return client.GetConsumerMock(t), nil
	}
}

// GetMultiplexerMock returns mock of Multiplexer that can be used for testing purposes.
func GetMultiplexerMock(t *testing.T) (*Multiplexer, *mocks.AsyncProducer, *mocks.SyncProducer) {
	asyncP, aMock := client.GetAsyncProducerMock(t)
	syncP, sMock := client.GetSyncProducerMock(t)
	return NewMultiplexer(getMockConsumerFactory(t), syncP, asyncP, "name", logroot.Logger()), aMock, sMock
}
