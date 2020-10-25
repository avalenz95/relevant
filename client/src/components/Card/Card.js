import React, { useState } from 'react'
import './Card.css'
import { useDispatch } from 'react-redux'
import { addKeywordToSub } from '../../actions'
import Tag from '../Tag/Tag.js'

// Displays a single subreddit along with associated keywords
function Card(props) {
    const { subName, username, keywords, banner } = props
    const [word, setWord] = useState("")
    const dispatch = useDispatch()

    let tags = []
    // TODO: Figure out what's going on here.
    tags = Object.entries(keywords).map(([_, word],index) => {
        return (
            <Tag word={word} key={word}/>
        )
    })

    return (
        <div className="card" key={subName}>
                <div className="container">
                    <div className="subName">
                    {subName}
                    </div>
                    
                    <div className="tagGrid">
                        {tags}
                    </div>
                    <div className="addWord">
                        <form onSubmit={e => {
                            e.preventDefault()
                            dispatch(addKeywordToSub(subName, username, word))
                            setWord("")
                        }}>
    

                            <input 
                                type="search"
                                name="word" 
                                value={word}
                                placeholder="add"
                                onChange={e => setWord(e.target.value)}
                            />
                            {/* <input name="submit" type="submit" placeholder='Submit'/> */}
                        </form>
                    </div>
                </div>
                {banner !== "" ? <img src={banner} loading="lazy" alt=""/>: "" }
            </div>
    )
}

export default Card