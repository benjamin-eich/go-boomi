package datahub

import "encoding/xml"

type ModelTags struct {
	XMLName xml.Name `xml:"tags"`
	Tags    []ModelTag
}

type ModelTag struct {
	XMLName xml.Name `xml:"tag"`
	Name    string   `xml:"name,attr"`
}

type ModelFields struct {
	XMLName     xml.Name          `xml:"fields"`
	Fields      []ModelField      `xml:"field"`
	FieldGroups []ModelFieldGroup `xml:"fieldGroup"`
}

type ModelFieldGroup struct {
	XMLName        xml.Name `xml:"fieldGroup"`
	Name           string   `xml:"name,attr,omitempty"`
	UniqueID       string   `xml:"uniqueId,attr,omitempty"`
	Repeatable     bool     `xml:"repeatable,attr,omitempty"`
	IdentifyBy     string   `xml:"identifyBy,attr,omitempty"`
	CollectionTag  string   `xml:"collectionTag,attr,omitempty"`
	CollectionKeys string   `xml:"collectionKeys,attr,omitempty"`
	Required       bool     `xml:"required,attr,omitempty"`

	Fields []ModelField `xml:"field"`

	// Future-proofing: capture unknown attributes
	Attrs []xml.Attr `xml:",any,attr"`
}

// Link: https://developer.boomi.com/docs/APIs/DataHub/PlatformAPIs/Get_Model/Model_XML_representation/mdm_fields/hub-mdm_fields_element_structure_6df92ffb-2902-4df7-ae06-2062dfdd8b95
type ModelField struct {
	//
	XMLName    xml.Name `xml:"field"`
	Name       string   `xml:"name,attr"`
	Repeatable bool     `xml:"repeatable,attr"`
	Required   bool     `xml:"required,attr"`
	Type       string   `xml:"type,attr"`
	Masked     string   `xml:"masked,attr"`

	UniqueID  string `xml:"uniqueId,attr,omitempty"`
	MaxLength int    `xml:"maxLength,attr,omitempty"`

	FirstMaskCharsCount int `xml:"firstMaskCharsCount,attr,omitempty"`
	LastMaskCharsCount  int `xml:"lastMaskCharsCount,attr,omitempty"`

	// Future-proofing: capture unknown attributes
	Attrs []xml.Attr `xml:",any,attr"`

	// Future-proofing: capture unexpected nested content
	InnerXML string `xml:",innerxml"`
}

type ModelSources struct {
	XMLName        xml.Name            `xml:"sources"`
	Sources        []ModelSource       `xml:"source"`
	SourceRankings ModelSourceRankings `xml:"sourceRankings"`
}

type ModelSource struct {
	XMLName            xml.Name `xml:"source"`
	Id                 string   `xml:"id,attr"`
	AllowMultipleLinks bool     `xml:"allowMultipleLinks,attr"`
	Default            bool     `xml:"default,attr"`
	Type               string   `xml:"type,attr"`

	EntityIdUrl string `xml:"entityIdUrl,omitempty"`

	// Future-proofing: capture unknown attributes
	Attrs []xml.Attr `xml:",any,attr"`

	// Future-proofing: capture unexpected nested content
	InnerXML string `xml:",innerxml"`

	// Docs state inbound and outbound child elements
	// https://developer.boomi.com/docs/APIs/DataHub/PlatformAPIs/Get_Model/Model_XML_representation/mdm_sources/hub-mdm_sources_element_structure_8ab8a61e-f31e-4456-a2ee-01cf7225abe2
	// If needed, they can be accessed via the `InnerXML` field
}

type ModelSourceRankings struct {
	XMLName        xml.Name             `xml:"sourceRankings"`
	SourceRankings []ModelSourceRanking `xml:"sourceRanking"`
}

type ModelSourceRanking struct {
	XMLName   xml.Name                     `xml:"sourceRanking"`
	UniqueId  string                       `xml:"uniqueId"`
	SourceIds []ModelSourceRankingSourceId `xml:"sourceId"`
}

type ModelSourceRankingSourceId struct {
	XMLName  xml.Name `xml:"sourceId"`
	Rank     int      `xml:"rank,attr"`
	SourceId string   `xml:",chardata"`
}

type ModelDataQualitySteps struct {
	XMLName xml.Name               `xml:"dataQualitySteps"`
	Steps   []ModelDataQualityStep `xml:"step"`
}

// Link: https://developer.boomi.com/docs/APIs/DataHub/PlatformAPIs/Get_Model/Model_XML_representation/mdm_dataQualitySteps/hub-mdm_dataQualitySteps_element_structure_15808dce-d6e1-4139-bf7f-d9572777e5e4
type ModelDataQualityStep struct {
	XMLName xml.Name `xml:"dataQualityStep"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`

	// type=BUSINESS_RULE
	BusinessRule string `xml:"businessRule,omitempty"`

	// type=PROCESS
	ProcessId       string `xml:"processId,attr,omitempty"`
	SourceCondition string `xml:"sourceCondition,attr,omitempty"`
	FieldCondition  string `xml:"fieldCondition,attr,omitempty"`

	// Future-proofing: capture unknown attributes
	Attrs []xml.Attr `xml:",any,attr"`

	// Future-proofing: capture unexpected nested content
	InnerXML string `xml:",innerxml"`
}

type ModelRecordTitle struct {
	XMLName         xml.Name                   `xml:"recordTitle"`
	TitleParameters ModelRecordTitleParameters `xml:"titleParameters"`
}

type ModelRecordTitleParameters struct {
	XMLName    xml.Name                    `xml:"titleParameters"`
	Parameters []ModelRecordTitleParameter `xml:"parameter"`
}

type ModelRecordTitleParameter struct {
	XMLName        xml.Name `xml:"parameter"`
	UniqueId       string   `xml:"uniqueId,attr"`
	StaticConstant string   `xml:"staticConstant,attr"`
}

type ModelMatchRules struct {
	XMLName xml.Name         `xml:"matchRules"`
	Rules   []ModelMatchRule `xml:"matchRule"`
}

type ModelMatchRule struct {
	XMLName          xml.Name `xml:"matchRule"`
	TopLevelOperator string   `xml:"topLevelOperator,attr"`

	SimpleExpressions   []ModelMatchRuleSimpleExpression   `xml:"simpleExpression"`
	AdvancedExpressions []ModelMatchRuleAdvancedExpression `xml:"advancedExpression"`
	ExpressionGroups    []ModelMatchRuleExpressionGroup    `xml:"expressionGroup"`
}

type ModelMatchRuleSimpleExpression struct {
	XMLName       xml.Name `xml:"simpleExpression"`
	FieldUniqueId string   `xml:"fieldUniqueId"`
}

type ModelMatchRuleAdvancedExpression struct {
	XMLName      xml.Name                      `xml:"advancedExpression"`
	RuleOperator string                        `xml:"ruleOperator"`
	Tolerance    float64                       `xml:"tolerance"`
	FirstInput   ModelMatchRuleExpressionInput `xml:"firstInput"`
	SecondInput  ModelMatchRuleExpressionInput `xml:"secondInput"`
}

type ModelMatchRuleExpressionGroup struct {
	XMLName  xml.Name `xml:"expressionGroup"`
	Operator string   `xml:"operator,attr"`
}

type ModelMatchRuleExpressionInput struct {
	InputType string `xml:"inputType"`
	Value     string `xml:"value"`
}
