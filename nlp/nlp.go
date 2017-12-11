package nlp

import (
	"github.com/tom1193/language-api/proto"
	"github.com/tom1193/language-api/websearch"
	language "cloud.google.com/go/language/apiv1"
	"golang.org/x/net/context"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	"log"
	"fmt"
)

func AnalyzeEntitySentiment(text string) (*languagepb.AnalyzeEntitySentimentResponse, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client.AnalyzeEntitySentiment(ctx, &languagepb.AnalyzeEntitySentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
}
func GenerateEntity(analyze *languagepb.AnalyzeEntitySentimentResponse) ([]proto.Entity) {
	var result []proto.Entity
	for _, entity := range analyze.Entities {
		for _, mention := range entity.Mentions {
			res := websearch.ImageQuery(entity.Name)
			images := websearch.ParseImageQueryResponse(res)
			result = append(result, proto.Entity{
				Name: entity.Name,
				Order: mention.Text.BeginOffset,
				Sentiment: mention.Sentiment.Score,
				Images: images,
			})
		}
	}
	return result
}