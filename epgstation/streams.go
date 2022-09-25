// Copyright 2022 coord_e
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  	 http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package epgstation

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type GetStreamsOpts struct {
	IsHalfWidth bool `url:"isHalfWidth"`
}

type GetStreamsResponse StreamInfo

func (c *Client) GetStreams(ctx context.Context, opts GetStreamsOpts) (*GetStreamsResponse, error) {
	q, err := query.Values(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create query string: %w", err)
	}

	req, err := c.newRequest(ctx, "GET", "/api/streams", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to dispatch request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("non-success status code %d", resp.StatusCode)
	}

	var streams GetStreamsResponse
	if err := decodeBody(resp, &streams); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &streams, nil
}
