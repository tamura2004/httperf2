package infra_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/infra"
)

func ExampleNewConfig() {
	config := infra.NewConfig()
	fmt.Printf("%#v", config)
	// Output:
	// domain.Config{Client:domain.Client{InsecureSkipVerify:true, MaxIdleConnsPerHost:1024, Proxy:"", Bps:0}, Scinario:domain.Scinario{Worker:3, Count:3, RampUp:domain.Duration{Duration:100000000}, Interval:domain.Duration{Duration:3000000000}}, Target:domain.Target{Url:"https://ogisui.azurewebsites.net"}}

}
