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

// The basic structure needed by the tool:
// * a directed graph (tree) with the master (initial) requirement
//   at the top
// * a directed graph (tree) with the master topic at the top

package app

import (
	"fmt"
	"path/filepath"
	"os"
	
	"gopkg.in/yaml.v3"
	"go.uber.org/zap"
)

type TypeChecker struct {
	Type DataType `yaml:"type"`
}

type RqmgData struct {
	Requirements *RequirementGraph
	// Topics *simple.DirectedGraph
}

func NewRqmgData() *RqmgData {
	var rqmgdata *RqmgData
	rqmgdata = new(RqmgData)
	rqmgdata.Requirements = NewRequirementGraph()
	return rqmgdata
}

func filterFile(log zap.Logger, path string, info os.FileInfo) bool {
	if info.IsDir() {
		log.Debug("Skip because it is a directory",
			zap.String("path", path))
		return false
	}
	if filepath.Ext(path) == ".yaml" {
		log.Info("Reading",
			zap.String("path", path))
		return true
	}
	log.Debug("Skip because it does not end in .yaml",
		zap.String("path", path))
	return false
}

func peekType(path string) *TypeChecker {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	typeChecker := new(TypeChecker)
	err = yaml.Unmarshal(data, typeChecker)
	if err != nil {
		panic(err)
	}
	return typeChecker
}

func ProcessRqmgData(log zap.Logger, dataDir string) *RqmgData {
	rqmgdata := NewRqmgData()

	var reqnamesmap map[string]*Requirement
	reqnamesmap = make(map[string]*Requirement)


	var topics []*Topic
	
	err := filepath.Walk(dataDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if ! filterFile(log, path, info) {
				return nil;
			}

			fileType := peekType(path)

			switch fileType.Type {
			case DTRequirement:
				nreq := NewRequirement(path)
				// requirements = append(requirements, nreq)
				reqnamesmap[nreq.Name] = nreq
			case DTTopic:
				topics = append(topics, NewTopic(path))
			}

			log.Info("typeChecker",
				zap.String("type", fileType.Type.String()),
				zap.String("path", path))

			return nil
		})
		
	if err != nil {
		fmt.Println(err)
	}

	var reqnodemap map[string]*RequirementNode
	reqnodemap = make(map[string]*RequirementNode)

	// Place all the nodes in the graph
	for name, requirement := range reqnamesmap {
		dnode := rqmgdata.Requirements.DirectedGraph.NewNode()
		node := RequirementNode{id: dnode.ID(), Requirement: requirement}
		rqmgdata.Requirements.AddNode(node)
		reqnodemap[requirement.Name] = &node
		log.Debug("Add node to requirement graph",
			zap.String("name", name),
			zap.Int64("id", dnode.ID()))
	}

	// Add the edges
	for name, requirement := range reqnamesmap {
		from_node := reqnodemap[name]

		for _, to_name := range requirement.SolvedBy {
			if to_node, ok := reqnodemap[to_name]; ! ok {
				log.Error("Node with name not found",
					zap.String("name", to_name))
			} else {
				new_edge := rqmgdata.Requirements.DirectedGraph.NewEdge(
					from_node, to_node)
				rqmgdata.Requirements.DirectedGraph.SetEdge(new_edge)
				log.Debug("Add edge",
					zap.String("from", name),
					zap.String("to", to_node.Requirement.Name))
			}
		}
	}

	fmt.Println("------------------")
	fmt.Printf("%#v\n", *rqmgdata.Requirements)
	fmt.Printf("Number of nodes [%d]\n",
		rqmgdata.Requirements.DirectedGraph.Nodes().Len())
	fmt.Println("..................")
	fmt.Println(reqnodemap)
	fmt.Println("++++++++++++++++++")

	return rqmgdata
}

// Local Variables:
// tab-width: 4
// End:
