package job

import (
	"context"
	"urlShorter/infra/repository"
	"urlShorter/internal/cron"
)

type Executer struct {
	query        repository.URLRepository
	cronExecuter cron.ExecutionCron
}

func NewExecuter(query repository.URLRepository, cronExecuter cron.ExecutionCron) *Executer {
	return &Executer{
		query:        query,
		cronExecuter: cronExecuter,
	}
}

func (e *Executer) Execute(ctx context.Context) {
	_, err := e.cronExecuter.AddFunc("@every 10s", func() {
		err := e.query.DeleteDeprecatedURL(ctx)
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}

	e.cronExecuter.Start()
}
