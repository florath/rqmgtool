# (c) 2022 by flonatel GmbH & Co. KG / Andreas Florath
#
# SPDX-License-Identifier: GPL-3.0-or-later
#
# This file is part of rqmgtool.
#
# rqmgtool is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# at your option any later version.
#
# rqmgtool is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with rqmgtool.  If not, see <https://www.gnu.org/licenses/>.

RQMGTOOL_BINARY = ../../bin/linux/rqmgtool

RQMGTOOL_TARGETS = artifacts/rqmgtool-graph1.dot

YAML_FILES := $(shell find topics -name '*.yaml')

all: artifacts/rqmgtool-graph1.png

$(RQMGTOOL_TARGETS): $(RQMGTOOL_BINARY) $(YAML_FILES)
	$(RQMGTOOL_BINARY) --config=rqmgtool-config.yaml --dataDir=topics

artifacts/rqmgtool-graph1.png: artifacts/rqmgtool-graph1.dot
	unflatten -l 23 artifacts/rqmgtool-graph1.dot | \
		dot -Tpng -o artifacts/rqmgtool-graph1.png
