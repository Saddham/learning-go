package printer

import (
	"fmt"
	"learning-go/models"
	"os"
	"text/tabwriter"
)

type Printer struct {
	w *tabwriter.Writer
}

func New() *Printer {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)

	return &Printer{
		w: w,
	}
}

func (p *Printer) CityHeader() {
	fmt.Fprint(p.w, "Name\tTempC\tTempF\tBeach Ready?\tMountain Ready?\n")
}

func (p *Printer) CityDetails(c models.ICity) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\t%v\t%v\n", c.Name(), c.TempC(), c.TempF(), c.HasBeach(), c.HasMountain())
}

func (p *Printer) Cleanup() {
	p.w.Flush()
}
