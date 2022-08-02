import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {DisplayQuote} from "../wailsjs/go/main/App";

function App() {
    const [quoteText, setQuoteText] = useState('');
    const [author, setAuthor] = useState('');
    const updateQuote = (text: any) => setQuoteText(text);
    const updateAuthor = (text: any) => setAuthor(text);

    function displayQuote() {
        DisplayQuote().then( response => {
            var quote : any = response.quote;
            var author : any = response.author;
            updateQuote(quote);
            updateAuthor(author);
        });
    }

    return (
        <div id="App">
            <div className="container-fluid primary">
				<div className="row pt-3 pb-3">
					<div className="col">
						<h1 id="quote" className="text-responsive">{quoteText}</h1>
					</div>
				</div>
				<div className="row pt-3 pb-3">
					<div className="col">
					</div>
					<div className="col text-end">
						<h1 id="author" className="text-responsive">{author}</h1>
					</div>
				</div>
			</div>
        </div>
    )
}

export default App
