import React from "react";

class Item extends React.Component {
  constructor(props) {
    super(props)
    this.state = {item: {}}
  }
  componentDidMount() {
    this.setState({item: this.props.value})
  }
  render(){
    return (
      <tr>
        <td> {this.state.item.ID} </td>
        <td> {this.state.item.Name} </td>
        <td> {this.state.item.CreatedAt} </td>
        <td> {this.state.item.UpdatedAt} </td>
        <td>
          <input type="button" value="change" />
          <input type="button" value="delete" />
        </td>
      </tr>
    )
  }
}

export default Item;