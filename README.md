# [godaddygo](https://oze4.github.io/godaddygo/)

[Check us out on `pkg.go.dev`](https://pkg.go.dev/github.com/oze4/godaddygo?tab=doc) \*seems to be a little behind a lot of the time

---

# Table of Contents

- [Intro](#intro)
- [Installation](#installation)
- [Getting Started](#getting-started)
  - [API Structure](#our-api-structure)
- [Usage Details](#usage)
  - [Using Default Client](#default-client)
  - [Providing Custom Client](#custom-client)
- [Examples](https://github.com/oze4/godaddygo/tree/master/examples)
- [Features](#features)

---

## Intro

 - See [here](#features) for more on which features this package currently supports
 - Whenever we reference endpoints, [this is what we are referring to](https://developer.godaddy.com/doc)

<br /> 

<small>Pull requests welcome! We would like to eventually support each GoDaddy API endpoint, not just domain/DNS related tasks</small>

## Installation

 - `go get -u github.com/oze4/godaddygo`
 - See [here](https://developer.godaddy.com/) for more on how to obtain an API key and API secret from GoDaddy (click 'API Keys')

## Getting Started

Bare minimum what you need to get started (aka how you will typically use this package):

```golang
package main

import (
	"github.com/oze4/godaddygo"
)

func main() {
	prodKey := "-"
	prodSecret := "-"

	// Target version 1 of the production API
	api := godaddygo.ConnectProduction(prodKey, prodSecret)
	prodv1 := api.V1()

	// Now have access to all GoDaddy production V1
	// API endpoints (via `prodv1`)

	// eg: prodv1.Domain("x.com").Records().GetAll()
	//     prodv1.Domain("x.com").Records().Add(someDNSRecord)
	//     prodv1.Domain("x.com").GetDetails()
	//     prodv1.GetDomainAvailability("<some-domain-you-want-to-buy.com>")
	// etc...
}
```

### Our API Structure

Consider the following endpoint, which allows you to add a DNS record to a domain.

![screenshot_from_godaddy_docs](https://i.imgur.com/tN2IveY.png)

Programmatically, this would look like:

```golang
api.V1().Domain("dom.com").Records().Add(newDNSRecord)
```

## Usage 

GoDaddy's API currently has 2 versions, `v1` and `v2`. Within the `godaddygo` package we provide 2 helper functions, one for each version. These helper functions  simply "wrap" around our "core", which means you have the ability to create yor own client(s).

We recommend using one of the following:

 - `godaddygo.ConnectProduction(key, secret)`
 - `godaddygo.ConnectDevelopment(key, secret)`

### Default Client

 - If you would like, you may create a default client "manually", then pass it to `endpoints.Connect(<default_client_here>)`. 
 - At a high level, the `godaddygo` package is essentially just a wrapper for `endpoints` and `client`. 
 - All we do is "wrap" those packages within `godaddygo`, as shown below

```go
package main

import (
	"github.com/oze4/godaddygo/pkg/client"
	"github.com/oze4/godaddygo/pkg/endpoints"
)

func main() {
	// Options for client
	key := "api_key"
	secret := "api_secret"
	// See here for more on GoDaddy production vs development (OTE) API's
	// https://developer.godaddy.com/getstarted
	targetProductionAPI := true

	// Create default client
	client := client.Default(key, secret, targetProductionAPI)

	// Connect our client to endpoints
	api := endpoints.Connect(client)

	//
	// Use `api` here...
	//
	// ...for example:
	prodv1 := api.V1()
	// Target specific domain
	mydomain := prodv1.Domain("mydomain.com")
	// Get all DNS records for target domain
	records, err := mydomain.Records().GetAll()

	// ...
}
```

### Custom Client

 - If you wish to use your own client instead of the default client, this is how you would do so.

```go
package main

import (
	"github.com/oze4/godaddygo/pkg/endpoints"
)

func main() {
	myCustomClient := &myClient{
		key:    "api_key",
		secret: "api_secret",
		isprod: true,
	}

	api := endpoints.Connect(myCustomClient)

	//
	// Use `api` here!
	//
	// ...
}

// myClient is your custom client
// As long as your client satisfies `client.Interface`
// You can use it to connect to the `endpoints` Gateway
type myClient struct {
	key    string
	secret string
	isprod bool
	// ...your custom stuff
}

func (c *myClient) APIKey() string {
	return c.key
}

func (c *myClient) APISecret() string {
	return c.secret
}

func (c *myClient) IsProduction() string {
	return c.isprod
}

```

## Features

Please see [here](https://developer.godaddy.com/doc) for more information on GoDaddy API endpoints

- [ ] Abuse
- [ ] Aftermarket
- [ ] Agreements
- [ ] Certificates
- [ ] Countries
- Domains
  - [x] Check domain availability
  - [x] Get all DNS records
  - [x] Get all DNS records of specific type
  - [x] Get specific DNS record
  - [x] Set DNS record(s)
  - [x] Add/create DNS record(s)
  - [x] Purchase domain
  - [x] Purchase privacy for domain
  - [x] Remove privacy for domain
- [ ] Orders
- [ ] Shoppers
- [ ] Subscriptions

<br />
<br />
<br />

---

[mattoestreich.com](https://mattoestreich.com)

---
