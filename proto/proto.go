package proto

type Request struct {
	Text string `json:"text, omitempty"`
}

type Entity struct {
	Name string `json:"name,omitempty"`
	Order int32 `json:"order,omitempty"`
	Sentiment float32 `json:"sentiment,omitempty"`
	Images []Image `json:"images,omitempty"`
}
type Image struct {
	Iid string `json:"iid,omitempty"`
	Url string `json:"url,omitempty"`
}
