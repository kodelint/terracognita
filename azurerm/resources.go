package azurerm

import (
	"context"

	"github.com/pkg/errors"

	"github.com/cycloidio/terracognita/provider"
	"github.com/cycloidio/terracognita/tag"
)

// ResourceType is the type used to define all the Resources
// from the Provider
type ResourceType int

//go:generate enumer -type ResourceType -addprefix azurerm_ -transform snake -linecomment
const (
	ResourceGroup  ResourceType = iota
	VirtualMachine ResourceType = iota
	VirtualNetwork ResourceType = iota
)

type rtFn func(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error)

var (
	resources = map[ResourceType]rtFn{
		ResourceGroup:  resourceGroup,
		VirtualMachine: virtualMachines,
		VirtualNetwork: virtualNetworks,
	}
)

func resourceGroup(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	resourceGroup := a.azurer.GetResourceGroup()
	resources := make([]provider.Resource, 0)
	r := provider.NewResource(*resourceGroup.ID, resourceType, a)
	resources = append(resources, r)
	return resources, nil
}

func virtualMachines(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	virtualMachines, err := a.azurer.ListVirtualMachines(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list virtual machines from reader")
	}
	resources := make([]provider.Resource, 0)
	for _, virtualMachine := range virtualMachines {
		r := provider.NewResource(*virtualMachine.ID, resourceType, a)
		resources = append(resources, r)
	}
	return resources, nil
}

func virtualNetworks(ctx context.Context, a *azurerm, resourceType string, tags []tag.Tag) ([]provider.Resource, error) {
	virtualNetworks, err := a.azurer.ListVirtualNetworks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list virtual machines from reader")
	}
	resources := make([]provider.Resource, 0)
	for _, virtualNetwork := range virtualNetworks {
		r := provider.NewResource(*virtualNetwork.ID, resourceType, a)
		resources = append(resources, r)
	}
	return resources, nil
}