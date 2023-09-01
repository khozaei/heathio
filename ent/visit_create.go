// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khozaei/healthio/ent/attachment"
	"github.com/khozaei/healthio/ent/reception"
	"github.com/khozaei/healthio/ent/visit"
)

// VisitCreate is the builder for creating a Visit entity.
type VisitCreate struct {
	config
	mutation *VisitMutation
	hooks    []Hook
}

// SetVisitPrice sets the "visit_price" field.
func (vc *VisitCreate) SetVisitPrice(s string) *VisitCreate {
	vc.mutation.SetVisitPrice(s)
	return vc
}

// SetNillableVisitPrice sets the "visit_price" field if the given value is not nil.
func (vc *VisitCreate) SetNillableVisitPrice(s *string) *VisitCreate {
	if s != nil {
		vc.SetVisitPrice(*s)
	}
	return vc
}

// SetVisitedAt sets the "visited_at" field.
func (vc *VisitCreate) SetVisitedAt(t time.Time) *VisitCreate {
	vc.mutation.SetVisitedAt(t)
	return vc
}

// SetNillableVisitedAt sets the "visited_at" field if the given value is not nil.
func (vc *VisitCreate) SetNillableVisitedAt(t *time.Time) *VisitCreate {
	if t != nil {
		vc.SetVisitedAt(*t)
	}
	return vc
}

// SetPaymentType sets the "payment_type" field.
func (vc *VisitCreate) SetPaymentType(s string) *VisitCreate {
	vc.mutation.SetPaymentType(s)
	return vc
}

// SetNillablePaymentType sets the "payment_type" field if the given value is not nil.
func (vc *VisitCreate) SetNillablePaymentType(s *string) *VisitCreate {
	if s != nil {
		vc.SetPaymentType(*s)
	}
	return vc
}

// SetIsPaid sets the "is_paid" field.
func (vc *VisitCreate) SetIsPaid(b bool) *VisitCreate {
	vc.mutation.SetIsPaid(b)
	return vc
}

// SetNillableIsPaid sets the "is_paid" field if the given value is not nil.
func (vc *VisitCreate) SetNillableIsPaid(b *bool) *VisitCreate {
	if b != nil {
		vc.SetIsPaid(*b)
	}
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VisitCreate) SetCreatedAt(t time.Time) *VisitCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VisitCreate) SetNillableCreatedAt(t *time.Time) *VisitCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VisitCreate) SetUpdatedAt(t time.Time) *VisitCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VisitCreate) SetNillableUpdatedAt(t *time.Time) *VisitCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetDeletedAt sets the "deleted_at" field.
func (vc *VisitCreate) SetDeletedAt(t time.Time) *VisitCreate {
	vc.mutation.SetDeletedAt(t)
	return vc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vc *VisitCreate) SetNillableDeletedAt(t *time.Time) *VisitCreate {
	if t != nil {
		vc.SetDeletedAt(*t)
	}
	return vc
}

// SetReceptionID sets the "reception" edge to the Reception entity by ID.
func (vc *VisitCreate) SetReceptionID(id int) *VisitCreate {
	vc.mutation.SetReceptionID(id)
	return vc
}

// SetNillableReceptionID sets the "reception" edge to the Reception entity by ID if the given value is not nil.
func (vc *VisitCreate) SetNillableReceptionID(id *int) *VisitCreate {
	if id != nil {
		vc = vc.SetReceptionID(*id)
	}
	return vc
}

// SetReception sets the "reception" edge to the Reception entity.
func (vc *VisitCreate) SetReception(r *Reception) *VisitCreate {
	return vc.SetReceptionID(r.ID)
}

// AddAttachmentIDs adds the "attachment" edge to the Attachment entity by IDs.
func (vc *VisitCreate) AddAttachmentIDs(ids ...int) *VisitCreate {
	vc.mutation.AddAttachmentIDs(ids...)
	return vc
}

// AddAttachment adds the "attachment" edges to the Attachment entity.
func (vc *VisitCreate) AddAttachment(a ...*Attachment) *VisitCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return vc.AddAttachmentIDs(ids...)
}

// Mutation returns the VisitMutation object of the builder.
func (vc *VisitCreate) Mutation() *VisitMutation {
	return vc.mutation
}

// Save creates the Visit in the database.
func (vc *VisitCreate) Save(ctx context.Context) (*Visit, error) {
	vc.defaults()
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VisitCreate) SaveX(ctx context.Context) *Visit {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VisitCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VisitCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VisitCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := visit.DefaultCreatedAt
		vc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VisitCreate) check() error {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Visit.created_at"`)}
	}
	return nil
}

func (vc *VisitCreate) sqlSave(ctx context.Context) (*Visit, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VisitCreate) createSpec() (*Visit, *sqlgraph.CreateSpec) {
	var (
		_node = &Visit{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(visit.Table, sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.VisitPrice(); ok {
		_spec.SetField(visit.FieldVisitPrice, field.TypeString, value)
		_node.VisitPrice = value
	}
	if value, ok := vc.mutation.VisitedAt(); ok {
		_spec.SetField(visit.FieldVisitedAt, field.TypeTime, value)
		_node.VisitedAt = value
	}
	if value, ok := vc.mutation.PaymentType(); ok {
		_spec.SetField(visit.FieldPaymentType, field.TypeString, value)
		_node.PaymentType = value
	}
	if value, ok := vc.mutation.IsPaid(); ok {
		_spec.SetField(visit.FieldIsPaid, field.TypeBool, value)
		_node.IsPaid = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(visit.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(visit.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vc.mutation.DeletedAt(); ok {
		_spec.SetField(visit.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := vc.mutation.ReceptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   visit.ReceptionTable,
			Columns: []string{visit.ReceptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reception.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.reception_visit = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   visit.AttachmentTable,
			Columns: []string{visit.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VisitCreateBulk is the builder for creating many Visit entities in bulk.
type VisitCreateBulk struct {
	config
	builders []*VisitCreate
}

// Save creates the Visit entities in the database.
func (vcb *VisitCreateBulk) Save(ctx context.Context) ([]*Visit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Visit, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VisitMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VisitCreateBulk) SaveX(ctx context.Context) []*Visit {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VisitCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VisitCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
