package usecase_test

import (
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra/mock"
	"github.com/tamura2004/httperf2/interfaces/presenters"
	"github.com/tamura2004/httperf2/usecase"
)

func ExampleReportResult() {
	mock.InitTime()
	mock.InitFileFactory()
	mock.InitHost()

	presenters.InitResultEncoder()

	reporter := usecase.NewReporter()
	job := domain.NewJob(1, 2)
	result := job.TimeStart()

	reporter.ReportResult(result)
	// Output:
	// DATE,TIME,TYPE,HOSTNAME,WORKER,JOB,CHECK,DURATION,SDATE,STIME
	// 2018-02-01,08:00:00,RESULT,dummy,1,2,START,10s,2018-02-01,08:00:00
}
