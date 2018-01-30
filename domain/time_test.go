package domain_test

import (
	"github.com/tamura2004/httperf2/domain"
	"testing"
	"time"
)

func TestExpRand(t *testing.T) {
	d := &domain.Duration{
		1 * time.Second,
	}

	sum := time.Duration(0)

	for i := 0; i < 100; i++ {
		sum += d.ExpRand()
	}

	want := 100 * time.Second
	from := want - 10*time.Second
	to := want + 10*time.Second

	if sum < from || to < sum {
		t.Errorf("bad duration exp random got %s want about %s", sum, want)
	}
}
