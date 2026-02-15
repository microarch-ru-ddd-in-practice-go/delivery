package outbox

import (
	"delivery/internal/pkg/ddd"
	"delivery/internal/pkg/errs"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type EventRegistry interface {
	RegisterDomainEvent(eventType reflect.Type) error
	DecodeDomainEvent(event *Message) (ddd.DomainEvent, error)
}

var _ EventRegistry = &eventRegistry{}

type eventRegistry struct {
	EventRegistry map[string]reflect.Type
}

func NewEventRegistry() (EventRegistry, error) {
	return &eventRegistry{
		EventRegistry: make(map[string]reflect.Type),
	}, nil
}

func (r *eventRegistry) RegisterDomainEvent(eventType reflect.Type) error {
	if eventType == nil {
		return errs.NewValueIsRequired("eventType")
	}
	eventName := eventType.Name()
	r.EventRegistry[eventName] = eventType
	return nil
}

func EncodeDomainEvent(domainEvent ddd.DomainEvent) (Message, error) {
	payload, err := json.Marshal(domainEvent)
	if err != nil {
		return Message{}, fmt.Errorf("failed to marshal event: %w", err)
	}

	return Message{
		ID:             domainEvent.GetID(),
		Name:           domainEvent.GetName(),
		Payload:        payload,
		OccurredAtUtc:  time.Now().UTC(),
		ProcessedAtUtc: nil,
	}, nil
}

func (r *eventRegistry) DecodeDomainEvent(outboxMessage *Message) (ddd.DomainEvent, error) {
	t, ok := r.EventRegistry[outboxMessage.Name]
	if !ok {
		return nil, fmt.Errorf("unknown outboxMessage type: %s", outboxMessage.Name)
	}

	// Создаём новый указатель на нужный тип
	eventPtr := reflect.New(t).Interface()

	if err := json.Unmarshal(outboxMessage.Payload, eventPtr); err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	// Приводим к DomainEvent
	domainEvent, ok := eventPtr.(ddd.DomainEvent)
	if !ok {
		return nil, fmt.Errorf("decoded outboxMessage does not implement DomainEvent")
	}

	return domainEvent, nil
}
