package commands

import (
	"context"
	"delivery/internal/core/domain/model/courier"
	"delivery/internal/core/domain/model/kernel"
	"delivery/internal/core/ports"
	"delivery/internal/pkg/errs"
)

type CreateCourierCmd struct {
	name  string
	speed int

	isSet bool
}

func NewCreateCourierCmd(name string, speed int) (CreateCourierCmd, error) {
	if name == "" {
		return CreateCourierCmd{}, errs.NewValueIsRequiredError("name")
	}
	if speed <= 0 {
		return CreateCourierCmd{}, errs.NewValueIsRequiredError("speed")
	}

	return CreateCourierCmd{
		name:  name,
		speed: speed,

		isSet: true,
	}, nil
}

func (cmd CreateCourierCmd) Speed() int {
	return cmd.speed
}

func (cmd CreateCourierCmd) Name() string {
	return cmd.name
}

func (cmd CreateCourierCmd) IsEmpty() bool {
	return !cmd.isSet
}

type CreateCourierCommandHandler interface {
	Handle(context.Context, CreateCourierCmd) error
}

var _ CreateCourierCommandHandler = &createCourierCommandHandler{}

type createCourierCommandHandler struct {
	unitOfWork       ports.UnitOfWork
	courseRepository ports.CourierRepository
}

func NewCreateCourierCommandHandler(uow ports.UnitOfWork, repo ports.CourierRepository) (CreateCourierCommandHandler, error) {
	if uow == nil {
		return nil, errs.NewValueIsRequiredError("uow")
	}

	if repo == nil {
		return nil, errs.NewValueIsRequiredError("repo")
	}

	return &createCourierCommandHandler{
		unitOfWork:       uow,
		courseRepository: repo,
	}, nil
}

func (ch *createCourierCommandHandler) Handle(ctx context.Context, cmd CreateCourierCmd) error {
	if cmd.IsEmpty() {
		return errs.NewValueIsRequiredError("cmd")
	}

	location := kernel.CreateRandomLocation()
	courierAggregate, err := courier.NewCourier(cmd.Name(), cmd.Speed(), location)
	if err != nil {
		return err
	}
	err = courierAggregate.AddStoragePlace("bag", 40)
	if err != nil {
		return err
	}

	err = ch.courseRepository.Add(ctx, courierAggregate)
	if err != nil {
		return err
	}
	return nil
}
