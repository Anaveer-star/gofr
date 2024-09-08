// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -source=interfaces.go -destination=mock_interfaces.go -package=nats
//

// Package nats is a generated GoMock package.
package nats

import (
	context "context"
	reflect "reflect"

	nats "github.com/nats-io/nats.go"
	gomock "go.uber.org/mock/gomock"
	pubsub "gofr.dev/pkg/gofr/datasource/pubsub"
)

// MockConsumer is a mock of Consumer interface.
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer.
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance.
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockConsumer) Consume(ctx context.Context, handler func(*nats.Msg)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", ctx, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Consume indicates an expected call of Consume.
func (mr *MockConsumerMockRecorder) Consume(ctx, handler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockConsumer)(nil).Consume), ctx, handler)
}

// FetchMessage mocks base method.
func (m *MockConsumer) FetchMessage(ctx context.Context) (*nats.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMessage", ctx)
	ret0, _ := ret[0].(*nats.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMessage indicates an expected call of FetchMessage.
func (mr *MockConsumerMockRecorder) FetchMessage(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMessage", reflect.TypeOf((*MockConsumer)(nil).FetchMessage), ctx)
}

// GetInfo mocks base method.
func (m *MockConsumer) GetInfo() (*nats.ConsumerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(*nats.ConsumerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockConsumerMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockConsumer)(nil).GetInfo))
}

// MockPublisher is a mock of Publisher interface.
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher.
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance.
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockPublisher) Publish(ctx context.Context, subject string, data []byte) (*nats.PubAck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, subject, data)
	ret0, _ := ret[0].(*nats.PubAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockPublisherMockRecorder) Publish(ctx, subject, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPublisher)(nil).Publish), ctx, subject, data)
}

// PublishAsync mocks base method.
func (m *MockPublisher) PublishAsync(subject string, data []byte) (nats.PubAckFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishAsync", subject, data)
	ret0, _ := ret[0].(nats.PubAckFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishAsync indicates an expected call of PublishAsync.
func (mr *MockPublisherMockRecorder) PublishAsync(subject, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAsync", reflect.TypeOf((*MockPublisher)(nil).PublishAsync), subject, data)
}

// MockStreamManager is a mock of StreamManager interface.
type MockStreamManager struct {
	ctrl     *gomock.Controller
	recorder *MockStreamManagerMockRecorder
}

// MockStreamManagerMockRecorder is the mock recorder for MockStreamManager.
type MockStreamManagerMockRecorder struct {
	mock *MockStreamManager
}

// NewMockStreamManager creates a new mock instance.
func NewMockStreamManager(ctrl *gomock.Controller) *MockStreamManager {
	mock := &MockStreamManager{ctrl: ctrl}
	mock.recorder = &MockStreamManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamManager) EXPECT() *MockStreamManagerMockRecorder {
	return m.recorder
}

// AddStream mocks base method.
func (m *MockStreamManager) AddStream(config *nats.StreamConfig) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStream", config)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddStream indicates an expected call of AddStream.
func (mr *MockStreamManagerMockRecorder) AddStream(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStream", reflect.TypeOf((*MockStreamManager)(nil).AddStream), config)
}

// DeleteStream mocks base method.
func (m *MockStreamManager) DeleteStream(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStream", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockStreamManagerMockRecorder) DeleteStream(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockStreamManager)(nil).DeleteStream), name)
}

// StreamInfo mocks base method.
func (m *MockStreamManager) StreamInfo(name string) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamInfo", name)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamInfo indicates an expected call of StreamInfo.
func (mr *MockStreamManagerMockRecorder) StreamInfo(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamInfo", reflect.TypeOf((*MockStreamManager)(nil).StreamInfo), name)
}

// UpdateStream mocks base method.
func (m *MockStreamManager) UpdateStream(config *nats.StreamConfig) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStream", config)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStream indicates an expected call of UpdateStream.
func (mr *MockStreamManagerMockRecorder) UpdateStream(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStream", reflect.TypeOf((*MockStreamManager)(nil).UpdateStream), config)
}

// MockJetStreamContext is a mock of JetStreamContext interface.
type MockJetStreamContext struct {
	ctrl     *gomock.Controller
	recorder *MockJetStreamContextMockRecorder
}

// MockJetStreamContextMockRecorder is the mock recorder for MockJetStreamContext.
type MockJetStreamContextMockRecorder struct {
	mock *MockJetStreamContext
}

// NewMockJetStreamContext creates a new mock instance.
func NewMockJetStreamContext(ctrl *gomock.Controller) *MockJetStreamContext {
	mock := &MockJetStreamContext{ctrl: ctrl}
	mock.recorder = &MockJetStreamContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJetStreamContext) EXPECT() *MockJetStreamContextMockRecorder {
	return m.recorder
}

// AddStream mocks base method.
func (m *MockJetStreamContext) AddStream(config *nats.StreamConfig) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStream", config)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddStream indicates an expected call of AddStream.
func (mr *MockJetStreamContextMockRecorder) AddStream(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStream", reflect.TypeOf((*MockJetStreamContext)(nil).AddStream), config)
}

// Consume mocks base method.
func (m *MockJetStreamContext) Consume(ctx context.Context, handler func(*nats.Msg)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", ctx, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Consume indicates an expected call of Consume.
func (mr *MockJetStreamContextMockRecorder) Consume(ctx, handler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockJetStreamContext)(nil).Consume), ctx, handler)
}

// DeleteStream mocks base method.
func (m *MockJetStreamContext) DeleteStream(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStream", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockJetStreamContextMockRecorder) DeleteStream(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockJetStreamContext)(nil).DeleteStream), name)
}

// FetchMessage mocks base method.
func (m *MockJetStreamContext) FetchMessage(ctx context.Context) (*nats.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMessage", ctx)
	ret0, _ := ret[0].(*nats.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMessage indicates an expected call of FetchMessage.
func (mr *MockJetStreamContextMockRecorder) FetchMessage(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMessage", reflect.TypeOf((*MockJetStreamContext)(nil).FetchMessage), ctx)
}

// GetInfo mocks base method.
func (m *MockJetStreamContext) GetInfo() (*nats.ConsumerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(*nats.ConsumerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockJetStreamContextMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockJetStreamContext)(nil).GetInfo))
}

// Publish mocks base method.
func (m *MockJetStreamContext) Publish(ctx context.Context, subject string, data []byte) (*nats.PubAck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, subject, data)
	ret0, _ := ret[0].(*nats.PubAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockJetStreamContextMockRecorder) Publish(ctx, subject, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockJetStreamContext)(nil).Publish), ctx, subject, data)
}

// PublishAsync mocks base method.
func (m *MockJetStreamContext) PublishAsync(subject string, data []byte) (nats.PubAckFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishAsync", subject, data)
	ret0, _ := ret[0].(nats.PubAckFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishAsync indicates an expected call of PublishAsync.
func (mr *MockJetStreamContextMockRecorder) PublishAsync(subject, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAsync", reflect.TypeOf((*MockJetStreamContext)(nil).PublishAsync), subject, data)
}

// StreamInfo mocks base method.
func (m *MockJetStreamContext) StreamInfo(name string) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamInfo", name)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamInfo indicates an expected call of StreamInfo.
func (mr *MockJetStreamContextMockRecorder) StreamInfo(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamInfo", reflect.TypeOf((*MockJetStreamContext)(nil).StreamInfo), name)
}

// UpdateStream mocks base method.
func (m *MockJetStreamContext) UpdateStream(config *nats.StreamConfig) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStream", config)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStream indicates an expected call of UpdateStream.
func (mr *MockJetStreamContextMockRecorder) UpdateStream(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStream", reflect.TypeOf((*MockJetStreamContext)(nil).UpdateStream), config)
}

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConnection) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// JetStream mocks base method.
func (m *MockConnection) JetStream(opts ...nats.JSOpt) (JetStreamContext, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "JetStream", varargs...)
	ret0, _ := ret[0].(JetStreamContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JetStream indicates an expected call of JetStream.
func (mr *MockConnectionMockRecorder) JetStream(opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JetStream", reflect.TypeOf((*MockConnection)(nil).JetStream), opts...)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Publish mocks base method.
func (m *MockClient) Publish(ctx context.Context, stream string, message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, stream, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockClientMockRecorder) Publish(ctx, stream, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockClient)(nil).Publish), ctx, stream, message)
}

// Subscribe mocks base method.
func (m *MockClient) Subscribe(ctx context.Context, stream string) (*pubsub.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, stream)
	ret0, _ := ret[0].(*pubsub.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockClientMockRecorder) Subscribe(ctx, stream any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockClient)(nil).Subscribe), ctx, stream)
}
