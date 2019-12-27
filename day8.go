package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

func count(arr []int, target int) (num int) {
	for _, v := range arr {
		if v == target {
			num++
		}
	}
	return
}

func flatten_layer(layer [][]int, heigth int) (layer_flat []int) {
	for y := 0; y < heigth; y++ {
		layer_flat = append(layer_flat, layer[y]...)
	}

	return
}

func main() {
	width := 25
	heigth := 6

	input_raw, _ := ioutil.ReadFile("day8_input.txt")
	//input := string(input_raw)
	//convert the read text to list of integers
	var input []int
	for _, c := range input_raw {
		number, _ := strconv.Atoi(string(c))
		input = append(input, number)
	}
	// build the layers
	var layers [][][]int
	l := 0
	for len(input) > 0 {
		//create new emtpy layer
		layers = append(layers, [][]int{})
		for y := 0; y < heigth; y++ {
			//create new empty line in layer
			layers[l] = append(layers[l], []int{})
			for x := 0; x < width; x++ {
				// pop integer from input
				c := input[0]
				input = input[1:]
				//put it on the current line in the layer
				layers[l][y] = append(layers[l][y], c)
			}
		}
		l++
	}

	// print the result
	fmt.Printf("Found %d layers\n", l)

	//for li:=0;li<l;li++{
	//	fmt.Printf("layer %d:\n", li)
	//	for y := 0; y < heigth; y++ {
	//		fmt.Println(layers[li][y])
	//	}
	//}

	// count all 0 digits in a layer
	l_selected := -1
	lowest_num_zero := math.MaxInt32
	for li, layer := range layers {
		// first flatten the layer
		layer_nums := flatten_layer(layer, heigth)
		num_zero := count(layer_nums, 0)
		if num_zero < lowest_num_zero {
			l_selected = li
			lowest_num_zero = num_zero
		}
	}
	fmt.Println("Layer with lowest number of 0: ", l_selected)
	// get the answer for part 1, multiply number of 1 digits with number of 2 digits in selected layer
	fmt.Println("Part 1: ", count(flatten_layer(layers[l_selected], heigth), 1)*count(flatten_layer(layers[l_selected], heigth), 2))

	//part 2
	//iterate through the grid, then select the pixels from all layers for that pixel and calculate its final color
	final_grid := make([][]int, heigth)
	for i := range final_grid {
		final_grid[i] = make([]int, width)
	}

	for y := 0; y < heigth; y++ {
		for x := 0; x < width; x++ {
			//iterate through all layers and find the first non transparent pixel
			for _, layer := range layers {
				if layer[y][x] != 2 {
					final_grid[y][x] = layer[y][x]
					break
				}
			}
		}
	}

	// print the resulting final_grid
	fmt.Println("part 2:")
	for y := 0; y < heigth; y++ {
		//fmt.Println(final_grid[y])
		for x:=0; x<width; x++ {
			switch final_grid[y][x] {
			case 0:
				fmt.Printf(" ")
			case 1:
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}
}
