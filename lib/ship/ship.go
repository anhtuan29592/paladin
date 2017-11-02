package ship

import (
	"github.com/anhtuan29592/battleship-ai/lib"
	"github.com/anhtuan29592/battleship-ai/lib/util"
)

type ShipAction interface {
	GetType() constant.ShipType
	GetSize(orientation constant.Orientation) lib.Size
	GetPositions(location lib.Point, orientation constant.Orientation) []lib.Point
}

type Ship struct {
	Location    lib.Point
	Orientation constant.Orientation
	Action      ShipAction
}

func (s *Ship) GetType() constant.ShipType {
	return s.Action.GetType()
}

func (s *Ship) GetSize() lib.Size {
	return s.Action.GetSize(s.Orientation)
}

func (s *Ship) GetPositions() []lib.Point {
	return s.Action.GetPositions(s.Location, s.Orientation)
}

func (s *Ship) ConflictWith(other Ship) bool {
	otherPositions := other.Action.GetPositions(other.Location, other.Orientation)
	myPositions := s.Action.GetPositions(s.Location, s.Orientation)

	for i := 0; i < len(otherPositions); i++ {
		for j := 0; j < len(myPositions); j++ {
			if otherPositions[i].X == myPositions[j].X && otherPositions[i].Y == myPositions[j].Y {
				return true
			}
		}
	}

	return false
}

func (s *Ship) UpdateLocation(orientation constant.Orientation, point lib.Point) {
	s.Location = point
	s.Orientation = orientation
}

func (s *Ship) IsValid(boardSize lib.Size) bool {
	if &s.Location == nil {
		return false
	}

	if s.Location.X < 0 || s.Location.Y < 0 {
		return false
	}

	if s.Location.X >= boardSize.Width || s.Location.Y >= boardSize.Height {
		return false
	}

	size := s.Action.GetSize(s.Orientation)

	if s.Location.X+size.Width > boardSize.Width || s.Location.Y+size.Height > boardSize.Height {
		return false
	}

	return true
}
