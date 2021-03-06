// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTableVersionRequest
type GetTableVersionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the tables reside. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The database in the catalog in which the table resides. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// The name of the table. For Hive compatibility, this name is entirely lowercase.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`

	// The ID value of the table version to be retrieved. A VersionID is a string
	// representation of an integer. Each version is incremented by 1.
	VersionId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetTableVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTableVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTableVersionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}
	if s.VersionId != nil && len(*s.VersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTableVersionResponse
type GetTableVersionOutput struct {
	_ struct{} `type:"structure"`

	// The requested table version.
	TableVersion *TableVersion `type:"structure"`
}

// String returns the string representation
func (s GetTableVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTableVersion = "GetTableVersion"

// GetTableVersionRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a specified version of a table.
//
//    // Example sending a request using GetTableVersionRequest.
//    req := client.GetTableVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetTableVersion
func (c *Client) GetTableVersionRequest(input *GetTableVersionInput) GetTableVersionRequest {
	op := &aws.Operation{
		Name:       opGetTableVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTableVersionInput{}
	}

	req := c.newRequest(op, input, &GetTableVersionOutput{})
	return GetTableVersionRequest{Request: req, Input: input, Copy: c.GetTableVersionRequest}
}

// GetTableVersionRequest is the request type for the
// GetTableVersion API operation.
type GetTableVersionRequest struct {
	*aws.Request
	Input *GetTableVersionInput
	Copy  func(*GetTableVersionInput) GetTableVersionRequest
}

// Send marshals and sends the GetTableVersion API request.
func (r GetTableVersionRequest) Send(ctx context.Context) (*GetTableVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTableVersionResponse{
		GetTableVersionOutput: r.Request.Data.(*GetTableVersionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTableVersionResponse is the response type for the
// GetTableVersion API operation.
type GetTableVersionResponse struct {
	*GetTableVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTableVersion request.
func (r *GetTableVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
