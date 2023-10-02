import { useState } from "react";
import Modal from "./Modal";

const Table = ({data, onDelete, onUpdate}) => {

    const [showModal, setShowModal] = useState()
    return (
        <>
            {showModal && <Modal data={showModal} onUpdate={onUpdate} clearModal ={() => setShowModal()}/>}
            <table>
                <thead>
                    <tr>
                        {Object.keys(data[0]).map(x=>(
                            <th>
                                {x}
                            </th>
                        ))}
                    </tr>
                </thead>
                <tbody>
                    {data.map(x=>(
                        <tr>
                            {Object.keys(x).map(key => (
                                <td>
                                    {x[key]}
                                </td>
                            ))}
                            <td>
                                <button onClick={() => setShowModal(x)}>
                                    Update
                                </button>
                            </td>
                            <td>
                                <button type="button" onClick={() => onDelete(x["id"])}>
                                    Delete
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </>
    )
}

export default Table;