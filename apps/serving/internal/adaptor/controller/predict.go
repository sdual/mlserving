package controller

import (
	"context"

	"github.com/sdual/mlserving/apps/serving/internal/usecase"
	pb "github.com/sdual/mlserving/proto/grpc/serving/predict"
)

type PredictServiceServer struct {
	prediction usecase.Prediction
	pb.UnimplementedPredictServiceServer
}

// Predict method responses predicted results.
// grpcurl -plaintext -d '{"id": 1001}' localhost:8080 serving.predict.PredictService.Predict
func (p PredictServiceServer) Predict(ctx context.Context, req *pb.PredictRequest) (*pb.PredictResponse, error) {
	p.prediction.Predict()
	return &pb.PredictResponse{}, nil
}
