package controller

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/sdual/mlserving/apps/serving/internal/adaptor/controller/convert"
	"github.com/sdual/mlserving/apps/serving/internal/usecase"
	pb "github.com/sdual/mlserving/proto/grpc/serving/predict"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PredictServiceServer struct {
	prediction usecase.Prediction
	pb.UnimplementedPredictServiceServer
}

func NewPredictServiceServer(pred usecase.Prediction) *PredictServiceServer {
	return &PredictServiceServer{
		prediction: pred,
	}
}

// Predict method responses predicted results.
// grpcurl -plaintext -d '{"id": 1001}' localhost:8080 serving.predict.PredictService.Predict
func (p PredictServiceServer) Predict(ctx context.Context, req *pb.PredictRequest) (*pb.PredictResponse, error) {
	featureInput := convert.ConvertToFeatures(req)
	preds, err := p.prediction.Predict(featureInput)
	if err != nil {
		log.Err(err).Msgf("failed to predict: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.PredictResponse{
		Predictions: preds,
	}, nil
}
