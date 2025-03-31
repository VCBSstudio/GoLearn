package job

import (
	// "context"
	// "time"

	"github.com/robfig/cron/v3"
)

type JobManager struct {
	cron *cron.Cron
}

func NewJobManager() *JobManager {
	return &JobManager{
		cron: cron.New(cron.WithSeconds()), // 支持秒级精度
	}
}

func (m *JobManager) Start() {
	m.cron.Start()
}

func (m *JobManager) Stop() {
	ctx := m.cron.Stop()
	<-ctx.Done()
}

func (m *JobManager) AddJob(spec string, cmd func()) (cron.EntryID, error) {
	return m.cron.AddFunc(spec, cmd)
}
