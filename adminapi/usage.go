package adminapi

import (
	"context"
)

type UsageRequest struct {
	UID         string    `url:"uid,omitempty"`
	Start       RadosTime `url:"start,omitempty"`
	End         RadosTime `url:"end,omitempty"`
	ShowEntries bool      `url:"show-entries,omitempty"`
	ShowSummary bool      `url:"show-summary,omitempty"`
}

type TrimUsageRequest struct {
	UID       string    `url:"uid,omitempty"`
	Start     RadosTime `url:"start,omitempty"`
	End       RadosTime `url:"end,omitempty"`
	RemoveAll bool      `url:"remove-all,omitempty"`
}

func (aa *AdminAPI) TrimUsage(ctx context.Context, treq *TrimUsageRequest) error {
	err := aa.delete(ctx, "/usage", treq, nil)
	return err
}

func (aa *AdminAPI) GetUsage(ctx context.Context, ureq *UsageRequest) (*UsageResponse, error) {
	uresp := new(UsageResponse)

	err := aa.get(ctx, "/usage", ureq, uresp)
	return uresp, err
}

type UsageResponse struct {
	Entries []UsageEntry   `json:"entries"`
	Summary []UsageSummary `json:"summary"`
}

type UsageEntry struct {
	Buckets []UsageBucket `json:"buckets"`
	User    string        `json:"user"`
}

type UsageSummary struct {
	Categories []UsageCategory `json:"categories"`
	Total      *UsageTotal     `json:"total"`
}

type UsageBucket struct {
	Bucket     string          `json:"bucket"`
	Owner      string          `json:"owner"`
	Categories []UsageCategory `json:"categories"`
	Epoch      int             `json:"epoch"`
	Time       RadosTime       `json:"time"`
}

type UsageCategory struct {
	BytesSent     int    `json:"bytes_sent"`
	BytesReceived int    `json:"bytes_received"`
	Ops           int    `json:"ops"`
	SuccessfulOps int    `json:"successful_ops"`
	Category      string `json:"category"`
}

type UsageTotal struct {
	BytesSent     int `json:"bytes_sent"`
	BytesReceived int `json:"bytes_received"`
	Ops           int `json:"ops"`
	SuccessfulOps int `json:"successful_ops"`
}
