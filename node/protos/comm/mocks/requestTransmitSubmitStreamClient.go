// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	protos "github.com/hyperledger/fabric-x-orderer/node/protos/comm"

	"google.golang.org/grpc/metadata"
)

type FakeRequestTransmit_SubmitStreamClient struct {
	CloseSendStub        func() error
	closeSendMutex       sync.RWMutex
	closeSendArgsForCall []struct {
	}
	closeSendReturns struct {
		result1 error
	}
	closeSendReturnsOnCall map[int]struct {
		result1 error
	}
	ContextStub        func() context.Context
	contextMutex       sync.RWMutex
	contextArgsForCall []struct {
	}
	contextReturns struct {
		result1 context.Context
	}
	contextReturnsOnCall map[int]struct {
		result1 context.Context
	}
	HeaderStub        func() (metadata.MD, error)
	headerMutex       sync.RWMutex
	headerArgsForCall []struct {
	}
	headerReturns struct {
		result1 metadata.MD
		result2 error
	}
	headerReturnsOnCall map[int]struct {
		result1 metadata.MD
		result2 error
	}
	RecvStub        func() (*protos.SubmitResponse, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct {
	}
	recvReturns struct {
		result1 *protos.SubmitResponse
		result2 error
	}
	recvReturnsOnCall map[int]struct {
		result1 *protos.SubmitResponse
		result2 error
	}
	RecvMsgStub        func(any) error
	recvMsgMutex       sync.RWMutex
	recvMsgArgsForCall []struct {
		arg1 any
	}
	recvMsgReturns struct {
		result1 error
	}
	recvMsgReturnsOnCall map[int]struct {
		result1 error
	}
	SendStub        func(*protos.Request) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 *protos.Request
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	SendMsgStub        func(any) error
	sendMsgMutex       sync.RWMutex
	sendMsgArgsForCall []struct {
		arg1 any
	}
	sendMsgReturns struct {
		result1 error
	}
	sendMsgReturnsOnCall map[int]struct {
		result1 error
	}
	TrailerStub        func() metadata.MD
	trailerMutex       sync.RWMutex
	trailerArgsForCall []struct {
	}
	trailerReturns struct {
		result1 metadata.MD
	}
	trailerReturnsOnCall map[int]struct {
		result1 metadata.MD
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRequestTransmit_SubmitStreamClient) CloseSend() error {
	fake.closeSendMutex.Lock()
	ret, specificReturn := fake.closeSendReturnsOnCall[len(fake.closeSendArgsForCall)]
	fake.closeSendArgsForCall = append(fake.closeSendArgsForCall, struct {
	}{})
	fake.recordInvocation("CloseSend", []interface{}{})
	fake.closeSendMutex.Unlock()
	if fake.CloseSendStub != nil {
		return fake.CloseSendStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeSendReturns
	return fakeReturns.result1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) CloseSendCallCount() int {
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	return len(fake.closeSendArgsForCall)
}

func (fake *FakeRequestTransmit_SubmitStreamClient) CloseSendCalls(stub func() error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = stub
}

func (fake *FakeRequestTransmit_SubmitStreamClient) CloseSendReturns(result1 error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = nil
	fake.closeSendReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) CloseSendReturnsOnCall(i int, result1 error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = nil
	if fake.closeSendReturnsOnCall == nil {
		fake.closeSendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeSendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) Context() context.Context {
	fake.contextMutex.Lock()
	ret, specificReturn := fake.contextReturnsOnCall[len(fake.contextArgsForCall)]
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct {
	}{})
	fake.recordInvocation("Context", []interface{}{})
	fake.contextMutex.Unlock()
	if fake.ContextStub != nil {
		return fake.ContextStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.contextReturns
	return fakeReturns.result1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *FakeRequestTransmit_SubmitStreamClient) ContextCalls(stub func() context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = stub
}

func (fake *FakeRequestTransmit_SubmitStreamClient) ContextReturns(result1 context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 context.Context
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) ContextReturnsOnCall(i int, result1 context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = nil
	if fake.contextReturnsOnCall == nil {
		fake.contextReturnsOnCall = make(map[int]struct {
			result1 context.Context
		})
	}
	fake.contextReturnsOnCall[i] = struct {
		result1 context.Context
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) Header() (metadata.MD, error) {
	fake.headerMutex.Lock()
	ret, specificReturn := fake.headerReturnsOnCall[len(fake.headerArgsForCall)]
	fake.headerArgsForCall = append(fake.headerArgsForCall, struct {
	}{})
	fake.recordInvocation("Header", []interface{}{})
	fake.headerMutex.Unlock()
	if fake.HeaderStub != nil {
		return fake.HeaderStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.headerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRequestTransmit_SubmitStreamClient) HeaderCallCount() int {
	fake.headerMutex.RLock()
	defer fake.headerMutex.RUnlock()
	return len(fake.headerArgsForCall)
}

func (fake *FakeRequestTransmit_SubmitStreamClient) HeaderCalls(stub func() (metadata.MD, error)) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = stub
}

func (fake *FakeRequestTransmit_SubmitStreamClient) HeaderReturns(result1 metadata.MD, result2 error) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = nil
	fake.headerReturns = struct {
		result1 metadata.MD
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) HeaderReturnsOnCall(i int, result1 metadata.MD, result2 error) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = nil
	if fake.headerReturnsOnCall == nil {
		fake.headerReturnsOnCall = make(map[int]struct {
			result1 metadata.MD
			result2 error
		})
	}
	fake.headerReturnsOnCall[i] = struct {
		result1 metadata.MD
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) Recv() (*protos.SubmitResponse, error) {
	fake.recvMutex.Lock()
	ret, specificReturn := fake.recvReturnsOnCall[len(fake.recvArgsForCall)]
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct {
	}{})
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if fake.RecvStub != nil {
		return fake.RecvStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.recvReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvCalls(stub func() (*protos.SubmitResponse, error)) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = stub
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvReturns(result1 *protos.SubmitResponse, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *protos.SubmitResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvReturnsOnCall(i int, result1 *protos.SubmitResponse, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	if fake.recvReturnsOnCall == nil {
		fake.recvReturnsOnCall = make(map[int]struct {
			result1 *protos.SubmitResponse
			result2 error
		})
	}
	fake.recvReturnsOnCall[i] = struct {
		result1 *protos.SubmitResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvMsg(arg1 any) error {
	fake.recvMsgMutex.Lock()
	ret, specificReturn := fake.recvMsgReturnsOnCall[len(fake.recvMsgArgsForCall)]
	fake.recvMsgArgsForCall = append(fake.recvMsgArgsForCall, struct {
		arg1 any
	}{arg1})
	fake.recordInvocation("RecvMsg", []interface{}{arg1})
	fake.recvMsgMutex.Unlock()
	if fake.RecvMsgStub != nil {
		return fake.RecvMsgStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.recvMsgReturns
	return fakeReturns.result1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvMsgCallCount() int {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return len(fake.recvMsgArgsForCall)
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvMsgCalls(stub func(any) error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = stub
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvMsgArgsForCall(i int) any {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	argsForCall := fake.recvMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvMsgReturns(result1 error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = nil
	fake.recvMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) RecvMsgReturnsOnCall(i int, result1 error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = nil
	if fake.recvMsgReturnsOnCall == nil {
		fake.recvMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.recvMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) Send(arg1 *protos.Request) error {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 *protos.Request
	}{arg1})
	fake.recordInvocation("Send", []interface{}{arg1})
	fake.sendMutex.Unlock()
	if fake.SendStub != nil {
		return fake.SendStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendReturns
	return fakeReturns.result1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendCalls(stub func(*protos.Request) error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = stub
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendArgsForCall(i int) *protos.Request {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	argsForCall := fake.sendArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendReturns(result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendReturnsOnCall(i int, result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendMsg(arg1 any) error {
	fake.sendMsgMutex.Lock()
	ret, specificReturn := fake.sendMsgReturnsOnCall[len(fake.sendMsgArgsForCall)]
	fake.sendMsgArgsForCall = append(fake.sendMsgArgsForCall, struct {
		arg1 any
	}{arg1})
	fake.recordInvocation("SendMsg", []interface{}{arg1})
	fake.sendMsgMutex.Unlock()
	if fake.SendMsgStub != nil {
		return fake.SendMsgStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendMsgReturns
	return fakeReturns.result1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendMsgCallCount() int {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	return len(fake.sendMsgArgsForCall)
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendMsgCalls(stub func(any) error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = stub
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendMsgArgsForCall(i int) any {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	argsForCall := fake.sendMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendMsgReturns(result1 error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = nil
	fake.sendMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) SendMsgReturnsOnCall(i int, result1 error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = nil
	if fake.sendMsgReturnsOnCall == nil {
		fake.sendMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) Trailer() metadata.MD {
	fake.trailerMutex.Lock()
	ret, specificReturn := fake.trailerReturnsOnCall[len(fake.trailerArgsForCall)]
	fake.trailerArgsForCall = append(fake.trailerArgsForCall, struct {
	}{})
	fake.recordInvocation("Trailer", []interface{}{})
	fake.trailerMutex.Unlock()
	if fake.TrailerStub != nil {
		return fake.TrailerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.trailerReturns
	return fakeReturns.result1
}

func (fake *FakeRequestTransmit_SubmitStreamClient) TrailerCallCount() int {
	fake.trailerMutex.RLock()
	defer fake.trailerMutex.RUnlock()
	return len(fake.trailerArgsForCall)
}

func (fake *FakeRequestTransmit_SubmitStreamClient) TrailerCalls(stub func() metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = stub
}

func (fake *FakeRequestTransmit_SubmitStreamClient) TrailerReturns(result1 metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = nil
	fake.trailerReturns = struct {
		result1 metadata.MD
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) TrailerReturnsOnCall(i int, result1 metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = nil
	if fake.trailerReturnsOnCall == nil {
		fake.trailerReturnsOnCall = make(map[int]struct {
			result1 metadata.MD
		})
	}
	fake.trailerReturnsOnCall[i] = struct {
		result1 metadata.MD
	}{result1}
}

func (fake *FakeRequestTransmit_SubmitStreamClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	fake.headerMutex.RLock()
	defer fake.headerMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	fake.trailerMutex.RLock()
	defer fake.trailerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRequestTransmit_SubmitStreamClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ protos.RequestTransmit_SubmitStreamClient = new(FakeRequestTransmit_SubmitStreamClient)
