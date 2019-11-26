package ocr

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

// PDFToText はgcs上のpdfをText化する
func PDFToText(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Query()
	pdfName := f["pdf"][0]
	bucketName := f["bucket"][0]
	gcsSourceURI := fmt.Sprintf("gs://%s/%s.pdf", bucketName, pdfName)
	gcsDestinationURI := fmt.Sprintf("gs://%s/%s/", bucketName, pdfName)

	ctx := context.Background()
	cl, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		errorHundler(w, err)
		return
	}

	req := &visionpb.AsyncBatchAnnotateFilesRequest{
		Requests: []*visionpb.AsyncAnnotateFileRequest{
			{
				Features: []*visionpb.Feature{
					{
						Type: visionpb.Feature_DOCUMENT_TEXT_DETECTION,
					},
				},
				InputConfig: &visionpb.InputConfig{
					GcsSource: &visionpb.GcsSource{Uri: gcsSourceURI},
					MimeType:  "application/pdf",
				},
				OutputConfig: &visionpb.OutputConfig{
					GcsDestination: &visionpb.GcsDestination{Uri: gcsDestinationURI},
					BatchSize:      1,
				},
			},
		},
	}

	op, err := cl.AsyncBatchAnnotateFiles(ctx, req)
	if err != nil {
		errorHundler(w, err)
		return
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		errorHundler(w, err)
		return
	}
	fmt.Fprintf(w, "%v", resp.Responses[0].OutputConfig.GcsDestination)
}

func errorHundler(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusNotFound)
	log.Printf("%+v", e.Error())
	fmt.Fprintf(w, html.EscapeString(e.Error()))
}
