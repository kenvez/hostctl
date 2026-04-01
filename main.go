package main

import (
	"fmt"
	"os"

	"github.com/kenvez/hostctl/internal/blocker"
	"github.com/kenvez/hostctl/internal/reader"
	"github.com/spf13/cobra"
)

func main() {
	if os.Getuid() != 0 {
		fmt.Println("error: run with sudo")
		os.Exit(1)
	}

	hosts, err := reader.ParseHosts()

	if err != nil {
		fmt.Println("error: cannot open /etc/hosts file")
	}

	rootCmd := &cobra.Command{Use: "hostctl"}
	blockCmd := &cobra.Command{
		Use:   "block [domain]",
		Short: "block a website",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := blocker.Block(hosts, args[0]); err != nil {
				return err
			}

			return nil
		},
	}
	unblockCmd := &cobra.Command{
		Use:   "unblock [domain]",
		Short: "unblock a website",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := blocker.Unblock(hosts, args[0]); err != nil {
				return err
			}

			return nil
		},
	}
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list current blocked websites",
		RunE: func(cmd *cobra.Command, args []string) error {
			for k := range hosts.Entries {
				fmt.Println(k)
			}

			return nil
		},
	}

	rootCmd.AddCommand(blockCmd)
	rootCmd.AddCommand(unblockCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.Execute()
}
