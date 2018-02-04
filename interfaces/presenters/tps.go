package presenters

// type Counter struct {
// 	TPS   map[string]int
// 	TPM   map[string]int
// 	Count int
// 	Multi int
// 	Total time.Duration
// }

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/usecase"
	"os"
	"sort"
)

var Hostname = os.Hostname

type tpEncoder struct{}

func (e *tpEncoder) Header(name string) string {
	return fmt.Sprintf("DATE,TIME,TYPE,HOSTNAME,%s", name)
}

func (e *tpEncoder) Encode(c *domain.Counter, name string) []string {
	counter := c.TPM
	if name == "TPS" {
		counter = c.TPS
	}

	keys := []string{}
	for k := range counter {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	rows := []string{}
	for _, k := range keys {
		row := fmt.Sprintf("%s,%s,%s,%d", k, name, Hostname(), counter[k])
		rows = append(rows, row)
	}

	return rows
}

func InitTpEncoder() {
	usecase.InitTpEncoder(&tpEncoder{})
}
