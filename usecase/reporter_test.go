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

	job := domain.NewJob()
	result := job.TimeStart()

	reporter.ReportResult(result)
	// Output:
	// -
}
