package cmd

import (
	"fmt"

	"github.com/cycloidio/terracognita/azurerm"
	"github.com/spf13/cobra"
)

var (
	azureResourcesCmd = &cobra.Command{
		Use:   "resources",
		Short: "List of all the azurerm supported resources",
		Run: func(cmd *cobra.Command, args []string) {
			for _, r := range azurerm.ResourceTypeStrings() {
				fmt.Println(r)
			}
		},
	}
)
