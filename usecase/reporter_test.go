package usecase_test

import (
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra"
	"github.com/tamura2004/httperf2/infra/mock"
	"github.com/tamura2004/httperf2/interfaces/presenters"
	"github.com/tamura2004/httperf2/usecase"
)

func ExampleReportResult() {
	reporter := usecase.NewReporter(
		mock.NewFileFactory(),
		presenters.NewResultEncoder(),
		presenters.NewTpEncoder(),
		infra.GetHostname(),
	)

	job := domain.NewJob(1)
	result := job.TimeStart()

	reporter.ReportResult(result)
	// Output:
	// DATE,TIME,TYPE,HOSTNAME,WORKER,JOB,CHECK,DURATION,SDATE,STIME
	// 2018-02-03,22:17:49,RESULT,scout,0,1,START,0s,2018-02-03,22:17:49
}
