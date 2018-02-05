package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"strings"
)

type ResultEncoder struct{}

func (e *ResultEncoder) Header() string {
	return "DATE,TIME,TYPE,HOSTNAME,WORKER,JOB,CHECK,DURATION,SDATE,STIME"
}

func (e *ResultEncoder) Encode(r *domain.Result) string {
	a := []string{
		r.Time.Format("2006-01-02"),
		r.Time.Format("15:04:05"),
		"RESULT",
		host.Name(),
		fmt.Sprint(r.WorkerId),
		fmt.Sprint(r.JobId),
		r.Check.String(),
		r.Duration.String(),
		r.Stime.Format("2006-01-02"),
		r.Stime.Format("15:04:05"),
	}
	// data,time,host,worker,job,check,duration,sdate,stime
	return strings.Join(a, ",")
}
