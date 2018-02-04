package infra_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/infra"
)

func ExampleNewConfigClient() {
	config := infra.NewConfig()
	fmt.Printf("%#v", config.Client)
	// Output:
	// domain.Client{InsecureSkipVerify:true, MaxIdleConnsPerHost:1024, Proxy:"", Bps:0}
}

func ExampleNewConfigScinario() {
	config := infra.NewConfig()
	fmt.Printf("%#v", config.Scinario)
	// Output:
	// domain.Scinario{Worker:3, Count:3, RampUp:domain.Duration{Duration:100}, Interval:domain.Duration{Duration:3000000}}
}

func ExampleNewConfigTarget() {
	config := infra.NewConfig()
	fmt.Printf("%#v", config.Target)
	// Output:
	//domain.Target{Url:"https://ogisui.azurewebsites.net"}
}
