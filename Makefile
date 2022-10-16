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

export GOPATH=$(shell go env GOPATH)
export GOOS=$(shell go env GOOS)
export GOARCH=$(shell go env GOARCH)

export GO_BUILD=go build

# List of binary cmds to build
CMDS := \
	bin/$(GOOS)/rqmgtool

SOURCES := $(shell find . -name '*.go' -not -name '*_test.go') go.mod go.sum

all: generate $(CMDS)

bin/$(GOOS)/rqmgtool: $(SOURCES)
	$(GO_BUILD) -o $@ ./cmd/$(shell basename "$@")

# Ignore for now
generate:
