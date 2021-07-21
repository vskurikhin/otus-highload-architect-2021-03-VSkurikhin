import React, {useState, useEffect, useRef } from 'react';
import NavBar from './containers/NavBar';
import ChatWindow from './containers/ChatWindow';
import ChatEntry from './containers/ChatEntry';

export default () => {
    const [newMessage, setNewMessage] = useState("");
    const [messages, setMessages] = useState(["Test message"]);

    const socket = useRef(new WebSocket("wss://echo.websocket.org"))

    useEffect(() => {
        socket.current.onmessage = (msg) => {
            const incomingMessage = `Message from WebSocket: ${msg.data}`;
            setMessages(messages.concat([incomingMessage]));
        }
    });

    useEffect(() => () => socket.current.close(), [socket])

    const onMessageChange = (e) => {;
        setNewMessage(e.target.value);
    }

    const onMessageSubmit = () => {
        socket.current.send(newMessage);
        setNewMessage("")
    }

    return (
        <>
            <NavBar title={"WebSocket Hook Component"} />
            <ChatWindow messages={messages} />
            <ChatEntry text={newMessage}
                       onChange={onMessageChange}
                       onSubmit={onMessageSubmit}/>
        </>
    );
}