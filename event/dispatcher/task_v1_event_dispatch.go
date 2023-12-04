// Package dispatcher code generated by oapi sdk gen
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

package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/service/task/v1"
)
// 任务信息变更（租户维度）
//
// - APP 订阅此事件后可接收到该 APP 所在租户的所有来源接口创建的任务的变更事件。事件体为发生变更任务的相关用户的 open_id，可用此 open_id ，通过 获取任务列表接口获取与该用户相关的所有任务。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/events/update_tenant
func ( dispatcher * EventDispatcher ) OnP2TaskUpdateTenantV1(handler func(ctx context.Context, event *larktask.P2TaskUpdateTenantV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["task.task.update_tenant_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "task.task.update_tenant_v1")
	}
	dispatcher.eventType2EventHandler["task.task.update_tenant_v1"] = larktask.NewP2TaskUpdateTenantV1Handler(handler)
	return dispatcher
}
// 任务信息变更
//
// - 当 APP 订阅此事件后可以接收到由该 APP 创建的任务发生的变更，包括任务标题、描述、截止时间、协作者、关注者、提醒时间、状态（完成或取消完成）。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/events/updated
func ( dispatcher * EventDispatcher ) OnP2TaskUpdatedV1(handler func(ctx context.Context, event *larktask.P2TaskUpdatedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["task.task.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "task.task.updated_v1")
	}
	dispatcher.eventType2EventHandler["task.task.updated_v1"] = larktask.NewP2TaskUpdatedV1Handler(handler)
	return dispatcher
}
// 任务评论信息变更
//
// - 当 APP 创建的任务评论信息发生变更时触发此事件，包括任务评论的创建、回复、更新、删除。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/events/updated
func ( dispatcher * EventDispatcher ) OnP2TaskCommentUpdatedV1(handler func(ctx context.Context, event *larktask.P2TaskCommentUpdatedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["task.task.comment.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "task.task.comment.updated_v1")
	}
	dispatcher.eventType2EventHandler["task.task.comment.updated_v1"] = larktask.NewP2TaskCommentUpdatedV1Handler(handler)
	return dispatcher
}