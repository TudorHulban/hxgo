package tailwindcssgeneration

import (
	"fmt"
	"os"
	"sort"
)

type SafelistConfig struct {
	Path string
}

// WriteSafelistJS writes a CommonJS module exporting a safelist array.
func WriteSafelistJS(cfg SafelistConfig, classes []string) error {
	sort.Strings(classes)

	f, err := os.Create(cfg.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "module.exports = {")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "  safelist: [")
	if err != nil {
		return err
	}

	for _, c := range classes {
		_, err = fmt.Fprintf(f, "    \"%s\",\n", c)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintln(f, "  ],")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "};")
	if err != nil {
		return err
	}

	return nil
}
