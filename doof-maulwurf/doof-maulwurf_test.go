func Testfunc Testdoofmaulwurf(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, false},
		{[]int{2}, true},
		{[]int{88}, true},
		{[]int{13877}, false},
		{[]int{56, 2, 8}, true},
		{[]int{7, 9, 15}, false},
		{[]int{48, 100, 9}, true},
		{[]int{33, 11, 2}, true},
		{[]int{4, 9, 3}, true},
		{[]int{3, 7, 2}, true},
		{[]int{-3, 1, 8}, true},
		{[]int{-1, -3, -5}, false},
		{[]int{-2, -3, 6}, true},
		{[]int{-1, -5, 11}, false},
	} {
		got := AnyIsEven(tc.s)
		if got != tc.want {
			t.Errorf("Negative integer(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}