package dbx

import (
	"fmt"

	"github.com/kcmvp/got/internal/str"
	"github.com/samber/lo"
)

type IEntity interface {
	Table() string
}

type Key interface {
	string | int64
}

type Sql interface {
	SqlStr() string
}

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

type Attr str.Mapper

func (a Attr) SqlStr() string {
	return a.B
}

type Set lo.Tuple2[Attr, any]

// Criteria attribute predicate
type Criteria struct {
	expression string
}

func (criteria Criteria) SqlStr() string {
	return criteria.expression
}

func (criteria Criteria) Or(rp Criteria) Criteria {
	return Criteria{
		expression: fmt.Sprintf("(%s or %s)", criteria.expression, rp.SqlStr()),
	}
}

func (criteria Criteria) Add(rp Criteria) Criteria {
	return Criteria{
		expression: fmt.Sprintf("(%s and %s)", criteria.expression, rp.SqlStr()),
	}
}

func LT(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s < ?", attr.SqlStr())}
}

func LTE(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s <= ?", attr.SqlStr())}
}

func GT(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s > ?", attr.SqlStr())}
}

func GTE(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s >= ?", attr.SqlStr())}
}

func Null(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s is null", attr.SqlStr())}
}

func NotNull(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s is not null", attr.SqlStr())}
}

func Like(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s like '%%?%%'", attr.SqlStr())}
}

func NotLike(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s not like '%%?%%'", attr.SqlStr())}
}

func Prefix(attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s like '?%%'", attr.SqlStr())}
}

func Suffix[E IEntity](attr Attr) Criteria {
	return Criteria{expression: fmt.Sprintf("%s like '%%?'", attr.SqlStr())}
}

type OrderBy lo.Tuple2[Attr, Order]

func (orderBy OrderBy) SqlStr() string {
	return fmt.Sprintf("%s %s", orderBy.A.SqlStr(), orderBy.B)
}

var (
	_ Sql = (*Attr)(nil)
	_ Sql = (*Criteria)(nil)
	_ Sql = (*OrderBy)(nil)
)
