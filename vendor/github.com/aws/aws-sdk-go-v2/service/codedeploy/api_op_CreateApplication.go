// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateApplication operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/CreateApplicationInput
type CreateApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of the application. This name must be unique with the applicable
	// IAM user or AWS account.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`

	// The destination platform type for the deployment (Lambda, Server, or ECS).
	ComputePlatform ComputePlatform `locationName:"computePlatform" type:"string" enum:"true"`

	// The metadata that you apply to CodeDeploy applications to help you organize
	// and categorize them. Each tag consists of a key and an optional value, both
	// of which you define.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateApplication operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/CreateApplicationOutput
type CreateApplicationOutput struct {
	_ struct{} `type:"structure"`

	// A unique application ID.
	ApplicationId *string `locationName:"applicationId" type:"string"`
}

// String returns the string representation
func (s CreateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateApplication = "CreateApplication"

// CreateApplicationRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Creates an application.
//
//    // Example sending a request using CreateApplicationRequest.
//    req := client.CreateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/CreateApplication
func (c *Client) CreateApplicationRequest(input *CreateApplicationInput) CreateApplicationRequest {
	op := &aws.Operation{
		Name:       opCreateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateApplicationInput{}
	}

	req := c.newRequest(op, input, &CreateApplicationOutput{})
	return CreateApplicationRequest{Request: req, Input: input, Copy: c.CreateApplicationRequest}
}

// CreateApplicationRequest is the request type for the
// CreateApplication API operation.
type CreateApplicationRequest struct {
	*aws.Request
	Input *CreateApplicationInput
	Copy  func(*CreateApplicationInput) CreateApplicationRequest
}

// Send marshals and sends the CreateApplication API request.
func (r CreateApplicationRequest) Send(ctx context.Context) (*CreateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationResponse{
		CreateApplicationOutput: r.Request.Data.(*CreateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationResponse is the response type for the
// CreateApplication API operation.
type CreateApplicationResponse struct {
	*CreateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplication request.
func (r *CreateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
