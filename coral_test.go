package coral

import (
	"testing"
)

func TestFoundMatchInArray(t *testing.T) {
	input := `{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"ppu": 0.55,
		"batters":
			{
				"batter":
					[
						{ "id": "1001", "type": "Regular" },
						{ "id": "1002", "type": "Chocolate" },
						{ "id": "1003", "type": "Blueberry" },
						{ "id": "1004", "type": "Devil's Food" }
					]
			},
		"topping":
			[
				{ "id": "5001", "type": "None" },
				{ "id": "5002", "type": "Glazed" },
				{ "id": "5005", "type": "Sugar" },
				{ "id": "5007", "type": "Powdered Sugar" },
				{ "id": "5006", "type": "Chocolate with Sprinkles" },
				{ "id": "5003", "type": "Chocolate" },
				{ "id": "5004", "type": "Maple" }
			]
	}`

	match, err := Filter(input, "topping.type", "Powdered Sugar", false)
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}
	if !match {
		t.Log("expected match, got no match")
		t.Fail()
	}
}

func TestNoMatchInArray(t *testing.T) {
	input := `{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"ppu": 0.55,
		"batters":
			{
				"batter":
					[
						{ "id": "1001", "type": "Regular" },
						{ "id": "1002", "type": "Chocolate" }
					]
			},
		"topping":
			[
				{ "id": "5001", "type": "None" },
				{ "id": "5005", "type": "Sugar" },
				{ "id": "5007", "type": "Powdered Sugar" },
				{ "id": "5006", "type": "Chocolate with Sprinkles" }
			]
	}`

	match, err := Filter(input, "topping.type", "Blueberry", false)
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}
	if match {
		t.Log("expected no match, got match")
		t.Fail()
	}
}

func TestMatchedEmptyString(t *testing.T) {
	input := `{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"ppu": 0.55,
		"batters":
			{
				"batter":
					[
						{ "id": "1001", "type": "Regular" },
						{ "id": "1002", "type": "Chocolate" }
					]
			},
		"topping":
			[
				{ "id": "5001", "type": "None" },
				{ "id": "5005", "type": "Sugar" },
				{ "id": "5007", "type": "Powdered Sugar" },
				{ "id": "5006", "type": "" }
			]
	}`

	match, err := Filter(input, "topping.type", "", false)
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}
	if !match {
		t.Log("expected match, got no match")
		t.Fail()
	}
}

func TestMatchedBoolean(t *testing.T) {
	input := `{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"ppu": 0.55,
		"favorite": true,
		"batters":
			{
				"batter":
					[
						{ "id": "1001", "type": "Regular" },
						{ "id": "1002", "type": "Chocolate" }
					]
			},
		"topping":
			[
				{ "id": "5001", "type": "None" },
				{ "id": "5005", "type": "Sugar" },
				{ "id": "5007", "type": "Powdered Sugar" },
				{ "id": "5006", "type": "Chocolate with Sprinkles" }
			]
	}`

	match, err := Filter(input, "favorite", "true", false)
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}
	if !match {
		t.Log("expected match, got no match")
		t.Fail()
	}
}

func TestMatchedNumber(t *testing.T) {
	input := `{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"ppu": 0.55,
		"batters":
			{
				"batter":
					[
						{ "id": "1001", "type": "Regular" },
						{ "id": "1002", "type": "Chocolate" }
					]
			},
		"topping":
			[
				{ "id": "5001", "type": "None" },
				{ "id": "5005", "type": "Sugar" },
				{ "id": "5007", "type": "Powdered Sugar" },
				{ "id": "5006", "type": "Chocolate with Sprinkles" }
			]
	}`

	match, err := Filter(input, "ppu", "0.55", false)
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}
	if !match {
		t.Log("expected match, got no match")
		t.Fail()
	}
}

func TestMatchedNull(t *testing.T) {
	input := `{
		"id": "0001",
		"type": "donut",
		"name": "Cake",
		"ppu": 0.55,
		"batters":
			{
				"batter":
					[
						{ "id": "1001", "type": "Regular" },
						{ "id": "1002", "type": "Chocolate" }
					]
			},
		"topping":
			[
				{ "id": "5001", "type": null },
				{ "id": "5005", "type": "Sugar" },
				{ "id": "5007", "type": "Powdered Sugar" },
				{ "id": "5006", "type": "Chocolate with Sprinkles" }
			]
	}`

	match, err := Filter(input, "topping.type", "", true)
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}
	if !match {
		t.Log("expected match, got no match")
		t.Fail()
	}
}
