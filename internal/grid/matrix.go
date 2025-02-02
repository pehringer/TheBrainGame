package grid

type (
	Matrix struct {
		Layers  int
		Rows    int
		Columns int
		data    []int32
	}
	Coordinate struct {
		X int `json:"x"`
		Y int `json:"y"`
		Z int `json:"z"`
	}
	Coordinates []Coordinate
)

func New(layers, rows, columns int, elements ...int32) Matrix {
	m := Matrix{
		Layers:  layers,
		Rows:    rows,
		Columns: columns,
		data:    make([]int32, layers*rows*columns),
	}
	copy(m.data, elements)
	return m
}

func (m *Matrix) Index(layer, row, column int) *int32 {
	return &m.data[layer*m.Rows*m.Columns+row*m.Columns+column]
}

func (m *Matrix) Row(layer, row int) []int32 {
	start := layer*m.Rows*m.Columns + row*m.Columns
	stop := start + m.Columns
	return m.data[start:stop]
}
