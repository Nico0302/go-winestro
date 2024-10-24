package winestro_test

import (
	"encoding/xml"
	"os"
	"testing"
)

const testdataDir = "testdata"

func decodeXMLFile(t *testing.T, filename string, v interface{}) {
	t.Helper()

	f, err := os.Open(testdataDir + "/" + filename)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := xml.NewDecoder(f).Decode(v); err != nil {
		t.Fatal(err)
	}
}
