import 'core-js/stable';
const runtime = require('@wailsapp/runtime');

// Main entry point
function start() {    

	// Ensure the default app div is 100% wide/high
	let app = document.getElementById('app');
	app.style.width = '100%';
	app.style.height = '100%';

	// Inject html
	app.innerHTML = `
	<<!DOCTYPE html>
	<html>
		<head>
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<!-- CSS only -->
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-0evHe/X+R7YkIZDRvuzKMRqM+OrBnVFBL6DOitfPri4tjfHxaWutUpFmBp4vmVor" crossorigin="anonymous">
			<!-- JavaScript Bundle with Popper -->
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/js/bootstrap.bundle.min.js" integrity="sha384-pprn3073KE6tl6bjs2QrFaJGz5/SUsLqktiwsUTF55Jfv3qYSDhgCecCxMW52nD2" crossorigin="anonymous"></script>
			<link href="index.css" rel="stylesheet" crossorigin="anonymous">
		</head>
		<body>
			<div class="container-fluid primary">
				<div class="row pt-3 pb-3">
					<div class="col">
						<h1 id="quote" class="text-responsive">No</h1>
					</div>
				</div>
				<div class="row pt-3 pb-3">
					<div class="col">
						<!-- Icon? Settings? -->
					</div>
					<div class="col text-end">
						<h1 id="author" class="text-responsive">No</h1>
					</div>
				</div>
			</div>
		</body>
	</html>
	`;

    window.backend.DisplayQuote().then(result => 
        {
            console.log(result)
            document.getElementById('quote').innerText = result.Quote
            document.getElementById('author').innerText = result.AuthorID
        })
	
};

// We provide our entrypoint as a callback for runtime.Init
runtime.Init(start);