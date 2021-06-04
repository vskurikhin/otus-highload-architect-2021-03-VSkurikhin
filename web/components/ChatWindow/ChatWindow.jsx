import * as React from 'react';

export default ({messages}) => {
    return (
        <>
            <div className="container chatWindow">
                <div className="card">
                    {messages.map(m => <div>{m}</div>)}
                </div>
            </div>
        </>
    )
}