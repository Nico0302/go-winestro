package winestro

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

type Customer struct {
	ID               int32   `xml:"adr_id" json:"id"`
	Number           string  `xml:"adr_nr" json:"number"`
	FirstName        string  `xml:"adr_vorname" json:"first_name"`
	LastName         string  `xml:"adr_nachname" json:"last_name"`
	Company          string  `xml:"adr_firma" json:"company"`
	ZIP              string  `xml:"adr_plz" json:"zip"`
	City             string  `xml:"adr_ort" json:"city"`
	WWW              string  `xml:"adr_www" json:"www"`
	Email            string  `xml:"adr_email" json:"email"`
	Street           string  `xml:"adr_str" json:"street"`
	HouseNum         string  `xml:"adr_str_nr" json:"house_number"`
	Country          string  `xml:"adr_land" json:"country"`
	Landline         string  `xml:"adr_festnetz" json:"landline"`
	Mobile           string  `xml:"adr_mobil" json:"mobile"`
	Fax              string  `xml:"adr_fax" json:"fax"`
	Note1            string  `xml:"adr_note1" json:"note1"`
	Note2            string  `xml:"adr_note2" json:"note2"`
	Note3            string  `xml:"adr_note3" json:"note3"`
	Note4            string  `xml:"adr_note4" json:"note4"`
	Discount         float64 `xml:"adr_rabatt" json:"discount"`
	PriceCategory    int     `xml:"adr_id_preiskategorie" json:"price_category"`
	NewsletterActive bool    `xml:"adr_newsletter_aktiv" json:"newsletter_active"`
	TaxType          int     `xml:"adr_kunden_mwst" json:"tax_type"`
	Salutation       string  `xml:"adr_anrede" json:"salutation"`
	SalutationType   int     `xml:"adr_anredenart" json:"salutation_type"`
	PaymentType      int     `xml:"adr_id_zahlungsart" json:"payment_type"`
}

func (c Customer) FullName() string {
	if c.FirstName == "" && c.LastName == "" {
		return c.Company
	}
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func (c Customer) FullSalutation() string {
	re := regexp.MustCompile(`,\s*$`)
	trimSaluation := strings.TrimSpace(re.ReplaceAllString(c.Salutation, ""))
	switch c.SalutationType {
	case 0:
		return fmt.Sprintf("%s %s", trimSaluation, strings.TrimSpace(c.FirstName))
	case 1:
		return fmt.Sprintf("%s %s", trimSaluation, strings.TrimSpace(c.LastName))
	default:
		return trimSaluation
	}
}

type CustomersResp struct {
	XMLName   xml.Name   `xml:"items"`
	Customers []Customer `xml:"item"`
}

func (c *Client) FetchCustomersForGroup(groupID int) ([]Customer, error) {
	params := map[string]string{
		"id_grp": fmt.Sprintf("%d", groupID),
	}
	var customers CustomersResp
	err := c.do("getKundenGruppe", params, &customers)
	return customers.Customers, err
}
