package datahub

import "encoding/xml"

type CreateModelRequest struct {
	XMLName xml.Name `xml:"CreateModelRequest"`

	Name string `xml:"name,omitempty"`

	Fields ModelFields `xml:"fields"`

	Sources ModelSources `xml:"sources,omitempty"`

	DataQualitySteps ModelDataQualitySteps `xml:"dataQualitySteps,omitempty"`

	RecordTitle ModelRecordTitle `xml:"recordTitle,omitempty"`

	MatchRules ModelMatchRules `xml:"matchRules,omitempty"`

	Tags ModelTags `xml:"tags,omitempty"`
}

type UpdateModelRequest = CreateModelRequest
