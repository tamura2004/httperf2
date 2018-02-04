package presenters_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra/mock"
	"github.com/tamura2004/httperf2/interfaces/presenters"
)

func ExampleResultEncode() {
	mock.InitTime()
	mock.InitHost()

	encoder := presenters.ResultEncoder{}
	job := domain.NewJob(1, 2)
	result := job.TimeStart()
	str := encoder.Encode(result)
	fmt.Println(str)
	// Output:
	// 2018-02-01,08:00:00,RESULT,dummy,1,2,START,10s,2018-02-01,08:00:00

}
