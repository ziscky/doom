/*
Copyright (C) 2016  Eric Ziscky

    This program is free software; you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation; either version 2 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License along
    with this program; if not, write to the Free Software Foundation, Inc.,
    51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/

package policy

import (
	"io/ioutil"
	"strings"
)

//policy gets the default behaviour of the OOM killer

func oomKillerEnabled() string {
	out, err := ioutil.ReadFile("/proc/sys/vm/overcommit_memory")
	if err != nil {
		return ""
	}

	switch strings.Trim(string(out), "\n") {
	case "0":
		return "kernel attempts to estimate the amount of free memory left when userspace requests more memory"
	case "1":
		return "kernel pretends there is always enough memory until it actually runs out."
	case "2":
		return "kernel uses a never overcommit policy that attempts to prevent any overcommit of memory"
	default:
		return "invalid vm.overcommit_memory value"
	}

}

func oomKillAllocatingTask() string {
	out, err := ioutil.ReadFile("/proc/sys/vm/oom_kill_allocating_task")
	if err != nil {
		return ""
	}
	if strings.Trim(string(out), "\n") != "0" {
		return "kernel kills overcommiting task"
	}
	return "kernel does heuristics on all the tasks choosing the best to kill"
}

func oomMinFreeKbytes() {}

func oomDumpTasks() string {
	out, err := ioutil.ReadFile("/proc/sys/vm/oom_dump_tasks")
	if err != nil {
		return ""
	}
	if strings.Trim(string(out), "\n") != "0" {
		return "system-wide task dump (excluding kernel threads) to be produced when the kernel performs an OOM-killing"
	}
	return "kernel does not dump all tasks, may save performance on large systems"
}

func oomPanic() string {
	out, err := ioutil.ReadFile("/proc/sys/vm/panic_on_oom")
	if err != nil {
		return ""
	}

	switch strings.Trim(string(out), "\n") {
	case "0":
		return "kernel does not panic,calls the oom_killer"
	case "1":
		return "kernel panics incase of an oom except if a process limits using nodes by mempolicy/cpusets,and those nodes become memory exhaustion status, one process may be killed by oom-killer"
	case "2":
		return "kernel panics regardless of mempolicy/cpusets limiting"
	default:
		return "invalid vm.panic_on_oom value"
	}
}

//BuildReport builds a report of all relevant information regarding OOM behaviour
func BuildReport() map[string]string {
	return map[string]string{
		"Overcommit Memory Policy": oomKillerEnabled(),
		"Kill Allocating Task":     oomKillAllocatingTask(),
		"Dump Tasks":               oomDumpTasks(),
		"Kernel Panic":             oomPanic(),
	}
}
