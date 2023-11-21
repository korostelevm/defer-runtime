const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms))



const run = async function(){
    await sleep(5000)
    console.log('run user process')
}

run()