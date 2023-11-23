// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"sort"

	"go.opentelemetry.io/collector/pdata/internal"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
)

// ResourceSpansSlice logically represents a slice of ResourceSpans.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewResourceSpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpansSlice struct {
	orig  *[]*otlptrace.ResourceSpans
	state *internal.State
}

func newResourceSpansSlice(orig *[]*otlptrace.ResourceSpans, state *internal.State) ResourceSpansSlice {
	return ResourceSpansSlice{orig: orig, state: state}
}

// NewResourceSpansSlice creates a ResourceSpansSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewResourceSpansSlice() ResourceSpansSlice {
	orig := []*otlptrace.ResourceSpans(nil)
	state := internal.StateMutable
	return newResourceSpansSlice(&orig, &state)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewResourceSpansSlice()".
func (es ResourceSpansSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es ResourceSpansSlice) At(i int) ResourceSpans {
	return newResourceSpans((*es.orig)[i], es.state)
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ResourceSpansSlice can be initialized:
//
//	es := NewResourceSpansSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es ResourceSpansSlice) EnsureCapacity(newCap int) {
	es.state.AssertMutable()
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlptrace.ResourceSpans, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty ResourceSpans.
// It returns the newly added ResourceSpans.
func (es ResourceSpansSlice) AppendEmpty() ResourceSpans {
	es.state.AssertMutable()
	*es.orig = append(*es.orig, &otlptrace.ResourceSpans{})
	return es.At(es.Len() - 1)
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ResourceSpansSlice) MoveAndAppendTo(dest ResourceSpansSlice) {
	es.state.AssertMutable()
	dest.state.AssertMutable()
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ResourceSpansSlice) RemoveIf(f func(ResourceSpans) bool) {
	es.state.AssertMutable()
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.orig = (*es.orig)[:newLen]
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es ResourceSpansSlice) CopyTo(dest ResourceSpansSlice) {
	dest.state.AssertMutable()
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newResourceSpans((*es.orig)[i], es.state).CopyTo(newResourceSpans((*dest.orig)[i], dest.state))
		}
		return
	}
	origs := make([]otlptrace.ResourceSpans, srcLen)
	wrappers := make([]*otlptrace.ResourceSpans, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newResourceSpans((*es.orig)[i], es.state).CopyTo(newResourceSpans(wrappers[i], dest.state))
	}
	*dest.orig = wrappers
}

// Sort sorts the ResourceSpans elements within ResourceSpansSlice given the
// provided less function so that two instances of ResourceSpansSlice
// can be compared.
func (es ResourceSpansSlice) Sort(less func(a, b ResourceSpans) bool) {
	es.state.AssertMutable()
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
}
