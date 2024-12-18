package game

type Time struct{
	currentTicks int 
	targetTicks int 
}

func NewTime(targetTicks int) *Time{
	return &Time{
		currentTicks: 0,
		targetTicks: targetTicks,
	}
}

func (t *Time) Update(){
	if (t.currentTicks < t.targetTicks){
		t.currentTicks++
	}
}

func (t *Time) Ready() bool{
	return t.currentTicks >= t.targetTicks
}

func (l *Time) Reset(){
	l.currentTicks = 0
}