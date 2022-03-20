package stack

import "com.github/MadhavaAdiga/GolangDS/standard"

type Stack interface {
	Push(interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	Size() int
	standard.Collection
}
