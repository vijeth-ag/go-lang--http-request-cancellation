let controller = new AbortController();
let signal = controller.signal;


function getMovies(){
    console.log('getMovies');

    fetch('https://0bbb-2401-4900-1cc4-b383-50bc-5963-df6c-939.in.ngrok.io/movie/pulpfiction/ticket/actions/search', {
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