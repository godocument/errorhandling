package main

// a piece of error handling codelet
//
// this code only demonstrates how to hide sql.ErrNoRows to the caller
//
// CAUTION: it will throw segmentation fault error since the required database connection is uninitialized
//

import (
	"context"
	"fmt"

	"github.com/godocument/errorhandling/pkg/project"
)

func main() {
	mgr := project.New()
	ctxt := context.Background()
	if _, err := mgr.Get(ctxt, 1); err != nil {
		fmt.Printf("%v", err)
	}
	return
}
