// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package snowballiface provides an interface to enable mocking the Amazon Import/Export Snowball service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package snowballiface

import (
	"github.com/aws/aws-sdk-go-v2/service/snowball"
)

// ClientAPI provides an interface to enable mocking the
// snowball.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Snowball.
//    func myFunc(svc snowballiface.ClientAPI) bool {
//        // Make svc.CancelCluster request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := snowball.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        snowballiface.ClientPI
//    }
//    func (m *mockClientClient) CancelCluster(input *snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
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
type ClientAPI interface {
	CancelClusterRequest(*snowball.CancelClusterInput) snowball.CancelClusterRequest

	CancelJobRequest(*snowball.CancelJobInput) snowball.CancelJobRequest

	CreateAddressRequest(*snowball.CreateAddressInput) snowball.CreateAddressRequest

	CreateClusterRequest(*snowball.CreateClusterInput) snowball.CreateClusterRequest

	CreateJobRequest(*snowball.CreateJobInput) snowball.CreateJobRequest

	DescribeAddressRequest(*snowball.DescribeAddressInput) snowball.DescribeAddressRequest

	DescribeAddressesRequest(*snowball.DescribeAddressesInput) snowball.DescribeAddressesRequest

	DescribeClusterRequest(*snowball.DescribeClusterInput) snowball.DescribeClusterRequest

	DescribeJobRequest(*snowball.DescribeJobInput) snowball.DescribeJobRequest

	GetJobManifestRequest(*snowball.GetJobManifestInput) snowball.GetJobManifestRequest

	GetJobUnlockCodeRequest(*snowball.GetJobUnlockCodeInput) snowball.GetJobUnlockCodeRequest

	GetSnowballUsageRequest(*snowball.GetSnowballUsageInput) snowball.GetSnowballUsageRequest

	ListClusterJobsRequest(*snowball.ListClusterJobsInput) snowball.ListClusterJobsRequest

	ListClustersRequest(*snowball.ListClustersInput) snowball.ListClustersRequest

	ListCompatibleImagesRequest(*snowball.ListCompatibleImagesInput) snowball.ListCompatibleImagesRequest

	ListJobsRequest(*snowball.ListJobsInput) snowball.ListJobsRequest

	UpdateClusterRequest(*snowball.UpdateClusterInput) snowball.UpdateClusterRequest

	UpdateJobRequest(*snowball.UpdateJobInput) snowball.UpdateJobRequest
}

var _ ClientAPI = (*snowball.Client)(nil)
