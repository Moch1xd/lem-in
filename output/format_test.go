package output

import "testing"

// TestFormatTurn checks single-turn formatting against small, hand-built
// cases plus the turns from the subject's Example 2 (3 ants, straight line
// of rooms 0(start)->2->3->1(end)):
//
//	L1-2
//	L1-3 L2-2
//	L1-1 L2-3 L3-2
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

// TestFormatAllTurns reproduces the full output of the subject's Example 2
// end to end, including that a fully-empty turn should be skipped.
func TestFormatAllTurns(t *testing.T) {
	turns := [][]Move{
		{{AntID: 1, Room: "2"}},
		{{AntID: 1, Room: "3"}, {AntID: 2, Room: "2"}},
		{{AntID: 1, Room: "1"}, {AntID: 2, Room: "3"}, {AntID: 3, Room: "2"}},
		{{AntID: 2, Room: "1"}, {AntID: 3, Room: "3"}},
		{{AntID: 3, Room: "1"}},
		{}, // e.g. a stray empty turn should not produce a blank line
	}

	want := []string{
		"L1-2",
		"L1-3 L2-2",
		"L1-1 L2-3 L3-2",
		"L2-1 L3-3",
		"L3-1",
	}

	got := FormatAllTurns(turns)

	if len(got) != len(want) {
		t.Fatalf("FormatAllTurns() returned %d lines, want %d\ngot:  %v\nwant: %v",
			len(got), len(want), got, want)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("line %d = %q, want %q", i, got[i], want[i])
		}
	}
}
