package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func get_place(x int, y int, room *[]byte, width int, heigth int) *byte {
	if x < 0 || x >= width || y < 0 || y >= heigth {
		return nil
	}
	return &((*room)[y*width+x])
}

func seat_in_direction(x_ int, y_ int, room *[]byte, width int, heigth int, dir string) bool {
	x := x_
	y := y_
	for {
		switch dir {
		case "LU":
			x--
			y--
			break
		case "U":
			y--
			break
		case "RU":
			x++
			y--
			break
		case "R":
			x++
			break
		case "RD":
			x++
			y++
			break
		case "D":
			y++
			break
		case "LD":
			x--
			y++
			break
		case "L":
			x--
			break
		}
		p := get_place(x, y, room, width, heigth)
		if p == nil {
			break
		}
		switch *p {
		case '#':
			return true
		case 'L':
			return false
		}
	}

	return false
}

func num_adjecent_seats(x int, y int, room *[]byte, width int, heigth int) int {
	num := 0
	adjecent := []*byte{
		get_place(x-1, y-1, room, width, heigth),
		get_place(x, y-1, room, width, heigth),
		get_place(x+1, y-1, room, width, heigth),
		get_place(x-1, y, room, width, heigth),
		get_place(x+1, y, room, width, heigth),
		get_place(x-1, y+1, room, width, heigth),
		get_place(x, y+1, room, width, heigth),
		get_place(x+1, y+1, room, width, heigth),
	}
	for _, adj := range adjecent {
		if adj != nil {
			if *adj == '#' {
				num++
			}
		}
	}
	return num
}

func num_seeing_seats(x int, y int, room *[]byte, width int, heigth int) int {
	num := 0

	adjecent := []bool{
		seat_in_direction(x, y, room, width, heigth, "LU"),
		seat_in_direction(x, y, room, width, heigth, "U"),
		seat_in_direction(x, y, room, width, heigth, "RU"),
		seat_in_direction(x, y, room, width, heigth, "R"),
		seat_in_direction(x, y, room, width, heigth, "RD"),
		seat_in_direction(x, y, room, width, heigth, "D"),
		seat_in_direction(x, y, room, width, heigth, "LD"),
		seat_in_direction(x, y, room, width, heigth, "L"),
	}

	for _, flag := range adjecent {
		if flag {
			num++
		}
	}

	return num
}

func room_to_str(room *[]byte, width int, heigth int) string {
	str := ""
	for y := 0; y < heigth; y++ {
		for x := 0; x < width; x++ {
			str += fmt.Sprintf(" %c", *get_place(x, y, room, width, heigth))
		}
		str += "\n"
	}

	return str
}

func main() {
	file, err := os.Open("day11_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	part, _ := strconv.Atoi(os.Args[1])

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var room1 []byte

	width := 0
	heigth := 0
	for scanner.Scan() {
		line := scanner.Text()
		if width == 0 {
			width = len(line)
		}
		room1 = append(room1, []byte(line)...)
		heigth++
	}
	room2 := make([]byte, len(room1))
	copy(room2, room1)
	selected_room := true // true: read 1 write 2, false vice verse
	i := 0
	for {
		changed := false
		var read_room, write_room *[]byte
		if selected_room {
			read_room = &room1
			write_room = &room2
		} else {
			read_room = &room2
			write_room = &room1
		}

		for y := 0; y < heigth; y++ {
			for x := 0; x < width; x++ {
				r_place := get_place(x, y, read_room, width, heigth)
				if *r_place == '.' { // Floor
					continue
				}
				w_place := get_place(x, y, write_room, width, heigth)
				var num int
				if part == 1 {
					num = num_adjecent_seats(x, y, read_room, width, heigth)
				} else if part == 2 {
					num = num_seeing_seats(x, y, read_room, width, heigth)
				}
				switch *r_place {
				case 'L': // Empty seat
					if num == 0 {
						*w_place = '#'
						changed = true
					} else {
						*w_place = *r_place
					}
					break
				case '#': // Occupied seat
					if (num >= 4 && part == 1) || (num >= 5 && part == 2) {
						*w_place = 'L'
						changed = true
					} else {
						*w_place = *r_place
					}
					break
				}
			}
		}

		if !changed {
			break
		}

		i++
		selected_room = !selected_room
		//fmt.Println(i)
		//fmt.Println(room_to_str(write_room, width, heigth))
	}
	var read_room *[]byte
	if selected_room {
		read_room = &room1
	} else {
		read_room = &room2
	}
	count := 0
	for _, c := range *read_room {
		if c == '#' {
			count++
		}
	}
	fmt.Println("Part", part)
	fmt.Println("Occupied seats when stable:", count)

}
