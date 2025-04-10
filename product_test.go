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

	assert.Equal(t, "5+2024", p.SKU)
	assert.Equal(t, "2024er Weißherbst Blauer Portugieser", p.Name)
	assert.Equal(t, "< 0.5", p.Fat)
	assert.NotEmpty(t, p.ProductGroups)
}

func TestParseBundle(t *testing.T) {
	resp := winestro.ProductResp{}
	decodeXMLFile(t, "bundle_product.xml", &resp)
	p := resp.Products[0]

	assert.Equal(t, "PR-FJ+2025", p.SKU)
	assert.Equal(t, "Probepaket Terroir & Classic", p.Name)
	assert.NotEmpty(t, p.BundleItems)

	assert.Equal(t, "5+2023", p.BundleItems[0].SKU)
}
