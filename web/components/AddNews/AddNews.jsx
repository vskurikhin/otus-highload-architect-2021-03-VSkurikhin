import './AddNews.css'

import {POST} from "../../lib/consts";

import React, {useState} from 'react'
import {Input} from 'semantic-ui-react'
import {useHistory} from "react-router-dom"

async function addNews(value) {
    return fetch('/news/add', {
        body: JSON.stringify(value),
        ...POST
    }).then(data => data.json())
}

export default function AddNews() {

    const [title, setTitle] = useState("Title")
    const [content, setContent] = useState("Content")
    const [publicAt, setPublicAt] = useState("2021-06-02 11:39:30")
    const history = useHistory()

    const handleSubmit = async e => {
        e.preventDefault()

        const token = await addNews({
            Title: title,
            Content: content,
            PublicAt: publicAt
        })
        if (token) {
            if (token.Code > 399 && token.Message) {
                history.push('/error/' + token.Message)
            } else {
                history.push('/newslist')
            }
        }
    }

    return (
        <div className="signin-wrapper">
            <form onSubmit={handleSubmit}>
                <div className="my-divTable">
                    <div className="my-divTableBody">
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <h1>For register Sign in please</h1>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>

                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Title:</p>
                                <Input type="text" name="title" value={title} onChange={e => setTitle(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>

                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Content:</p>
                                <Input type="text" name="title" value={content} onChange={e => setContent(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>

                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Public At:</p>
                                <Input type="text" name="publicAt" value={publicAt} onChange={e => setPublicAt(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>

                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <button
                                    type="submit"
                                >Submit
                                </button>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    )
}
