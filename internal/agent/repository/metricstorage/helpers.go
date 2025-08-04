package metricstorage

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func compress(data []byte) ([]byte, error) {
	var b bytes.Buffer

	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return nil, fmt.Errorf("failed compress: %v", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed write data to compress temporary buffer: %v", err)
	}

	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("failed compress data: %v", err)
	}

	return b.Bytes(), nil
}

func send(req *resty.Request) error {
	r, err := req.Send()
	if err != nil {
		return err
	}

	if r.StatusCode() != http.StatusOK {
		return fmt.Errorf("%s", r.Status())
	}

	return nil
}

func craftURL(srvAddr, edpoint string) string {
	if len(srvAddr) == 5 {
		//Если длина 5, это значит что хост не указан. А для агента важно знать хост
		srvAddr = fmt.Sprintf("http://localhost%s", srvAddr)
	}

	edpointPath := fmt.Sprintf("%s/%s", srvAddr, edpoint)
	if edpointPath[:4] != "http" {
		edpointPath = fmt.Sprintf("http://%s", edpointPath)
	}

	return edpointPath
}
