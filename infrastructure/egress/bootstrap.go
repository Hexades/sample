package egress

import "github.com/hexades/sqlite"

func New() {
	sqlite.NewRepository()
}
