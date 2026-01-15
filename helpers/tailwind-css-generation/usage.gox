package tailwindcssgeneration

// ResolveUsedMethods filters the full method list based on the set of
// method names detected in the source tree. Only methods whose names
// appear in usedNames are returned. This component is independent of
// CSS source and scanning implementation.
func ResolveUsedMethods(allMethods []MethodInfo, usedNames map[string]bool) []MethodInfo {
	out := make([]MethodInfo, 0, len(usedNames))

	for _, m := range allMethods {
		if usedNames[m.Name] {
			out = append(out, m)
		}
	}

	return out
}

// ClassesFromMethods extracts the Tailwind class strings from the
// resolved method list. Duplicate classes are removed. The returned
// slice preserves no specific ordering guarantees.
func ClassesFromMethods(methods []MethodInfo) []string {
	seen := make(map[string]bool)
	out := make([]string, 0, len(methods))

	for _, m := range methods {
		if m.Class == "" {
			continue
		}
		if seen[m.Class] {
			continue
		}
		seen[m.Class] = true
		out = append(out, m.Class)
	}

	return out
}
