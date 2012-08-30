package goquery

// First() reduces the set of matched elements to the first in the set.
// It returns a new Selection object.
func (this *Selection) First() *Selection {
	return this.Eq(0)
}

// Last() reduces the set of matched elements to the last in the set.
// It returns a new Selection object.
func (this *Selection) Last() *Selection {
	return this.Eq(-1)
}

// Eq() reduces the set of matched elements to the one at the specified index.
// If a negative index is given, it counts backwards starting at the end of the set.
// It returns a new Selection object, and an empty Selection object if the index is invalid.
func (this *Selection) Eq(index int) *Selection {
	var l = len(this.Nodes)

	if index < 0 {
		index += l
	}
	return this.Slice(index, index+1)
}

// Slice() reduces the set of matched elements to a subset specified by a range of indices.
// At the moment, negative indices are not supported.
func (this *Selection) Slice(start int, end int) *Selection {
	return pushStack(this, this.Nodes[start:end])
}