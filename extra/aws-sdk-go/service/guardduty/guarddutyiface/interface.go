// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package guarddutyiface provides an interface to enable mocking the Amazon GuardDuty service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package guarddutyiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/guardduty"
)

// GuardDutyAPI provides an interface to enable mocking the
// guardduty.GuardDuty service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon GuardDuty.
//    func myFunc(svc guarddutyiface.GuardDutyAPI) bool {
//        // Make svc.AcceptAdministratorInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := guardduty.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGuardDutyClient struct {
//        guarddutyiface.GuardDutyAPI
//    }
//    func (m *mockGuardDutyClient) AcceptAdministratorInvitation(input *guardduty.AcceptAdministratorInvitationInput) (*guardduty.AcceptAdministratorInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGuardDutyClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type GuardDutyAPI interface {
	AcceptAdministratorInvitation(*guardduty.AcceptAdministratorInvitationInput) (*guardduty.AcceptAdministratorInvitationOutput, error)
	AcceptAdministratorInvitationWithContext(aws.Context, *guardduty.AcceptAdministratorInvitationInput, ...request.Option) (*guardduty.AcceptAdministratorInvitationOutput, error)
	AcceptAdministratorInvitationRequest(*guardduty.AcceptAdministratorInvitationInput) (*request.Request, *guardduty.AcceptAdministratorInvitationOutput)

	AcceptInvitation(*guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error)
	AcceptInvitationWithContext(aws.Context, *guardduty.AcceptInvitationInput, ...request.Option) (*guardduty.AcceptInvitationOutput, error)
	AcceptInvitationRequest(*guardduty.AcceptInvitationInput) (*request.Request, *guardduty.AcceptInvitationOutput)

	ArchiveFindings(*guardduty.ArchiveFindingsInput) (*guardduty.ArchiveFindingsOutput, error)
	ArchiveFindingsWithContext(aws.Context, *guardduty.ArchiveFindingsInput, ...request.Option) (*guardduty.ArchiveFindingsOutput, error)
	ArchiveFindingsRequest(*guardduty.ArchiveFindingsInput) (*request.Request, *guardduty.ArchiveFindingsOutput)

	CreateDetector(*guardduty.CreateDetectorInput) (*guardduty.CreateDetectorOutput, error)
	CreateDetectorWithContext(aws.Context, *guardduty.CreateDetectorInput, ...request.Option) (*guardduty.CreateDetectorOutput, error)
	CreateDetectorRequest(*guardduty.CreateDetectorInput) (*request.Request, *guardduty.CreateDetectorOutput)

	CreateFilter(*guardduty.CreateFilterInput) (*guardduty.CreateFilterOutput, error)
	CreateFilterWithContext(aws.Context, *guardduty.CreateFilterInput, ...request.Option) (*guardduty.CreateFilterOutput, error)
	CreateFilterRequest(*guardduty.CreateFilterInput) (*request.Request, *guardduty.CreateFilterOutput)

	CreateIPSet(*guardduty.CreateIPSetInput) (*guardduty.CreateIPSetOutput, error)
	CreateIPSetWithContext(aws.Context, *guardduty.CreateIPSetInput, ...request.Option) (*guardduty.CreateIPSetOutput, error)
	CreateIPSetRequest(*guardduty.CreateIPSetInput) (*request.Request, *guardduty.CreateIPSetOutput)

	CreateMembers(*guardduty.CreateMembersInput) (*guardduty.CreateMembersOutput, error)
	CreateMembersWithContext(aws.Context, *guardduty.CreateMembersInput, ...request.Option) (*guardduty.CreateMembersOutput, error)
	CreateMembersRequest(*guardduty.CreateMembersInput) (*request.Request, *guardduty.CreateMembersOutput)

	CreatePublishingDestination(*guardduty.CreatePublishingDestinationInput) (*guardduty.CreatePublishingDestinationOutput, error)
	CreatePublishingDestinationWithContext(aws.Context, *guardduty.CreatePublishingDestinationInput, ...request.Option) (*guardduty.CreatePublishingDestinationOutput, error)
	CreatePublishingDestinationRequest(*guardduty.CreatePublishingDestinationInput) (*request.Request, *guardduty.CreatePublishingDestinationOutput)

	CreateSampleFindings(*guardduty.CreateSampleFindingsInput) (*guardduty.CreateSampleFindingsOutput, error)
	CreateSampleFindingsWithContext(aws.Context, *guardduty.CreateSampleFindingsInput, ...request.Option) (*guardduty.CreateSampleFindingsOutput, error)
	CreateSampleFindingsRequest(*guardduty.CreateSampleFindingsInput) (*request.Request, *guardduty.CreateSampleFindingsOutput)

	CreateThreatIntelSet(*guardduty.CreateThreatIntelSetInput) (*guardduty.CreateThreatIntelSetOutput, error)
	CreateThreatIntelSetWithContext(aws.Context, *guardduty.CreateThreatIntelSetInput, ...request.Option) (*guardduty.CreateThreatIntelSetOutput, error)
	CreateThreatIntelSetRequest(*guardduty.CreateThreatIntelSetInput) (*request.Request, *guardduty.CreateThreatIntelSetOutput)

	DeclineInvitations(*guardduty.DeclineInvitationsInput) (*guardduty.DeclineInvitationsOutput, error)
	DeclineInvitationsWithContext(aws.Context, *guardduty.DeclineInvitationsInput, ...request.Option) (*guardduty.DeclineInvitationsOutput, error)
	DeclineInvitationsRequest(*guardduty.DeclineInvitationsInput) (*request.Request, *guardduty.DeclineInvitationsOutput)

	DeleteDetector(*guardduty.DeleteDetectorInput) (*guardduty.DeleteDetectorOutput, error)
	DeleteDetectorWithContext(aws.Context, *guardduty.DeleteDetectorInput, ...request.Option) (*guardduty.DeleteDetectorOutput, error)
	DeleteDetectorRequest(*guardduty.DeleteDetectorInput) (*request.Request, *guardduty.DeleteDetectorOutput)

	DeleteFilter(*guardduty.DeleteFilterInput) (*guardduty.DeleteFilterOutput, error)
	DeleteFilterWithContext(aws.Context, *guardduty.DeleteFilterInput, ...request.Option) (*guardduty.DeleteFilterOutput, error)
	DeleteFilterRequest(*guardduty.DeleteFilterInput) (*request.Request, *guardduty.DeleteFilterOutput)

	DeleteIPSet(*guardduty.DeleteIPSetInput) (*guardduty.DeleteIPSetOutput, error)
	DeleteIPSetWithContext(aws.Context, *guardduty.DeleteIPSetInput, ...request.Option) (*guardduty.DeleteIPSetOutput, error)
	DeleteIPSetRequest(*guardduty.DeleteIPSetInput) (*request.Request, *guardduty.DeleteIPSetOutput)

	DeleteInvitations(*guardduty.DeleteInvitationsInput) (*guardduty.DeleteInvitationsOutput, error)
	DeleteInvitationsWithContext(aws.Context, *guardduty.DeleteInvitationsInput, ...request.Option) (*guardduty.DeleteInvitationsOutput, error)
	DeleteInvitationsRequest(*guardduty.DeleteInvitationsInput) (*request.Request, *guardduty.DeleteInvitationsOutput)

	DeleteMembers(*guardduty.DeleteMembersInput) (*guardduty.DeleteMembersOutput, error)
	DeleteMembersWithContext(aws.Context, *guardduty.DeleteMembersInput, ...request.Option) (*guardduty.DeleteMembersOutput, error)
	DeleteMembersRequest(*guardduty.DeleteMembersInput) (*request.Request, *guardduty.DeleteMembersOutput)

	DeletePublishingDestination(*guardduty.DeletePublishingDestinationInput) (*guardduty.DeletePublishingDestinationOutput, error)
	DeletePublishingDestinationWithContext(aws.Context, *guardduty.DeletePublishingDestinationInput, ...request.Option) (*guardduty.DeletePublishingDestinationOutput, error)
	DeletePublishingDestinationRequest(*guardduty.DeletePublishingDestinationInput) (*request.Request, *guardduty.DeletePublishingDestinationOutput)

	DeleteThreatIntelSet(*guardduty.DeleteThreatIntelSetInput) (*guardduty.DeleteThreatIntelSetOutput, error)
	DeleteThreatIntelSetWithContext(aws.Context, *guardduty.DeleteThreatIntelSetInput, ...request.Option) (*guardduty.DeleteThreatIntelSetOutput, error)
	DeleteThreatIntelSetRequest(*guardduty.DeleteThreatIntelSetInput) (*request.Request, *guardduty.DeleteThreatIntelSetOutput)

	DescribeMalwareScans(*guardduty.DescribeMalwareScansInput) (*guardduty.DescribeMalwareScansOutput, error)
	DescribeMalwareScansWithContext(aws.Context, *guardduty.DescribeMalwareScansInput, ...request.Option) (*guardduty.DescribeMalwareScansOutput, error)
	DescribeMalwareScansRequest(*guardduty.DescribeMalwareScansInput) (*request.Request, *guardduty.DescribeMalwareScansOutput)

	DescribeMalwareScansPages(*guardduty.DescribeMalwareScansInput, func(*guardduty.DescribeMalwareScansOutput, bool) bool) error
	DescribeMalwareScansPagesWithContext(aws.Context, *guardduty.DescribeMalwareScansInput, func(*guardduty.DescribeMalwareScansOutput, bool) bool, ...request.Option) error

	DescribeOrganizationConfiguration(*guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationWithContext(aws.Context, *guardduty.DescribeOrganizationConfigurationInput, ...request.Option) (*guardduty.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationRequest(*guardduty.DescribeOrganizationConfigurationInput) (*request.Request, *guardduty.DescribeOrganizationConfigurationOutput)

	DescribeOrganizationConfigurationPages(*guardduty.DescribeOrganizationConfigurationInput, func(*guardduty.DescribeOrganizationConfigurationOutput, bool) bool) error
	DescribeOrganizationConfigurationPagesWithContext(aws.Context, *guardduty.DescribeOrganizationConfigurationInput, func(*guardduty.DescribeOrganizationConfigurationOutput, bool) bool, ...request.Option) error

	DescribePublishingDestination(*guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error)
	DescribePublishingDestinationWithContext(aws.Context, *guardduty.DescribePublishingDestinationInput, ...request.Option) (*guardduty.DescribePublishingDestinationOutput, error)
	DescribePublishingDestinationRequest(*guardduty.DescribePublishingDestinationInput) (*request.Request, *guardduty.DescribePublishingDestinationOutput)

	DisableOrganizationAdminAccount(*guardduty.DisableOrganizationAdminAccountInput) (*guardduty.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountWithContext(aws.Context, *guardduty.DisableOrganizationAdminAccountInput, ...request.Option) (*guardduty.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountRequest(*guardduty.DisableOrganizationAdminAccountInput) (*request.Request, *guardduty.DisableOrganizationAdminAccountOutput)

	DisassociateFromAdministratorAccount(*guardduty.DisassociateFromAdministratorAccountInput) (*guardduty.DisassociateFromAdministratorAccountOutput, error)
	DisassociateFromAdministratorAccountWithContext(aws.Context, *guardduty.DisassociateFromAdministratorAccountInput, ...request.Option) (*guardduty.DisassociateFromAdministratorAccountOutput, error)
	DisassociateFromAdministratorAccountRequest(*guardduty.DisassociateFromAdministratorAccountInput) (*request.Request, *guardduty.DisassociateFromAdministratorAccountOutput)

	DisassociateFromMasterAccount(*guardduty.DisassociateFromMasterAccountInput) (*guardduty.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountWithContext(aws.Context, *guardduty.DisassociateFromMasterAccountInput, ...request.Option) (*guardduty.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountRequest(*guardduty.DisassociateFromMasterAccountInput) (*request.Request, *guardduty.DisassociateFromMasterAccountOutput)

	DisassociateMembers(*guardduty.DisassociateMembersInput) (*guardduty.DisassociateMembersOutput, error)
	DisassociateMembersWithContext(aws.Context, *guardduty.DisassociateMembersInput, ...request.Option) (*guardduty.DisassociateMembersOutput, error)
	DisassociateMembersRequest(*guardduty.DisassociateMembersInput) (*request.Request, *guardduty.DisassociateMembersOutput)

	EnableOrganizationAdminAccount(*guardduty.EnableOrganizationAdminAccountInput) (*guardduty.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountWithContext(aws.Context, *guardduty.EnableOrganizationAdminAccountInput, ...request.Option) (*guardduty.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountRequest(*guardduty.EnableOrganizationAdminAccountInput) (*request.Request, *guardduty.EnableOrganizationAdminAccountOutput)

	GetAdministratorAccount(*guardduty.GetAdministratorAccountInput) (*guardduty.GetAdministratorAccountOutput, error)
	GetAdministratorAccountWithContext(aws.Context, *guardduty.GetAdministratorAccountInput, ...request.Option) (*guardduty.GetAdministratorAccountOutput, error)
	GetAdministratorAccountRequest(*guardduty.GetAdministratorAccountInput) (*request.Request, *guardduty.GetAdministratorAccountOutput)

	GetCoverageStatistics(*guardduty.GetCoverageStatisticsInput) (*guardduty.GetCoverageStatisticsOutput, error)
	GetCoverageStatisticsWithContext(aws.Context, *guardduty.GetCoverageStatisticsInput, ...request.Option) (*guardduty.GetCoverageStatisticsOutput, error)
	GetCoverageStatisticsRequest(*guardduty.GetCoverageStatisticsInput) (*request.Request, *guardduty.GetCoverageStatisticsOutput)

	GetDetector(*guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error)
	GetDetectorWithContext(aws.Context, *guardduty.GetDetectorInput, ...request.Option) (*guardduty.GetDetectorOutput, error)
	GetDetectorRequest(*guardduty.GetDetectorInput) (*request.Request, *guardduty.GetDetectorOutput)

	GetFilter(*guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error)
	GetFilterWithContext(aws.Context, *guardduty.GetFilterInput, ...request.Option) (*guardduty.GetFilterOutput, error)
	GetFilterRequest(*guardduty.GetFilterInput) (*request.Request, *guardduty.GetFilterOutput)

	GetFindings(*guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error)
	GetFindingsWithContext(aws.Context, *guardduty.GetFindingsInput, ...request.Option) (*guardduty.GetFindingsOutput, error)
	GetFindingsRequest(*guardduty.GetFindingsInput) (*request.Request, *guardduty.GetFindingsOutput)

	GetFindingsStatistics(*guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error)
	GetFindingsStatisticsWithContext(aws.Context, *guardduty.GetFindingsStatisticsInput, ...request.Option) (*guardduty.GetFindingsStatisticsOutput, error)
	GetFindingsStatisticsRequest(*guardduty.GetFindingsStatisticsInput) (*request.Request, *guardduty.GetFindingsStatisticsOutput)

	GetIPSet(*guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error)
	GetIPSetWithContext(aws.Context, *guardduty.GetIPSetInput, ...request.Option) (*guardduty.GetIPSetOutput, error)
	GetIPSetRequest(*guardduty.GetIPSetInput) (*request.Request, *guardduty.GetIPSetOutput)

	GetInvitationsCount(*guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error)
	GetInvitationsCountWithContext(aws.Context, *guardduty.GetInvitationsCountInput, ...request.Option) (*guardduty.GetInvitationsCountOutput, error)
	GetInvitationsCountRequest(*guardduty.GetInvitationsCountInput) (*request.Request, *guardduty.GetInvitationsCountOutput)

	GetMalwareScanSettings(*guardduty.GetMalwareScanSettingsInput) (*guardduty.GetMalwareScanSettingsOutput, error)
	GetMalwareScanSettingsWithContext(aws.Context, *guardduty.GetMalwareScanSettingsInput, ...request.Option) (*guardduty.GetMalwareScanSettingsOutput, error)
	GetMalwareScanSettingsRequest(*guardduty.GetMalwareScanSettingsInput) (*request.Request, *guardduty.GetMalwareScanSettingsOutput)

	GetMasterAccount(*guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error)
	GetMasterAccountWithContext(aws.Context, *guardduty.GetMasterAccountInput, ...request.Option) (*guardduty.GetMasterAccountOutput, error)
	GetMasterAccountRequest(*guardduty.GetMasterAccountInput) (*request.Request, *guardduty.GetMasterAccountOutput)

	GetMemberDetectors(*guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error)
	GetMemberDetectorsWithContext(aws.Context, *guardduty.GetMemberDetectorsInput, ...request.Option) (*guardduty.GetMemberDetectorsOutput, error)
	GetMemberDetectorsRequest(*guardduty.GetMemberDetectorsInput) (*request.Request, *guardduty.GetMemberDetectorsOutput)

	GetMembers(*guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error)
	GetMembersWithContext(aws.Context, *guardduty.GetMembersInput, ...request.Option) (*guardduty.GetMembersOutput, error)
	GetMembersRequest(*guardduty.GetMembersInput) (*request.Request, *guardduty.GetMembersOutput)

	GetRemainingFreeTrialDays(*guardduty.GetRemainingFreeTrialDaysInput) (*guardduty.GetRemainingFreeTrialDaysOutput, error)
	GetRemainingFreeTrialDaysWithContext(aws.Context, *guardduty.GetRemainingFreeTrialDaysInput, ...request.Option) (*guardduty.GetRemainingFreeTrialDaysOutput, error)
	GetRemainingFreeTrialDaysRequest(*guardduty.GetRemainingFreeTrialDaysInput) (*request.Request, *guardduty.GetRemainingFreeTrialDaysOutput)

	GetThreatIntelSet(*guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error)
	GetThreatIntelSetWithContext(aws.Context, *guardduty.GetThreatIntelSetInput, ...request.Option) (*guardduty.GetThreatIntelSetOutput, error)
	GetThreatIntelSetRequest(*guardduty.GetThreatIntelSetInput) (*request.Request, *guardduty.GetThreatIntelSetOutput)

	GetUsageStatistics(*guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error)
	GetUsageStatisticsWithContext(aws.Context, *guardduty.GetUsageStatisticsInput, ...request.Option) (*guardduty.GetUsageStatisticsOutput, error)
	GetUsageStatisticsRequest(*guardduty.GetUsageStatisticsInput) (*request.Request, *guardduty.GetUsageStatisticsOutput)

	GetUsageStatisticsPages(*guardduty.GetUsageStatisticsInput, func(*guardduty.GetUsageStatisticsOutput, bool) bool) error
	GetUsageStatisticsPagesWithContext(aws.Context, *guardduty.GetUsageStatisticsInput, func(*guardduty.GetUsageStatisticsOutput, bool) bool, ...request.Option) error

	InviteMembers(*guardduty.InviteMembersInput) (*guardduty.InviteMembersOutput, error)
	InviteMembersWithContext(aws.Context, *guardduty.InviteMembersInput, ...request.Option) (*guardduty.InviteMembersOutput, error)
	InviteMembersRequest(*guardduty.InviteMembersInput) (*request.Request, *guardduty.InviteMembersOutput)

	ListCoverage(*guardduty.ListCoverageInput) (*guardduty.ListCoverageOutput, error)
	ListCoverageWithContext(aws.Context, *guardduty.ListCoverageInput, ...request.Option) (*guardduty.ListCoverageOutput, error)
	ListCoverageRequest(*guardduty.ListCoverageInput) (*request.Request, *guardduty.ListCoverageOutput)

	ListCoveragePages(*guardduty.ListCoverageInput, func(*guardduty.ListCoverageOutput, bool) bool) error
	ListCoveragePagesWithContext(aws.Context, *guardduty.ListCoverageInput, func(*guardduty.ListCoverageOutput, bool) bool, ...request.Option) error

	ListDetectors(*guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error)
	ListDetectorsWithContext(aws.Context, *guardduty.ListDetectorsInput, ...request.Option) (*guardduty.ListDetectorsOutput, error)
	ListDetectorsRequest(*guardduty.ListDetectorsInput) (*request.Request, *guardduty.ListDetectorsOutput)

	ListDetectorsPages(*guardduty.ListDetectorsInput, func(*guardduty.ListDetectorsOutput, bool) bool) error
	ListDetectorsPagesWithContext(aws.Context, *guardduty.ListDetectorsInput, func(*guardduty.ListDetectorsOutput, bool) bool, ...request.Option) error

	ListFilters(*guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error)
	ListFiltersWithContext(aws.Context, *guardduty.ListFiltersInput, ...request.Option) (*guardduty.ListFiltersOutput, error)
	ListFiltersRequest(*guardduty.ListFiltersInput) (*request.Request, *guardduty.ListFiltersOutput)

	ListFiltersPages(*guardduty.ListFiltersInput, func(*guardduty.ListFiltersOutput, bool) bool) error
	ListFiltersPagesWithContext(aws.Context, *guardduty.ListFiltersInput, func(*guardduty.ListFiltersOutput, bool) bool, ...request.Option) error

	ListFindings(*guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error)
	ListFindingsWithContext(aws.Context, *guardduty.ListFindingsInput, ...request.Option) (*guardduty.ListFindingsOutput, error)
	ListFindingsRequest(*guardduty.ListFindingsInput) (*request.Request, *guardduty.ListFindingsOutput)

	ListFindingsPages(*guardduty.ListFindingsInput, func(*guardduty.ListFindingsOutput, bool) bool) error
	ListFindingsPagesWithContext(aws.Context, *guardduty.ListFindingsInput, func(*guardduty.ListFindingsOutput, bool) bool, ...request.Option) error

	ListIPSets(*guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error)
	ListIPSetsWithContext(aws.Context, *guardduty.ListIPSetsInput, ...request.Option) (*guardduty.ListIPSetsOutput, error)
	ListIPSetsRequest(*guardduty.ListIPSetsInput) (*request.Request, *guardduty.ListIPSetsOutput)

	ListIPSetsPages(*guardduty.ListIPSetsInput, func(*guardduty.ListIPSetsOutput, bool) bool) error
	ListIPSetsPagesWithContext(aws.Context, *guardduty.ListIPSetsInput, func(*guardduty.ListIPSetsOutput, bool) bool, ...request.Option) error

	ListInvitations(*guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error)
	ListInvitationsWithContext(aws.Context, *guardduty.ListInvitationsInput, ...request.Option) (*guardduty.ListInvitationsOutput, error)
	ListInvitationsRequest(*guardduty.ListInvitationsInput) (*request.Request, *guardduty.ListInvitationsOutput)

	ListInvitationsPages(*guardduty.ListInvitationsInput, func(*guardduty.ListInvitationsOutput, bool) bool) error
	ListInvitationsPagesWithContext(aws.Context, *guardduty.ListInvitationsInput, func(*guardduty.ListInvitationsOutput, bool) bool, ...request.Option) error

	ListMembers(*guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error)
	ListMembersWithContext(aws.Context, *guardduty.ListMembersInput, ...request.Option) (*guardduty.ListMembersOutput, error)
	ListMembersRequest(*guardduty.ListMembersInput) (*request.Request, *guardduty.ListMembersOutput)

	ListMembersPages(*guardduty.ListMembersInput, func(*guardduty.ListMembersOutput, bool) bool) error
	ListMembersPagesWithContext(aws.Context, *guardduty.ListMembersInput, func(*guardduty.ListMembersOutput, bool) bool, ...request.Option) error

	ListOrganizationAdminAccounts(*guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsWithContext(aws.Context, *guardduty.ListOrganizationAdminAccountsInput, ...request.Option) (*guardduty.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsRequest(*guardduty.ListOrganizationAdminAccountsInput) (*request.Request, *guardduty.ListOrganizationAdminAccountsOutput)

	ListOrganizationAdminAccountsPages(*guardduty.ListOrganizationAdminAccountsInput, func(*guardduty.ListOrganizationAdminAccountsOutput, bool) bool) error
	ListOrganizationAdminAccountsPagesWithContext(aws.Context, *guardduty.ListOrganizationAdminAccountsInput, func(*guardduty.ListOrganizationAdminAccountsOutput, bool) bool, ...request.Option) error

	ListPublishingDestinations(*guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error)
	ListPublishingDestinationsWithContext(aws.Context, *guardduty.ListPublishingDestinationsInput, ...request.Option) (*guardduty.ListPublishingDestinationsOutput, error)
	ListPublishingDestinationsRequest(*guardduty.ListPublishingDestinationsInput) (*request.Request, *guardduty.ListPublishingDestinationsOutput)

	ListPublishingDestinationsPages(*guardduty.ListPublishingDestinationsInput, func(*guardduty.ListPublishingDestinationsOutput, bool) bool) error
	ListPublishingDestinationsPagesWithContext(aws.Context, *guardduty.ListPublishingDestinationsInput, func(*guardduty.ListPublishingDestinationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *guardduty.ListTagsForResourceInput, ...request.Option) (*guardduty.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*guardduty.ListTagsForResourceInput) (*request.Request, *guardduty.ListTagsForResourceOutput)

	ListThreatIntelSets(*guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error)
	ListThreatIntelSetsWithContext(aws.Context, *guardduty.ListThreatIntelSetsInput, ...request.Option) (*guardduty.ListThreatIntelSetsOutput, error)
	ListThreatIntelSetsRequest(*guardduty.ListThreatIntelSetsInput) (*request.Request, *guardduty.ListThreatIntelSetsOutput)

	ListThreatIntelSetsPages(*guardduty.ListThreatIntelSetsInput, func(*guardduty.ListThreatIntelSetsOutput, bool) bool) error
	ListThreatIntelSetsPagesWithContext(aws.Context, *guardduty.ListThreatIntelSetsInput, func(*guardduty.ListThreatIntelSetsOutput, bool) bool, ...request.Option) error

	StartMalwareScan(*guardduty.StartMalwareScanInput) (*guardduty.StartMalwareScanOutput, error)
	StartMalwareScanWithContext(aws.Context, *guardduty.StartMalwareScanInput, ...request.Option) (*guardduty.StartMalwareScanOutput, error)
	StartMalwareScanRequest(*guardduty.StartMalwareScanInput) (*request.Request, *guardduty.StartMalwareScanOutput)

	StartMonitoringMembers(*guardduty.StartMonitoringMembersInput) (*guardduty.StartMonitoringMembersOutput, error)
	StartMonitoringMembersWithContext(aws.Context, *guardduty.StartMonitoringMembersInput, ...request.Option) (*guardduty.StartMonitoringMembersOutput, error)
	StartMonitoringMembersRequest(*guardduty.StartMonitoringMembersInput) (*request.Request, *guardduty.StartMonitoringMembersOutput)

	StopMonitoringMembers(*guardduty.StopMonitoringMembersInput) (*guardduty.StopMonitoringMembersOutput, error)
	StopMonitoringMembersWithContext(aws.Context, *guardduty.StopMonitoringMembersInput, ...request.Option) (*guardduty.StopMonitoringMembersOutput, error)
	StopMonitoringMembersRequest(*guardduty.StopMonitoringMembersInput) (*request.Request, *guardduty.StopMonitoringMembersOutput)

	TagResource(*guardduty.TagResourceInput) (*guardduty.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *guardduty.TagResourceInput, ...request.Option) (*guardduty.TagResourceOutput, error)
	TagResourceRequest(*guardduty.TagResourceInput) (*request.Request, *guardduty.TagResourceOutput)

	UnarchiveFindings(*guardduty.UnarchiveFindingsInput) (*guardduty.UnarchiveFindingsOutput, error)
	UnarchiveFindingsWithContext(aws.Context, *guardduty.UnarchiveFindingsInput, ...request.Option) (*guardduty.UnarchiveFindingsOutput, error)
	UnarchiveFindingsRequest(*guardduty.UnarchiveFindingsInput) (*request.Request, *guardduty.UnarchiveFindingsOutput)

	UntagResource(*guardduty.UntagResourceInput) (*guardduty.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *guardduty.UntagResourceInput, ...request.Option) (*guardduty.UntagResourceOutput, error)
	UntagResourceRequest(*guardduty.UntagResourceInput) (*request.Request, *guardduty.UntagResourceOutput)

	UpdateDetector(*guardduty.UpdateDetectorInput) (*guardduty.UpdateDetectorOutput, error)
	UpdateDetectorWithContext(aws.Context, *guardduty.UpdateDetectorInput, ...request.Option) (*guardduty.UpdateDetectorOutput, error)
	UpdateDetectorRequest(*guardduty.UpdateDetectorInput) (*request.Request, *guardduty.UpdateDetectorOutput)

	UpdateFilter(*guardduty.UpdateFilterInput) (*guardduty.UpdateFilterOutput, error)
	UpdateFilterWithContext(aws.Context, *guardduty.UpdateFilterInput, ...request.Option) (*guardduty.UpdateFilterOutput, error)
	UpdateFilterRequest(*guardduty.UpdateFilterInput) (*request.Request, *guardduty.UpdateFilterOutput)

	UpdateFindingsFeedback(*guardduty.UpdateFindingsFeedbackInput) (*guardduty.UpdateFindingsFeedbackOutput, error)
	UpdateFindingsFeedbackWithContext(aws.Context, *guardduty.UpdateFindingsFeedbackInput, ...request.Option) (*guardduty.UpdateFindingsFeedbackOutput, error)
	UpdateFindingsFeedbackRequest(*guardduty.UpdateFindingsFeedbackInput) (*request.Request, *guardduty.UpdateFindingsFeedbackOutput)

	UpdateIPSet(*guardduty.UpdateIPSetInput) (*guardduty.UpdateIPSetOutput, error)
	UpdateIPSetWithContext(aws.Context, *guardduty.UpdateIPSetInput, ...request.Option) (*guardduty.UpdateIPSetOutput, error)
	UpdateIPSetRequest(*guardduty.UpdateIPSetInput) (*request.Request, *guardduty.UpdateIPSetOutput)

	UpdateMalwareScanSettings(*guardduty.UpdateMalwareScanSettingsInput) (*guardduty.UpdateMalwareScanSettingsOutput, error)
	UpdateMalwareScanSettingsWithContext(aws.Context, *guardduty.UpdateMalwareScanSettingsInput, ...request.Option) (*guardduty.UpdateMalwareScanSettingsOutput, error)
	UpdateMalwareScanSettingsRequest(*guardduty.UpdateMalwareScanSettingsInput) (*request.Request, *guardduty.UpdateMalwareScanSettingsOutput)

	UpdateMemberDetectors(*guardduty.UpdateMemberDetectorsInput) (*guardduty.UpdateMemberDetectorsOutput, error)
	UpdateMemberDetectorsWithContext(aws.Context, *guardduty.UpdateMemberDetectorsInput, ...request.Option) (*guardduty.UpdateMemberDetectorsOutput, error)
	UpdateMemberDetectorsRequest(*guardduty.UpdateMemberDetectorsInput) (*request.Request, *guardduty.UpdateMemberDetectorsOutput)

	UpdateOrganizationConfiguration(*guardduty.UpdateOrganizationConfigurationInput) (*guardduty.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationWithContext(aws.Context, *guardduty.UpdateOrganizationConfigurationInput, ...request.Option) (*guardduty.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationRequest(*guardduty.UpdateOrganizationConfigurationInput) (*request.Request, *guardduty.UpdateOrganizationConfigurationOutput)

	UpdatePublishingDestination(*guardduty.UpdatePublishingDestinationInput) (*guardduty.UpdatePublishingDestinationOutput, error)
	UpdatePublishingDestinationWithContext(aws.Context, *guardduty.UpdatePublishingDestinationInput, ...request.Option) (*guardduty.UpdatePublishingDestinationOutput, error)
	UpdatePublishingDestinationRequest(*guardduty.UpdatePublishingDestinationInput) (*request.Request, *guardduty.UpdatePublishingDestinationOutput)

	UpdateThreatIntelSet(*guardduty.UpdateThreatIntelSetInput) (*guardduty.UpdateThreatIntelSetOutput, error)
	UpdateThreatIntelSetWithContext(aws.Context, *guardduty.UpdateThreatIntelSetInput, ...request.Option) (*guardduty.UpdateThreatIntelSetOutput, error)
	UpdateThreatIntelSetRequest(*guardduty.UpdateThreatIntelSetInput) (*request.Request, *guardduty.UpdateThreatIntelSetOutput)
}

var _ GuardDutyAPI = (*guardduty.GuardDuty)(nil)
