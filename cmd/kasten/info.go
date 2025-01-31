// Copyright (c) arkade author(s) 2021. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// kasten contains a suite of Sponsored Apps for arkade
package kasten

import (
	"fmt"

	"github.com/spf13/cobra"
)

func MakeInfo() *cobra.Command {

	command := &cobra.Command{
		Use:          "info",
		Short:        "Info for an app",
		Long:         `Info for an app`,
		Example:      `  arkade kasten info [APP]`,
		SilenceUsage: true,
	}

	command.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("give an app as an argument")
		}

		info := "None found."

		if args[0] == "k10" {
			info = k10InfoMsg
		}

		fmt.Printf("Info for your app: %s\n\n%s\n\n", args[0], info)

		return nil
	}

	return command
}
