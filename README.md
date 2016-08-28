# doom
[![Build Status](https://goreportcard.com/badge/github.com/ziscky/zist)](https://goreportcard.com/report/github.com/ziscky/mock-pesa)
[![Build Status](https://travis-ci.org/ziscky/doom.svg?branch=master)](https://travis-ci.org/ziscky/doom)

A simple program to find PIDs with the best/worst OOM scores

### Precompiled binaries

Precompiled binaries for released versions are available in the
[*releases* section](https://github.com/ziscky/mock-pesa/releases)
of the GitHub repository. Supported OS/Arch:

 1. Linux amd64
 
### Getting Started
` Note: Works by reading the proc pseudo fs so you may need to run the cmds as root `
PID with best OOM score: ` doom best `
PID with worst OOM score: ` doom worst `
Top 10 worst OOM scores: ` doom worst 10`
Top 10 best OOM scores: ` doom best 10`

### Building From Source
`sudo make install` 

### Contiributing
I'm very open to PRs.  

 - Fork
 - Create Branch
 - Do magic
 - Initiate PR

