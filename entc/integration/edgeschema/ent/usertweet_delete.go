// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
)

// UserTweetDelete is the builder for deleting a UserTweet entity.
type UserTweetDelete struct {
	config
	hooks    []Hook
	mutation *UserTweetMutation
}

// Where appends a list predicates to the UserTweetDelete builder.
func (utd *UserTweetDelete) Where(ps ...predicate.UserTweet) *UserTweetDelete {
	utd.mutation.Where(ps...)
	return utd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (utd *UserTweetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UserTweetMutation](ctx, utd.sqlExec, utd.mutation, utd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (utd *UserTweetDelete) ExecX(ctx context.Context) int {
	n, err := utd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (utd *UserTweetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usertweet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertweet.FieldID,
			},
		},
	}
	if ps := utd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, utd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	utd.mutation.done = true
	return affected, err
}

// UserTweetDeleteOne is the builder for deleting a single UserTweet entity.
type UserTweetDeleteOne struct {
	utd *UserTweetDelete
}

// Where appends a list predicates to the UserTweetDelete builder.
func (utdo *UserTweetDeleteOne) Where(ps ...predicate.UserTweet) *UserTweetDeleteOne {
	utdo.utd.mutation.Where(ps...)
	return utdo
}

// Exec executes the deletion query.
func (utdo *UserTweetDeleteOne) Exec(ctx context.Context) error {
	n, err := utdo.utd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usertweet.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (utdo *UserTweetDeleteOne) ExecX(ctx context.Context) {
	if err := utdo.Exec(ctx); err != nil {
		panic(err)
	}
}
