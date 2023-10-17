package model

type (
	FFMModelField struct {
		FieldIndex  int
		TargetField string
	}

	CacheFFMParameter struct {
		Field         string
		LatentVectors map[int][]float64
	}

	FFMParameter struct {
		FieldIndex    int
		LatentVectors map[int][]float64
	}
)
