package dal

type SimpleSelect[E IEntity] struct {
	cols      Attr[E]
	predicate Predicate
	orders    []Order
}

func (s SimpleSelect[E]) SqlStr() string {
	//TODO implement me
	panic("implement me")
}

var _ Sql = (*SimpleSelect[IEntity])(nil)

func Select[E IEntity](cols ...Attr[E]) SimpleSelect[E] {
	return SimpleSelect[E]{}
}

func (s SimpleSelect[E]) Where(predicate Predicate) SimpleSelect[E] {
	s.predicate = predicate
	return s
}
func (s SimpleSelect[E]) OrderBy(orders ...Order) SimpleSelect[E] {
	s.orders = orders
	return s
}

type JoinSelect[E1 IEntity, E2 IEntity] struct {
	cols1     []Attr[E1]
	cols2     []Attr[E2]
	joins     []TableJoin[E1, E2]
	predicate Predicate
	orderBy   []Order
}

func (s JoinSelect[E1, E2]) SqlStr() string {
	//TODO implement me
	panic("implement me")
}

var _ Sql = (*JoinSelect[IEntity, IEntity])(nil)

func SelectJoin[E1 IEntity, E2 IEntity](joins []TableJoin[E1, E2], cols1 []Attr[E1], cols2 ...Attr[E2]) JoinSelect[E1, E2] {
	return JoinSelect[E1, E2]{}
}

func (s JoinSelect[E1, E2]) Where(predicate Predicate) JoinSelect[E1, E2] {
	return JoinSelect[E1, E2]{}
}

func (s JoinSelect[E1, E2]) OrderBy(orders ...Order) JoinSelect[E1, E2] {
	return JoinSelect[E1, E2]{}
}
