var QuoteBoard = React.createClass({
  mixins: [ReactFireMixin],
  getInitialState: function() {
    return {quotes: []};
  },
  componentWillMount: function() {
    var firebaseRef = new Firebase("https://quoteboard.firebaseio.com/testquotes/");
    firebaseRef = firebaseRef.orderByChild("added");
    this.bindAsArray(firebaseRef, "quotes");
  },
  handleQuoteSubmit: function(quote) {
    quote.owner = 'wkirschbaum@gmail.com';
    this.firebaseRefs["quotes"].push(quote);
  },
  render: function() {
    return (
      <div>
        <LoginForm />
        <QuoteForm onQuoteSubmit={this.handleQuoteSubmit} />
        <QuoteList quotes={this.state.quotes} />
      </div>
    );
  }
});

var LoginForm = React.createClass({
  render: function() {
    return (
      <div>LoginForm</div>
    );
  }
});

var QuoteForm = React.createClass({
  handleSubmit: function(e) {
    e.preventDefault();
    var author = this.refs.author.getDOMNode().value.trim();
    var content = this.refs.content.getDOMNode().value.trim();
    if (!content || !author) {
      return;
    }
    this.props.onQuoteSubmit({author: author, content: content});
    this.refs.author.getDOMNode().value = '';
    this.refs.content.getDOMNode().value = '';
    this.refs.content.getDOMNode().focus();
    return;
  },
  render: function() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label className="visuallyhidden" htmlFor="quotefield">Quote</label>
        <div className="quotebox">
          <textarea placeholder="your mamma is so fat..." ref='content'></textarea>
        </div>
        <div className="authorbox" >
          <label htmlFor="authorfield">Author</label>
          <input type="text" ref='author' />
        </div>
        <div className="submitbox">
          <button type="submit">Submit</button>
        </div>
      </form>
    );
  }
});

var QuoteList = React.createClass({
  render: function() {
    var quotes = this.props.quotes.reverse().map(function (quote) {
      return (
        <div className='quote-container' >
          <Quote owner={quote.owner} content={quote.content} author={quote.author} />
        </div>
      );
    });
    return (
      <div>
        {quotes}
      </div>
    );
  }
});

var Quote = React.createClass({
  render: function() {
    return (
      <a className="quote" href="#">
          {this.props.content} - {this.props.author}
      </a>
    );
  }
});