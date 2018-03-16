import { observable, action } from 'mobx'
import { Observable } from 'rxjs'


const empty = {
  name: 'root',
  children: [],
}

class TreeModel {

  @observable data
  
  constructor(){
    this.data = [empty]

    this.ws = new WebSocket('ws://127.0.0.1:3000/ws')

    this.ws.onopen = () => {
    }
    this.ws.onclose = () => {
      this.ws = null
    }
    this.ws.onmessage = (e) => this.receiveData(e)
    this.ws.onerror = (evt) => {
      console.log('ERROR: ' + evt.data) // eslint-disable-line
    }

  }

  receiveData(e){
    let strData = e.data.replace(/"Name"/g, '"name"')
    strData = strData.replace(/"Children"/g, '"children"')
    var data = JSON.parse(strData)
    this.data = [
      data,
    ]
  }

  @action
  initialize(){
    Observable
      .ajax('/data')
      .map(r => r.response)
      .subscribe(l => {
        this.data = []
        this.data = l
      })
  }


}

export default TreeModel
