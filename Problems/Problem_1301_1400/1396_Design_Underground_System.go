package Problem_1301_1400

type CheckinInfo struct {
	CheckinStation string
	CheckinTime    int
}

type StatisticsInfo struct {
	AverageTime float64
	Count       int
}

type UndergroundSystem struct {
	Customers   map[int]*CheckinInfo
	AverageTime map[string]*StatisticsInfo
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{make(map[int]*CheckinInfo), make(map[string]*StatisticsInfo)}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.Customers[id] = &CheckinInfo{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	customInfo, _ := this.Customers[id]
	stayTime := t - customInfo.CheckinTime
	delete(this.Customers, id)
	stationKey := customInfo.CheckinStation + "_" + stationName
	channelInfo, exist := this.AverageTime[stationKey]
	if !exist {
		channelInfo = &StatisticsInfo{0.0, 0}
		this.AverageTime[stationKey] = channelInfo
	}
	channelInfo.AverageTime += float64(stayTime)
	channelInfo.Count += 1
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	stationKey := startStation + "_" + endStation
	channelInfo := this.AverageTime[stationKey]
	return channelInfo.AverageTime / float64(channelInfo.Count)
}
