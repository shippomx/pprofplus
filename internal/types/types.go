package types

type MemProfile struct {
	CreatedAt	int64 `json:"created_at"`
	FilePath	string `json:"file_path"`
	ProMemRss	float64 `json:"pro_mem_rss"`
	Unit 		string `json:"unit"`
}
