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
	"sort"
)

type tpEncoder struct{}

func (e *tpEncoder) HeaderTp(name string) string {
	return fmt.Sprintf("DATE,TIME,TYPE,HOSTNAME,%s", name)
}

func (e *tpEncoder) EncodeTp(c *domain.Counter, name, hostname string) []string {
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
		row := fmt.Sprintf("%s,%s,%s,%d", k, name, hostname, counter[k])
		rows = append(rows, row)
	}

	return rows
}

func NewTpEncoder() *tpEncoder {
	return &tpEncoder{}
}
