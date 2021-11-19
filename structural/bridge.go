package structural

import "fmt"

// 桥接模式用来解决什么问题？抽象层包含实现层的引用。（抽象层有接口和实现、实现层也有接口和实现，抽象层的实现包含对实现层的引用，因为可以将抽象层的大部分调用委派给实现层）
// 抽象层和实现层如果平铺组合关系，会导致类的指数级膨胀，使用桥接模式通过增加层次结构，降低类的数量。
// 桥接模式包含：实施层、抽象层

// 实现层
type printer interface {
	printFile()
}

type epson struct {}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct {}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}

// 抽象层
type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}
type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

func test(){
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}

	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	winComputer := &windows{}

	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}