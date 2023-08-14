package app

import (
	"github.com/hexades/sample/infrastructure/egress"
	"github.com/hexades/sample/infrastructure/ingress"
)

func New() {
	go egress.New()
	go ingress.New()
	doKeepAlive()
}

// Temporary until controls are put in place...
func doKeepAlive() {
	keepAlive := true
	for keepAlive == true {
	}

}
