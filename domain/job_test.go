package domain_test

import (
	"github.com/tamura2004/httperf2/domain"
	"testing"
)

func TestNewJob(t *testing.T) {
	j := domain.NewJob(2)

	got := j.Id
	want := 2

	if got != want {
		t.Errorf("bad job id got %s want %s", got, want)
	}

}
