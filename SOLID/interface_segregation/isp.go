package interfacesegregation

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}
type MultifunctionPriner struct {
}

func (m *MultifunctionPriner) Print(d Document) {

}
func (m *MultifunctionPriner) Fax(d Document) {

}
func (m *MultifunctionPriner) Scan(d Document) {

}

// This type does not implement Machine interface
// If we want to add it laer on, it doesn't comply with the Machine
// We are forced to implement two methods that we don't use
// Better solution is to split the Machine interface into 3 interfaces
type OldFashionPrinter struct {
}

func (m *OldFashionPrinter) Print(d Document) {

}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (p *MyPrinter) Print(d Document) {

}

// Composition of interfaces
type MultifunctionDevice interface {
	Printer
	Scanner
}

// Implements the MultifunctionDevice
type Photocopier struct{}

func (p *Photocopier) Print(d Document) {
}
func (p *Photocopier) Scan(d Document) {
}

// Decorator pattern
type MultifunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultifunctionMachine) Print(d Document) {
	m.printer.Print(d)
}
func (m MultifunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
