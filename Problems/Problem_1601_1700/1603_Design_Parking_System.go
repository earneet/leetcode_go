package Problem_1601_1700

type ParkingSystem struct {
	big     int
	medium  int
	small   int
	bigC    int
	mediumC int
	smallC  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small, 0, 0, 0}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > this.bigC {
			this.bigC += 1
			return true
		} else {
			return false
		}
	case 2:
		if this.medium > this.mediumC {
			this.mediumC += 1
			return true
		} else {
			return false
		}
	case 3:
		if this.small > this.smallC {
			this.smallC += 1
			return true
		} else {
			return false
		}
	default:
		return false
	}
}
