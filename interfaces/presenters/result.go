package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"os"
	"strings"
)

var Hostname = os.Hostname

type resultEncoder struct{}

func (e *resultEncoder) Header() string {
	return "DATE,TIME,TYPE,HOSTNAME,WORKER,JOB,CHECK,DURATION,SDATE,STIME"
}

func (e *resultEncoder) Encode(r *domain.Result) string {
	a := []string{
		r.Time.Format("2006-01-02"),
		r.Time.Format("15:04:05"),
		"RESULT",
		Hostname(),
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

func InitResultEnconder() {
	presenters.InitResultEncoder(&resultEncoder{})
}
