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

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"go.uber.org/zap"
)

type RequirementsInputConfig struct {
	Default_language string
}

type RequirementsConfig struct {
	Input RequirementsInputConfig
	Inventors []string `yaml:"inventors,flow"`
	Stakeholders []string `yaml:"stakeholders,flow"`
}

type OutputTypeConfig struct {
	Type string `yaml:"type"`
	Params map[string]string
}

type Config struct {
	Type string
	Requirements RequirementsConfig
	Logging zap.Config
	Output []*OutputTypeConfig `yaml:"output,flow"`
}

func NewConfig(configFile string) *Config {
	data, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	config := new(Config)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return config
}

// Local Variables:
// tab-width: 4
// End:
