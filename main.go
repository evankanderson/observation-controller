package main

import (
	"fmt"

	"github.com/evankanderson/observation-controller/pkg/reconciler"
)

func main() {
	r := reconciler.ObservationReconciler{}
	fmt.Print(r)
}
