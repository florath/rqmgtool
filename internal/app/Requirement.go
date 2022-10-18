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
	"os"

	"gopkg.in/yaml.v3"
)

// Requirement

type Requirement struct {
	SubType ReqSubType `yaml:"subtype"`
	Name string `yaml:"name"`
	SolvedBy []string `yaml:"solved-by,flow"`
}

func NewRequirement(path string) *Requirement {
	requirement := new(Requirement)

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, requirement)
	if err != nil {
		panic(err)
	}
	
	return requirement
}

// Local Variables:
// tab-width: 4
// End:
