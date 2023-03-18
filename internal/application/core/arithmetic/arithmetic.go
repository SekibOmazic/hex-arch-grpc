package arithmetic

type Arith struct {
}

func New() *Arith {
	return &Arith{}
}

func (arith Arith) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith Arith) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith Arith) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith Arith) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
