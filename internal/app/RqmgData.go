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

import (
	"github.com/gonum/gonum/graph/simple"
)

type RqmgData struct {
	Requirements simple.DirectedGraph
	Topics simple.DirectedGraph
}

// Local Variables:
// tab-width: 4
// End:
