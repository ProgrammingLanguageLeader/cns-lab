package main

import (
	"flag"
	"fmt"
)

type LabArgs struct {
	Variant        uint
	Output         string
	StudentsNumber uint
	SamplesNumber  uint
}

func parseArgs() *LabArgs {
	args := LabArgs{}
	flag.UintVar(&args.Variant, "var", 1, "Номер варианта")
	flag.StringVar(&args.Output, "o", "output.csv", "Путь к выходному файлу")
	flag.UintVar(&args.StudentsNumber, "s", 30, "Количество студентов в группе")
	flag.UintVar(&args.SamplesNumber, "n", 10, "Количество генерируемых строк")
	flag.Parse()
	return &args
}

func (args LabArgs) ToString() string {
	return fmt.Sprintf(
		"Номер варианта: %d,\n"+
			"Путь к выходному файлу: %s,\n"+
			"Количество студентов в группе: %d,\n"+
			"Количество генерируемых строк: %d",
		args.Variant,
		args.Output,
		args.StudentsNumber,
		args.SamplesNumber)
}
