// Copyright (c) 2016 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// +build linux

package main

import "github.com/ciao-project/ciao/deviceinfo"

func getOnlineCPUs() int {
	return deviceinfo.GetOnlineCPUs()
}

func getTotalMemory() int {
	total, _ := deviceinfo.GetMemoryInfo()
	total /= 1024
	return total
}

func getMemAndCpus() (mem int, cpus int) {
	cpus = getOnlineCPUs() / 2
	if cpus < 0 {
		cpus = 1
	} else if cpus > 8 {
		cpus = 8
	}

	mem = getTotalMemory() / 2
	if mem < 0 {
		mem = 1
	} else if mem > 8 {
		mem = 8
	}

	return mem, cpus
}
