package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"strings"
)

type resultEncoder struct{}

func (e *resultEncoder) HeaderResult() string {
	return "DATE,TIME,TYPE,HOSTNAME,WORKER,JOB,CHECK,DURATION,SDATE,STIME"
}

func (e *resultEncoder) EncodeResult(r *domain.Result, hostname string) string {
	a := []string{
		r.Time.Format("2006-01-02"),
		r.Time.Format("15:04:05"),
		"RESULT",
		hostname,
		fmt.Sprint(r.Job.Worker.Id),
		fmt.Sprint(r.Job.Id),
		r.Check.String(),
		r.Duration.String(),
		r.Job.Stime.Format("2006-01-02"),
		r.Job.Stime.Format("15:04:05"),
	}
	// data,time,host,worker,job,check,duration,sdate,stime
	return strings.Join(a, ",")
}

func NewResultEncoder() *resultEncoder {
	return &resultEncoder{}
}
