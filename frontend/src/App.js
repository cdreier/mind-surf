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

// const svgSquare = {
//   shape: 'rect',
//   shapeProps: {
//     width: 130,
//     height: 30,
//     x: -10,
//     y: -10,
//   },
// }

const textLayout = {
  textAnchor: 'start', 
  x: 0, 
  y: 10, 
  transform: undefined,
}

const treeStyle = {
  nodes: {
    node: {
      circle: {
        fill: '#6ab7ff',
        stroke: '#005cb2',
        strokeWidth: 1,
      },
      name: {
        fontSize: 14,
        width: 130,
        style: {
          overflow: 'hidden',
          'text-overflow': 'ellipsis',
        },
      },
    },
    leafNode: {
      circle: {
        fill: '#b2fab4',
        stroke: '#519657',
        strokeWidth: 1,
      },
      name: {
        fontSize: 14,
        width: 130,
        style: {
          overflow: 'hidden',
          'text-overflow': 'ellipsis',
        },
      },
    },
  },
}

const separation = { 
  siblings: 1.3, 
  nonSiblings: 3,
}


@inject('tree')
@observer
class App extends Component {
  render(){
    return (
      <Wrapper >
        <Tree 
          separation={separation}
          styles={treeStyle}
          textLayout={textLayout}
          data={mobx.toJS(this.props.tree.data)} />
      </Wrapper>
    )
  }
}

export default App
