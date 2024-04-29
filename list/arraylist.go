package list

// watch the primeagen video on arryalist
type ArrayList struct {
	data   []int
	length int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		data:   []int{},
		length: 0,
	}
}

func (a *ArrayList) push(value int) {
	a.data = append(a.data, value)
	a.length++
}

func (a *ArrayList) pop() int {
	popedEl := a.data[a.length-1]
	a.data = a.data[:a.length-1]

	a.length--

	return popedEl
}

func (a *ArrayList) get(index int) int {
	if index < 0 || index >= a.length {
		return -1
	}

	el := a.data[index]

	return el
}

func (a *ArrayList) delete(index int) int {
	if index < 0 || index >= a.length {
		return -1
	}

	deletedEl := a.data[index]

	a.data = append(a.data[:index], a.data[index+1:]...)

	a.length--

	return deletedEl
}
