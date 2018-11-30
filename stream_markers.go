package helix

// Marker ...
type Marker struct {
	ID              string `json:"id"`
	CreatedAt       Time   `json:"created_at"`
	Description     string `json:"description"`
	PositionSeconds int    `json:"position_seconds"`
	URL             string `json:"URL"`
}

// VideoMarker ...
type VideoMarker struct {
	VideoID string   `json:"video_id"`
	Markers []Marker `json:"markers"`
}

// StreamMarker ...
type StreamMarker struct {
	UserID   string        `json:"user_id"`
	UserName string        `json:"user_name"`
	Videos   []VideoMarker `json:"videos"`
}

// ManyStreamMarkers ...
type ManyStreamMarkers struct {
	StreamMarkers []StreamMarker `json:"data"`
	Pagination    Pagination     `json:"pagination"`
}

// StreamMarkersResponse ...
type StreamMarkersResponse struct {
	ResponseCommon
	Data ManyStreamMarkers
}

// StreamMarkersParams requires _either_ UserID or VideoID set
//
// UserID: fetches stream markers of the current livestream of the given user
// (VOD recording must be enabled).
// VideoID: fetches streams markers of the VOD.
type StreamMarkersParams struct {
	UserID  string `query:"user_id"`
	VideoID string `query:"video_id"`

	// Optional
	After  string `query:"after"`
	Before string `query:"before"`
	First  int    `query:"first,20"` // Limit 100
}

// GetStreamMarkers gets stream markers of a VOD or of the current live stream
// of an user being recorded as VOD.
//
// Required Scope: user:read:broadcast
func (c *Client) GetStreamMarkers(params *StreamMarkersParams) (*StreamMarkersResponse, error) {
	resp, err := c.get("/streams/markers", &ManyStreamMarkers{}, params)
	if err != nil {
		return nil, err
	}

	markers := &StreamMarkersResponse{}
	markers.StatusCode = resp.StatusCode
	markers.Header = resp.Header
	markers.Error = resp.Error
	markers.ErrorStatus = resp.ErrorStatus
	markers.ErrorMessage = resp.ErrorMessage
	markers.Data.StreamMarkers = resp.Data.(*ManyStreamMarkers).StreamMarkers
	markers.Data.Pagination = resp.Data.(*ManyStreamMarkers).Pagination

	return markers, nil
}
