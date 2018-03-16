import React, { Component } from 'react'
import { inject, observer } from 'mobx-react'
import styled from 'styled-components'
import Tree from 'react-d3-tree'
import * as mobx from 'mobx'

const Wrapper = styled.div`
  width: 100%;
  height: 100%;
  top: 0px;
  bottom: 0px;
  left: 0px;
  right: 0px;
  position: absolute;
`

@inject('tree')
@observer
class App extends Component {
  render(){
    return (
      <Wrapper >
        <Tree data={mobx.toJS(this.props.tree.data)} />
      </Wrapper>
    )
  }
}

export default App
