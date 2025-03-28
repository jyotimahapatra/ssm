// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package chimesdkmediapipelinesiface provides an interface to enable mocking the Amazon Chime SDK Media Pipelines service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package chimesdkmediapipelinesiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/chimesdkmediapipelines"
)

// ChimeSDKMediaPipelinesAPI provides an interface to enable mocking the
// chimesdkmediapipelines.ChimeSDKMediaPipelines service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Chime SDK Media Pipelines.
//    func myFunc(svc chimesdkmediapipelinesiface.ChimeSDKMediaPipelinesAPI) bool {
//        // Make svc.CreateMediaCapturePipeline request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := chimesdkmediapipelines.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockChimeSDKMediaPipelinesClient struct {
//        chimesdkmediapipelinesiface.ChimeSDKMediaPipelinesAPI
//    }
//    func (m *mockChimeSDKMediaPipelinesClient) CreateMediaCapturePipeline(input *chimesdkmediapipelines.CreateMediaCapturePipelineInput) (*chimesdkmediapipelines.CreateMediaCapturePipelineOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockChimeSDKMediaPipelinesClient{}
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
type ChimeSDKMediaPipelinesAPI interface {
	CreateMediaCapturePipeline(*chimesdkmediapipelines.CreateMediaCapturePipelineInput) (*chimesdkmediapipelines.CreateMediaCapturePipelineOutput, error)
	CreateMediaCapturePipelineWithContext(aws.Context, *chimesdkmediapipelines.CreateMediaCapturePipelineInput, ...request.Option) (*chimesdkmediapipelines.CreateMediaCapturePipelineOutput, error)
	CreateMediaCapturePipelineRequest(*chimesdkmediapipelines.CreateMediaCapturePipelineInput) (*request.Request, *chimesdkmediapipelines.CreateMediaCapturePipelineOutput)

	CreateMediaConcatenationPipeline(*chimesdkmediapipelines.CreateMediaConcatenationPipelineInput) (*chimesdkmediapipelines.CreateMediaConcatenationPipelineOutput, error)
	CreateMediaConcatenationPipelineWithContext(aws.Context, *chimesdkmediapipelines.CreateMediaConcatenationPipelineInput, ...request.Option) (*chimesdkmediapipelines.CreateMediaConcatenationPipelineOutput, error)
	CreateMediaConcatenationPipelineRequest(*chimesdkmediapipelines.CreateMediaConcatenationPipelineInput) (*request.Request, *chimesdkmediapipelines.CreateMediaConcatenationPipelineOutput)

	CreateMediaInsightsPipeline(*chimesdkmediapipelines.CreateMediaInsightsPipelineInput) (*chimesdkmediapipelines.CreateMediaInsightsPipelineOutput, error)
	CreateMediaInsightsPipelineWithContext(aws.Context, *chimesdkmediapipelines.CreateMediaInsightsPipelineInput, ...request.Option) (*chimesdkmediapipelines.CreateMediaInsightsPipelineOutput, error)
	CreateMediaInsightsPipelineRequest(*chimesdkmediapipelines.CreateMediaInsightsPipelineInput) (*request.Request, *chimesdkmediapipelines.CreateMediaInsightsPipelineOutput)

	CreateMediaInsightsPipelineConfiguration(*chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationInput) (*chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationOutput, error)
	CreateMediaInsightsPipelineConfigurationWithContext(aws.Context, *chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationInput, ...request.Option) (*chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationOutput, error)
	CreateMediaInsightsPipelineConfigurationRequest(*chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationInput) (*request.Request, *chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationOutput)

	CreateMediaLiveConnectorPipeline(*chimesdkmediapipelines.CreateMediaLiveConnectorPipelineInput) (*chimesdkmediapipelines.CreateMediaLiveConnectorPipelineOutput, error)
	CreateMediaLiveConnectorPipelineWithContext(aws.Context, *chimesdkmediapipelines.CreateMediaLiveConnectorPipelineInput, ...request.Option) (*chimesdkmediapipelines.CreateMediaLiveConnectorPipelineOutput, error)
	CreateMediaLiveConnectorPipelineRequest(*chimesdkmediapipelines.CreateMediaLiveConnectorPipelineInput) (*request.Request, *chimesdkmediapipelines.CreateMediaLiveConnectorPipelineOutput)

	DeleteMediaCapturePipeline(*chimesdkmediapipelines.DeleteMediaCapturePipelineInput) (*chimesdkmediapipelines.DeleteMediaCapturePipelineOutput, error)
	DeleteMediaCapturePipelineWithContext(aws.Context, *chimesdkmediapipelines.DeleteMediaCapturePipelineInput, ...request.Option) (*chimesdkmediapipelines.DeleteMediaCapturePipelineOutput, error)
	DeleteMediaCapturePipelineRequest(*chimesdkmediapipelines.DeleteMediaCapturePipelineInput) (*request.Request, *chimesdkmediapipelines.DeleteMediaCapturePipelineOutput)

	DeleteMediaInsightsPipelineConfiguration(*chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationInput) (*chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationOutput, error)
	DeleteMediaInsightsPipelineConfigurationWithContext(aws.Context, *chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationInput, ...request.Option) (*chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationOutput, error)
	DeleteMediaInsightsPipelineConfigurationRequest(*chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationInput) (*request.Request, *chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationOutput)

	DeleteMediaPipeline(*chimesdkmediapipelines.DeleteMediaPipelineInput) (*chimesdkmediapipelines.DeleteMediaPipelineOutput, error)
	DeleteMediaPipelineWithContext(aws.Context, *chimesdkmediapipelines.DeleteMediaPipelineInput, ...request.Option) (*chimesdkmediapipelines.DeleteMediaPipelineOutput, error)
	DeleteMediaPipelineRequest(*chimesdkmediapipelines.DeleteMediaPipelineInput) (*request.Request, *chimesdkmediapipelines.DeleteMediaPipelineOutput)

	GetMediaCapturePipeline(*chimesdkmediapipelines.GetMediaCapturePipelineInput) (*chimesdkmediapipelines.GetMediaCapturePipelineOutput, error)
	GetMediaCapturePipelineWithContext(aws.Context, *chimesdkmediapipelines.GetMediaCapturePipelineInput, ...request.Option) (*chimesdkmediapipelines.GetMediaCapturePipelineOutput, error)
	GetMediaCapturePipelineRequest(*chimesdkmediapipelines.GetMediaCapturePipelineInput) (*request.Request, *chimesdkmediapipelines.GetMediaCapturePipelineOutput)

	GetMediaInsightsPipelineConfiguration(*chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationInput) (*chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationOutput, error)
	GetMediaInsightsPipelineConfigurationWithContext(aws.Context, *chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationInput, ...request.Option) (*chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationOutput, error)
	GetMediaInsightsPipelineConfigurationRequest(*chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationInput) (*request.Request, *chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationOutput)

	GetMediaPipeline(*chimesdkmediapipelines.GetMediaPipelineInput) (*chimesdkmediapipelines.GetMediaPipelineOutput, error)
	GetMediaPipelineWithContext(aws.Context, *chimesdkmediapipelines.GetMediaPipelineInput, ...request.Option) (*chimesdkmediapipelines.GetMediaPipelineOutput, error)
	GetMediaPipelineRequest(*chimesdkmediapipelines.GetMediaPipelineInput) (*request.Request, *chimesdkmediapipelines.GetMediaPipelineOutput)

	ListMediaCapturePipelines(*chimesdkmediapipelines.ListMediaCapturePipelinesInput) (*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, error)
	ListMediaCapturePipelinesWithContext(aws.Context, *chimesdkmediapipelines.ListMediaCapturePipelinesInput, ...request.Option) (*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, error)
	ListMediaCapturePipelinesRequest(*chimesdkmediapipelines.ListMediaCapturePipelinesInput) (*request.Request, *chimesdkmediapipelines.ListMediaCapturePipelinesOutput)

	ListMediaCapturePipelinesPages(*chimesdkmediapipelines.ListMediaCapturePipelinesInput, func(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, bool) bool) error
	ListMediaCapturePipelinesPagesWithContext(aws.Context, *chimesdkmediapipelines.ListMediaCapturePipelinesInput, func(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, bool) bool, ...request.Option) error

	ListMediaInsightsPipelineConfigurations(*chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsInput) (*chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsOutput, error)
	ListMediaInsightsPipelineConfigurationsWithContext(aws.Context, *chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsInput, ...request.Option) (*chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsOutput, error)
	ListMediaInsightsPipelineConfigurationsRequest(*chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsInput) (*request.Request, *chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsOutput)

	ListMediaInsightsPipelineConfigurationsPages(*chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsInput, func(*chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsOutput, bool) bool) error
	ListMediaInsightsPipelineConfigurationsPagesWithContext(aws.Context, *chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsInput, func(*chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsOutput, bool) bool, ...request.Option) error

	ListMediaPipelines(*chimesdkmediapipelines.ListMediaPipelinesInput) (*chimesdkmediapipelines.ListMediaPipelinesOutput, error)
	ListMediaPipelinesWithContext(aws.Context, *chimesdkmediapipelines.ListMediaPipelinesInput, ...request.Option) (*chimesdkmediapipelines.ListMediaPipelinesOutput, error)
	ListMediaPipelinesRequest(*chimesdkmediapipelines.ListMediaPipelinesInput) (*request.Request, *chimesdkmediapipelines.ListMediaPipelinesOutput)

	ListMediaPipelinesPages(*chimesdkmediapipelines.ListMediaPipelinesInput, func(*chimesdkmediapipelines.ListMediaPipelinesOutput, bool) bool) error
	ListMediaPipelinesPagesWithContext(aws.Context, *chimesdkmediapipelines.ListMediaPipelinesInput, func(*chimesdkmediapipelines.ListMediaPipelinesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*chimesdkmediapipelines.ListTagsForResourceInput) (*chimesdkmediapipelines.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *chimesdkmediapipelines.ListTagsForResourceInput, ...request.Option) (*chimesdkmediapipelines.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*chimesdkmediapipelines.ListTagsForResourceInput) (*request.Request, *chimesdkmediapipelines.ListTagsForResourceOutput)

	TagResource(*chimesdkmediapipelines.TagResourceInput) (*chimesdkmediapipelines.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *chimesdkmediapipelines.TagResourceInput, ...request.Option) (*chimesdkmediapipelines.TagResourceOutput, error)
	TagResourceRequest(*chimesdkmediapipelines.TagResourceInput) (*request.Request, *chimesdkmediapipelines.TagResourceOutput)

	UntagResource(*chimesdkmediapipelines.UntagResourceInput) (*chimesdkmediapipelines.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *chimesdkmediapipelines.UntagResourceInput, ...request.Option) (*chimesdkmediapipelines.UntagResourceOutput, error)
	UntagResourceRequest(*chimesdkmediapipelines.UntagResourceInput) (*request.Request, *chimesdkmediapipelines.UntagResourceOutput)

	UpdateMediaInsightsPipelineConfiguration(*chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationInput) (*chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationOutput, error)
	UpdateMediaInsightsPipelineConfigurationWithContext(aws.Context, *chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationInput, ...request.Option) (*chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationOutput, error)
	UpdateMediaInsightsPipelineConfigurationRequest(*chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationInput) (*request.Request, *chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationOutput)

	UpdateMediaInsightsPipelineStatus(*chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusInput) (*chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusOutput, error)
	UpdateMediaInsightsPipelineStatusWithContext(aws.Context, *chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusInput, ...request.Option) (*chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusOutput, error)
	UpdateMediaInsightsPipelineStatusRequest(*chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusInput) (*request.Request, *chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusOutput)
}

var _ ChimeSDKMediaPipelinesAPI = (*chimesdkmediapipelines.ChimeSDKMediaPipelines)(nil)
