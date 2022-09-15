package main

import "fmt"

type QMGRNameNotFoundErr struct{}

func (e QMGRNameNotFoundErr) Error() string {
	return fmt.Sprintf("%T: Qmgr name not set in env variable", e)
}
