package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	Name string
	A    [2]int
	B    [2]int
}

type Ticket struct {
	Values []int
	Valid  bool
}

func parse_ticket_values(line string) []int {
	var data []int
	for _, v := range strings.Split(line, ",") {
		V, _ := strconv.Atoi(v)
		data = append(data, V)
	}
	return data
}

// return the ticket scanning error rate
func validate_ticket(t *Ticket, fields *[]Field) int {
	error_rate := 0

	for _, v := range t.Values {
		valid := false
		for _, f := range *fields {
			if (v >= f.A[0] && v <= f.A[1]) || (v >= f.B[0] && v <= f.B[1]) {
				valid = true
				break
			}
		}
		if !valid {
			error_rate += v
		}
	}

	if error_rate == 0 {
		t.Valid = true
	}

	return error_rate
}

func remove(F []Field, f Field) []Field {
	if len(F) == 0 {
		return F
	}
	for i, f_ := range F {
		if f_.Name == f.Name {
			F[i] = F[len(F)-1]
			break
		}
	}
	return F[:len(F)-1]
}

func main() {
	file, err := os.Open("day16_input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	state := 0 // 0: reading fields 1: reading your ticket 2: reading other tickets
	var fields []Field
	your_ticket := Ticket{}
	var other_tickets []Ticket

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			state++
			scanner.Scan() // skip the next line because its a header
			continue
		}
		switch state {
		case 0:
			parts := strings.Split(line, ":")
			var a1, a2, b1, b2 int
			_, _ = fmt.Sscanf(strings.TrimSpace(parts[1]), "%d-%d or %d-%d", &a1, &a2, &b1, &b2)
			f := Field{
				Name: strings.TrimSpace(parts[0]),
				A:    [2]int{a1, a2},
				B:    [2]int{b1, b2},
			}
			fields = append(fields, f)
			break
		case 1:
			your_ticket.Values = parse_ticket_values(line)
			break
		case 2:
			other_tickets = append(other_tickets, Ticket{
				Values: parse_ticket_values(line),
			})
			break
		}
	}
	ticket_scanning_error_rate := 0

	values_per_field := make([][]int, len(fields))
	for _, t := range other_tickets {
		ticket_scanning_error_rate += validate_ticket(&t, &fields)
		if !t.Valid {
			continue
		}
		for i, v := range t.Values {
			values_per_field[i] = append(values_per_field[i], v)
		}
	}
	for i, v := range your_ticket.Values {
		values_per_field[i] = append(values_per_field[i], v)
	}

	fmt.Println("ticket_scanning_error_rate:", ticket_scanning_error_rate)
	possible_fields := make([][]Field, len(fields))

	// for each field in the ticket make a list of the defined fields that are valid for all values in that field
	for i, V := range values_per_field {
		for _, f := range fields {
			valid := true
			for _, v := range V {
				if !((v >= f.A[0] && v <= f.A[1]) || (v >= f.B[0] && v <= f.B[1])) {
					valid = false
					break
				}
			}
			if valid {
				possible_fields[i] = append(possible_fields[i], f)
			}
		}
	}

	// start defining the fields, see where there is only one possibility then define that one and remove it from all other fields
	// repeat this until all fields are defined
	defined_fields := make([]string, len(fields))
	for {
		// find a field with only one posibility
		var selected_field *Field
		for i, F := range possible_fields {
			if len(F) == 1 {
				selected_field = &(F[0])
				defined_fields[i] = F[0].Name
				break
			}
		}
		if selected_field == nil {
			break
		}
		for i, F := range possible_fields {
			possible_fields[i] = remove(F, *selected_field)
		}
	}

	// map the defined fields to your ticket
	ticket_filled := map[string]int{}
	puzzleanswer := int64(0)
	for i, v := range your_ticket.Values {
		ticket_filled[defined_fields[i]] = v
		if strings.Contains(defined_fields[i], "departure") {
			if puzzleanswer == 0 {
				puzzleanswer = int64(v)
			} else {
				puzzleanswer *= int64(v)
			}
		}
	}

	fmt.Println("Solved your ticket:", ticket_filled)
	fmt.Println("Puzzleanswer:", puzzleanswer)
}
