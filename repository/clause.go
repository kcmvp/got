package repository

import (
	"fmt"
	"github.com/kcmvp/got/internal/repository"
	"github.com/samber/lo"
)

type Sql interface {
	SqlStr() string
}

type Attr[E IEntity] internal.Mapper

var _ Sql = (*Attr[IEntity])(nil)

func (a Attr[E]) SqlStr() string {
	return a.B
}

// Order sql order clause
type Order Attr[IEntity]

func DESC[E IEntity](attr Attr[E]) Order {
	return Order{
		A: fmt.Sprintf("%s desc", attr.B),
		B: attr.B,
	}
}

func ASC[E IEntity](attr Attr[E]) Order {
	return Order{
		A: fmt.Sprintf("%s asc", attr.B),
		B: attr.B,
	}
}

// Predicate attribute predicate
type Predicate struct {
	expression string
}

var _ Sql = (*Predicate)(nil)

func (p Predicate) SqlStr() string {
	return p.expression
}

func (p Predicate) And(rp Predicate) Predicate {
	return Predicate{
		expression: fmt.Sprintf("(%s and %s)", p.expression, rp.SqlStr()),
	}
}

func (p Predicate) Or(rp Predicate) Predicate {
	return Predicate{
		expression: fmt.Sprintf("(%s or %s)", p.expression, rp.SqlStr()),
	}
}

func EQ[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func NotEQ[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func LT[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func LTE[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func GT[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func GTE[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func Null[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func NotNull[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func Like[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func NotLike[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func Prefix[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}
func Suffix[E IEntity](attr Attr[E]) Predicate {
	// todo
	return Predicate{""}
}

type TableJoin[E1 IEntity, E2 IEntity] lo.Tuple3[Attr[E1], Attr[E2], string]
