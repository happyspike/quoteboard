<quoteboard>
  <login-form />
  <quote-section />
</quoteboard>

<login-form>
  <div>
    <input type="text" placeholder="email"/>
    <input type="text" placeholder="password"/>
    <button>Sign In</button>
  </div>
</login-form>

<quote-section>
  <quote-form />
  <quote-list name="quotelist" quotes={ quotes } />

  var self = this
  self.quotes = []

  this.on('mount', function() {
    self.firebaseRef = new Firebase('https://quoteboard.firebaseio.com/quotestest');
    self.firebaseRef.on('value', function(snapshot) {
      self.quotes = []
      snapshot.forEach(function(item) {
        self.quotes.push(item.val())
      })
      self.update()
    });
  })

  addQuote(text, author) {
    var quote = {owner: 'wkirschbaum@gmail.com', text: text, author: author}
    self.firebaseRef.push(quote)
  }
</quote-section>

<quote-form>
  <form onsubmit={ submit }>
    <div>
      <textarea name="text" placeholder="Your mamma is so fat..." autofocus />
    </div>
    <div>
      <input name="author" type="text" placeholder="Author" />
    </div>
    <button>Submit</button>
  </form>

  submit(e) {
    this.parent.addQuote(this.text.value, this.author.value)
    this.resetForm()
  }

  resetForm() {
    this.text.value = ''
    this.author.value = ''
    this.text.focus()
  }
</quote-form>

<quote-list>
  <div each={ this.opts.quotes }>
    <quote-item owner={ this.owner } text={ this.text } author={ this.author } />
  </div>
</quote-list>

<quote-item>
  <div>
    <span>{ this.opts.owner }</span> |
    <span>{ this.opts.text }</span> -
    <span>{ this.opts.author }</span>
  </div>
</quote-item>

