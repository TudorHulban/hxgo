package dsl

import "io"

type Node func(io.Writer) (bool, []Style, error)
