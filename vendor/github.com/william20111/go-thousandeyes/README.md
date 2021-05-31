[![GoDoc](https://godoc.org/github.com/william20111/go-thousandeyes?status.svg)](http://godoc.org/github.com/william20111/go-thousandeyes) [![Go Report Card](https://goreportcard.com/badge/github.com/william20111/go-thousandeyes)](https://goreportcard.com/report/github.com/william20111/go-thousandeyes) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE) [![codecov](https://codecov.io/gh/william20111/go-thousandeyes/branch/master/graph/badge.svg)](https://codecov.io/gh/william20111/go-thousandeyes)
[![Gitter](https://badges.gitter.im/go-thousandeyes/community.svg)](https://gitter.im/go-thousandeyes/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
# go-thousandeyes

go-thousandeyes is a [go](https://golang.org/) client library for the [Thousandeyes v6 API](https://developer.thousandeyes.com/v6). This repo
also contains a supporting CLI tool tectl that uses the client lib.

## Installation

First, download the source code
```cli
go get github.com/william20111/go-thousandeyes
```

## Usage

### Using `tectl`

```shell script
tectl get agents           
AGENT NAME                              AGENTID ENABLED LOCATION                        IPADDRESSES                                                                                                                                                                                                     
Dallas, TX (Trial)                      4492    0       Dallas Area                     104.130.154.136,104.130.156.108,104.130.141.203,104.130.155.161                                                                                                                                                        
Chicago, IL (Trial)                     4495    0       Chicago Area                    23.253.167.114,23.253.41.151,23.253.46.129,23.253.47.96                                                                                                                                                                
Ashburn, VA (Trial)                     4497    0       Ashburn Area                    23.253.149.202,162.242.253.124,23.253.149.138,162.242.252.22                                                                                                                                                           
Sydney, Australia (Trial)               4500    0       New South Wales, Australia      119.9.24.178,119.9.24.218,119.9.24.167,119.9.24.176,2401:1801:7801:101:ec6e:d653:b713:68cb 
```

### Using client library

```go
package main

import (
	"fmt"
	"github.com/william20111/go-thousandeyes"
)

func main() {
	client := thousandeyes.NewClient("foobar")
	agents, err := client.GetAgents()
	if err != nil {
		panic(err)
	}
	for _, a := range *agents {
		fmt.Println(a.AgentName)
	}
}
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## License

[Apache 2](http://www.apache.org/licenses/LICENSE-2.0)

## Mentions

This is my first golang library and i have heavily leaned on [go-pagerduty](https://github.com/PagerDuty/go-pagerduty)
library for inspiration. Its a library I use and I have implemented a similar client interface based on their HttpClient
wrapper. So thanks for the great oss lib!
