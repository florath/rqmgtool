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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/florath/rqmgtool/internal/pkg/config"
)

func main() {
	fmt.Println("+++ rqmgtool starting")

	var configFile string
	var dataDir string

	flag.StringVar(&configFile, "config", "", "config file name")
	flag.StringVar(&dataDir, "dataDir", "", "directory of the requirements data")
	flag.Parse()

	if len(configFile) == 0 || len(dataDir) == 0 {
		fmt.Println("Usage: rqmgtool")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("ConfigFile [%s]\n", configFile)
	fmt.Printf("DataDir [%s]\n", dataDir)

	cfg := config.NewConfig(configFile)

	fmt.Printf("Type [%s]\n", cfg.Type)
	fmt.Printf("Config [%v]\n", cfg)

	fmt.Println("+++ rqmgtool This is the End")
}

// Local Variables:
// tab-width: 4
// End:
