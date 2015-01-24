<quoteboard>
  <session-section firebase={ firebase } />
  <quote-section firebase={ firebase } />

  this.firebase = new Firebase('https://quoteboard.firebaseio.com')

</quoteboard>

<session-section>
  <div if={ authenticated }><login-form /></div>
  <div if={ !authenticated }><logout-form /></div>

  var self = this
  var firebase = this.opts.firebase
  self.authenticated = false

  this.on('mount', function() {
    firebase.onAuth(function(authData) {
      if (authData) {
        self.authenticated = false
      } else {
        self.authenticated = true
      }
      self.update()
    })
  })

  login(email, password) {
    firebaseRef.authWithPassword({
      email    : email,
      password : password
    }, function(authdata){

    });
  }

  logout() {
    firebase.unauth()
  }

</session-section>

<register-form>
  register
</register-form>

<logout-form>
  <form onsubmit={ submit }>
    <button>Sign Out</button>
  </form>

  submit(e) {
    e.preventDefault()
    this.parent.logout()
  }

</logout-form

<login-form>
  <form onsubmit={ submit }>
    <input type="text" name="email" placeholder="email"/>
    <input type="password" name="password" placeholder="password"/>
    <button>Sign In</button>
  </form>

  submit(e) {
    e.preventDefault()
    this.parent.login(this.email.value, this.password.value)
  }
</login-form>

<quote-section>
  <quote-form />
  <quote-list name="quotelist" quotes={ quotes } />

  var self = this
  self.quotes = []
  firebaseRef = this.opts.firebase
  firebaseQuotesRef = this.opts.firebase.child('quotestest')

  this.on('mount', function() {
    firebaseQuotesRef.on('value', function(snapshot) {
      self.quotes = []
      snapshot.forEach(function(item) {
        self.quotes.push(item.val())
      })
      self.update()
    });
  })

  addQuote(text, author) {
    var quote = {owner: 'wkirschbaum@gmail.com', text: text, author: author}
    var authData = firebaseRef.getAuth();
    if (authData) {
      firebaseQuotesRef.push(quote)
    } else {
      console.log('please login first')
    }
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
    e.preventDefault()
    var text = this.text.value
    var author = this.author.value

    if (text && author) {
      this.parent.addQuote(this.text.value, this.author.value)
      this.resetForm()
    }
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

