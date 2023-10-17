package model

type (
	FFMModelField struct {
		FieldIndex  int
		TargetField string
	}

	CacheFFMParameter struct {
		Field         string            `json:"key"`
		LatentVectors map[int][]float64 `json:"value"`
	}

	FFMParameter struct {
		FieldIndex    int
		LatentVectors map[int][]float64
	}
)
