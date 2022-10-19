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
	"github.com/florath/rqmgtool/internal/app/data"
)

type TypeChecker struct {
	Type DataType `yaml:"type"`
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

func generateRequirementsGraph(
	log zap.Logger,
	reqnamesmap map[string]*data.Requirement) *data.RequirementGraph {

	var reqGraph *data.RequirementGraph
	reqGraph = data.NewRequirementGraph()

	// Place all the nodes in the graph
	for name, requirement := range reqnamesmap {
		new_id := reqGraph.AddRequirement(requirement)
		log.Debug("Add node to requirement graph",
			zap.String("name", name),
			zap.Int64("id", new_id))
	}

	reqGraph.CreateAllEdges(&log)
	
	return reqGraph
}

func ProcessRqmgData(log zap.Logger, dataDir string) *data.RqmgData {
	rqmgdata := data.NewRqmgData()

	var reqnamesmap map[string]*data.Requirement
	reqnamesmap = make(map[string]*data.Requirement)

	var topics []*data.Topic
	
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
				nreq := data.NewRequirement(path)
				reqnamesmap[nreq.Name] = nreq
			case DTTopic:
				topics = append(topics, data.NewTopic(path))
			}

			log.Info("typeChecker",
				zap.String("type", fileType.Type.String()),
				zap.String("path", path))

			return nil
		})
		
	if err != nil {
		fmt.Println(err)
	}

	rqmgdata.Requirements = generateRequirementsGraph(log, reqnamesmap)

	return rqmgdata
}

// Local Variables:
// tab-width: 4
// End:
