package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

func main() {
	resp := EnvelopeReq{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Cbs:     "http://www.kcs.co.th/schema/CBS",
		Body: BodyReq{
			InquiryControlFileRequest: InquiryControlFileReq{
				InquiryControlFileReqHeader: InquiryControlFileReqHeader{
					ChannelId: "12345-67890",
					TransDate: "01-02-2018",
					TransTime: "12:13",
					RefId:     "AA0VORw4CkiTKsPrTYgGLaIUg9Ufnw3N",
				},
				TableName:   "ZUTBLACOBJ",
				SelectField: "ACOBJ|DES|PCFLG",
			},
		},
	}

	s, err := xml.MarshalIndent(&resp, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	// END OMIT

	fmt.Printf("%s", string(s))

}

type EnvelopeReq struct {
	XMLName xml.Name  `xml:"soapenv:Envelope"`
	Soapenv string    `xml:"xmlns:soapenv,attr"`
	Cbs     string    `xml:"xmlns:cbs,attr"`
	Header  HeaderReq `xml:"soapenv:Header"`
	Body    BodyReq   `xml:"soapenv:Body"`
}

type HeaderReq struct {
}

type BodyReq struct {
	InquiryControlFileRequest InquiryControlFileReq `xml:"cbs:InquiryControlFileRequest"`
}

type InquiryControlFileReq struct {
	InquiryControlFileReqHeader InquiryControlFileReqHeader `xml:"cbs:InquiryControlFileReqHeader"`
	TableName                   string                      `xml:"cbs:TableName"`
	SelectField                 string                      `xml:"cbs:SelectField"`
}

type InquiryControlFileReqHeader struct {
	ChannelId string `xml:"cbs:ChannelId"`
	TransDate string `xml:"cbs:TransDate"`
	TransTime string `xml:"cbs:TransTime"`
	RefId     string `xml:"cbs:RefId"`
}

type ReqClient struct {
	ReqHeader ReqHeader `xml:"ReqHeader"`
}

type ReqHeader struct {
	ReqID      string `xml:"reqID"`
	ReqChannel string `xml:"reqChannel"`
	ReqDtm     string `xml:"reqDtm"`
	ReqBy      string `xml:"reqBy"`
	HostCif    string `xml:"hostCif"`
	Locale     string `xml:"locale"`
	Service    string `xml:"service"`
}
