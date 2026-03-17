package html

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		name           string
		categories     []GroceryCategory
		wantedContains []string
		wantedExcludes []string
	}{
		{
			name:       "empty_categories",
			categories: []GroceryCategory{},
			wantedContains: []string{
				"<!doctype html>",
				"<html lang=\"en\">",
				"<title>Weekly Meals</title>",
				"<h1>Meals for this week:</h1>",
				"<th>Grocery Type</th>",
				"<th>Ingredients</th>",
			},
		},
		{
			name: "single_category",
			categories: []GroceryCategory{
				{
					Name: "Produce",
					Ingredients: map[string]float64{
						"Apples":  5,
						"Bananas": 3,
					},
				},
			},
			wantedContains: []string{
				"<td>Produce</td>",
				"Apples",
				"Bananas",
				"5",
				"3",
			},
		},
		{
			name: "multiple_categories",
			categories: []GroceryCategory{
				{
					Name: "Produce",
					Ingredients: map[string]float64{
						"Apples": 5,
					},
				},
				{
					Name: "Dairy",
					Ingredients: map[string]float64{
						"Milk":   1,
						"Cheese": 0.5,
					},
				},
			},
			wantedContains: []string{
				"<td>Produce</td>",
				"<td>Dairy</td>",
				"Apples",
				"Milk",
				"Cheese",
				"0.5",
			},
		},
		{
			name: "special_characters_escaped",
			categories: []GroceryCategory{
				{
					Name: "<script>alert('xss')</script>",
					Ingredients: map[string]float64{
						"Milk & Cream": 1,
					},
				},
			},
			wantedContains: []string{
				"&lt;script&gt;",
				"Milk &amp; Cream",
			},
			wantedExcludes: []string{
				"<script>alert",
				"Milk & Cream",
			},
		},
		{
			name: "float_formatting",
			categories: []GroceryCategory{
				{
					Name: "Test",
					Ingredients: map[string]float64{
						"Integer":    5,
						"Decimal":    1.5,
						"Long":       3.14159,
						"Zero":       0,
						"Negative":   -2.5,
						"Large":      1000.5,
						"Scientific": 1e-06,
					},
				},
			},
			wantedContains: []string{
				"5",
				"1.5",
				"3.14159",
				"0",
				"-2.5",
				"1000.5",
				"0.000001",
			},
		},
		{
			name: "empty_ingredients",
			categories: []GroceryCategory{
				{
					Name:        "Empty",
					Ingredients: map[string]float64{},
				},
			},
			wantedContains: []string{
				"<td>Empty</td>",
				"<td></td>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			component := Email(tt.categories)
			err := component.Render(t.Context(), &buf)
			if err != nil {
				t.Fatalf("Email.Render() error = %v", err)
			}

			got := buf.String()

			for _, wanted := range tt.wantedContains {
				if !strings.Contains(got, wanted) {
					t.Errorf("Email.Render() output missing expected content: %q\nGot: %s", wanted, got)
				}
			}

			for _, exclude := range tt.wantedExcludes {
				if strings.Contains(got, exclude) {
					t.Errorf("Email.Render() output should not contain: %q\nGot: %s", exclude, got)
				}
			}
		})
	}
}

func TestEmailWriteToFile(t *testing.T) {
	categories := []GroceryCategory{
		{
			Name: "Produce",
			Ingredients: map[string]float64{
				"Apples":  5,
				"Bananas": 3,
				"Carrots": 2,
			},
		},
		{
			Name: "Dairy",
			Ingredients: map[string]float64{
				"Milk":   1,
				"Cheese": 0.5,
				"Yogurt": 4,
			},
		},
		{
			Name: "Meat",
			Ingredients: map[string]float64{
				"Chicken": 2,
				"Beef":    1.5,
			},
		},
	}

	var buf bytes.Buffer
	component := Email(categories)
	err := component.Render(t.Context(), &buf)
	if err != nil {
		t.Fatalf("Email.Render() error = %v", err)
	}

	outputPath := "../../output.html"
	err = os.WriteFile(outputPath, buf.Bytes(), 0644)
	if err != nil {
		t.Fatalf("Failed to write output.html: %v", err)
	}

	t.Logf("Successfully wrote email HTML to %s", outputPath)

	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read back output.html: %v", err)
	}

	if len(content) == 0 {
		t.Error("output.html is empty")
	}

	if !strings.Contains(string(content), "Produce") {
		t.Error("output.html missing expected content")
	}
}
