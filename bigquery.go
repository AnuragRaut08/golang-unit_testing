package bigquery

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

type Client struct {
	Conn *bigquery.Client
}

func NewClient(config *Config) (*Client, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, config.ProjectID, option.WithCredentialsFile(config.CredsPath))
	if err != nil {
		return nil, fmt.Errorf("Error: Failed to create BigQuery client: %v", err)
	}
	return &Client{Conn: client}, nil
	
}


func (c *Client) TestConnection() error {
	ctx := context.Background()
	datasets := c.Conn.Datasets(ctx)
	_, err := datasets.Next()
	if err != nil {
		return fmt.Errorf("Error: Failed to connect to BigQuery: %v", err)
	}
	return nil
}
