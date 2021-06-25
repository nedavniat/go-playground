module example.com/hw1

go 1.16

replace example.com/chess => ../chess

require (
	example.com/card v0.0.0-00010101000000-000000000000
	example.com/chess v0.0.0-00010101000000-000000000000
	example.com/countnums v0.0.0-00010101000000-000000000000
	example.com/fib v0.0.0-00010101000000-000000000000
)

replace example.com/countnums => ../countnums

replace example.com/card => ../card

replace example.com/fib => ../fib
