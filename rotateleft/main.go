package main

func findPos(pos, rot, size int) int {

	refRot := rot

	if rot > size {
		refRot = size % rot
	}

	if pos-refRot < 0 {
		return pos - refRot + size
	}

	return pos - refRot

}

func rotLeft(a []int32, rotations int32) []int32 {
	arrayLeft := make([]int32, len(a))

	for i, value := range a {

		pos := findPos(i, int(rotations), len(a))
		arrayLeft[pos] = value
	}

	return arrayLeft
}

func main() {

}
