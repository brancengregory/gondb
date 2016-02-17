package gondb

import "testing"

const APIKEY = "DEMO_KEY"

func Test_GetFoodReport(t *testing.T) {
	api := NewClient(nil, APIKEY)
	p := &Parameters{
		NdbNo: "01009",
		Type:  "f",
	}
	_, err := api.GetFoodReport(p)
	if err != nil {
		t.Fatal(err)
	}

	p.Type = "s"
	_, err = api.GetFoodReport(p)
	if err != nil {
		t.Fatal(err)
	}

	p.Type = "b"
	_, err = api.GetFoodReport(p)
	if err != nil {
		t.Fatal(err)
	}

}