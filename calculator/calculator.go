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

package calculator

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//calculator calculates the OOM score given an array of pids and orders them

type pair struct {
	Key int
	Val int
}

//pairList is a list of pairs that satisfys sort.Interface
type pairList []pair

//Len gets the length of the pairlist
func (pl pairList) Len() int {
	return len(pl)
}

//Less checks if the first param ,i, is less than the second param,j,
//Using > to get descending order
func (pl pairList) Less(i, j int) bool {
	return pl[i].Val > pl[j].Val
}

//Swap exchanges the positions of the selected pairList indices
func (pl pairList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

//SortPIDs takes in:
//num-> the number of PIDs to return
//best-> whether to return best or worst PID OOM scores,
//pids-> a list of PIDs to calculate the oom scores
func SortPIDs(num int, best bool, pids ...int) pairList {
	//if the requested number is larger than the passed pids readjust to the number
	//of PIDs
	if num > len(pids) {
		num = len(pids)
	}
	var list pairList
	for _, pid := range pids {
		score := GetOOMScore(pid)
		if score == -2 {
			num--
			continue
		}
		if score == -1 {
			num--
			fmt.Println(fmt.Sprintf("Calculating OOM score for %d failed", pid))
			continue
		}
		list = append(list, pair{
			Key: pid,
			Val: score,
		})
	}
	if len(list) == 0 {
		return list
	}
	//First items have the highest(worst) OOM score
	//Last items have the lowest(best) OOM score
	sort.Sort(list)
	if best {
		return list[len(list)-num:]
	}
	return list[:num]
}

//GetOOMScore gets the OOM score for a passed pid
//returns -2 if the pid does not have an accessible directory at /proc fs
//returns -1 if any other errors occur
func GetOOMScore(pid int) int {
	if err := os.Chdir(fmt.Sprintf("/proc/%d", pid)); err != nil {
		if os.IsPermission(err) {
			fmt.Println("Permission error when reading dir:", pid)
			return -1
		}
		if os.IsNotExist(err) {
			return -2 //ignore non existent pid dirs
		}
		fmt.Println(err)
		return -1
	}

	//check if the process has OOM_DISABLE set as defined in: Linux/include/uapi/linux/oom.h
	data, err := ioutil.ReadFile("oom_adj")
	if err != nil {

		if os.IsPermission(err) {
			fmt.Println("Permission error when reading file oom_adj")
		}
		fmt.Println(err)
		return -1
	}
	if string(data) == "-17" {
		return -17
	}

	//get the OOM score
	data, err = ioutil.ReadFile("oom_score")
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("Permission error when reading file oom_score")
		}
		fmt.Println(err)
		return -1
	}
	strData := strings.Trim(string(data), "\n")

	if val, err := strconv.Atoi(strData); err != nil {
		return -1
	} else {
		return val
	}

}
