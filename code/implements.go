// myVec is essentially treated as a parameter to Magnitude
// Now Vec2f implements the Magnituder interface
func (myVec Vec2f) Magnitude() float {
	return math.Sqrt(myVec.x * myVec.x + myVec.y * myVec.y)
}
