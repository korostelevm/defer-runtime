const { spawn } = require('child_process')

const lambdaHandler = async (event, context) => {
    // console.log("EVENT", JSON.stringify(event))
    // execute user process as subprocess
    const proc = spawn('node', ['./lambda/user_proc.js'], {stdio: 'pipe'})
    let proc_res = null
    proc.stdout.on('data', (data) => {
        proc_res = data.toString()
    })

    await new Promise((resolve, reject) => {
        
        proc.on('close', (code) => {
            console.log(`child process exited with code ${code}`)
            resolve()
        })
    })



    return {
        statusCode: 200,
        body: JSON.stringify({
            message: proc_res,
        }),
    }
}


module.exports = {
    lambdaHandler,
}


