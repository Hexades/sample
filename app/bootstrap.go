package app

import (
	"github.com/hexades/sample/infrastructure/egress"
	"github.com/hexades/sample/infrastructure/ingress"
	"time"
)

func New() {
	go egress.New()
	go ingress.New()
	time.Sleep(60 * time.Second)
}
