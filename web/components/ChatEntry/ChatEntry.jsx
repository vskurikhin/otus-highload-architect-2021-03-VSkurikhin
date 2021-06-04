import * as React from 'react';

export default ({text, onChange, onSubmit}) => {
    return (
        <>
            <div className="container chatEntry">
                <input className="form-control chatEntryText"
                       placeholder="Enter a message to send"
                       type="text"
                       value={text}
                       onChange={onChange} />
                <button className="btn btn-primary" onClick={onSubmit}>Send</button>
            </div>
        </>
    )
}