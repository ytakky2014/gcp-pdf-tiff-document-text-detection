package ocr

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// ShowJSON はgcs上のjsonを表示する
func ShowJSON(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Query()
	pdfName := f["pdf"][0]
	bucketName := f["bucket"][0]

	ctx := context.Background()
	cli, err := storage.NewClient(ctx)

	if err != nil {
		errorHundler(w, err)
		return
	}

	q := storage.Query{Prefix: fmt.Sprintf("%s/", pdfName)}
	it := cli.Bucket(bucketName).Objects(ctx, &q)

	for {
		obj, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			errorHundler(w, err)
			return
		}

		r, err := cli.Bucket(bucketName).Object(obj.Name).NewReader(ctx)
		if err != nil {
			errorHundler(w, err)
			return
		}

		j, err := ioutil.ReadAll(r)
		if err != nil {
			errorHundler(w, err)
			return
		}

		var t Text
		if err := json.Unmarshal(j, &t); err != nil {
			errorHundler(w, err)
			return
		}
		fmt.Fprintf(w, "%+v", t.Responses[0].FullTextAnnotation.Text)
	}
	return
}

type Text struct {
	Responses []struct {
		FullTextAnnotation struct {
			Pages []struct {
				Blocks []struct {
					Paragraphs []struct {
						Words []struct {
							Symbols []struct {
								Text string `json:"text"`
							} `json:"symbols"`
						} `json:"words"`
					} `json:"paragraphs"`
				} `json:"blocks"`
			} `json:"pages"`
			Text string `json:"text"`
		} `json:"fullTextAnnotation"`
	} `json:"responses"`
}
