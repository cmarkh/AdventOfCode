package day8

func (m Map) Starts() (starts []string) {
	for key := range m {
		if key[len(key)-1] == 'A' {
			starts = append(starts, key)
		}
	}
	return
}

type Line struct {
	CurrentPosition    string
	CurrentInstruction int
	Steps              uint64
	Map                *Map
	Instructions       *Instructions
}

func NewLine(start string, instructions *Instructions, mapIn *Map) (line Line) {
	return Line{
		CurrentPosition:    start,
		CurrentInstruction: 0,
		Steps:              0,
		Map:                mapIn,
		Instructions:       instructions,
	}
}

func (line *Line) StepOne() {
	switch line.CurrentInstruction {
	case Left:
		line.CurrentPosition = (*line.Map)[line.CurrentPosition].Left
	case Right:
		line.CurrentPosition = (*line.Map)[line.CurrentPosition].Right
	}
	line.CurrentInstruction = (line.CurrentInstruction + 1) % len(*line.Instructions)
	line.Steps++
}

func (line *Line) StepToZIter() (steps uint64) {
	for steps = 0; line.CurrentPosition[len(line.CurrentPosition)-1] != 'Z' || steps == 0; steps++ {
		line.StepOne()
	}
	return
}

func (line *Line) StepToZ(stepsToZ StepsToZ) {
	start := Start{line.CurrentPosition, line.CurrentInstruction}

	end, ok := stepsToZ[start]
	if !ok {
		steps := line.StepToZIter()
		end = End{steps, line.CurrentPosition, line.CurrentInstruction}
		stepsToZ[start] = end
	} else {
		line.Steps += end.Steps
		line.CurrentPosition = end.Position
		line.CurrentInstruction = end.InstructionIndex
	}
}

type StepsToZ map[Start]End

type Start struct {
	Position         string
	InstructionIndex int
}

type End struct {
	Steps            uint64
	Position         string
	InstructionIndex int
}

type Lines struct {
	Lines    []Line
	StepsToZ StepsToZ
}

func NewLines(instructions Instructions, mapIn Map) Lines {
	starts := mapIn.Starts()
	lines := make([]Line, len(starts))
	for i, start := range starts {
		lines[i] = NewLine(start, &instructions, &mapIn)
	}
	return Lines{lines, make(StepsToZ)}
}

func (lines *Lines) StepAll() bool {
	maxStep := uint64(0)
	for _, line := range lines.Lines {
		if line.Steps > maxStep {
			maxStep = line.Steps
		}
	}

	equal := true
	for _, line := range lines.Lines {
		if line.Steps != maxStep {
			equal = false
			break
		}
	}

	for i, line := range lines.Lines {
		if line.Steps < maxStep || equal {
			lines.Lines[i].StepToZ(lines.StepsToZ)
		}
	}

	stepsTest := lines.Lines[0].Steps
	for _, line := range lines.Lines {
		if line.Steps != stepsTest || line.CurrentPosition[len(line.CurrentPosition)-1] != 'Z' {
			return false
		}
	}
	return true
}

func Part2(instructions Instructions, mapIn Map) (steps uint64) {
	lines := NewLines(instructions, mapIn)

	for !lines.StepAll() {
	}

	return lines.Lines[0].Steps
}
