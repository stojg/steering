package steering

// Steering is the interface for all steering behaviour
type Steering interface {
	Get() *SteeringOutput
}

//func NewFlee(character, target *BodyComponent) *Flee {
//	return &Flee{
//		character: character,
//		target:    target,
//	}
//}

//// Flee makes the character to flee from the target
//type Flee struct {
//	character *BodyComponent
//	target *BodyComponent
//}

//// GetSteering returns a linear steering
//func (s *Flee) Get() *SteeringOutput {
//	steering := &SteeringOutput{}
//	steering.linear = s.character.Position.NewSub(s.target.Position)
//	steering.linear.Normalize()
//	steering.linear.NewHadamardProduct(s.character.MaxAcceleration)
//	steering.angular = &vector.Vector3{}
//	return steering
//}

//// Wander lets the character wander around
//type Wander struct {
//	character *BodyComponent
//	// Holds the radius and offset of the wander circle. The
//	// offset is now a full 3D vector
//	offset         *vector.Vector3
//	WanderRadiusXZ float64
//	WanderRadiusY  float64
//
//	// holds the maximum rate at which the wander orientation
//	// can change. Should be strictly less than 1/sqrt(3) = 0.577
//	// to avoid the chance of ending up with a zero length wander vector
//	rate float64
//
//	// Holds the current offset of the wander target
//	Vector *vector.Vector3
//
//	// holds the max acceleration for this character, this
//	// again should be a 3D vector, typically with only a
//	// non zero z value
//	maxAcceleration *vector.Vector3
//}
//
//// NewWander returns a new Wander behaviour
//func NewWander(character *BodyComponent, offset, radiusXZ, radiusY, rate float64) *Wander {
//	w := &Wander{}
//	w.character = character
//	w.offset = &vector.Vector3{offset, 0, 0}
//	w.WanderRadiusXZ = radiusXZ
//	w.WanderRadiusY = radiusY
//	w.rate = rate
//
//	w.maxAcceleration = &vector.Vector3{1, 0, 0}
//	// start by wandering straight forward
//	w.Vector = &vector.Vector3{1, 0, 0}
//
//	return w
//}
//
//// GetSteering returns a new linear and angular steering for wander
//func (wander *Wander) Get() *SteeringOutput {
//
//	// 1. make a target that looks ahead
//	charOffset := wander.character.Position.NewAdd(wander.offset.NewRotate(wander.character.Orientation))
//	target := NewEntity()
//	target.Position.Add(charOffset)
//
//	// 2. randomise the wander vector a bit, this represents the "small" sphere at the center of the
//	// target
//	wander.Vector[0] += (wander.randomBinomial() * wander.rate)
//	wander.Vector[1] += (wander.randomBinomial() * wander.rate)
//	wander.Vector[2] += (wander.randomBinomial() * wander.rate)
//	wander.Vector.Normalize()
//
//	// 3. offset the target with the scaled "small" sphere
//	target.Position[0] += wander.Vector[0] * wander.WanderRadiusXZ
//	target.Position[1] += wander.Vector[1] * wander.WanderRadiusY
//	target.Position[2] += wander.Vector[2] * wander.WanderRadiusXZ
//
//	// 4. Delegate to face
//	face := NewFace(wander.character, target)
//
//	// 5. Now set the linear acceleration to be at full
//	// acceleration in the direction of the orientation
//	steering := face.Get()
//
//	steering.linear = wander.maxAcceleration.NewRotate(wander.character.Orientation)
//
//	return steering
//}
//
//// randomBinomial get a random number between -1 and + 1
//func (s *Wander) randomBinomial() float64 {
//	return rand.Float64() - rand.Float64()
//}
//
//func NewFollowPath(character *BodyComponent, path *Path) *FollowPath {
//	return &FollowPath{
//		character:    character,
//		path:         path,
//		pathOffset:   1,
//		currentParam: 0,
//	}
//}
//
//type FollowPath struct {
//	character    *BodyComponent
//	path         *Path
//	pathOffset   int
//	currentParam int
//}
//
//func (follow *FollowPath) Get() *SteeringOutput {
//
//	// find the current position on the path
//	follow.currentParam = follow.path.getParam(follow.character.Position, follow.currentParam)
//
//	// offset it
//	targetParam := follow.currentParam + follow.pathOffset
//
//	target := NewEntity()
//	target.Position = follow.path.getPosition(targetParam)
//
//	seek := NewSeek(follow.character, target)
//	return seek.Get()
//}
//
//type Path struct {
//	points []*vector.Vector3
//}
//
//func (p *Path) getParam(position *vector.Vector3, lastparam int) int {
//	closest := 0
//	distance := math.MaxFloat64
//	for i := range p.points {
//		sqrDist := position.NewSub(p.points[i]).SquareLength()
//		if sqrDist < distance {
//			closest = i
//			distance = sqrDist
//		}
//	}
//	return closest
//}
//
//func (p *Path) getPosition(param int) *vector.Vector3 {
//	if param > len(p.points)-1 {
//		param = len(p.points) - 1
//	}
//	if param < 0 {
//		param = 0
//	}
//	if len(p.points) == 0 {
//		Println("Getting a request for a Path.getPosition when Path.points is empty")
//		return &vector.Vector3{0, 0, 0}
//	}
//
//	//fmt.Println(param, len(p.points))
//
//	return p.points[param]
//}
