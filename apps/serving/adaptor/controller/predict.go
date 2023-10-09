package controller

import (
	"context"

	pb "github.com/sdual/mlserving/proto/grpc/serving/predict"
)

type PredictServiceServer struct {
	pb.UnimplementedPredictServiceServer
}

// Predict method responses predicted results.
// grpcurl -plaintext -d '{"id": 1001}' localhost:8080 serving.predict.PredictService.Predict
func (p PredictServiceServer) Predict(ctx context.Context, req *pb.PredictRequest) (*pb.PredictResponse, error) {
	return &pb.PredictResponse{
		Id: req.Id,
	}, nil
}
