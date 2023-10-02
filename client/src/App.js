import logo from './logo.svg';
import './App.css';
import {useEffect, useState} from 'react';
import axios from 'axios';
import Table from './component/Table';

function App() {
  const [data,setData] = useState([]);
  const [isValidated,setIsValidated] = useState(true);
  const [formData,setFormData] = useState({
    title:"",
    description:"",
    reporter:"Myself",
    watchers:["Myself"]
  })

  useEffect(()=>{
    if(isValidated){
      axios.get('http://localhost:8080/issues').then(res => {
        setIsValidated(false)
        setData(res.data.data)
      })
    }
  },[isValidated]);

  const deleteIssue = (id) => {
    axios.delete(`http://localhost:8080/issues/${id}`).then(() => setIsValidated(true)).catch(data => alert(data))
  }

  const updateIssue = (issue) => {
    axios.patch(`http://localhost:8080/issues/${issue.id}`,JSON.stringify(issue)).then(() => setIsValidated(true))
  }

  const addIssue = (issue) => {
    axios.post(`http://localhost:8080/issues`,JSON.stringify(issue)).then(() => setIsValidated(true))
  }

  const updateForm = (event) => {
    setFormData((prev) => ({
      ...prev,
      [event.target.name]:event.target.value,
    }))
  }

  const addData = (event) => {
    event.preventDefault()
    addIssue(formData)
    setData([...data],formData)
  }

  return (
    <header>
      <h1>Issue</h1>
      <form onSubmit={addData} method='post'>
        <label htmlFor ="title">Title</label>
        <input type='text' id='title' value={formData.title} name ="title" onChange={updateForm}/>
        <br/>
        <label htmlFor ="description">Description</label>
        <input type='text' id='description' value={formData.description} name ="description" onChange={updateForm}/>
        <br/>
        <input type='submit'/>
      </form>
      {data.length > 0 ? <Table data = {data} onDelete={deleteIssue} onUpdate={updateIssue} /> : <h1>No Data</h1>}
    </header>
  );
}

export default App;
