import './App.css';
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import Top from './components/Top';
import ItemForm from './components/ItemForm';

function App() {
  return (
    <>
      <BrowserRouter>
        <h1>Item Register</h1>
        <Link to='/'>TOP</Link>
        <Link to='/create'>CREATE</Link>
        <Routes> 
          <Route exact path='/' element={<Top/>} />
          <Route path="create" element={<ItemForm />} />
        </Routes>
          
      </BrowserRouter>
    </>
  );
}

export default App;
