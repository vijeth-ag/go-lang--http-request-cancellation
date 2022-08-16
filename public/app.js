let controller = new AbortController();
let signal = controller.signal;


function getMovies(){
    console.log('getMovies');

    fetch('<go_server_url>/movie/pulpfiction/ticket/actions/search', {
        signal: signal
    })
    .then(response => {
         response.json().then(function(result){
            console.log(result);
         })
    })
    .then(json =>{
        console.log(json)
    }) 
}

function cancelGetMovies(){
    controller.abort();
    console.log('cancelGetMovies')    
}