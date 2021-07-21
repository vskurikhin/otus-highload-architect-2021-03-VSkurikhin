
export const throwResultCode = ({Code, Message}) => {
    if (Code && Message) {
        console.error('Code:' + Code + ', Message: ' + Message)
        throw {
            code: Code,
            message: Message
        }
    }
}

export default throwResultCode
