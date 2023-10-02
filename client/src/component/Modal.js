import { useEffect, useState } from "react";
import './modal.css';

const Modal = ({data,onUpdate,clearModal}) => {
    const [formData,setFormData] = useState()
    useEffect(()=>{
        setFormData(data)
    },[])
    const onChange = (event) => {
        setFormData(prevState => ({
            ...prevState,
            [event.target.name] : event.target.value
        }))
    }

    const update = (event) => {
        event.preventDefault()
        onUpdate(formData)
        clearModal()
    }

    return (
        <>
            {formData && (
            <div className="modal">
                <div className="modal-body">
                    <form onSubmit={update} method="patch">
                        <label htmlFor="title">Title</label>
                        <input type="text" name="title" value={formData.title} onChange={onChange}/>
                        <br/>
                        <label htmlFor="description">Description</label>
                        <input type="text" name="description" value={formData.description} onChange={onChange}/>
                        <br/>
                        <input type="submit"/>
                    </form>
                </div>
            </div>
            )}
        </>
    )
}

export default Modal;