package batchprocessor

import (
	. "github.com/deqing/batching-api/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess(t *testing.T) {
	jobs := []Job{{Status: QUEUED}, {Status: QUEUED}}
	bp := BatchProcessor{}
	bp.Process(jobs)

	assert.Equal(t, []Job{{Status: SUBMITTED}, {Status: SUBMITTED}}, jobs)
}
