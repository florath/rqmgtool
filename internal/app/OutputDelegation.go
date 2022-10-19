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
	"go.uber.org/zap"
	"github.com/florath/rqmgtool/internal/app/data"
	"github.com/florath/rqmgtool/internal/app/output"
)

func OutputDelegation(log *zap.Logger, rqmgdata *data.RqmgData, oname string,
	vals map[string]string) {

	log.Info("OutputDelegation",
		zap.String("name", oname))

	switch oname {
	case "requirements-graph":
		output.RequirementsGraph(log, rqmgdata, vals)
	default:
		log.Error("Output module not found",
			zap.String("name", oname))
	}
}

// Local Variables:
// tab-width: 4
// End:
