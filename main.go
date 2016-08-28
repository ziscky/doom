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

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/ziscky/doom/calculator"
	"github.com/ziscky/doom/policy"

	"gopkg.in/urfave/cli.v1"
)

//getAllProcs gets all process directories in the/ proc pseudo filesystem
func getAllProcs() []int {
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		panic(err)
	}
	var procs []int
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		i, err := strconv.Atoi(file.Name())
		if err != nil { //not a PID file
			continue
		}
		procs = append(procs, i)
	}
	return procs
}

//getPIDCmd gets the command the pid was run with
//by reading the /proc/{pid}/cmdline
func getPIDCmd(pid int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return "*Error fetching xmd*"
	}
	cmd := strings.Trim(string(data), "\n")
	if len(cmd) == 0 {
		return "*none*"
	}
	return cmd
}

//getPIDName uses the ps utility to get the name of the process
func getPIDName(pid int) string {
	data, err := exec.Command("/bin/ps", "-p", fmt.Sprintf("%d", pid), "-o", "comm=").Output()
	if err != nil {
		fmt.Println(err)
		return "*Error fetching Name*"
	}
	return strings.Trim(string(data), "\n")
}

//resolveProcess resolves a process name to a list of pids using the pgrep utility
func resolveProcess(pname string) []int {
	var selected []int
	data, err := exec.Command("pgrep", "-f", pname).Output()
	if err != nil {
		return selected
	}

	pids := strings.Split(string(data), "\n")

	for _, pid := range pids {
		num, _ := strconv.Atoi(pid)
		if num == os.Getpid() { //don't report ourselves
			continue
		}
		selected = append(selected, num)
	}
	return selected
}

//calculate computes the OOM scores of all the passed PIDs and returns an ordered
//pairList of the key-value pair of PID - OOM score
func calculate(num int, best bool, procs ...int) {
	list := calculator.SortPIDs(num, best, procs...)
	for _, p := range list {
		fmt.Println("------------------------------------------")
		fmt.Print(fmt.Sprintf("PID:%d\nName:%s\nOOM Score:%d\nCMD:%s\n", p.Key, getPIDName(p.Key), p.Val, getPIDCmd(p.Key)))

	}
}

//_policy gets the policies that govern the operation of the OOM killer
func _policy() {
	fmt.Println("More info:", "www.kernel.org/doc/Documentation/sysctl/vm.txt")
	fmt.Println("------------------------------------------------------")
	for k, v := range policy.BuildReport() {
		fmt.Println(fmt.Sprintf("%s: %s", k, v))
	}
	fmt.Println("------------------------------------------------------")
}

func main() {
	app := cli.NewApp()
	app.Name = "doom"
	app.Version = "0.1"
	app.Usage = "An OOM navigator. More info: www.kernel.org/doc/Documentation/sysctl/vm.txt"
	app.Commands = []cli.Command{
		{
			Name:  "best",
			Usage: "Get best/lowest OOM scores",
			Action: func(c *cli.Context) error {
				var num int = 10
				var err error
				if c.NArg() > 0 {
					num, err = strconv.Atoi(c.Args().Get(0))
					if err != nil {
						fmt.Println("Invalid number:", c.Args().Get(0))
						return err
					}
					calculate(num, true, getAllProcs()...)
					return nil
				}
				//show all if number is not specified
				procs := getAllProcs()
				calculate(len(procs), true, procs...)
				return nil
			},
		},
		{
			Name:  "worst",
			Usage: "Get worst/highest OOM scores",
			Action: func(c *cli.Context) error {
				var num int = 10
				var err error
				if c.NArg() > 0 {
					num, err = strconv.Atoi(c.Args().Get(0))
					if err != nil {
						fmt.Println("Invalid number:", c.Args().Get(0))
						return err
					}
					calculate(num, false, getAllProcs()...)
					return nil
				}
				//show all if number is not specified
				procs := getAllProcs()
				calculate(len(procs), true, procs...)
				return nil
			},
		},
		{
			Name:  "policy",
			Usage: "Get OOM policy",
			Action: func(c *cli.Context) error {
				_policy()
				return nil
			},
		},
		{
			Name:  "inspect",
			Usage: "Get OOM score and details of a particular process and all its children",
			Action: func(c *cli.Context) error {
				var num int
				var pids []int
				var err error
				if c.NArg() < 1 {
					fmt.Println("doom inspect {pid|processname}")
					return nil
				}
				num, err = strconv.Atoi(c.Args().Get(0))
				if err != nil {
					//assume process name is passed
					pids = resolveProcess(c.Args().Get(0))
					if len(pids) == 0 {
						fmt.Println(fmt.Sprintf("No process with the name %s found", c.Args().Get(0)))
						return nil
					}
				} else { //add valid pid to calculation list
					pids = append(pids, num)
				}
				calculate(len(pids), true, pids...)
				return nil

			},
		},
		{
			Name:  "next",
			Usage: "Next process to be OOM killed",
			Action: func(c *cli.Context) error {
				calculate(1, false, getAllProcs()...)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
