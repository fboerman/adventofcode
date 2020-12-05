package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	num    int
	recipe *Recipe
}

type Recipe struct {
	name       string
	num_output int
	inputs     []*Input
}

func is_storage_empty(storage map[string]int) bool {
	for k, v := range storage {
		if k != "ORE" && v != 0 {
			return false
		}
	}

	return true
}

func execute_recipe(r *Recipe, num int, storage map[string]int) map[string]int {
	// calculate how much we need when taking into account stock
	num_required := num - storage[r.name]
	// round up since we can only execute the recipe in round integers
	runs := int(math.Ceil(float64(num_required) / float64(r.num_output)))

	if runs == 0 {
		storage[r.name] -= num
		return storage
	}
	// execute all input recipes
	for _, inp := range r.inputs {
		if inp.recipe.name == "ORE" {
			storage["ORE"] += runs * inp.num
		} else {
			storage = execute_recipe(inp.recipe, runs*inp.num, storage)
		}
	}
	storage[r.name] = storage[r.name] + runs*r.num_output - num
	//fmt.Printf("%dx%s\n", num, r.name)
	//fmt.Println(storage)
	return storage
}

func main() {
	file, err := os.Open("day14_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	recipe_book := map[string]*Recipe{}

	// parse the recipes and build a tree structure
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " => ")
		parts_out := strings.Fields(parts[1])
		num_output, _ := strconv.Atoi(parts_out[0])
		r := recipe_book[parts_out[1]]
		if r == nil {
			r = &Recipe{
				name: parts_out[1],
			}
			recipe_book[parts_out[1]] = r
		}
		r.num_output = num_output

		var parts_ins []string
		if !strings.Contains(parts[0], "ORE") {
			parts_ins = strings.Split(parts[0], ",")
		} else {
			parts_ins = []string{parts[0]}
		}
		for _, input := range parts_ins {
			parts_in := strings.Fields(input)
			r_i := recipe_book[parts_in[1]]
			if r_i == nil {
				r_i = &Recipe{
					name: parts_in[1],
				}
				recipe_book[parts_in[1]] = r_i
			}
			num_input, _ := strconv.Atoi(parts_in[0])
			i := &Input{num_input, r_i}
			r.inputs = append(r.inputs, i)
		}

	}

	fmt.Println("num recipes:", len(recipe_book)-1)

	storage := map[string]int{} // storage of left over material, one exception: ORE is the total ORE needed
	storage = execute_recipe(recipe_book["FUEL"], 1, storage)
	ore_per_fuel := storage["ORE"]
	fmt.Println("ORE used for 1 FUEL:", ore_per_fuel)
	storage = map[string]int{}
	// put in the lower bound (at least 1^12/ore per fuel)
	// then iterate up towards the actual number
	fuel := int(math.Pow10(12) / float64(ore_per_fuel))
	storage = execute_recipe(recipe_book["FUEL"], fuel, storage)
	for {
		execute_recipe(recipe_book["FUEL"], 1, storage)
		if storage["ORE"] >= int(math.Pow10(12)) {
			break
		} else {
			fuel += 1
		}
	}

	fmt.Printf("Produced %d fuel", fuel)
}
