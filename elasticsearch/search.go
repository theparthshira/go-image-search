package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type PhotoTag struct {
	Tag string `json:"tag"`
	Id  string `json:"id"`
}

func ConnectElasticSearch() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err
}

func IndexElasticData(client *elastic.Client, photoTag PhotoTag) error {
	ctx := context.Background()
	photoTagJSON, err := json.Marshal(photoTag)

	if err != nil {
		fmt.Println("Error in PhotoTag Marshal")
		return err
	}

	photoTagJSONString := string(photoTagJSON)

	result, err := client.Index().Index("photos").BodyJson(photoTagJSONString).Do(ctx)

	fmt.Println("result", result)

	if err != nil {
		fmt.Println("Error in elastic index")
		return err
	}

	return nil
}

func QueryElasticData(client *elastic.Client, search string) []PhotoTag {
	ctx := context.Background()

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("tag", search))

	searchService := client.Search().Index("photos").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
	}

	var photos []PhotoTag

	for _, hit := range searchResult.Hits.Hits {
		var photo PhotoTag
		err := json.Unmarshal(hit.Source, &photo)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		photos = append(photos, photo)
	}

	return photos
}
