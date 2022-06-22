// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"

	"github.com/larksuite/oapi-sdk-go/service/attendance/v1"
)

func (dispatcher *EventDispatcher) OnUserFlowCreatedV1(handler func(ctx context.Context, event *attendance.UserFlowCreatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["attendance.user_flow.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "attendance.user_flow.created_v1")
	}
	dispatcher.eventType2EventHandler["attendance.user_flow.created_v1"] = attendance.NewUserFlowCreatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnUserTaskUpdatedV1(handler func(ctx context.Context, event *attendance.UserTaskUpdatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["attendance.user_task.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "attendance.user_task.updated_v1")
	}
	dispatcher.eventType2EventHandler["attendance.user_task.updated_v1"] = attendance.NewUserTaskUpdatedEventHandler(handler)
	return dispatcher
}
