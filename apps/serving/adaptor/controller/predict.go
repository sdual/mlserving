package controller

import (
	"context"

	pb "github.com/sdual/mlserving/proto/serving/predict"
)

type PredictServiceServer struct {
}

func (p PredictServiceServer) Predict(ctx context.Context, in *pb.PredictRequest) (*pb.PredictResponse, error) {
	return nil, nil
}
