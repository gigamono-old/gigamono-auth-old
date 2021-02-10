module github.com/sageflow/sageauth

go 1.15

require (
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/sageflow/sageengine v0.0.0-20210209170653-d3d0fb4ec2a2
	github.com/sageflow/sageflow v0.0.0-20210209165522-8b9455bbe20d
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/sageflow/sageflow v0.0.0-20210209165522-8b9455bbe20d => ../sageflow
