package gr_test

import (
	"log"
	"testing"

	"github.com/xzip/gr"
)

func TestTransactionBegin(t *testing.T) {
	log.Println("[Testing Transaction]")
}

func TestTransaction(t *testing.T) {

	safeTestContext(func() {

		var get *gr.RespString

		err := redis.Pipelined(func(p *gr.Pipeline) {
			p.Multi()
			p.Set("gr::multikey", "multivalue")
			get = p.Get("gr::multikey")
			p.Exec()
		})

		if err != nil {
			t.Fail()
		}

	})
}

func TestTransactionEnd(t *testing.T) {
	println("[OK]")
}
