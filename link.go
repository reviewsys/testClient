package main

import (
	"github.com/reviewsys/testClient/cmd"

	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
)

func init() {
	// Add client generated commands to cobra's root cmd.
	cmd.RootCmd.AddCommand(pb.UserClientCommand)
	cmd.RootCmd.AddCommand(pb.ServiceClientCommand)
}
