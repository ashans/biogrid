package config

import "biogrid/internal/entities"

func InitConfig() entities.AlignConfig {
	return entities.AlignConfig{
		Seq1: "ATGCT",
		Seq2: "AGCT",
		Mode: "Global",
		Scheme: entities.Scheme{
			Match:    1,
			Mismatch: -1,
			Gap:      -2,
		},
	}
}
