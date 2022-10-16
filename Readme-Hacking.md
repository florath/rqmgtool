[//]: # (copyright 2022 by flonatel GmbH & Co. KG / Andreas Florath)
[//]: # ( )
[//]: # (SPDX-License-Identifier: GPL-3.0-or-later)
[//]: # ( )
[//]: # (This file is part of rqmgtool.)
[//]: # ( )  
[//]: # (rqmgtool is free software: you can redistribute it and/or modify)
[//]: # (it under the terms of the GNU General Public License as published by)
[//]: # (the Free Software Foundation, either version 3 of the License, or)
[//]: # (at your option any later version.)
[//]: # ( )
[//]: # (rqmgtool is distributed in the hope that it will be useful,)
[//]: # (but WITHOUT ANY WARRANTY; without even the implied warranty of)
[//]: # (MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the)
[//]: # (GNU General Public License for more details.)
[//]: # ( )  
[//]: # (You should have received a copy of the GNU General Public License)
[//]: # (along with rqmgtool.  If not, see <https://www.gnu.org/licenses/>.)

# Readme for rqmgtool Development and Hacking

rqmgtool is written in golang.

## Installation of golang

Go to

    https://go.dev/dl/

and download the lastest version. At least version 1.19.2 is needed.
Untar it and set the PATH to go/bin.

    export PATH=${PWD}/go/bin:${PATH}
	
## Addition of Modules

rqmgtool contains some modules - each handling a specific task like
configuration, reading of files, processing files, output modules, ...

Please check out the documentation about the 'standard' golang
directory layout:

    https://github.com/golang-standards/project-layout

## Initial Setup

The initial setup was done using

    go mod init github.com/florath/rqmgtool
	
The following was logged:

    go: to add module requirements and sums:
	    go mod tidy

## Build / Run

