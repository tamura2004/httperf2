package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"sort"
)

type TpEncoder struct{}

func (e *TpEncoder) Header(name string) string {
	return fmt.Sprintf("DATE,TIME,TYPE,HOSTNAME,%s", name)
}

func (e *TpEncoder) Encode(c *domain.Counter, name string) []string {
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
		row := fmt.Sprintf("%s,%s,%s,%d", k, name, host.Name(), counter[k])
		rows = append(rows, row)
	}

	return rows
}
