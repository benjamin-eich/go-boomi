package datahub

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestUnmarshalCreateModelRequest(t *testing.T) {
	input := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<mdm:CreateModelRequest xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:mdm="http://mdm.api.platform.boomi.com/">
    <mdm:name>employer</mdm:name>
    <mdm:fields>
        <mdm:field name="employer" repeatable="false" required="false" type="STRING" uniqueId="EMPLOYER" maxLength="100"/>
        <mdm:field name="name" repeatable="false" required="false" type="STRING" uniqueId="NAME" maxLength="100"/>
        <mdm:field name="isEmployed" repeatable="false" required="false" type="BOOLEAN" uniqueId="ISEMPLOYED" maxLength="100" masked="ALL"/>
        <mdm:field name="email" repeatable="false" required="false" type="STRING" uniqueId="EMAIL" maxLength="100" masked="ALL"/>
    <mdm:fieldGroup name="address" repeatable="false" required="false" uniqueId="ADDRESS">
            <mdm:field name="street" repeatable="false" required="false" type="STRING" uniqueId="STREET" masked="PARTIAL MASK" firstMaskCharsCount="2" lastMaskCharsCount ="3"/>
            <mdm:field name="city" repeatable="false" required="false" type="STRING" uniqueId="CITY"/>
    </mdm:fieldGroup>
    </mdm:fields>
    <mdm:sources>
        <mdm:source id="SF" type="Both" allowMultipleLinks="false" default="true">
            <mdm:inbound>
                <mdm:createApproval required="false"/>
                <mdm:updateApproval required="false"/>
                <mdm:updateApprovalWithBaseValue>false</mdm:updateApprovalWithBaseValue>
                <mdm:endDateApproval required="false"/>
                <mdm:earlyChangeDetectionEnabled>true</mdm:earlyChangeDetectionEnabled>
            </mdm:inbound>
            <mdm:outbound>
                <mdm:channelUpdatesFields>All</mdm:channelUpdatesFields>
                <mdm:sendCreates>true</mdm:sendCreates>
            </mdm:outbound>
        </mdm:source>
    </mdm:sources>
    <mdm:dataQualitySteps/>
      <mdm:recordTitle>
            <mdm:titleParameters>
            <mdm:parameter uniqueId="EMPLOYER"/>
            <mdm:parameter staticConstant="."/>
            <mdm:parameter uniqueId="NAME"/>
        </mdm:titleParameters>
    </mdm:recordTitle>
    <mdm:matchRules>
        <mdm:matchRule topLevelOperator="AND">
            <mdm:simpleExpression>
                <mdm:fieldUniqueId>EMPLOYER</mdm:fieldUniqueId>
            </mdm:simpleExpression>
        </mdm:matchRule>
    </mdm:matchRules>
    <mdm:tags/>
</mdm:CreateModelRequest>`

	var model CreateModelRequest

	err := xml.Unmarshal([]byte(input), &model)
	if err != nil {
		t.Errorf("Error unmarshalling: %v", err)
	}

	if model.Name != "employer" {
		t.Errorf("Expected name to be 'employer', got '%s'", model.Name)
	}
}

func TestMarshalCreateModelRequest(t *testing.T) {
	testStrings := StringsWithCounter{}
	randomNumbers := RandomLongNumbers{}
	model := CreateModelRequest{
		Name: testStrings.Get(),
		Fields: ModelFields{
			Fields: []ModelField{
				ModelField{
					Name:       testStrings.Get(),
					Repeatable: true,
					Required:   false,
					Type:       "STRING",
					UniqueID:   testStrings.Get(),
					MaxLength:  randomNumbers.Get(),
				},
				ModelField{
					Name:       testStrings.Get(),
					Repeatable: false,
					Required:   false,
					Type:       "STRING",
					UniqueID:   testStrings.Get(),
					MaxLength:  randomNumbers.Get(),
				},
			},
			FieldGroups: []ModelFieldGroup{
				ModelFieldGroup{
					Name: testStrings.Get(),
					Fields: []ModelField{
						ModelField{
							Name:       testStrings.Get(),
							Repeatable: true,
							Required:   false,
							Type:       "STRING",
							UniqueID:   testStrings.Get(),
							MaxLength:  randomNumbers.Get(),
						},
						ModelField{
							Name:       testStrings.Get(),
							Repeatable: false,
							Required:   false,
							Type:       "STRING",
							UniqueID:   testStrings.Get(),
							MaxLength:  randomNumbers.Get(),
						},
						ModelField{
							Name:                testStrings.Get(),
							Repeatable:          false,
							Required:            false,
							Type:                "STRING",
							UniqueID:            testStrings.Get(),
							MaxLength:           randomNumbers.Get(),
							FirstMaskCharsCount: randomNumbers.Get(),
							LastMaskCharsCount:  randomNumbers.Get(),
						},
					},
				},
			},
		},
		Sources:          ModelSources{},
		DataQualitySteps: ModelDataQualitySteps{},
		RecordTitle:      ModelRecordTitle{},
		MatchRules:       ModelMatchRules{},
		Tags: ModelTags{
			Tags: []ModelTag{
				ModelTag{Name: testStrings.Get()},
				ModelTag{Name: testStrings.Get()},
			},
		},
	}

	xmlBytes, err := xml.Marshal(model)
	if err != nil {
		t.Errorf("Error marshalling: %v", err)
	}
	for _, s := range testStrings.All() {
		if !strings.Contains(string(xmlBytes), s) {
			t.Errorf("Expected to find '%s' in %s", s, xmlBytes)
		}
	}
	for _, d := range randomNumbers.All() {
		if !strings.Contains(string(xmlBytes), strconv.Itoa(d)) {
			t.Errorf("Expected to find '%d' in %s", d, xmlBytes)
		}
	}
}

type Counter struct {
	val int
}

func (c *Counter) GetNext() int {
	c.val++
	return c.val
}
func (c *Counter) GetCurrent() int {
	return c.val
}

type StringsWithCounter struct {
	counter Counter
}

func (c *StringsWithCounter) formatTestString(i int) string {
	return fmt.Sprintf("testString%d", i)
}
func (c *StringsWithCounter) Get() string {
	return c.formatTestString(c.counter.GetNext())
}
func (c *StringsWithCounter) All() []string {
	result := make([]string, c.counter.GetCurrent())
	for i := 1; i <= c.counter.GetCurrent(); i++ {
		result[i-1] = c.formatTestString(i)
	}
	return result
}

type RandomLongNumbers struct {
	numbers []int
}

func (r *RandomLongNumbers) Get() int {
	number := rand.Intn(899999) + 100000
	for ; slices.Contains(r.numbers, number); number = rand.Intn(899999) + 100000 {
	}

	r.numbers = append(r.numbers, number)
	return number
}

func (r *RandomLongNumbers) All() []int {
	return r.numbers
}
