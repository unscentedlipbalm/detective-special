package main

import (
	"fmt"

	"github.com/scentedlipbalm/detective-special/internal/db"
)

func main() {
	fmt.Println(db.GetEntry("foo"))
}
