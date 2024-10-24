package winestro_test

import (
	"testing"

	"github.com/nico0302/go-winestro"
	"github.com/stretchr/testify/assert"
)

func TestParseSingleProduct(t *testing.T) {
	resp := winestro.ProductResp{}
	decodeXMLFile(t, "single_product.xml", &resp)
	p := resp.Products[0]

	assert.Equal(t, "AP+2023", p.Number)
	assert.Equal(t, "<0.5", p.Fat)
	assert.NotEmpty(t, p.ProductGroups)
}
