// (c) 2022 by flonatel GmbH & Co. KG / Andreas Florath
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This file is part of rqmgtool.
//
// rqmgtool is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// at your option any later version.
//
// rqmgtool is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with rqmgtool.  If not, see <https://www.gnu.org/licenses/>.

package app

import (
	"gonum.org/v1/gonum/graph/simple"
)	

type RequirementGraph struct {
	*simple.DirectedGraph
}

func NewRequirementGraph() *RequirementGraph {
	var reqGraph *RequirementGraph
	reqGraph = new(RequirementGraph)
	reqGraph.DirectedGraph = simple.NewDirectedGraph()
	return reqGraph
}


type RequirementNode struct {
	id int64
	*Requirement
}

func (n RequirementNode) ID() int64 { return n.id }

// Local Variables:
// tab-width: 4
// End:
