package cmd

import (
	"delivery/internal/adapters/in/jobs"
	kafkain "delivery/internal/adapters/in/kafka"
	"delivery/internal/adapters/out/grpc/geo"
	"delivery/internal/adapters/out/postgres/courierrepo"
	"delivery/internal/adapters/out/postgres/orderrepo"
	"delivery/internal/adapters/out/postgres/shared"
	"delivery/internal/core/application/usecases/commands"
	"delivery/internal/core/application/usecases/queries"
	"delivery/internal/core/domain/services"
	"delivery/internal/core/ports"
	"github.com/robfig/cron/v3"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type CompositionRoot struct {
	container *dig.Container
	configs   Config

	closers []Closer
}

func NewCompositionRoot(config Config, db *gorm.DB) CompositionRoot {
	container := dig.New()

	// Provide common singletons
	if err := container.Provide(func() *gorm.DB { return db }); err != nil {
		panic(err)
	}
	if err := container.Provide(func() Config { return config }); err != nil {
		panic(err)
	}
	if err := container.Provide(shared.NewTxManagerFactory); err != nil {
		panic(err)
	}
	if err := container.Provide(services.NewOrderDispatcher); err != nil {
		panic(err)
	}
	if err := container.Provide(func(config Config) (ports.GeoLocationGateway, error) {
		return geo.NewGeoLocationService(config.GeoServiceGrpcHost)
	}); err != nil {
		panic(err)
	}

	return CompositionRoot{
		container: container,
		configs:   config,
	}
}

func (cr *CompositionRoot) NewAssignOrderJob() cron.Job {
	job, err := jobs.NewAssignOrderJob(cr.buildAssignOrderHandler())
	if err != nil {
		panic(err)
	}
	return job
}

func (cr *CompositionRoot) NewMoveCouriersJob() cron.Job {
	job, err := jobs.NewMoveCouriersJob(cr.buildMoveCouriersHandler())
	if err != nil {
		panic(err)
	}
	return job
}

func (cr *CompositionRoot) buildAssignOrderHandler() commands.AssignOrderCommandHandler {
	var handler commands.AssignOrderCommandHandler
	err := cr.container.Invoke(func(
		factory shared.TxManagerFactory,
		dispatcher services.OrderDispatcher,
	) {
		txManager := factory.New()
		uow := ports.UnitOfWork(txManager)

		orderRepo, err := orderrepo.NewOrderRepository(txManager)
		if err != nil {
			panic(err)
		}
		courierRepo, err := courierrepo.NewCourierRepository(txManager)
		if err != nil {
			panic(err)
		}
		handler, err = commands.NewAssignOrderCommandHandler(uow, dispatcher, orderRepo, courierRepo)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
	return handler
}

func (cr *CompositionRoot) buildCreateCourierHandler() commands.CreateCourierCommandHandler {
	var handler commands.CreateCourierCommandHandler
	err := cr.container.Invoke(func(
		factory shared.TxManagerFactory,
	) {
		txManager := factory.New()
		uow := ports.UnitOfWork(txManager)

		repo, err := courierrepo.NewCourierRepository(txManager)
		if err != nil {
			panic(err)
		}
		handler, err = commands.NewCreateCourierCommandHandler(uow, repo)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
	return handler
}

func (cr *CompositionRoot) buildCreateOrderHandler() commands.CreateOrderCommandHandler {
	var handler commands.CreateOrderCommandHandler
	err := cr.container.Invoke(func(
		factory shared.TxManagerFactory,
		geoClient ports.GeoLocationGateway,
	) {
		txManager := factory.New()
		uow := ports.UnitOfWork(txManager)

		orderRepo, err := orderrepo.NewOrderRepository(txManager)
		if err != nil {
			panic(err)
		}
		handler, err = commands.NewCreateOrderCommandHandler(uow, orderRepo, geoClient)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
	return handler
}

func (cr *CompositionRoot) buildMoveCouriersHandler() commands.MoveCouriersCommandHandler {
	var handler commands.MoveCouriersCommandHandler
	err := cr.container.Invoke(func(
		factory shared.TxManagerFactory,
	) {
		txManager := factory.New()
		uow := ports.UnitOfWork(txManager)

		orderRepo, err := orderrepo.NewOrderRepository(txManager)
		if err != nil {
			panic(err)
		}
		courierRepo, err := courierrepo.NewCourierRepository(txManager)
		if err != nil {
			panic(err)
		}
		handler, err = commands.NewMoveCouriersCommandHandler(uow, orderRepo, courierRepo)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
	return handler
}

func (cr *CompositionRoot) NewCreateCourierCommandHandler() commands.CreateCourierCommandHandler {
	return cr.buildCreateCourierHandler()
}

func (cr *CompositionRoot) NewCreateOrderCommandHandler() commands.CreateOrderCommandHandler {
	return cr.buildCreateOrderHandler()
}

func (cr *CompositionRoot) NewAssignOrderCommandHandler() commands.AssignOrderCommandHandler {
	return cr.buildAssignOrderHandler()
}

func (cr *CompositionRoot) NewMoveCouriersCommandHandler() commands.MoveCouriersCommandHandler {
	return cr.buildMoveCouriersHandler()
}

func (cr *CompositionRoot) NewGetAllCouriersQueryHandler() queries.GetAllCouriersQueryHandler {
	handler, err := queries.NewGetAllCouriersQueryHandler(cr.getDB())
	if err != nil {
		panic(err)
	}
	return handler
}

func (cr *CompositionRoot) NewGetNotCompletedOrdersQueryHandler() queries.GetNotCompletedOrdersQueryHandler {
	handler, err := queries.NewGetNotCompletedOrdersQueryHandler(cr.getDB())
	if err != nil {
		panic(err)
	}
	return handler
}

func (cr *CompositionRoot) NewBasketConfirmedConsumer() kafkain.BasketConfirmedConsumer {
	consumer, err := kafkain.NewBasketConfirmedConsumer(
		[]string{cr.configs.KafkaHost},
		cr.configs.KafkaConsumerGroup,
		cr.configs.KafkaBasketConfirmedTopic,
		cr.NewCreateOrderCommandHandler(),
	)
	if err != nil {
		panic(err)
	}
	return consumer
}

func (cr *CompositionRoot) getDB() *gorm.DB {
	var db *gorm.DB
	err := cr.container.Invoke(func(d *gorm.DB) {
		db = d
	})
	if err != nil {
		panic(err)
	}
	return db
}

func (cr *CompositionRoot) Container() *dig.Container {
	return cr.container
}
