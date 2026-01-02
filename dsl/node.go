package dsl

import "io"

type IsAttribute func() bool
type Render func(wr io.Writer) ([]Style, error)

type Node func(io.Writer) (IsAttribute, Render)
