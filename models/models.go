package models

type IndexInfo struct {
	Name              string
	SizeGB            float64
	Shards            int
	BalanceRatio      int
	RecommendedShards int
}

type RawIndex struct {
	Name      string `json:"index"`
	SizeStr   string `json:"pri.store.size"`
	ShardsStr string `json:"pri"`
}
