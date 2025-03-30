package main

import "testing"

func TestApartmentHunting(t *testing.T) {
	blocks := make([]map[string]bool, 0)

	blocks = append(blocks, map[string]bool{"gym": false, "school": true, "store": false})
	blocks = append(blocks, map[string]bool{"gym": true, "school": false, "store": false})
	blocks = append(blocks, map[string]bool{"gym": true, "school": true, "store": false})
	blocks = append(blocks, map[string]bool{"gym": false, "school": true, "store": false})
	blocks = append(blocks, map[string]bool{"gym": false, "school": true, "store": true})

	testTable := []struct {
		name   string
		blocks []map[string]bool
		reqs   []string
		exp    int
	}{
		{
			name: "happy path",
			blocks: []map[string]bool{
				{"gym": false, "school": true, "store": false},
				{"gym": true, "school": false, "store": false},
				{"gym": true, "school": true, "store": false},
				{"gym": false, "school": true, "store": false},
				{"gym": false, "school": true, "store": true},
			},
			reqs: []string{"gym", "school", "store"},
			exp:  3,
		},
		{
			name: "school and store only check distance should return 4",
			blocks: []map[string]bool{
				{"gym": false, "school": true, "store": false},
				{"gym": true, "school": false, "store": false},
				{"gym": true, "school": true, "store": false},
				{"gym": false, "school": true, "store": false},
				{"gym": false, "school": true, "store": true},
			},
			reqs: []string{"school", "store"},
			exp:  4,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := apartmentHunting(subtest.blocks, subtest.reqs)

			if res != subtest.exp {
				t.Errorf("expected %v, got %v", subtest.exp, res)
			}
		})
	}
}
