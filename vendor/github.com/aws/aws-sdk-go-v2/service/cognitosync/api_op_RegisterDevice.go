// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to RegisterDevice.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/RegisterDeviceRequest
type RegisterDeviceInput struct {
	_ struct{} `type:"structure"`

	// The unique ID for this identity.
	//
	// IdentityId is a required field
	IdentityId *string `location:"uri" locationName:"IdentityId" min:"1" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. Here, the ID of the pool that the identity belongs
	// to.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `location:"uri" locationName:"IdentityPoolId" min:"1" type:"string" required:"true"`

	// The SNS platform type (e.g. GCM, SDM, APNS, APNS_SANDBOX).
	//
	// Platform is a required field
	Platform Platform `type:"string" required:"true" enum:"true"`

	// The push token.
	//
	// Token is a required field
	Token *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterDeviceInput"}

	if s.IdentityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityId"))
	}
	if s.IdentityId != nil && len(*s.IdentityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityId", 1))
	}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}
	if len(s.Platform) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Platform"))
	}

	if s.Token == nil {
		invalidParams.Add(aws.NewErrParamRequired("Token"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterDeviceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.Platform) > 0 {
		v := s.Platform

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Platform", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityId != nil {
		v := *s.IdentityId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Response to a RegisterDevice request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/RegisterDeviceResponse
type RegisterDeviceOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID generated for this device by Cognito.
	DeviceId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s RegisterDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterDeviceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeviceId != nil {
		v := *s.DeviceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeviceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opRegisterDevice = "RegisterDevice"

// RegisterDeviceRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Registers a device to receive push sync notifications.
//
// This API can only be called with temporary credentials provided by Cognito
// Identity. You cannot call this API with developer credentials.
//
//    // Example sending a request using RegisterDeviceRequest.
//    req := client.RegisterDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/RegisterDevice
func (c *Client) RegisterDeviceRequest(input *RegisterDeviceInput) RegisterDeviceRequest {
	op := &aws.Operation{
		Name:       opRegisterDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/identitypools/{IdentityPoolId}/identity/{IdentityId}/device",
	}

	if input == nil {
		input = &RegisterDeviceInput{}
	}

	req := c.newRequest(op, input, &RegisterDeviceOutput{})
	return RegisterDeviceRequest{Request: req, Input: input, Copy: c.RegisterDeviceRequest}
}

// RegisterDeviceRequest is the request type for the
// RegisterDevice API operation.
type RegisterDeviceRequest struct {
	*aws.Request
	Input *RegisterDeviceInput
	Copy  func(*RegisterDeviceInput) RegisterDeviceRequest
}

// Send marshals and sends the RegisterDevice API request.
func (r RegisterDeviceRequest) Send(ctx context.Context) (*RegisterDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterDeviceResponse{
		RegisterDeviceOutput: r.Request.Data.(*RegisterDeviceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterDeviceResponse is the response type for the
// RegisterDevice API operation.
type RegisterDeviceResponse struct {
	*RegisterDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterDevice request.
func (r *RegisterDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
