Library webhooks with Gitee support 12345
================
![Project status](https://img.shields.io/badge/version-6.0.2-green.svg)
[![Test](https://github.com/go-playground/webhooks/workflows/Test/badge.svg?branch=master)](https://github.com/go-playground/webhooks/actions)
[![Coverage Status](https://coveralls.io/repos/go-playground/webhooks/badge.svg?branch=master&service=github)](https://coveralls.io/github/go-playground/webhooks?branch=master)
[![Go Report Card](https://goreportcard.com/badge/robekeane/webhooks)](https://goreportcard.com/report/robekeane/webhooks)
[![GoDoc](https://godoc.org/github.com/robekeane/webhooks/v6?status.svg)](https://godoc.org/github.com/robekeane/webhooks/v6)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Thanks for go-playground/webhooks, this is a fork project to add gitee support.

Installation
------------

Use go get.

```shell
go get -u github.com/robekeane/webhooks/v6
```

Then import the package into your own code.

	import "github.com/robekeane/webhooks/v6"

Usage and Documentation
------

Please see http://godoc.org/github.com/go-playground/webhooks/v6 for detailed usage docs.

##### Examples:
```go
package main

import (
	"fmt"

	"net/http"

	"github.com/robekeane/webhooks/v6/gitee"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := gitee.New(gitee.Options.Secret("test12345"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, gitee.IssuesEvents)
		if err != nil {
			if err == gitee.ErrEventNotFound {
			}
		}
		switch payload.(type) {

		case gitee.IssueEventPayload:
			release := payload.(gitee.IssueEventPayload)
			fmt.Printf("%+v", release.Issue.Title)
		}
	})
	http.ListenAndServe(":3000", nil)
}


```

License
------
Distributed under MIT License, please see license file in code for more details.
