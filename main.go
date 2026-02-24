package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/itrs-group/terraform-provider-itrs-uptrends/provider"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()
	// Create cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Serve with fully-qualified address:
	err := providerserver.Serve(ctx, provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/ITRS-Group/itrs-uptrends",
		Debug:   debug,
	})
	if err != nil {
		log.Fatalf("Failed to serve the provider: %v", err)
	}
}
