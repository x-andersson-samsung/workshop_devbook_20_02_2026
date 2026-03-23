import fs from 'fs'

fs.appendFile('my-file.txt', "File was been created by Node.js", (err) => {
    if (err) throw err
    console.log("File saved!")
})

// setTimeout(() => console.log("The end"), 40000)