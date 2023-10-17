package model

type (
	FFMModelField struct {
		FieldIndex  int
		TargetField string
	}

	RawFFMParameter struct {
		Field         string               `json:"key"`
		LatentVectors map[string][]float64 `json:"value"`
	}

	FFMParameter struct {
		FieldIndex    int
		LatentVectors map[string][]float64
	}
)
