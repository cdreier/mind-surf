import { observable, action } from 'mobx'
import { Observable } from 'rxjs'

class TreeModel {

  @observable data
  
  constructor(){
    this.data = [
      {
        name: 'Top Level',
        attributes: {
          keyA: 'val A',
          keyB: 'val B',
          keyC: 'val C',
        },
        children: [
          {
            name: 'Level 2: A',
            attributes: {
              keyA: 'val A',
              keyB: 'val B',
              keyC: 'val C',
            },
          },
          {
            name: 'Level 2: B',
          },
        ],
      },
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
