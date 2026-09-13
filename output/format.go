// Package output handles formatting of ant movements into the
// per-turn "Lx-y Lz-w ..." output format required by lem-in.
package output

import (
	"fmt"
	"sort"
	"strings"
)

// Move represents a single ant's movement during one turn:
// ant number AntID moves into room Room.
type Move struct {
	AntID int
	Room  string
}

// FormatTurn takes all the moves that happened in a single turn
// and returns the formatted line, e.g. "L1-3 L2-2".
// Ants are ordered by ascending AntID, matching the subject's examples.
// If moves is empty, it returns an empty string (no ants moved this turn).
func FormatTurn(moves []Move) string {
	if len(moves) == 0 {
		return ""
	}

	// Copy before sorting so we never mutate the caller's slice.
	sorted := make([]Move, len(moves))
	copy(sorted, moves)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].AntID < sorted[j].AntID
	})

	parts := make([]string, len(sorted))
	for i, m := range sorted {
		parts[i] = fmt.Sprintf("L%d-%s", m.AntID, m.Room)
	}

	return strings.Join(parts, " ")
}

// FormatAllTurns takes the full schedule (one []Move per turn, in order)
// and returns the formatted lines ready to be printed, one per turn.
// Turns where no ant moved are skipped, since the subject only asks us
// to display turns where movement happened.
func FormatAllTurns(turns [][]Move) []string {
	lines := make([]string, 0, len(turns))
	for _, turn := range turns {
		line := FormatTurn(turn)
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}
