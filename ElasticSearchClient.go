package elasticsearch

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
)

var EsClient esClientInterface = &esClient{}

//we can use without interface but while mocking and changine behaviour while testing we use interface
type esClientInterface interface {
	setClient(client *elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

//instaead of providing client package to outside world we are going to work through an interface

// Obtain a client for an Elasticsearch cluster of two nodes,
// running on 10.0.1.1 and 10.0.1.2. Do not run the sniffer.
// Set the healthcheck interval to 10s. When requests fail,
// retry 5 times. Print error messages to os.Stderr and informational
// messages to os.Stdout.

func Init() {

	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetRetrier(NewCustomRetrier()),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetHeaders(http.Header{
			"X-Caller-Id": []string{"..."},
		}),
	)

	if err != nil {
		panic(err)
	}
	EsClient.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}
func (c *esClient) Index(interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}
