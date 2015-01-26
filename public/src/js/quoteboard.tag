<quoteboard>
  <session-section firebase={ firebase } />
  <quote-section firebase={ firebase } />

  this.firebase = new Firebase('https://quoteboard.firebaseio.com')

</quoteboard>

<session-section>
  <div if={ showLogin }>
    <form onsubmit={ submitLogin }>
      <input type="text" name="loginEmail" placeholder="email"/>
      <input type="password" name="loginPassword" placeholder="password"/>
      <button>Sign In</button>
    </form>
  </div>

  <div if={ showRegister }>
    <form onsubmit={ submitRegister }>
      <input type="text" name="registerEmail" placeholder="email"/>
      <input type="password" name="registerPassword" placeholder="password"/>
      <input type="password" name="registerConfirmPassword" placeholder="confirm password"/>
      <button>Register</button>
    </form>
  </div>

  <div if={ showLogout }>
    <form onsubmit={ submitLogout }>
      { this.userEmail } <button>Sign Out</button>
    </form>
  </div>

  var self = this
  var firebase = this.opts.firebase

  this.on('mount', function() {
    firebase.onAuth(function(authData) {
      if (authData) {
        self.handleAuth({email: authData.password.email})
      } else {
        self.handleAuth()
      }
    })
  })

  handleAuth(userData) {
    if (userData) {
      this.showLogout = true
      this.showLogin = false
      this.userEmail = userData.email
    } else {
      this.showLogout = false
      this.showLogin = true
    }
    this.update()
  }


  submitLogin(e) {
    var email = e.item.loginEmail.value
    var password = e.item.loginPassword.value
    if (email && password) {
      this.login(email, password, function(success) {
        e.item.loginPassword.value = ''
        if (success) {
          e.item.loginEmail.value
        }
      })
    }
  }

  submitRegister(e) {
    var email = e.item.registerEmail.value
    var password = e.item.registerPassword.value
    var confirmPassword = e.item.registerConfirmPassword.value

    if (email && password && confirmPassword && (password == confirmPassword)) {
      this.register(email, password, function(success) {
        e.item.registerPassword.value = ''
        e.item.registerConfirmPassword.value = ''
        if (success) {
          e.item.registerEmail.value
        }
      })
    }
  }

  submitLogout(e) {
    this.logout()
  }

  register(email, password, success) {
    firebaseRef.createUser({
      email    : email,
      password : password
    }, function(error){
      if (error === null) {
        login(email, password, function(success){})
        success(true)
      } else {
        success(false)
        console.log(error)
      }
    });
  }

  login(email, password, success) {
    firebaseRef.authWithPassword({
      email    : email,
      password : password
    }, function(error){
      if (error) {
        console.log(error)
        success(false)
      } else {
        success(true)
      }
    });
  }

  logout() {
    firebase.unauth()
  }

</session-section>

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

  addQuote(text, author, success) {
    var quote = {owner: 'wkirschbaum@gmail.com', text: text, author: author}
    var authData = firebaseRef.getAuth();
    if (authData) {
      firebaseQuotesRef.push(quote)
      if (success) {
        success()
      }
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
      this.parent.addQuote(this.text.value, this.author.value, function() {
        this.resetForm()
      })
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

