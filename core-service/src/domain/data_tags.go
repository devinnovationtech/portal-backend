package domain

import "context"

// DataTag ..
type DataTag struct {
	ID       int64  `json:"id"`
	DataID   int64  `json:"data_id"`
	TagID    int64  `json:"tag_id"`
	TagsName string `json:"tags_name"`
	Type     string `json:"type"`
}

// DataTagRepository ..
type DataTagRepository interface {
	FetchDataTags(ctx context.Context, id int64) ([]DataTag, error)
	StoreDataTag(ctx context.Context, dt *DataTag) error
}
