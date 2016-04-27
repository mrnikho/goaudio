package goaudio

import "fmt"
import "time"

type Oscillator struct {
	
	Frequency AudioParam
	Detune AudioParam
	OscType string
	node Node
	wave *Wave
	on bool
	timestamp time.Time
	startStamp float64
	stopStamp float64
}

func (o *Oscillator) Start(x float64){
	
	//todo
	o.startStamp = x
}

func (o *Oscillator) Stop(x float64){
	
	//todo
	o.stopStamp = x

}



func (o *Oscillator) Connect(c Component){
	
	cnode := c.getNode()
	o.node.output = c
	(*cnode).input = o
}

func (o *Oscillator) process(data *[]float32){
	
	if !o.isOn() {
		return
	}
	if o.Frequency.valueChanged() {
		o.wave.step = float64(o.Frequency.Value/44100)
	}
	switch o.OscType{
		case "SINE":
			o.getSine(data)
		case "SAW":
			o.getSaw(data)
		case "TRI":
			o.getTri(data)
		case "SQR":
			o.getSqr(data)
	}
}

func (o *Oscillator) getNode() *Node{
	
	return &o.node
}

func (o *Oscillator) isOn() bool {
	
	//todo: multiple start stops
	
	//fmt.Println(time.Since(o.timestamp).Seconds(), len(o.startStamp), len(o.stopStamp))
	//if len(o.startStamp) > 0 && time.Since(o.timestamp).Seconds() >= o.startStamp[0] {
		//o.startStamp = o.startStamp[1:]
		//o.on = true
		//return true
	//} else if len(o.stopStamp) > 0 &&time.Since(o.timestamp).Seconds() >= o.stopStamp[0] {
		//o.stopStamp = o.startStamp[1:]
			//o.on = false
			//return false
		//}
	//return false
	
	
	fmt.Println(time.Since(o.timestamp).Seconds())
	if time.Since(o.timestamp).Seconds() >= o.startStamp {
		o.on = true
		return true
	} else if time.Since(o.timestamp).Seconds() >= o.stopStamp {
			o.on = false
			return false
		}
	return false
}