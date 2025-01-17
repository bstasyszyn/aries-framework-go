/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package aries-agentd (Aries Agent Server) of aries-framework-go.
//
//
// Terms Of Service:
//
//
//     Schemes: http, https
//     Host: 127.0.0.1:8080
//     Version: 0.1.0
//     License: SPDX-License-Identifier: Apache-2.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"github.com/spf13/cobra"

	"github.com/hyperledger/aries-framework-go/cmd/aries-agentd/startcmd"
	"github.com/hyperledger/aries-framework-go/pkg/common/log"
)

// This is an application which starts Aries agent controller API on given port
func main() {
	rootCmd := &cobra.Command{
		Use: "aries-agentd",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	logger := log.New("aries-framework/agentd")

	startCmd, err := startcmd.Cmd(&startcmd.HTTPServer{})
	if err != nil {
		logger.Fatalf(err.Error())
	}
	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Fatalf("Failed to run aries-agentd: %s", err)
	}
}
