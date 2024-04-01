/*
Package slicesx provides generic slice utility functions.
*/
package slicesx

// Props returns a slice of a property, which is accessed using the given getter.
func Props[E, P any](vals []E, getter func(v E) P) []P {
	props := make([]P, len(vals))
	for i, v := range vals {
		props[i] = getter(v)
	}
	return props
}

// PropMap returns a map, mapping from a property, which is accessed using the given getter.
func PropMap[E any, P comparable](vals []E, getter func(v E) P) map[P]E {
	m := make(map[P]E, len(vals))
	for _, v := range vals {
		m[getter(v)] = v
	}
	return m
}

// PropsMap returns a map, mapping from a property, which is accessed using the given getter.
// It is allowed / normal that multiple elements have the same property value, so map values are slices collecting those elements.
func PropsMap[S ~[]E, E any, P comparable](vals S, getter func(v E) P) map[P]S {
	m := make(map[P]S)
	for _, v := range vals {
		p := getter(v)
		m[p] = append(m[p], v)
	}
	return m
}

// Filter returns a new slice holding only the filtered elements.
func Filter[S ~[]E, E any](vals S, f func(v E) bool) S {
	var out S
	for _, v := range vals {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}
