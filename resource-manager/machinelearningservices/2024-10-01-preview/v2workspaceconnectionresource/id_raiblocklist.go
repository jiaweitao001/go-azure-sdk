package v2workspaceconnectionresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RaiBlocklistId{})
}

var _ resourceids.ResourceId = &RaiBlocklistId{}

// RaiBlocklistId is a struct representing the Resource ID for a Rai Blocklist
type RaiBlocklistId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	ConnectionName    string
	RaiBlocklistName  string
}

// NewRaiBlocklistID returns a new RaiBlocklistId struct
func NewRaiBlocklistID(subscriptionId string, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string) RaiBlocklistId {
	return RaiBlocklistId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		ConnectionName:    connectionName,
		RaiBlocklistName:  raiBlocklistName,
	}
}

// ParseRaiBlocklistID parses 'input' into a RaiBlocklistId
func ParseRaiBlocklistID(input string) (*RaiBlocklistId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaiBlocklistId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaiBlocklistId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRaiBlocklistIDInsensitively parses 'input' case-insensitively into a RaiBlocklistId
// note: this method should only be used for API response data and not user input
func ParseRaiBlocklistIDInsensitively(input string) (*RaiBlocklistId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaiBlocklistId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaiBlocklistId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RaiBlocklistId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.ConnectionName, ok = input.Parsed["connectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "connectionName", input)
	}

	if id.RaiBlocklistName, ok = input.Parsed["raiBlocklistName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "raiBlocklistName", input)
	}

	return nil
}

// ValidateRaiBlocklistID checks that 'input' can be parsed as a Rai Blocklist ID
func ValidateRaiBlocklistID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRaiBlocklistID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rai Blocklist ID
func (id RaiBlocklistId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/connections/%s/raiBlocklists/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.ConnectionName, id.RaiBlocklistName)
}

// Segments returns a slice of Resource ID Segments which comprise this Rai Blocklist ID
func (id RaiBlocklistId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticConnections", "connections", "connections"),
		resourceids.UserSpecifiedSegment("connectionName", "connectionName"),
		resourceids.StaticSegment("staticRaiBlocklists", "raiBlocklists", "raiBlocklists"),
		resourceids.UserSpecifiedSegment("raiBlocklistName", "raiBlocklistName"),
	}
}

// String returns a human-readable description of this Rai Blocklist ID
func (id RaiBlocklistId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Connection Name: %q", id.ConnectionName),
		fmt.Sprintf("Rai Blocklist Name: %q", id.RaiBlocklistName),
	}
	return fmt.Sprintf("Rai Blocklist (%s)", strings.Join(components, "\n"))
}
