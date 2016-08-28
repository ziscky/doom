# doom
[![Build Status](https://goreportcard.com/badge/github.com/ziscky/doom)](https://goreportcard.com/report/github.com/ziscky/doom)
[![Build Status](https://travis-ci.org/ziscky/doom.svg?branch=master)](https://travis-ci.org/ziscky/doom)

A simple program to find PIDs with the best/worst OOM scores. Also gives policy information on how the system
would behave in the case of an OOM Error e.g(Kernel Panic,Kill overcommiting process e.t.c). Support for
oom_killer_controller patch (https://lkml.org/lkml/2009/1/29/220), contributed by Nikanth Karthikesan is WIP.

### Installation: Precompiled binaries and Packages

Precompiled binaries for released versions are available in the
[*releases* section](https://github.com/ziscky/doom/releases)
of the GitHub repository. Supported OS/Arch:

 1. Linux amd64 binary
 2. RPM
 3. Debian packages for Debian 7, Ubuntu 14.04 and Ubuntu 16.04
 
### Getting Started
```
Works by reading the /proc pseudo fs so you may need to run the cmds as root
Also make sure the package: procps-ng is installed.(provides ps|pgrep etc)  

Help: doom
Rank all PIDs by best OOM score:  doom best   
Rank all PIDs by worst OOM score:  doom worst   
PID with worst OOM score:  doom next
Top 10 worst OOM scores:  doom worst 10  
Top 10 best OOM scores:  doom best 10 
Inspect a particular process by name(in this case chrome):  doom inspect chrome   
Inspect a particular process by PID:  doom inspect 23456   
Show your system's relevant OOM behaviour:  doom policy  
```


### Building From Source
```
make 
sudo make install
```

### Installation

#### CentOS/RHEL
Execute as root:  
`sudo rpm --import https://rpm.packager.io/key`  

Add the following to /etc/yum.repos.d/doom.repo:   
```
[doom]
name=Repository for ziscky/doom application.
baseurl=https://rpm.packager.io/gh/ziscky/doom/centos6/master
enabled=1
```

`sudo yum/dnf install doom`

#### Debian 7
Add the following to /etc/apt/sources.list.d/doom.list:  
```
wget -qO - https://deb.packager.io/key | sudo apt-key add - 
echo "deb https://deb.packager.io/gh/ziscky/doom wheezy master" | sudo tee /etc/apt/sources.list.d/doom.list  
sudo apt-get update  
sudo apt-get install doom*=0.0.0-1472392791.57d938e.wheezy
```

#### Ubuntu 14.04
```
wget -qO - https://deb.packager.io/key | sudo apt-key add -  
echo "deb https://deb.packager.io/gh/ziscky/doom trusty master" | sudo tee /etc/apt/sources.list.d/doom.list  
sudo apt-get update  
sudo apt-get install doom*=0.0.0-1472392791.57d938e.trusty
```

#### Ubuntu 16.04
```
wget -qO - https://deb.packager.io/key | sudo apt-key add -  
echo "deb https://deb.packager.io/gh/ziscky/doom xenial master" | sudo tee /etc/apt/sources.list.d/doom.list  
sudo apt-get update  
sudo apt-get install doom*=0.0.0-1472392791.57d938e.xenial
```  



### Contiributing
I'm very open to PRs.  

 - Fork
 - Create Branch
 - Do magic
 - Initiate PR

