package winestro

import (
	"encoding/xml"
	"fmt"
	"html"
	"net/url"
)

type Product struct {
	Number             string   `xml:"artikel_nr" json:"number"`
	Name               string   `xml:"artikel_name" json:"name"`
	Description        string   `xml:"artikel_beschreibung" json:"description"`
	Vintage            string   `xml:"artikel_jahrgang" json:"vintage"`
	Variety            string   `xml:"artikel_sorte" json:"variety"`
	Quality            string   `xml:"artikel_qualitaet" json:"quality"`
	Taste              string   `xml:"artikel_geschmack" json:"taste"`
	Sugar              string   `xml:"artikel_zucker" json:"sugar"`
	Alcohol            string   `xml:"artikel_alkohol" json:"alcohol"`
	Acid               string   `xml:"artikel_saeure" json:"acid"`
	Volume             string   `xml:"artikel_liter" json:"volume"`
	Weight             string   `xml:"artikel_gewicht" json:"weight"`
	Sulfites           bool     `xml:"artikel_sulfite" json:"sulfites"`
	Picture            string   `xml:"artikel_bild" json:"picture"`
	PictureLarge       string   `xml:"artikel_bild_big" json:"picture_large"`
	Picture2           string   `xml:"artikel_bild_2" json:"picture_2"`
	Picture2Large      string   `xml:"artikel_bild_big_2" json:"picture_2_large"`
	Picture3           string   `xml:"artikel_bild_3" json:"picture_3"`
	Picture3Large      string   `xml:"artikel_bild_big_3" json:"picture_3_large"`
	Picture4           string   `xml:"artikel_bild_4" json:"picture_4"`
	Picture4Large      string   `xml:"artikel_bild_big_4" json:"picture_4_large"`
	ShippingQuantity   int      `xml:"artikel_versandmenge" json:"shipping_quantity"`
	BundleQuantity     int      `xml:"artikel_sort_anzahl" json:"bundle_quantity"`
	Price              float32  `xml:"artikel_preis" json:"price"`
	PriceLiter         float32  `xml:"artikel_literpreis" json:"price_liter"`
	VAT                float32  `xml:"artikel_mwst" json:"vat"`
	Calories           string   `xml:"artikel_brennwert" json:"calories"`
	Protein            string   `xml:"artikel_eiweiss" json:"protein"`
	FreeShipping       bool     `xml:"artikel_versandfrei" json:"free_shipping"`
	HidePriceLiter     bool     `xml:"artikel_keinliterpreis" json:"hide_price_liter"`
	FillWeight         int      `xml:"artikel_fuellgewicht" json:"fill_weight"`
	PriceKg            float32  `xml:"artikel_kilopreis" json:"price_kg"`
	OutOfStockDate     string   `xml:"artikel_ausgetrunken" json:"out_of_stock_date"`
	APNR               string   `xml:"artikel_apnr" json:"apnr"`
	Vineyard           string   `xml:"artikel_lage" json:"vineyardn"`
	Expertise          string   `xml:"artikel_expertise" json:"expertise"`
	TypeName           string   `xml:"artikel_typ" json:"type_name"`
	TypeID             int      `xml:"artikel_typ_id" json:"type_id"`
	Color              string   `xml:"artikel_farbe" json:"color"`
	DrinkTemperature   string   `xml:"artikel_trinktemperatur" json:"drink_temperature"`
	StorageTemperature string   `xml:"artikel_lagertemperatur" json:"storage_temperature"`
	Elaboration        string   `xml:"artikel_ausbau" json:"elaboration"`
	Soil               string   `xml:"artikel_boden" json:"soil"`
	StorableYears      string   `xml:"artikel_lagerfaehigkeit" json:"storable_years"`
	VideoURL           string   `xml:"artikel_video" json:"video_url"`
	Country            string   `xml:"artikel_land" json:"country"`
	Region             string   `xml:"artikel_region" json:"region"`
	Appellation        string   `xml:"artikel_anbaugebiet" json:"appellation"`
	StockWarningLevel  int      `xml:"artikel_bestand_warnung_ab	" json:"stock_warning_level"`
	ProducerID         int      `xml:"artikel_erzeuger" json:"producer_id"`
	ProducerName       string   `xml:"artikel_erzeuger_name" json:"producer_name"`
	ProducerNumber     string   `xml:"artikel_erzeuger_nr" json:"producer_number"`
	Category           string   `xml:"artikel_kategorie" json:"category"`
	PackagingID        int      `xml:"artikel_verpackung" json:"packaging_id"`
	PackagingName      string   `xml:"artikel_verpackung_bezeichnung	" json:"packaging_name"`
	PackagingQuantity  int      `xml:"artikel_verpackung_inhalt" json:"packaging_quantity"`
	EAN13              string   `xml:"artikel_ean13" json:"ean13"`
	EAN13Packaging     string   `xml:"artikel_ean13_kiste" json:"ean13_packaging"`
	Ingredients        string   `xml:"artikel_zutaten" json:"ingredients"`
	BestByDate         string   `xml:"artikel_mhd" json:"best_by_date"`
	Fat                string   `xml:"artikel_fett" json:"fat"`
	UnsaturatedFat     string   `xml:"artikel_fetts" json:"unsaturated_fat"`
	Carbohydrates      string   `xml:"artikel_kohlenhydrate" json:"carbohydrates"`
	Salt               string   `xml:"artikel_salz" json:"salt"`
	Fibre              string   `xml:"artikel_ballast" json:"fibre"`
	Vitamins           string   `xml:"artikel_vitamine" json:"vitamins"`
	SulfuricAcids      string   `xml:"artikel_gesamt_schwefelsaeure" json:"sulfuric_acids"`
	FreeSulfuricAcids  string   `xml:"artikel_frei_schwefelsaeure" json:"free_sulfuric_acids"`
	Histamine          string   `xml:"artikel_histamin" json:"histamine"`
	Glycerin           string   `xml:"artikel_glycerin" json:"glycerin"`
	ELabelText         string   `xml:"artikel_labeltext" json:"elabel_text"`
	ELabelURL          string   `xml:"artikel_qr" json:"elabel_url"`
	ProductGroups      []string `xml:"artikel_warengruppen>warengruppe" json:"product_groups"`
	Stock              float32  `xml:"artikel_bestand" json:"stock"`
	CompanyStock       float32  `xml:"artikel_bestand_firmenverbund" json:"company_stock"`
	WebshopStock       float32  `xml:"artikel_bestand_webshop" json:"webshop_stock"`
}

func (p *Product) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type productAlias Product
	var v productAlias
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	v.Number, err = url.QueryUnescape(v.Number)
	if err != nil {
		return err
	}

	v.Fat = html.UnescapeString(v.Fat)
	v.UnsaturatedFat = html.UnescapeString(v.UnsaturatedFat)
	v.Carbohydrates = html.UnescapeString(v.Carbohydrates)
	v.Salt = html.UnescapeString(v.Salt)
	v.Fibre = html.UnescapeString(v.Fibre)
	v.Vitamins = html.UnescapeString(v.Vitamins)
	v.SulfuricAcids = html.UnescapeString(v.SulfuricAcids)
	v.FreeSulfuricAcids = html.UnescapeString(v.FreeSulfuricAcids)
	v.Histamine = html.UnescapeString(v.Histamine)
	v.Glycerin = html.UnescapeString(v.Glycerin)
	v.Protein = html.UnescapeString(v.Protein)
	v.Calories = html.UnescapeString(v.Calories)

	*p = Product(v)
	return nil
}

type ProductResp struct {
	XMLName  xml.Name  `xml:"items"`
	Products []Product `xml:"item"`
}

type ProductOptions struct {
	GroupID             int
	Query               string
	Number              string
	IncludeCompanyStock bool
}

func (c *Client) FetchProducts(opt ProductOptions) ([]Product, error) {
	params := map[string]string{}
	if opt.GroupID > 0 {
		params["id_grp"] = fmt.Sprintf("%d", opt.GroupID)
	}
	if opt.Query != "" {
		params["suchstring"] = opt.Query
	}
	if opt.Number != "" {
		params["artikelnr"] = opt.Number
	}
	if opt.IncludeCompanyStock {
		params["firmenverbund"] = "true"
	}
	var products ProductResp
	err := c.do("getArtikel", params, &products)
	return products.Products, err
}
