[![Build Status](https://travis-ci.org/bharathshetty4/InstaMore.svg?branch=master)](https://travis-ci.org/bharathshetty4/InstaMore)
[![Go Report Card](https://goreportcard.com/badge/github.com/bharathshetty4/InstaMore)](https://goreportcard.com/report/github.com/bharathshetty4/InstaMore)
[![GoDoc](https://godoc.org/github.com/bharathshetty4/InstaMore?status.svg)](https://godoc.org/github.com/bharathshetty4/InstaMore)
[![Coverage Status](https://coveralls.io/repos/github/bharathshetty4/InstaMore/badge.svg?branch=master)](https://coveralls.io/github/bharathshetty4/InstaMore?branch=master)

# List Of Features
*Scroll down for the CLI and API reference*
1. Like every photos of a given Instagram Profile

# Some Good Habits to Follow in a Project
1. Use Package/Repo Management tool for a Project. In this project we are
Using [Glide](https://github.com/Masterminds/glide) which is taking care of it.
You can find the *glide.yaml* of this project to know more.

2. Maintain a file to describe and document RESTful APIs in a Project. We are
Using [Swagger](https://github.com/go-swagger/go-swagger) Tool, Which can also
create the CLI for us using the *swagger.yaml*.
You can refer the *swagger.yaml* of this project to know more.

3. There are many libraries available in go, Use it whenever possible.
To build a CLI, In this project we are
using [cli](https://github.com/urfave/cli) library,
[viper](https://github.com/spf13/viper) for config file management and the
lists will go on as the feature increases.

4. Make use of free tools to keep your code clean and unbreakable whenever
possible to maintain the sanity of the project,
We are using [Travis CI](https://travis-ci.org/bharathshetty4/InstaMore) in
this project, which will be running on every PR's created. Many more tools are
available in this opensource world and try to make use of it.

# CLI Reference
### Pre-Requisites:
**Login to Instagram**

InstaMore login --username <your_instagram_username> --password password

### Features
1. Like every photos of a given Instagram Profile


# API Reference
### Pre-Requisites:
**Start the API Server**

InstaMore server start --port 5555

**Login to Instagram**

*Put login api reference*

### Features
1. Like every photos of a given Instagram Profile


