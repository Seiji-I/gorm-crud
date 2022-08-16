import React from "react";
import axios from "axios";

class ItemForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      item: ""
    }
    this.submitHandler = this.submitHandler.bind(this)
    this.changeFieldHandler = this.changeFieldHandler.bind(this)
  }
  render(){
    return (
      <form onSubmit={this.submitHandler}>
        <label>
          Item Name:
          <input onChange={this.changeFieldHandler} type="text" name="item" value={this.state.item}/>
          <button>Register</button>
        </label>
      </form>
    )
  }
  submitHandler(e) {
    e.preventDefault();
    const formData = new FormData()
    formData.set("item", this.state.item)
    return axios.post("http://localhost:3333/api/create", formData)
    .then(response => {
      console.log(response)
    })
    .catch(err => console.log(err))
  }
  changeFieldHandler(e){
    this.setState({ item : e.target.value })
  }
}

export default ItemForm;