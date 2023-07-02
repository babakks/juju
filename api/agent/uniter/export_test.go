// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package uniter

import (
	"github.com/juju/names/v4"

	"github.com/juju/juju/core/life"
	"github.com/juju/juju/rpc/params"
)

var (
	NewSettings = newSettings
)

// CreateUnit creates uniter.Unit for tests.
func CreateUnit(client *Client, tag names.UnitTag) *Unit {
	return &Unit{
		client:       client,
		tag:          tag,
		life:         life.Alive,
		resolvedMode: params.ResolvedNone,
	}
}

// CreateRelation creates uniter.Relation for tests.
func CreateRelation(client *Client, tag names.RelationTag) *Relation {
	return &Relation{
		client: client,
		tag:    tag,
		id:     666,
	}
}

// CreateRelationUnit creates uniter.RelationUnit for tests.
func CreateRelationUnit(client *Client, relationTag names.RelationTag, unitTag names.UnitTag) *RelationUnit {
	return &RelationUnit{
		client:   client,
		unitTag:  unitTag,
		relation: &Relation{tag: relationTag},
	}
}
