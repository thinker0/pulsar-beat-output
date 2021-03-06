// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListHyperParameterTuningJobsRequest
type ListHyperParameterTuningJobsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only tuning jobs that were created after the specified
	// time.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only tuning jobs that were created before the specified
	// time.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns only tuning jobs that were modified after the specified
	// time.
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only tuning jobs that were modified before the specified
	// time.
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of tuning jobs to return. The default value is 10.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the tuning job name. This filter returns only tuning jobs whose
	// name contains the specified string.
	NameContains *string `type:"string"`

	// If the result of the previous ListHyperParameterTuningJobs request was truncated,
	// the response includes a NextToken. To retrieve the next set of tuning jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`

	// The field to sort results by. The default is Name.
	SortBy HyperParameterTuningJobSortByOptions `type:"string" enum:"true"`

	// The sort order for results. The default is Ascending.
	SortOrder SortOrder `type:"string" enum:"true"`

	// A filter that returns only tuning jobs with the specified status.
	StatusEquals HyperParameterTuningJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListHyperParameterTuningJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListHyperParameterTuningJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListHyperParameterTuningJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListHyperParameterTuningJobsResponse
type ListHyperParameterTuningJobsOutput struct {
	_ struct{} `type:"structure"`

	// A list of HyperParameterTuningJobSummary objects that describe the tuning
	// jobs that the ListHyperParameterTuningJobs request returned.
	//
	// HyperParameterTuningJobSummaries is a required field
	HyperParameterTuningJobSummaries []HyperParameterTuningJobSummary `type:"list" required:"true"`

	// If the result of this ListHyperParameterTuningJobs request was truncated,
	// the response includes a NextToken. To retrieve the next set of tuning jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHyperParameterTuningJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListHyperParameterTuningJobs = "ListHyperParameterTuningJobs"

// ListHyperParameterTuningJobsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Gets a list of HyperParameterTuningJobSummary objects that describe the hyperparameter
// tuning jobs launched in your account.
//
//    // Example sending a request using ListHyperParameterTuningJobsRequest.
//    req := client.ListHyperParameterTuningJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListHyperParameterTuningJobs
func (c *Client) ListHyperParameterTuningJobsRequest(input *ListHyperParameterTuningJobsInput) ListHyperParameterTuningJobsRequest {
	op := &aws.Operation{
		Name:       opListHyperParameterTuningJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListHyperParameterTuningJobsInput{}
	}

	req := c.newRequest(op, input, &ListHyperParameterTuningJobsOutput{})
	return ListHyperParameterTuningJobsRequest{Request: req, Input: input, Copy: c.ListHyperParameterTuningJobsRequest}
}

// ListHyperParameterTuningJobsRequest is the request type for the
// ListHyperParameterTuningJobs API operation.
type ListHyperParameterTuningJobsRequest struct {
	*aws.Request
	Input *ListHyperParameterTuningJobsInput
	Copy  func(*ListHyperParameterTuningJobsInput) ListHyperParameterTuningJobsRequest
}

// Send marshals and sends the ListHyperParameterTuningJobs API request.
func (r ListHyperParameterTuningJobsRequest) Send(ctx context.Context) (*ListHyperParameterTuningJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHyperParameterTuningJobsResponse{
		ListHyperParameterTuningJobsOutput: r.Request.Data.(*ListHyperParameterTuningJobsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListHyperParameterTuningJobsRequestPaginator returns a paginator for ListHyperParameterTuningJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListHyperParameterTuningJobsRequest(input)
//   p := sagemaker.NewListHyperParameterTuningJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListHyperParameterTuningJobsPaginator(req ListHyperParameterTuningJobsRequest) ListHyperParameterTuningJobsPaginator {
	return ListHyperParameterTuningJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListHyperParameterTuningJobsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListHyperParameterTuningJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListHyperParameterTuningJobsPaginator struct {
	aws.Pager
}

func (p *ListHyperParameterTuningJobsPaginator) CurrentPage() *ListHyperParameterTuningJobsOutput {
	return p.Pager.CurrentPage().(*ListHyperParameterTuningJobsOutput)
}

// ListHyperParameterTuningJobsResponse is the response type for the
// ListHyperParameterTuningJobs API operation.
type ListHyperParameterTuningJobsResponse struct {
	*ListHyperParameterTuningJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHyperParameterTuningJobs request.
func (r *ListHyperParameterTuningJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
