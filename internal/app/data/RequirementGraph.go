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

package data

import (
	"go.uber.org/zap"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

// Graph
// As there is no (real) inheritance in golang, it was
// chosen to have a 'simple' graph and a dict to map
// ids to requirements.
// It looks that it is not possible to 'inherit' the graph.Node
// and add more fields there.

type RequirementGraph struct {
	*simple.DirectedGraph
	// Different 'views' towards the graph node content
	// ID -> Requirement mapping
	ReqsIDsMap map[int64]*Requirement
	// Name -> Requirement mapping
	ReqsNameMap map[string]*Requirement
	// Name -> ID mapping
	IDsNameMap map[string]int64
	// ID -> graph.Node
	NodesIDMap map[int64]*graph.Node
}

func NewRequirementGraph() *RequirementGraph {
	var reqGraph *RequirementGraph
	reqGraph = new(RequirementGraph)
	reqGraph.DirectedGraph = simple.NewDirectedGraph()
	reqGraph.ReqsIDsMap = make(map[int64]*Requirement)
	reqGraph.ReqsNameMap = make(map[string]*Requirement)
	reqGraph.IDsNameMap = make(map[string]int64)
	reqGraph.NodesIDMap = make(map[int64]*graph.Node)
	return reqGraph
}

func (self RequirementGraph) AddRequirement(
	requirement *Requirement) int64 {
	
	dnode := self.DirectedGraph.NewNode()
	self.AddNode(dnode)
	self.ReqsIDsMap[dnode.ID()] = requirement
	self.ReqsNameMap[requirement.Name] = requirement
	self.IDsNameMap[requirement.Name] = dnode.ID()
	self.NodesIDMap[dnode.ID()] = &dnode
	return dnode.ID()
}

func (self RequirementGraph) CreateAllEdges(log *zap.Logger) {
	for from_id, requirement := range self.ReqsIDsMap {
		// from_node := reqnodemap[name]

		for _, to_name := range requirement.SolvedBy {
			if to_id, ok := self.IDsNameMap[to_name]; ! ok {
				log.Error("Node with name not found",
					zap.String("name", to_name),
					zap.String("from", self.ReqsIDsMap[from_id].Name))
			} else {
				new_edge := self.NewEdge(*self.NodesIDMap[from_id],
					*self.NodesIDMap[to_id])
				self.SetEdge(new_edge)
				log.Debug("Add edge",
					zap.String("from", self.ReqsIDsMap[from_id].Name),
					zap.String("to", to_name))
			}
		}
	}
}

func (self RequirementGraph) GetReqByID(id int64) *Requirement {
	return self.ReqsIDsMap[id]
}

// Local Variables:
// tab-width: 4
// End:
