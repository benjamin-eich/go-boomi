package datahub

import "encoding/xml"

type CreateModelRequest struct {
	XMLName xml.Name `xml:"CreateModelRequest"`

	Name string `xml:"name"`

	Fields ModelFields `xml:"fields"`

	Sources ModelSources `xml:"sources"`

	DataQualitySteps ModelDataQualitySteps `xml:"dataQualitySteps"`

	RecordTitle ModelRecordTitle `xml:"recordTitle"`

	MatchRules ModelMatchRules `xml:"matchRules"`

	Tags ModelTags `xml:"tags"`
}

type UpdateModelRequest = CreateModelRequest
