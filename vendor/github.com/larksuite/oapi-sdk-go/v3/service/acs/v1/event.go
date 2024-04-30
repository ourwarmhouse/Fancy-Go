// Package acs code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkacs

import (
	"context"
)

// 消息处理器定义
type P2AccessRecordCreatedV1Handler struct {
	handler func(context.Context, *P2AccessRecordCreatedV1) error
}

func NewP2AccessRecordCreatedV1Handler(handler func(context.Context, *P2AccessRecordCreatedV1) error) *P2AccessRecordCreatedV1Handler {
	h := &P2AccessRecordCreatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2AccessRecordCreatedV1Handler) Event() interface{} {
	return &P2AccessRecordCreatedV1{}
}

// 回调开发者注册的handle
func (h *P2AccessRecordCreatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2AccessRecordCreatedV1))
}

// 消息处理器定义
type P2UserUpdatedV1Handler struct {
	handler func(context.Context, *P2UserUpdatedV1) error
}

func NewP2UserUpdatedV1Handler(handler func(context.Context, *P2UserUpdatedV1) error) *P2UserUpdatedV1Handler {
	h := &P2UserUpdatedV1Handler{handler: handler}
	return h
}

// 返回事件的消息体的实例，用于反序列化用
func (h *P2UserUpdatedV1Handler) Event() interface{} {
	return &P2UserUpdatedV1{}
}

// 回调开发者注册的handle
func (h *P2UserUpdatedV1Handler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*P2UserUpdatedV1))
}
