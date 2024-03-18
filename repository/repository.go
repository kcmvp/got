package dal

import (
	"fmt"
	"github.com/kcmvp/got/boot"
	"github.com/samber/do/v2"
)

type IEntity interface {
	Table() string
}

type Key interface {
	string | int64
}

type AndOr int

const (
	And AndOr = iota
	Or
)

type Repository[E IEntity, K Key] interface {
	Insert(entity E) (int64, error)
	BatchInsert(entities []E) (int64, error)
	Delete(Key K) (int64, error)
	Find(Key K) (E, error)
	FindKeys(keys []K) ([]E, error)
	DeleteSample(sample E, andOr AndOr) (int64, error)
	DeleteSampleCriteria(sample E, where Predicate) (int64, error)
	UpdateSample(sample E, andOr AndOr) (int64, error)
	UpdateSampleCriteria(sample E, update []Attr[E], where Predicate) (int64, error)
	SearchSample(sample E, andOr AndOr, orderBy ...Order) ([]E, error)
	SearchSampleCriteria(sample E, where Predicate, orderBy ...Order) ([]E, error)
	Count() (int, error)
	All() ([]E, error)
}

type MustBeStructError struct {
	msg string
}

func (e MustBeStructError) Error() string {
	return fmt.Sprintf("must be struct %s", e.msg)
}

// defaultRepository default Repository implementation
type defaultRepository[E IEntity, K Key] struct {
	zeroK K
	zeroE E
	//sqlBuilder *SqlBuilder
	dbx DBX
}

func (d defaultRepository[E, K]) Insert(entity E) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) BatchInsert(entities []E) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) Delete(Key K) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) Find(Key K) (E, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) FindKeys(keys []K) ([]E, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) DeleteSample(entity E, andOr AndOr) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) DeleteSampleCriteria(entity E, where Predicate) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) UpdateSample(entity E, andOr AndOr) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) UpdateSampleCriteria(entity E, updating []Attr[E], where Predicate) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) SearchSample(entity E, andOr AndOr, orderBy ...Order) ([]E, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) SearchSampleCriteria(entity E, where Predicate, orderBy ...Order) ([]E, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) Count() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultRepository[E, K]) All() ([]E, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository[E IEntity, K Key]() Repository[E, K] {
	return NewRepositoryWithDS[E, K](DefaultDS)
}

func NewRepositoryWithDS[E IEntity, K Key](dsName string) Repository[E, K] {
	repo := &defaultRepository[E, K]{
		dbx: do.MustInvokeNamed[DBX](boot.Container(), dsName),
	}
	//@todo validate with Attr
	return repo
}
