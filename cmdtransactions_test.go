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

		err := redis.Multi(func(m *gr.Transaction) {
			m.Set("gr::multikey", "multivalue")
			get = m.Get("gr::multikey")
		})

		if err != nil {
			t.Fail()
		}

	})
}

func TestTransactionEnd(t *testing.T) {
	println("[OK]")
}
