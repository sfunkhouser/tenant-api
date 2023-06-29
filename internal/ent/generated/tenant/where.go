// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package tenant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/tenant-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldDescription, v))
}

// ParentTenantID applies equality check predicate on the "parent_tenant_id" field. It's identical to ParentTenantIDEQ.
func ParentTenantID(v gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldParentTenantID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Tenant {
	return predicate.Tenant(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Tenant {
	return predicate.Tenant(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Tenant {
	return predicate.Tenant(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Tenant {
	return predicate.Tenant(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Tenant {
	return predicate.Tenant(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Tenant {
	return predicate.Tenant(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Tenant {
	return predicate.Tenant(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Tenant {
	return predicate.Tenant(sql.FieldContainsFold(FieldDescription, v))
}

// ParentTenantIDEQ applies the EQ predicate on the "parent_tenant_id" field.
func ParentTenantIDEQ(v gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldEQ(FieldParentTenantID, v))
}

// ParentTenantIDNEQ applies the NEQ predicate on the "parent_tenant_id" field.
func ParentTenantIDNEQ(v gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldNEQ(FieldParentTenantID, v))
}

// ParentTenantIDIn applies the In predicate on the "parent_tenant_id" field.
func ParentTenantIDIn(vs ...gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldIn(FieldParentTenantID, vs...))
}

// ParentTenantIDNotIn applies the NotIn predicate on the "parent_tenant_id" field.
func ParentTenantIDNotIn(vs ...gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldNotIn(FieldParentTenantID, vs...))
}

// ParentTenantIDGT applies the GT predicate on the "parent_tenant_id" field.
func ParentTenantIDGT(v gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldGT(FieldParentTenantID, v))
}

// ParentTenantIDGTE applies the GTE predicate on the "parent_tenant_id" field.
func ParentTenantIDGTE(v gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldGTE(FieldParentTenantID, v))
}

// ParentTenantIDLT applies the LT predicate on the "parent_tenant_id" field.
func ParentTenantIDLT(v gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldLT(FieldParentTenantID, v))
}

// ParentTenantIDLTE applies the LTE predicate on the "parent_tenant_id" field.
func ParentTenantIDLTE(v gidx.PrefixedID) predicate.Tenant {
	return predicate.Tenant(sql.FieldLTE(FieldParentTenantID, v))
}

// ParentTenantIDContains applies the Contains predicate on the "parent_tenant_id" field.
func ParentTenantIDContains(v gidx.PrefixedID) predicate.Tenant {
	vc := string(v)
	return predicate.Tenant(sql.FieldContains(FieldParentTenantID, vc))
}

// ParentTenantIDHasPrefix applies the HasPrefix predicate on the "parent_tenant_id" field.
func ParentTenantIDHasPrefix(v gidx.PrefixedID) predicate.Tenant {
	vc := string(v)
	return predicate.Tenant(sql.FieldHasPrefix(FieldParentTenantID, vc))
}

// ParentTenantIDHasSuffix applies the HasSuffix predicate on the "parent_tenant_id" field.
func ParentTenantIDHasSuffix(v gidx.PrefixedID) predicate.Tenant {
	vc := string(v)
	return predicate.Tenant(sql.FieldHasSuffix(FieldParentTenantID, vc))
}

// ParentTenantIDIsNil applies the IsNil predicate on the "parent_tenant_id" field.
func ParentTenantIDIsNil() predicate.Tenant {
	return predicate.Tenant(sql.FieldIsNull(FieldParentTenantID))
}

// ParentTenantIDNotNil applies the NotNil predicate on the "parent_tenant_id" field.
func ParentTenantIDNotNil() predicate.Tenant {
	return predicate.Tenant(sql.FieldNotNull(FieldParentTenantID))
}

// ParentTenantIDEqualFold applies the EqualFold predicate on the "parent_tenant_id" field.
func ParentTenantIDEqualFold(v gidx.PrefixedID) predicate.Tenant {
	vc := string(v)
	return predicate.Tenant(sql.FieldEqualFold(FieldParentTenantID, vc))
}

// ParentTenantIDContainsFold applies the ContainsFold predicate on the "parent_tenant_id" field.
func ParentTenantIDContainsFold(v gidx.PrefixedID) predicate.Tenant {
	vc := string(v)
	return predicate.Tenant(sql.FieldContainsFold(FieldParentTenantID, vc))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
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
func Not(p predicate.Tenant) predicate.Tenant {
	return predicate.Tenant(func(s *sql.Selector) {
		p(s.Not())
	})
}
