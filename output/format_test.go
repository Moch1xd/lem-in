package output

import "testing"

// TestFormatTurn checks single-turn formatting against small, hand-built
// cases plus the turns from the subject's Example 2 (3 ants, straight line
// of rooms 0(start)->2->3->1(end)):
//
//	L1-2
//	L1-3 L2-2
//	L1-1 L2-3 L3-2
//	L2-1 L3-3
//	L3-1
func TestFormatTurn(t *testing.T) {
	tests := []struct {
		name  string
		moves []Move
		want  string
	}{
		{
			name:  "no moves",
			moves: []Move{},
			want:  "",
		},
		{
			name:  "single ant",
			moves: []Move{{AntID: 1, Room: "2"}},
			want:  "L1-2",
		},
		{
			name:  "two ants, already in order",
			moves: []Move{{AntID: 1, Room: "3"}, {AntID: 2, Room: "2"}},
			want:  "L1-3 L2-2",
		},
		{
			name:  "two ants, given out of order",
			moves: []Move{{AntID: 2, Room: "1"}, {AntID: 1, Room: "5"}},
			want:  "L1-5 L2-1",
		},
		{
			name: "three ants at once",
			moves: []Move{
				{AntID: 1, Room: "1"},
				{AntID: 2, Room: "3"},
				{AntID: 3, Room: "2"},
			},
			want: "L1-1 L2-3 L3-2",
		},
		{
			name:  "last ant finishing alone",
			moves: []Move{{AntID: 3, Room: "1"}},
			want:  "L3-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatTurn(tt.moves)
			if got != tt.want {
				t.Errorf("FormatTurn(%+v) = %q, want %q", tt.moves, got, tt.want)
			}
		})
	}
}

// TestFormatAllTurns_SubjectExamples reproduces the full expected output of
// all three worked examples from the subject, end to end.
func TestFormatAllTurns_SubjectExamples(t *testing.T) {
	tests := []struct {
		name  string
		turns [][]Move
		want  []string
	}{
		{
			// Example 1: 3 ants, farm with rooms 0(end),1..7(##start=1),
			// expected output:
			//   L1-3 L2-2
			//   L1-4 L2-5 L3-3
			//   L1-0 L2-6 L3-4
			//   L2-0 L3-0
			name: "Example 1",
			turns: [][]Move{
				{{AntID: 1, Room: "3"}, {AntID: 2, Room: "2"}},
				{{AntID: 1, Room: "4"}, {AntID: 2, Room: "5"}, {AntID: 3, Room: "3"}},
				{{AntID: 1, Room: "0"}, {AntID: 2, Room: "6"}, {AntID: 3, Room: "4"}},
				{{AntID: 2, Room: "0"}, {AntID: 3, Room: "0"}},
			},
			want: []string{
				"L1-3 L2-2",
				"L1-4 L2-5 L3-3",
				"L1-0 L2-6 L3-4",
				"L2-0 L3-0",
			},
		},
		{
			// Example 2: 3 ants, straight line ##start(0)->2->3->##end(1),
			// expected output:
			//   L1-2
			//   L1-3 L2-2
			//   L1-1 L2-3 L3-2
			//   L2-1 L3-3
			//   L3-1
			name: "Example 2",
			turns: [][]Move{
				{{AntID: 1, Room: "2"}},
				{{AntID: 1, Room: "3"}, {AntID: 2, Room: "2"}},
				{{AntID: 1, Room: "1"}, {AntID: 2, Room: "3"}, {AntID: 3, Room: "2"}},
				{{AntID: 2, Room: "1"}, {AntID: 3, Room: "3"}},
				{{AntID: 3, Room: "1"}},
				{}, // a stray empty turn should not produce a blank line
			},
			want: []string{
				"L1-2",
				"L1-3 L2-2",
				"L1-1 L2-3 L3-2",
				"L2-1 L3-3",
				"L3-1",
			},
		},
		{
			// Example 3: 3 ants, ##start(0) with two disjoint paths to ##end(1),
			// expected output:
			//   L1-2 L2-3
			//   L1-1 L2-1 L3-2
			//   L3-1
			name: "Example 3",
			turns: [][]Move{
				{{AntID: 1, Room: "2"}, {AntID: 2, Room: "3"}},
				{{AntID: 1, Room: "1"}, {AntID: 2, Room: "1"}, {AntID: 3, Room: "2"}},
				{{AntID: 3, Room: "1"}},
			},
			want: []string{
				"L1-2 L2-3",
				"L1-1 L2-1 L3-2",
				"L3-1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatAllTurns(tt.turns)

			if len(got) != len(tt.want) {
				t.Fatalf("FormatAllTurns() returned %d lines, want %d\ngot:  %v\nwant: %v",
					len(got), len(tt.want), got, tt.want)
			}

			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("line %d = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}
