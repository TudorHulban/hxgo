package helpers

import (
	"unsafe"
)

type headerString struct {
	ptr unsafe.Pointer
	len int
}

type headerSlice struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func ptrString(s string) unsafe.Pointer {
	return (*headerString)(unsafe.Pointer(&s)).ptr
}

//go:linkname runtime_environ runtime.environ
var runtime_environ []string

func IsRunningInCI() bool {
	envs := *(*[]string)(unsafe.Pointer(&runtime_environ))

	varsCI := []struct {
		ptr unsafe.Pointer
		len int
	}{
		{ptrString("CI="), 3},
		{ptrString("GITHUB_ACTIONS="), 15},
		{ptrString("GITLAB_CI="), 10},
		{ptrString("TRAVIS="), 7},
		{ptrString("CIRCLECI="), 9},
		{ptrString("JENKINS_URL="), 13},
		{ptrString("BUILDKITE="), 10},
		{ptrString("TEAMCITY_VERSION="), 18},
	}

	sliceHdr := (*headerSlice)(unsafe.Pointer(&envs))

	for i := 0; i < sliceHdr.len; i++ {
		strPtr := *(**headerString)(unsafe.Add(sliceHdr.data, i*int(unsafe.Sizeof(uintptr(0)))))

		for _, v := range varsCI {
			if strPtr.len > v.len {
				match := true

				for j := 0; j < v.len; j++ {
					if *(*byte)(unsafe.Add(strPtr.ptr, j)) != *(*byte)(unsafe.Add(v.ptr, j)) {
						match = false

						break
					}
				}

				if match {
					return true
				}
			}
		}
	}

	return false
}
