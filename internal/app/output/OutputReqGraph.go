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

package output

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"github.com/florath/rqmgtool/internal/app/data"
)

func dotEscape(s string) string {
	t := strings.Replace(s, "-", "_", -1)
	t = strings.Replace(t, " ", "_", -1)
	return t
}

func RequirementsGraph(log *zap.Logger, rqmgdata *data.RqmgData,
	vals map[string]string) {
	ofile := vals["ofile"]
	
	log.Info("output.RequirementsGraph",
		zap.String("ofile", ofile))

	fd, err := os.Create(ofile)
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fd.Close(); err != nil {
            panic(err)
        }
    }()

	fd.Write([]byte("digraph reqdeps {\nrankdir=TB;\nmclimit=10.0;\n"))
	fd.Write([]byte("nslimit=10.0;ranksep=1;\n"))

	node_iter := rqmgdata.Requirements.Nodes()
	for {
		ok := node_iter.Next()
		if ! ok {
			break
		}

		node := node_iter.Node()
		requirement := rqmgdata.Requirements.GetReqByID(node.ID())

		var nattr string
		nattr = ""

		if requirement.SubType == data.RSTDesignDecision {
			nattr += "color=green"
		}
		if requirement.Status == data.RSTATEOpen {
			if len(nattr) > 0 {
				nattr += ","
			}
			nattr += "fontcolor=red"
		}
		if requirement.Status == data.RSTATEAssigned {
			if len(nattr) > 0 {
				nattr += ","
			}
			nattr += ",fontcolor=blue"
		}
		
		fmt.Fprintf(fd, "\"%s\" [%s];\n",
			dotEscape(requirement.Name), nattr)
	}

	edge_iter := rqmgdata.Requirements.Edges()
	for {
		ok := edge_iter.Next()
		if ! ok {
			break
		}

		edge := edge_iter.Edge()
		node_from := edge.From()
		node_to := edge.To()
		req_from := rqmgdata.Requirements.GetReqByID(node_from.ID())
		req_to := rqmgdata.Requirements.GetReqByID(node_to.ID())

		fmt.Fprintf(fd, "\"%s\" -> \"%s\";\n",
			dotEscape(req_from.Name), dotEscape(req_to.Name))
	}
	
	fd.Write([]byte("}\n"))
}

// Local Variables:
// tab-width: 4
// End:
