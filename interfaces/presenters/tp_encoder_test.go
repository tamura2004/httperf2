package presenters_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra/mock"
	"github.com/tamura2004/httperf2/interfaces/presenters"
	"time"
)

func ExampleTpsHeader() {
	e := presenters.TpEncoder{}
	fmt.Println(e.Header("TPS"))
	// Output:
	// DATE,TIME,TYPE,HOSTNAME,TPS

}

func ExampleTpmHeader() {
	e := presenters.TpEncoder{}
	fmt.Println(e.Header("TPM"))
	// Output:
	// DATE,TIME,TYPE,HOSTNAME,TPM

}

func ExampleTpsEncode() {
	mock.InitTime()
	mock.InitHost()

	c := domain.NewCounter()
	c.AddDuration(10 * time.Second)
	c.IncTp()
	c.IncMulti()

	encoder := presenters.TpEncoder{}
	str := encoder.Encode(c, "TPS")
	fmt.Println(str)
	// Output:
	// [2018-02-01,08:00:00,TPS,dummyhost,1]

}

func ExampleTpmEncode() {
	mock.InitTime()
	mock.InitHost()

	c := domain.NewCounter()
	c.AddDuration(10 * time.Second)
	c.IncTp()
	c.IncMulti()

	encoder := presenters.TpEncoder{}
	str := encoder.Encode(c, "TPM")
	fmt.Println(str)
	// Output:
	// [2018-02-01,08:00,TPM,dummyhost,1]

}
