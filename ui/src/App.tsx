import {useEffect, useState} from 'react'
import './App.css'
import {Board} from "./schema/board.ts";

function App() {
    const [board, setBoard] = useState<Board | null>(null);

    useEffect(() => {
        console.time('fetch')
        fetch('http://localhost:8086/board')
            .then(res => res.arrayBuffer())
            .then(buf => {
                console.timeEnd('fetch')
                console.time('decode')
                const deserializeBinary = Board.decode(new Uint8Array(buf))
                setBoard(deserializeBinary);
                console.timeEnd('decode')
            })
    }, [])

    if (!board)
        return <div>Loading...</div>

    return (
        <>
            <h1>Board</h1>
            <p>Width: {board.width}</p>
            <p>Height: {board.height}</p>
        </>
    )
}

export default App
