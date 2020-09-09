package godaddygo

import (
	"github.com/oze4/godaddygo/pkg/client"
	"github.com/oze4/godaddygo/pkg/endpoints"
)

// NewClient creates a new *default* client
// You do have the ability to create your own client
// Just a wrapper around `client.Default(...)` and is mostly for convenience
//  - If `prod` is true, we use the "production" GoDaddy API (https://api.godaddy.com)
//  - If it is false we use the "development" (OTE) GoDaddy API (https://api-ote.godaddy.com)
func NewClient(prod bool, key, secret string) client.Interface {
	return client.Default(key, secret, prod)
}

// Connect connects you to GoDaddy API endpoints
// Just a wrapper around `endpoints.Connect(...)`
// and is mostly for convenience
func Connect(godaddygoClient client.Interface) endpoints.Gateway {
	return endpoints.Connect(godaddygoClient)
}

// ConnectProduction is for convenience - create a new connection without 
// having to specify if "isproduction" or not
func ConnectProduction(apikey, apisecret string) endpoints.Gateway {
	return endpoints.ConnectProduction(apikey, apisecret)
}

// ConnectDevelopment is for convenience - create a new connection without 
// having to specify if "isproduction" or not
func ConnectDevelopment(apikey, apisecret string) endpoints.Gateway {
	return endpoints.ConnectDevelopment(apikey, apisecret)
}
