package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
	"regexp"
	"sort"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title         string
	Deadline      time.Time
	Completed     bool
	CreatedAt     time.Time
	CompletedAt   *time.Time
	UntilDeadline time.Duration
}

// Todo slice
type Todos []Todo

// ParseComplexDuration parses a complex duration string like "1d 2h 30m" and returns a time.Duration.
func ParseComplexDuration(durationStr string) (time.Duration, error) {
	var totalDuration time.Duration
	re := regexp.MustCompile(`(\d+)([dhm])`)
	matches := re.FindAllStringSubmatch(durationStr, -1)

	for _, match := range matches {
		value, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}

		switch match[2] {
		case "d":
			totalDuration += time.Duration(value*24) * time.Hour
		case "h":
			totalDuration += time.Duration(value) * time.Hour
		case "m":
			totalDuration += time.Duration(value) * time.Minute
		}
	}

	return totalDuration, nil
}


func (todos *Todos) add(title string, durationString ...string) {
	// Define a default deadline (e.g., one month from now)
	defaultDeadline := time.Now().AddDate(0, 1, 0)

	var todoDeadline time.Time
	if len(durationString) > 0 {
		duration, err := ParseComplexDuration(durationString[0])
		if err != nil {
			fmt.Println("Invalid duration format! Use a valid format (e.g., '1d 2h 30m').")
			return
		}
		todoDeadline = time.Now().Add(duration)
	} else {
		todoDeadline = defaultDeadline
	}

	// Calculate UntilDeadline
	now := time.Now()
	untilDeadline := todoDeadline.Sub(now)

	todo := Todo{
		Title:         title,
		Deadline:      todoDeadline,
		Completed:     false,
		CompletedAt:   nil,
		CreatedAt:     now,
		UntilDeadline: untilDeadline, // Store duration
	}
	// Add to existing list using pointer
	*todos = append(*todos, todo)
}

// Method that checks whether provided index is valid
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}

	return nil // Indicates valid index
}

// Delete function
func (todos *Todos) delete(index int) error {
	t := *todos

	// Verify using helper method
	if err := t.validateIndex(index); err != nil { 
		return err
	}

	// Split at specified index and join the sections before and after
	*todos = append(t[:index], t[index+1:]...)

	return nil
}

// Toggle completion
func (todos *Todos) toggle(index int) error {
	// Dereference pointer for todo slice
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	// Check completion
	isCompleted := t[index].Completed
	// If not completed, set CompletedAt to current time
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	// Toggle status
	t[index].Completed = !isCompleted

	return nil
}

// Edit method
func (todos *Todos) edit(index int, title string) error {
	t := *todos
	
	if err := t.validateIndex(index); err != nil {
		return err
	}
	
	t[index].Title = title
	
	return nil
}

func (todos *Todos) sort(criteria string, ascending bool) error {
	t := *todos

	switch criteria {
	case "title":
		if ascending {
			sort.SliceStable(t, func(i, j int) bool {return t[i].Title < t[j].Title })
		} else {
			sort.SliceStable(t, func(i, j int) bool {return t[i].Title > t[j].Title})
		}
	case "deadline":
		if ascending {
			sort.SliceStable(t, func(i, j int) bool {return t[i].Deadline.Before(t[j].Deadline)})
		} else {
			sort.SliceStable(t, func(i, j int) bool {return t[i].Deadline.After(t[j].Deadline)})
		}
	default: 
		return fmt.Errorf("Invalid sorting criteria: %s", criteria)
	}
	return nil
}

// Format time.Duration into a more readable format
func formatDuration(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60

	return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
}


func (todos *Todos) print() {
	// Calculate the current time once
	now := time.Now()

	// Recalculate UntilDeadline for each todo
	for i := range *todos {
		(*todos)[i].UntilDeadline = (*todos)[i].Deadline.Sub(now)
	}

	// Create and configure the table for output
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Deadline", "Completed", "Created At", "Completed At", "Until Deadline")

	// Print each todo item
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(
			strconv.Itoa(index),
			t.Title,
			t.Deadline.Format(time.RFC1123),
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
			formatDuration(t.UntilDeadline),
		)
	}

	table.Render()
}

