// Code generated by entc, DO NOT EDIT.

package resource

import (
	"lab-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// Parent applies equality check predicate on the "parent" field. It's identical to ParentEQ.
func Parent(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParent), v))
	})
}

// Router applies equality check predicate on the "router" field. It's identical to RouterEQ.
func Router(v bool) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRouter), v))
	})
}

// Nav applies equality check predicate on the "nav" field. It's identical to NavEQ.
func Nav(v bool) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNav), v))
	})
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v uint) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPath), v))
	})
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPath), v...))
	})
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPath), v...))
	})
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPath), v))
	})
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPath), v))
	})
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPath), v))
	})
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPath), v))
	})
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPath), v))
	})
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPath), v))
	})
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPath), v))
	})
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPath), v))
	})
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPath), v))
	})
}

// ParentEQ applies the EQ predicate on the "parent" field.
func ParentEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParent), v))
	})
}

// ParentNEQ applies the NEQ predicate on the "parent" field.
func ParentNEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParent), v))
	})
}

// ParentIn applies the In predicate on the "parent" field.
func ParentIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldParent), v...))
	})
}

// ParentNotIn applies the NotIn predicate on the "parent" field.
func ParentNotIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldParent), v...))
	})
}

// ParentGT applies the GT predicate on the "parent" field.
func ParentGT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParent), v))
	})
}

// ParentGTE applies the GTE predicate on the "parent" field.
func ParentGTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParent), v))
	})
}

// ParentLT applies the LT predicate on the "parent" field.
func ParentLT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParent), v))
	})
}

// ParentLTE applies the LTE predicate on the "parent" field.
func ParentLTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParent), v))
	})
}

// ParentContains applies the Contains predicate on the "parent" field.
func ParentContains(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldParent), v))
	})
}

// ParentHasPrefix applies the HasPrefix predicate on the "parent" field.
func ParentHasPrefix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldParent), v))
	})
}

// ParentHasSuffix applies the HasSuffix predicate on the "parent" field.
func ParentHasSuffix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldParent), v))
	})
}

// ParentEqualFold applies the EqualFold predicate on the "parent" field.
func ParentEqualFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldParent), v))
	})
}

// ParentContainsFold applies the ContainsFold predicate on the "parent" field.
func ParentContainsFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldParent), v))
	})
}

// RouterEQ applies the EQ predicate on the "router" field.
func RouterEQ(v bool) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRouter), v))
	})
}

// RouterNEQ applies the NEQ predicate on the "router" field.
func RouterNEQ(v bool) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRouter), v))
	})
}

// NavEQ applies the EQ predicate on the "nav" field.
func NavEQ(v bool) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNav), v))
	})
}

// NavNEQ applies the NEQ predicate on the "nav" field.
func NavNEQ(v bool) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNav), v))
	})
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIcon), v))
	})
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIcon), v))
	})
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIcon), v...))
	})
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIcon), v...))
	})
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIcon), v))
	})
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIcon), v))
	})
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIcon), v))
	})
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIcon), v))
	})
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIcon), v))
	})
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIcon), v))
	})
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIcon), v))
	})
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIcon)))
	})
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIcon)))
	})
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIcon), v))
	})
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIcon), v))
	})
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v uint) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v uint) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSort), v))
	})
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...uint) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSort), v...))
	})
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...uint) predicate.Resource {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resource(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSort), v...))
	})
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v uint) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSort), v))
	})
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v uint) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSort), v))
	})
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v uint) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSort), v))
	})
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v uint) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSort), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Resource) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Resource) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Resource) predicate.Resource {
	return predicate.Resource(func(s *sql.Selector) {
		p(s.Not())
	})
}
