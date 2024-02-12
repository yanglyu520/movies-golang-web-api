package debugutils

import "runtime/debug"

func CommitSHA() string {
	if v, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range v.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}
