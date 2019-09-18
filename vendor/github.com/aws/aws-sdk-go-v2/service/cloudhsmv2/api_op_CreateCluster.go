// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsmv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/CreateClusterRequest
type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	// The type of HSM to use in the cluster. Currently the only allowed value is
	// hsm1.medium.
	//
	// HsmType is a required field
	HsmType *string `type:"string" required:"true"`

	// The identifier (ID) of the cluster backup to restore. Use this value to restore
	// the cluster from a backup instead of creating a new cluster. To find the
	// backup ID, use DescribeBackups.
	SourceBackupId *string `type:"string"`

	// The identifiers (IDs) of the subnets where you are creating the cluster.
	// You must specify at least one subnet. If you specify multiple subnets, they
	// must meet the following criteria:
	//
	//    * All subnets must be in the same virtual private cloud (VPC).
	//
	//    * You can specify only one subnet per Availability Zone.
	//
	// SubnetIds is a required field
	SubnetIds []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClusterInput"}

	if s.HsmType == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmType"))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}
	if s.SubnetIds != nil && len(s.SubnetIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SubnetIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/CreateClusterResponse
type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	// Information about the cluster that was created.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCluster = "CreateCluster"

// CreateClusterRequest returns a request value for making API operation for
// AWS CloudHSM V2.
//
// Creates a new AWS CloudHSM cluster.
//
//    // Example sending a request using CreateClusterRequest.
//    req := client.CreateClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/CreateCluster
func (c *Client) CreateClusterRequest(input *CreateClusterInput) CreateClusterRequest {
	op := &aws.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	req := c.newRequest(op, input, &CreateClusterOutput{})
	return CreateClusterRequest{Request: req, Input: input, Copy: c.CreateClusterRequest}
}

// CreateClusterRequest is the request type for the
// CreateCluster API operation.
type CreateClusterRequest struct {
	*aws.Request
	Input *CreateClusterInput
	Copy  func(*CreateClusterInput) CreateClusterRequest
}

// Send marshals and sends the CreateCluster API request.
func (r CreateClusterRequest) Send(ctx context.Context) (*CreateClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterResponse{
		CreateClusterOutput: r.Request.Data.(*CreateClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterResponse is the response type for the
// CreateCluster API operation.
type CreateClusterResponse struct {
	*CreateClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCluster request.
func (r *CreateClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}