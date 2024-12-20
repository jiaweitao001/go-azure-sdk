package discoveredassetendpointprofiles

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DiscoveredAssetEndpointProfileId{})
}

var _ resourceids.ResourceId = &DiscoveredAssetEndpointProfileId{}

// DiscoveredAssetEndpointProfileId is a struct representing the Resource ID for a Discovered Asset Endpoint Profile
type DiscoveredAssetEndpointProfileId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	DiscoveredAssetEndpointProfileName string
}

// NewDiscoveredAssetEndpointProfileID returns a new DiscoveredAssetEndpointProfileId struct
func NewDiscoveredAssetEndpointProfileID(subscriptionId string, resourceGroupName string, discoveredAssetEndpointProfileName string) DiscoveredAssetEndpointProfileId {
	return DiscoveredAssetEndpointProfileId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		DiscoveredAssetEndpointProfileName: discoveredAssetEndpointProfileName,
	}
}

// ParseDiscoveredAssetEndpointProfileID parses 'input' into a DiscoveredAssetEndpointProfileId
func ParseDiscoveredAssetEndpointProfileID(input string) (*DiscoveredAssetEndpointProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiscoveredAssetEndpointProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiscoveredAssetEndpointProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDiscoveredAssetEndpointProfileIDInsensitively parses 'input' case-insensitively into a DiscoveredAssetEndpointProfileId
// note: this method should only be used for API response data and not user input
func ParseDiscoveredAssetEndpointProfileIDInsensitively(input string) (*DiscoveredAssetEndpointProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiscoveredAssetEndpointProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiscoveredAssetEndpointProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DiscoveredAssetEndpointProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DiscoveredAssetEndpointProfileName, ok = input.Parsed["discoveredAssetEndpointProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "discoveredAssetEndpointProfileName", input)
	}

	return nil
}

// ValidateDiscoveredAssetEndpointProfileID checks that 'input' can be parsed as a Discovered Asset Endpoint Profile ID
func ValidateDiscoveredAssetEndpointProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiscoveredAssetEndpointProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Discovered Asset Endpoint Profile ID
func (id DiscoveredAssetEndpointProfileId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DeviceRegistry/discoveredAssetEndpointProfiles/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DiscoveredAssetEndpointProfileName)
}

// Segments returns a slice of Resource ID Segments which comprise this Discovered Asset Endpoint Profile ID
func (id DiscoveredAssetEndpointProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDeviceRegistry", "Microsoft.DeviceRegistry", "Microsoft.DeviceRegistry"),
		resourceids.StaticSegment("staticDiscoveredAssetEndpointProfiles", "discoveredAssetEndpointProfiles", "discoveredAssetEndpointProfiles"),
		resourceids.UserSpecifiedSegment("discoveredAssetEndpointProfileName", "discoveredAssetEndpointProfileName"),
	}
}

// String returns a human-readable description of this Discovered Asset Endpoint Profile ID
func (id DiscoveredAssetEndpointProfileId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Discovered Asset Endpoint Profile Name: %q", id.DiscoveredAssetEndpointProfileName),
	}
	return fmt.Sprintf("Discovered Asset Endpoint Profile (%s)", strings.Join(components, "\n"))
}
