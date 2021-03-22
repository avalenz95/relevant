import React, { useEffect } from "react"
import Grid from '../Grid/Grid.js'
import Nav from '../Nav/Nav.js'
import { loadUsername } from "../../actions/index.js"
import { useDispatch } from 'react-redux'


function Dashboard(props) {
    const {ep} = props
    const dispatch = useDispatch() // Get the dispatcher

    // Attempt to load username on component mount
    useEffect(() => {
        dispatch(loadUsername())
    })

    return (
        <div className="dashboard">

            <Nav
                endpoint={ep} 
            />
                
            <Grid
                endpoint={ep}
            />
        </div>
    )
}

export default Dashboard