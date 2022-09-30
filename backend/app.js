const express = require('express')
const app = express()

app.get('/', (req,res) => {
    res.json({
        msg:'i am backend'
    })
})

app.listen(3000)